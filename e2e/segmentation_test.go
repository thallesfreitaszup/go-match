package e2e

import (
	"context"
	"fmt"
	"github.com/gavv/httpexpect/v2"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go-match/api/request"
	"go-match/cmd/http"
	"go-match/internal/segmentation/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestEchoClient(t *testing.T) {

	req := testcontainers.ContainerRequest{
		Image:        "mongo",
		ExposedPorts: []string{"27017/tcp"},
		WaitingFor:   wait.ForListeningPort("27017").WithStartupTimeout(time.Second),
	}
	container, err := testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if err != nil {
		panic(err)
	}
	host, _ := container.Host(context.Background())
	ports, _ := container.Ports(context.Background())

	port := ports["27017/tcp"][0].HostPort
	mongoUrl := fmt.Sprintf("%s%s:%s", "mongodb://", host, port)
	error := os.Setenv("MONGO_HOST", mongoUrl)
	assert.NoError(t, error)
	app := http.NewApp()
	handler := app.Server()

	server := httptest.NewServer(handler)

	defer server.Close()

	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  server.URL,
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})
	client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUrl))
	testCreateSegmentationRegular(e, container, client, t)
	//testIdentifySimpleKV(e, container, client, t)
}

func testCreateSegmentationRegular(e *httpexpect.Expect, container testcontainers.Container, client *mongo.Client, t *testing.T) {
	ruleNode := request.NodeRequest{}
	ruleNode.Clauses = nil
	ruleNode.Content = getEqualContent("name", "user")
	ruleNode.Type = request.Rule
	ruleNode.LogicalOperator = request.AND
	ruleNode2 := request.NodeRequest{}
	ruleNode2.Clauses = nil
	ruleNode2.Content = getEqualContent("city", "dummy-city")
	ruleNode.Type = request.Rule
	ruleNode.LogicalOperator = request.AND
	node := request.NodeRequest{}
	node.Type = request.Clause
	node.LogicalOperator = request.AND
	node.Clauses = []request.NodeRequest{ruleNode, ruleNode2}
	segmentationRequest := new(request.SegmentationRequest)
	segmentationRequest.Node = node
	segmentationRequest.CircleID = "dummy-circle-id"
	segmentationRequest.WorkspaceID = "dummy-workspace-id"
	segmentationRequest.Name = "circle"
	segmentationRequest.Type = request.Regular

	e.POST("/segmentation").WithJSON(segmentationRequest).Expect().Status(201)

	collection := client.Database("matcher").Collection("segmentation")
	filter := bson.D{
		{"circleId", segmentationRequest.CircleID},
	}
	nodeDB := entity.Segmentation{}
	find := collection.FindOne(context.Background(), filter)
	find.Decode(&nodeDB)
	assert.Equal(t, nodeDB.CircleID, segmentationRequest.CircleID)
	assert.Equal(t, nodeDB.Key, "name_city_")
	assert.Equal(t, nodeDB.Value, "equal(name,'user') && equal(city,'dummy-city')")
	defer container.Terminate(context.Background())
}

func getEqualContent(key, value string) request.Content {
	return request.Content{
		Key:       key,
		Condition: request.Equals,
		Value:     value,
	}
}

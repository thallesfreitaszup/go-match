package test_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go-match/internal/domain/segmentation"
)

var _ = Describe("Segmentation", func() {
	Context("when is a regular segmentation", func() {

		It("should create the correct expression", func() {
			var node = new(segmentation.Node)
			node.Clauses = createClausesWith2Rules()
			node.Type = segmentation.Clause
			node.LogicalOperator = segmentation.AND
			exp := node.Expression()
			fmt.Println(exp)
			Expect(exp).To(Equal("(equal(username,'user1') && equal(city,'dummy-city')) "))
		})

		It("should create the correct expression with multiple groups", func() {
			var node = new(segmentation.Node)
			node.Clauses = createGroupClauses()
			node.Type = segmentation.Clause
			node.LogicalOperator = segmentation.AND
			exp := node.Expression()
			fmt.Println(exp)
			Expect(exp).To(Equal("((equal(username,'user1') && equal(city,'dummy-city')) && (equal(email,'mail@email.com') && equal(some-key,'some-value') && equal(another-key,'another-value'))) "))
		})

	})
})

func createGroupClauses() []segmentation.Node {
	clauses := make([]segmentation.Node, 0)
	clauseWith2Rules := createClausesWith2Rules()
	clauseGroup := segmentation.Node{}
	clauseGroup.Type = segmentation.Clause
	clauseGroup.Clauses = clauseWith2Rules
	clauseGroup.LogicalOperator = segmentation.AND
	clauseWith3Rules := createClausesWith3Rules()
	clauseGroup2 := segmentation.Node{}
	clauseGroup2.Type = segmentation.Clause
	clauseGroup2.Clauses = clauseWith3Rules
	clauseGroup2.LogicalOperator = segmentation.AND
	clauses = append(clauses, clauseGroup)
	clauses = append(clauses, clauseGroup2)

	return clauses
}

func createClausesWith2Rules() []segmentation.Node {
	rule1 := createRule("username", "user1")
	rule2 := createRule("city", "dummy-city")
	clauses := []segmentation.Node{rule1, rule2}
	return clauses
}

func createClausesWith3Rules() []segmentation.Node {
	rule1 := createRule("email", "mail@email.com")
	rule2 := createRule("some-key", "some-value")
	rule3 := createRule("another-key", "another-value")
	clauses := []segmentation.Node{rule1, rule2, rule3}
	return clauses
}

func createRule(key, value string) segmentation.Node {
	var node = segmentation.Node{}
	node.Type = segmentation.Rule
	node.Content = createContent(key, value)
	node.LogicalOperator = segmentation.AND
	return node
}

func createContent(key, value string) segmentation.Content {
	return segmentation.Content{
		Key:       key,
		Condition: segmentation.Equal,
		Value:     value,
	}
}

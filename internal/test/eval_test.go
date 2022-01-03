package test_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
	"go-match/internal/domain/segmentation"
	"go-match/internal/eval"
)

var _ = Describe("Eval", func() {
	Context("when evaluating a contains expressions", func() {

		It("should evaluate successfully", func() {
			exp := "contains(city,'uber')"
			mapParameters := make(map[string]interface{})
			mapParameters["city"] = "uberlandia"
			bool, error := eval.Expression(exp, mapParameters)
			assert.Nil(GinkgoT(), error)
			assert.Equal(GinkgoT(), true, bool)
		})

		It("should evaluate successfully when parameter is a number", func() {
			exp := "contains(city,'44')"
			mapParameters := make(map[string]interface{})
			mapParameters["city"] = 444
			bool, error := eval.Expression(exp, mapParameters)
			assert.Nil(GinkgoT(), error)
			assert.Equal(GinkgoT(), true, bool)
		})

		It("should return false when string not contains another", func() {
			exp := "contains(city,'landia')"
			mapParameters := make(map[string]interface{})
			mapParameters["city"] = "uberaba"
			bool, error := eval.Expression(exp, mapParameters)
			assert.Nil(GinkgoT(), error)
			assert.Equal(GinkgoT(), false, bool)
		})
	})

	Context("when evaluating a equal expressions", func() {

		It("should evaluate successfully", func() {
			exp := "equal(name,'user')"
			mapParameters := make(map[string]interface{})
			mapParameters["name"] = "user"
			bool, error := eval.Expression(exp, mapParameters)
			assert.Nil(GinkgoT(), error)
			assert.Equal(GinkgoT(), true, bool)
		})

		It("should evaluate successfully when parameter is a number", func() {
			exp := "equal(number,5)"
			mapParameters := make(map[string]interface{})
			mapParameters["number"] = 5
			bool, error := eval.Expression(exp, mapParameters)
			assert.Nil(GinkgoT(), error)
			assert.Equal(GinkgoT(), true, bool)
		})

		It("should be not equal when are different numbers", func() {
			exp := "equal(number,6)"
			mapParameters := make(map[string]interface{})
			mapParameters["number"] = 5
			bool, error := eval.Expression(exp, mapParameters)
			assert.Nil(GinkgoT(), error)
			assert.Equal(GinkgoT(), false, bool)
		})

		It("should be not equal when are different strings", func() {
			exp := "equal(name,'user')"
			mapParameters := make(map[string]interface{})
			mapParameters["name"] = "user1"
			bool, error := eval.Expression(exp, mapParameters)
			assert.Nil(GinkgoT(), error)
			assert.Equal(GinkgoT(), false, bool)
		})

	})

	Context("when evaluating  node expression", func() {

		It("should evaluate successfully", func() {
			node := segmentation.Node{}
			node.Clauses = createClausesWith2Rules()
			node.Type = segmentation.Clause
			node.LogicalOperator = segmentation.AND
			mapParameters := make(map[string]interface{})
			mapParameters["username"] = "user1"
			mapParameters["city"] = "dummy-city"
			fmt.Println(node.Expression())
			bool, error := eval.Expression(node.Expression(), mapParameters)
			assert.Nil(GinkgoT(), error)
			assert.Equal(GinkgoT(), true, bool)
		})

		It("should return false when expression is false", func() {
			node := segmentation.Node{}
			node.Clauses = createClausesWith2Rules()
			node.Type = segmentation.Clause
			node.LogicalOperator = segmentation.AND
			mapParameters := make(map[string]interface{})
			mapParameters["username"] = "not-equal-username"
			mapParameters["city"] = "dummy-city"
			fmt.Println(node.Expression())
			bool, error := eval.Expression(node.Expression(), mapParameters)
			assert.Nil(GinkgoT(), error)
			assert.Equal(GinkgoT(), false, bool)
		})

	})
})

package goexpr

import (
	"testing"

	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
)

func TestParsing(t *testing.T) {
	RunSpecs(t, "Goexpr")
}

func testEvaluate(t GinkgoTInterface, expr string, scope map[string]float64, expected float64) {

	parsed, err := Parse(expr)

	assert.Nil(t, err)

	result, err := Evaluate(parsed, scope)

	assert.Nil(t, err)
	assert.Equal(t, result, expected)
}

func testEvaluateError(t GinkgoTInterface, expr string, scope map[string]float64) error {

	parsed, err := Parse(expr)

	assert.Nil(t, err)

	_, err = Evaluate(parsed, scope)

	assert.NotNil(t, err)
	return err
}

func testParseError(t GinkgoTInterface, expr string, scope map[string]float64) error {

	_, err := Parse(expr)

	assert.NotNil(t, err)
	return err
}

var _ = Describe("Sanity", func() {

	It("Evaluates a constant", func() {

		testEvaluate(GinkgoT(), "1", map[string]float64{}, 1)
	})

	It("Evaluates a value", func() {

		scope := map[string]float64{
			"a": 1,
		}

		testEvaluate(GinkgoT(), "1", scope, 1)
	})
})

var _ = Describe("Operations", func() {

	It("Adds values and constants", func() {

		scope := map[string]float64{
			"a": 1,
			"b": 2,
		}

		testEvaluate(GinkgoT(), "a + b + 1", scope, 4)
	})

	It("Subtracts values and constants", func() {

		scope := map[string]float64{
			"a": 3,
			"b": 2,
		}

		testEvaluate(GinkgoT(), "a - b - 1", scope, 0)
	})

	It("Multiplies values and constants", func() {

		scope := map[string]float64{
			"a": 1,
			"b": 2,
		}

		testEvaluate(GinkgoT(), "a * b * 3", scope, 6)
	})

	It("Divides values and constants", func() {

		scope := map[string]float64{
			"a": 12,
			"b": 6,
		}

		testEvaluate(GinkgoT(), "a / b / 2", scope, 1)
	})
})

var _ = Describe("Error conditions", func() {

	It("Errors on evaluating with a missing scope value", func() {

		t := GinkgoT()
		scope := map[string]float64{
			"a": 1,
		}

		err := testEvaluateError(t, "a + b", scope)
		assert.Regexp(t, "scope", err.Error())
	})

	It("Errors on unsupported binary operation", func() {

		t := GinkgoT()
		scope := map[string]float64{
			"a": 1,
			"b": 2,
		}

		err := testParseError(t, "a & b", scope)
		assert.Regexp(t, "operation", err.Error())
	})

})

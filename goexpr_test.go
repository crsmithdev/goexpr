package goexpr

import (
	"testing"

	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
)

func TestParsing(t *testing.T) {
	RunSpecs(t, "Goexpr")
}

type Struct1 struct {
	value float64
}

var _ = Describe("Goexpr", func() {

	It("Evaluates a simple integer constant", func() {

		parsed, err := Parse("1")
		t := GinkgoT()

		assert.Nil(t, err)

		scope := map[string]interface{}{}
		result, err := Evaluate(parsed, scope)

		assert.Nil(t, err)
		assert.Equal(t, result, 1)
	})

	It("Evaluates a simple integer value", func() {

		parsed, err := Parse("a")
		t := GinkgoT()

		assert.Nil(t, err)

		scope := map[string]interface{}{
			"a": 1,
		}

		result, err := Evaluate(parsed, scope)

		assert.Nil(t, err)
		assert.Equal(t, result, 1)
	})

	It("Evaluates a simple float constant", func() {

		parsed, err := Parse("1.0")
		t := GinkgoT()

		assert.Nil(t, err)

		scope := map[string]interface{}{}
		result, err := Evaluate(parsed, scope)

		assert.Nil(t, err)
		assert.Equal(t, result, 1.0)
	})

	It("Evaluates a simple float value", func() {

		parsed, err := Parse("a")
		t := GinkgoT()

		assert.Nil(t, err)

		scope := map[string]interface{}{
			"a": 1.0,
		}

		result, err := Evaluate(parsed, scope)

		assert.Nil(t, err)
		assert.Equal(t, result, 1.0)
	})

})

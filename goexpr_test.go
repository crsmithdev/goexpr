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

	It("Evaluates a simple integer expression", func() {

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

	It("Evaluates a simple float expression", func() {

		parsed, err := Parse("a")
		t := GinkgoT()

		assert.Nil(t, err)

		scope := map[string]interface{}{
			"a": 1.0,
		}

		result, err := Evaluate(parsed, scope)

		assert.Nil(t, err)
		assert.Equal(t, result, 1)
	})

	It("Evaluates a simple selection expression", func() {

		parsed, err := Parse("a.b")
		t := GinkgoT()

		assert.Nil(t, err)

		obj := struct {
			b float64
		}{1}
		scope := map[string]interface{}{
			"a": obj,
		}

		result, err := Evaluate(parsed, scope)

		assert.Nil(t, err)
		assert.Equal(t, result, 1)
	})
})

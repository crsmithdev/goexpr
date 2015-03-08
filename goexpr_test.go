package goexpr

import (
	"fmt"
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

	It("Evaluates a simple expression", func() {

		parsed, err := Parse("a + b")
		t := GinkgoT()

		assert.Nil(t, err)

		scope := map[string]float64{
			"a": 1,
			"b": 2,
		}

		result, err := Evaluate(parsed, scope)

		assert.Nil(t, err)
		assert.Equal(t, result, 3)
	})

	It("Evalutes an expression with selection", func() {

		parsed, err := Parse("a.b + c")
		t := GinkgoT()

		assert.Nil(t, err)

		scope := map[string]interface{}{
			"a": Struct1{1},
			"c": 2,
		}

		fmt.Println(parsed)
		fmt.Println(scope)

		//result, err := Evaluate(parsed, scope)

	})
})

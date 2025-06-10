package rim

import (
	"github.com/abhub23/go-utility/rim"
	"testing"
)

func TestFlames(t *testing.T) {

	type Tests struct {
		testName     string
		name1, name2 string
		Expected     string
	}

	tests := []Tests{
		{"Javascript and Typescript", "Javascript", "Typescript", "javascript and typescript have a relation of Affection"},
		{"Go and Rust", "Go", "Rust", "go and rust have a relation of Marriage"},
	}

	for _, el := range tests {
		t.Run(el.testName, func(t *testing.T) {
			result, err := rim.Flames(el.name1, el.name2)
			if err != nil {
				panic(err)
			}

			if result != el.Expected{
				t.Errorf("func was expecting %s but got %s ", el.Expected, result )
			}
		})
	}

}

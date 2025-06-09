package rim

import (
	"github.com/abhub23/go-utility/rim"
	"testing"
)

func TestFlames(t *testing.T) {
	got, err := rim.Flames("Javascript", "Typescript")
	if err != nil {
		t.Errorf("Error occured %v", err)
	}

	want := "javascript and typescript have a relation of Affection"

	if got != want {
		t.Errorf("Flames(Abdullah, Rimsha) got %s and want %s", got, want)
	}
}

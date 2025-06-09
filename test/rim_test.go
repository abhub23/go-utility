package rim

import (
	
	"testing"
	"github.com/abhub23/go-utility/rim"
)

func TestFlames(t *testing.T) {
	got, err := rim.Flames("Javascript", "Typescript")
	if err != nil {
		t.Errorf("Error occured %v", err)
	}
	
	want := "Javascript and Typescript have a relation of Affection"

	if got != want {
		t.Errorf("Flames(Abdullah, Rimsha) got %s and want %s", got, want)
	}
}

package rim

import (
	"errors"
	"fmt"
	"strings"
)


func Flames(name1, name2 string) (string, error) {

	if name1 == "" || name2 == "" {
		return "", errors.New("missing input strings")
	}

	name1 = strings.ToLower(strings.ReplaceAll(name1, " ", ""))
	name2 = strings.ToLower(strings.ReplaceAll(name2, " ", ""))

	slice1 := []rune(name1)
	slice2 := []rune(name2)

	i := 0
	for i < len(slice1) {
		found := false
		for j := 0; j < len(slice2); j++ {
			if slice1[i] == slice2[j] {
				slice1 = append(slice1[:i], slice1[i+1:]...)
				slice2 = append(slice2[:j], slice2[j+1:]...)
				found = true
				break
			}
		}
		if !found {
			i++
		}
	}

	total := len(slice1) + len(slice2)
    flames := []rune{'F','L','A','M','E','S'}
    
    relation := Findrelation(total, flames)



	return fmt.Sprintf("%s and %s have a relation of %s", name1, name2, relation), nil

}

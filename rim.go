package rim

import (
	"strings"
)

func Flames(name1 string, name2 string) string {
	name1 = strings.ToLower(strings.ReplaceAll(name1, " ",""))
	name2 = strings.ToLower(strings.ReplaceAll(name2, " ",""))
	Slice := []byte{}
	for i := 0; i< len(name1); i++ {
		for j := 1; j < len(name2); j++{
			if name1[i] != name2[j] {
				Slice = append(Slice, name1[i])
			}
		}
	} 

	flamesmap := make(map[rune]string)

	flamesmap['F'] = "Friends"
	flamesmap['L'] = "Lovers"
	flamesmap['A'] = "Affection"
	flamesmap['M'] = "Marriage"
	flamesmap['E'] = "Enemies"
	flamesmap['S'] = "Siblings"

	return string(Slice)
}

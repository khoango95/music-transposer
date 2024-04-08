package music_transposer

import (
	"slices"
	"strings"
)

func firstN(s string, n int) string {
	if len(s) > n {
		return s[:n]
	}
	return s
}

func Transpose(inputText string, transposeValue int) string {
	chordList := []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
	// inputText := "C# B A Gm7 F#m7 D# C E7 A#m"
	// var transposeValue int
	// fmt.Scan(&transposeValue)

	outputText := ""

	text := strings.Split(inputText, " ")

	for _, item := range text {
		itemFiltered := firstN(item, 1)

		if len(item) > 1 {
			if string(item[1]) == "#" {
				itemFiltered = firstN(item, 2)
			}
		}

		index := slices.Index(chordList, itemFiltered)
		newIndex := index + transposeValue

		if newIndex < 0 {
			newIndex += 12
		} else if newIndex > 11 {
			newIndex -= 12
		}

		item = strings.Replace(item, chordList[index], chordList[newIndex], -1)
		outputText += item + " "
	}

	// fmt.Println("-Chords transposed:\t", outputText)
	return outputText
}

package main

import (
	"fmt"
	transposer "music-transposer"
)

func main() {

	fmt.Print("Enter chords: ")
	// inputText, _ := reader.ReadString('\n')

	inputText := "C# B A Gm7 F#m7 D# C E7 A#m"
	fmt.Println("\t\t " + inputText)

	fmt.Print("Enter transpose value: ")
	var transposeValue int
	fmt.Scan(&transposeValue)

	outputText := transposer.Transpose(inputText, &transposeValue)

	fmt.Println("-Chords transposed:\t", outputText)
}

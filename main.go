package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/atotto/clipboard"
	"github.com/go-loremipsum/loremipsum"
)

func main() {
	var generatedType string
	var generatedLength int
	flag.StringVar(&generatedType, "t", "sentences", "Type of lorem to generate")
	flag.IntVar(&generatedLength, "n", 3, "Length of lorem to generate")
	flag.Parse()

	loremGenerator := loremipsum.New()
	generatedText := ""
	if generatedType == "sentences" {
		generatedText = loremGenerator.Sentences(generatedLength)
	}
	if generatedType == "words" {
		generatedText = loremGenerator.Words(generatedLength)
	}
	if generatedType == "paragraphs" {
		generatedText = loremGenerator.Paragraphs(generatedLength)
	}
	if generatedText == "" {
		fmt.Println("Invalid type - choose 'sentences', 'words', or 'paragraphs'")
		os.Exit(1)
	}

	err := clipboard.WriteAll(generatedText)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generatedText)
}

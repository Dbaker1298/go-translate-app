package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	sourceLang string
	targetLang string
	sourceText string
)

func init() {
	flag.StringVar(&sourceLang, "s", "en", "Source language[en]")
	flag.StringVar(&targetLang, "t", "fr", "Target language[fr]")
	flag.StringVar(&sourceText, "st", "", "Test to translate")
}

func main() {
	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Println("OPtions: ")
		flag.PrintDefaults()
		os.Exit(1)
	}

	reqBody := &cli.RequestBody{
		SourceLang: sourceLang,
		TargetLang: targetLang,
		SourceText: sourceText,
	}

	cli.RequestTranslate(reqBody)
}

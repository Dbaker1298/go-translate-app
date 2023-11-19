package cli

import (

)

type RequestBody struct {
	SourceLang string
	TargetLang string
	SourceText string
}

const translateURL = "https://translation.googleapis.com/translate_a/single"  // URL for Google Translate API, needs to be updated 

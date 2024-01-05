//
// Copyright (c) 2021 Markku Rossi.
//
// All rights reserved.
//

package main

import (
	"fmt"
	"log"

	"github.com/jdkato/prose/v2"
)

func main() {
	fmt.Println("Hello, world!")
	// Create a new document with the default configuration:
	doc, err := prose.NewDocument("Package aes implements the Advanced Encryption Standard (AES) block cipher operations.")
	if err != nil {
		log.Fatal(err)
	}

	// Iterate over the doc's tokens:
	fmt.Println("- Tokens:")
	for _, tok := range doc.Tokens() {
		fmt.Println(tok.Text, tok.Tag, tok.Label)
		// Go NNP B-GPE
		// is VBZ O
		// an DT O
		// ...
	}

	// Iterate over the doc's named-entities:
	fmt.Println("- Entities:")
	for _, ent := range doc.Entities() {
		fmt.Println(ent.Text, ent.Label)
		// Go GPE
		// Google GPE
	}

	// Iterate over the doc's sentences:
	fmt.Println("- Sentences:")
	for _, sent := range doc.Sentences() {
		fmt.Println(sent.Text)
		// Go is an open-source programming language created at Google.
	}
}

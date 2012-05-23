// Copyright 2010 Hannes Baldursson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file is part of go-idn

// +build ignore

// Stringprep table generator
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"net/http"
	"log"
	"strings"
)

func main() {
	flag.Parse()
	loadDefaultTables()
}

var url = flag.String("url", "http://www.ietf.org/rfc/rfc3454.txt", "URL of RFC 3453 (stringprep)")

var logger = log.New(os.Stderr, "", log.Lshortfile)

// This contains only the properties we're interested in.
type Char struct {
	codepoint rune // if zero, this index is not a valid code point.
}

var chars = make([]Char, unicode.MaxChar+1)

func loadDefaultTables() {
	// get rfc3454
	resp, err := http.Get(*url)
	if err != nil {
		logger.Fatal(err)
	}
	if resp.StatusCode != 200 {
		
		logger.Fatal("bad GET status for rfc3454.txt", resp.Status)
	}
	f := resp.Body
	defer f.Close()
	
	// read rfc3454
	input := bufio.NewReader(f)
	for {
		line, err  input.ReadString("\n")
		if err != nil {
			if err == io.EOF {
				break
			}
			logger.Fatal(err)
		}
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "----- Start Table") {
			parseTable(input)
		}
	}
	return
}


// parseTable returns after it reaches "---- End Table -----"
func parseTable(input *bufio.Reader) {
	
}
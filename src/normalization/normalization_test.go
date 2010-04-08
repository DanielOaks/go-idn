// Copyright 2010 Hannes Baldursson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file is part of go-idn

package normalization


import (
	"fmt"
	"reflect"
	"testing"
	"bufio"
	"http"
	"strconv"
	"os"
	"unicode"
	"log"
	"strings"
)

var die = log.New(os.Stderr, nil, "", log.Lexit|log.Lshortfile)

// Random samples from the decompose index
var decomposeTests = [][]int {
    []int{0x00A0, 0},
    []int{0x00A8, 1},
    []int{0x00AA, 2},
    []int{0x00AF, 3},
    []int{0x3161, 1518},
    []int{0x3162, 1519},
    []int{0x3163, 1520},
    []int{0x3164, 1521},
	[]int{0xFFEC, 2826},
    []int{0xFFED, 2827},
    []int{0xFFEE, 2828},
}

func TestDecomposeIndex(t *testing.T) {

	for _, tt := range decomposeTests {
		var out int
		
		out = decomposeIndex(tt[0])
		if !reflect.DeepEqual(out, tt[1]) {
			t.Errorf("decomposeIndex(%v) = %v; want %v", string(tt[0]), string(out), string(tt[1]))
		}
	}
}



type removeTest struct {
	Def []int
	Pos0 []int
	Pos6 []int
	Pos10 []int
}

var _removeTest = removeTest {
	[]int {0, 1, 2, 3, 4, 5, 6, 7, 8, 9 ,10},
	[]int { 1, 2, 3, 4, 5, 6, 7, 8, 9 ,10},
	[]int {0, 1, 2, 3, 4, 5, 7, 8, 9 ,10},
	[]int {0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
}

func TestRemove(t *testing.T) {
	var out []int
	
	out = remove(_removeTest.Def, 0)
	if !reflect.DeepEqual(_removeTest.Pos0, out) {
		t.Errorf("remove(%v) == %v; want %v", hex32(_removeTest.Def), hex32(_removeTest.Pos0), hex32(out))
	}
	
	out = remove(_removeTest.Def, 6)
	if !reflect.DeepEqual(_removeTest.Pos6, out) {
		t.Errorf("remove(%v) == %v; want %v", hex32(_removeTest.Def), hex32(_removeTest.Pos6), hex32(out))
	}
	
	out = remove(_removeTest.Def, 10)
	if !reflect.DeepEqual(_removeTest.Pos10, out) {
		t.Errorf("remove(%v) == %v; want %v", hex32(_removeTest.Def), hex32(_removeTest.Pos10), hex32(out))
	}
}

type hex32 []int

func (h hex32) Format(f fmt.State, c int) {
	fmt.Fprint(f, "[")
	for i, v := range h {
		if i > 0 {
			fmt.Fprint(f, " ")
		}
		fmt.Fprintf(f, "%x", v)
	}
	fmt.Fprint(f, "]")
}

type _nfkcTest struct {
	C1 []int
	C2 []int
	C3 []int
	C4 []int 
	C5 []int
}

// Downloads the NormalizationTest.txt from unicode.org and tests all rows as described in the file
func TestNKFC(t *testing.T) {
	resp, _, err:= http.Get("http://www.unicode.org/Public/"+ unicode.Version +"/ucd/NormalizationTest.txt")
	
		if err != nil {
		die.Log(err)
	}
	if resp.StatusCode != 200 {
		die.Log("bad GET status for NormalizationTest.txt", resp.Status)
	}
	input := bufio.NewReader(resp.Body)
	
		for i:=0; true; i++ {
		line, err := input.ReadString('\n')
		if err != nil {
			if err == os.EOF {
				break
			}
			die.Log(err)
		}
		
		if i == 5999 {
			break // 
		}
		
		line = strings.TrimSpace(line)
		
		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, "@Part"){
		 // ignore 
		} else {
			
			// Remove comments
			if strings.LastIndex(line, "#") != -1 {
				line = strings.Split(line, "#",0)[0]
			}
			
			
			columns_s := strings.Split(line, ";",0)
			
			var test _nfkcTest
			
			for i:=0; i < 5; i++ {
				cps:= strings.Fields(columns_s[i])
				column:= make([]int, len(cps))
				for j:=0; j < len(cps); j++ {
					column_j, err := strconv.Btoi64("0x" + cps[j], 0)
					if err != nil {
						die.Log("Line 147: " + err.String())
					}
					column[j] = int(column_j)
				}
				switch i {
					case 0:
						test.C1 = column
						break
					case 1: 
						test.C2 = column
						break
					case 2:
						test.C3 = column
						break
					case 3:
						test.C4 = column
						break
					case 4: 
						test.C5 = column
						break
				}
			}
			
			if !reflect.DeepEqual(test.C4, NFKC(test.C1)) {
				t.Errorf("NormalizeNFCK(%v) == %v; want %v", hex32(test.C1), hex32(NFKC(test.C1)), hex32(test.C4))
			}
			if !reflect.DeepEqual(test.C4, NFKC(test.C2)) {
				t.Errorf("NormalizeNFCK(%v) == %v; want %v", hex32(test.C2), hex32(NFKC(test.C2)), hex32(test.C4))
			}
			if !reflect.DeepEqual(test.C4, NFKC(test.C3)) {
				t.Errorf("NormalizeNFCK(%v) == %v; want %v", hex32(test.C3), hex32(NFKC(test.C3)), hex32(test.C4))
			}
			if !reflect.DeepEqual(test.C4, NFKC(test.C4)) {
				t.Errorf("NormalizeNFCK(%v) == %v; want %v", hex32(test.C4), hex32(NFKC(test.C4)), hex32(test.C4))
			}
			if !reflect.DeepEqual(test.C4, NFKC(test.C5)) {
				t.Errorf("NormalizeNFCK(%v) == %v; want %v", hex32(test.C5), hex32(NFKC(test.C5)), hex32(test.C4))
			}
		}	
		
	}
	resp.Body.Close()
}

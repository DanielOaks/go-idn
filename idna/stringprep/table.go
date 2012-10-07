package stringprep

type TableElement struct {
	Lo  int
	Hi  int
	Map d // can be empty
}

type Table []TableElement

// Exports tables
var Tables = map[string]Table{
	"A1":  _A1,
	"B1":  _B1,
	"B2":  _B2,
	"B3":  _B3,
	"C11": _C11,
	"C12": _C12,
	"C21": _C21,
	"C22": _C22,
	"C3":  _C3,
	"C4":  _C4,
	"C5":  _C5,
	"C6":  _C6,
	"C7":  _C7,
	"C8":  _C8,
	"C9":  _C9,
	"D1":  _D1,
	"D2":  _D2,
}

// Returns true if the rune is in table 
func in_table(c int, table Table) bool {
	for i := 0; i < len(table); i++ {
		if table[i].Lo <= c && c <= table[i].Hi {
			return true
		}
	}
	return false
}

// Returns a filtered rune sequence 
func filter(input []int, table Table) []int {
	output := make([]int, len(input))
	c := 0 // count

	for i := 0; i < len(input); i++ {
		if !in_table(input[i], table) {
			output[c] = input[i]
			c++
		}
	}

	return output[0:len(output)]
}

// Iterates over the input rune array and replaces runes with their maps
func map_table(input []int, table Table) []int {

	output := make([]int, len(input))
	c := 0 // count

	for i := 0; i < len(input); i++ {
		// If rune is in table, replace it with its map
		if in_table(input[i], table) {
			for k := 0; k < len(table); k++ {
				if input[i] == table[k].Lo {
					copy(output[c:], table[k].Map[0:len(table[k].Map)])
					c += len(table[k].Map)
					break
				}
			}
		} else {
			output[c] = input[i]
			c++
		}

	}
	return output[0:len(output)]
}
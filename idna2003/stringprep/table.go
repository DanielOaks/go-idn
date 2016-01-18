package stringprep

type TableElement struct {
	Lo  rune
	Hi  rune
	Map d // can be empty
}

type Table []TableElement

// Returns true if the rune is in table
func in_table(c rune, table Table) bool {
	for i := 0; i < len(table); i++ {
		if table[i].Lo <= c && c <= table[i].Hi {
			return true
		}
	}
	return false
}

// Returns a filtered rune sequence
func filter(input []rune, table Table) []rune {
	output := make([]rune, len(input))
	c := 0 // count

	for i := 0; i < len(input); i++ {
		if !in_table(input[i], table) {
			output[c] = input[i]
			c++
		}
	}

	return output[0:]
}

// Iterates over the input rune array and replaces runes with their maps
func map_table(input []rune, table Table) []rune {
	output := make([]rune, len(input))
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
	return output[0:]
}

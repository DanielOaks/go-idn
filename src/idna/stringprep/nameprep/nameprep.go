// Package nameprep contains the profile data required to implement Nameprep. 
// The package is typically only imported for the side effect of registering its tables.
//
// To use nameprep, link this package into your program:
//     import _ "idna/stringprep/nameprep"
// Then use stringprep.Prep("nameprep", b) 
package nameprep

import (
	"idna/stringprep"
)

func init() {
	stringprep.RegisterProfile(Alias, prof)
}

// Alias is the name that nameprep is registered to with stringprep
const Alias = "nameprep"

var prof = &stringprep.Profile{
	
}

/*
 * Mapping Table:
 * The data in mapping table is sorted according to the length of the mapping sequence.
 * If the type of the code point is USPREP_MAP and value in trie word is an index, the index
 * is compared with start indexes of sequence length start to figure out the length according to
 * the following algorithm:
 *
 *              if(       index >= indexes[_SPREP_ONE_UCHAR_MAPPING_INDEX_START] &&
 *                        index < indexes[_SPREP_TWO_UCHARS_MAPPING_INDEX_START]){
 *                   length = 1;
 *               }else if(index >= indexes[_SPREP_TWO_UCHARS_MAPPING_INDEX_START] &&
 *                        index < indexes[_SPREP_THREE_UCHARS_MAPPING_INDEX_START]){
 *                   length = 2;
 *               }else if(index >= indexes[_SPREP_THREE_UCHARS_MAPPING_INDEX_START] &&
 *                        index < indexes[_SPREP_FOUR_UCHARS_MAPPING_INDEX_START]){
 *                   length = 3;
 *               }else{
 *                   // The first position in the mapping table contains the length 
 *                   // of the sequence
 *                   length = mappingTable[index++];
 *        
 *               }
*/
mappingTable []uintt16
spreptrie 
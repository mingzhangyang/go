package utilities

import (
	"strings"
)

func splitLine(str string, sep rune) []string {
	elem := ""
	var st int
	res := make([]string, 0)
	for _, v := range str {
		switch v {
		case '"':
			if st == 0 {
				st++
			} else {
				st--
			}
		case sep:
			if st == 0 {
				res = append(res, strings.TrimSpace(elem))
				elem = ""
				continue
			}
		}
		elem += string(v)
	}

	res = append(res, strings.TrimSpace(elem))
	return res
}


func splitBytes(str []byte, sep byte) []string {
	elem := ""
	var st int
	res := make([]string, 0)
	for _, v := range str {
		switch v {
		case byte('"'):
			if st == 0 {
				st++
			} else {
				st--
			}
		case sep:
			if st == 0 {
				res = append(res, elem)
				elem = ""
				continue
			}
		}
		elem += string(v)
	}

	res = append(res, elem)
	return res
}

// This version of splitLine will reuse a 
func splitLine1(collector *[]string, str string, sep rune) {
	elem := ""
	var st int
	var i int
	
	for _, v := range str {
		switch v {
		case '"':
			if st == 0 {
				st++
			} else {
				st--
			}
		case sep:
			if st == 0 {
				(*collector)[i] = elem
				elem = ""
				i++
				continue
			}
		}
		elem += string(v)
	}

	(*collector)[i] = elem
}

// func main() {
// 	s := `10,2,tRNA (cytidine(34)-2'-O)-methyltransferase TrmL catalyzes the methyl transfer from S-adenosyl-L-methionine to the ribose at the nucleotide 34 wobble position in the two leucyl isoacceptors tRNA(Leu)(CmAA) and tRNA(Leu)(cmnm5UmAA),tRNA (cytidine(34)-2'-O)-methyltransferase,SpoU_methylase,COG0219,1,1,published,17:04.8,17:04.8,curated`
// 	//s := os.Args[1]
// 	v := SplitLine(s, ",")
// 	for i := 0; i < len(v); i++ {
// 		fmt.Printf("i = %d, substr: %s\n", i, string(v[i]))
// 	}
// }
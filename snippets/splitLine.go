package main

import(
	"fmt"
	"os"
	"strings"
)

type State struct{
	open bool
	sign string
	ready bool
}

func SplitLine(str, sep string) []string {
	elem := ""
	state := State{false, "", false}
	res := []string{}

	for i := 0; i < len(str); i++ {
		if state.open {
			elem += string(str[i])
			if state.sign == string(str[i]) {
				state.open = false
				state.ready = true
			}
		} else {
			if string(str[i]) == sep {
				if state.ready {
					res = append(res, strings.TrimSpace(elem))
					state.sign = ""
					state.ready = false
					elem = ""
					continue
				}
				res = append(res, strings.TrimSpace(elem))
				elem = ""
				continue
			}
			if string(str[i]) == "\"" {
				state.open = true
				state.sign = "\""
				elem += "\""
			}
			elem += string(str[i])
		}
	}

	res = append(res, strings.TrimSpace(elem))
	return res
}

func main() {
	//s := `10,2,tRNA (cytidine(34)-2'-O)-methyltransferase TrmL catalyzes the methyl transfer from S-adenosyl-L-methionine to the ribose at the nucleotide 34 wobble position in the two leucyl isoacceptors tRNA(Leu)(CmAA) and tRNA(Leu)(cmnm5UmAA),tRNA (cytidine(34)-2'-O)-methyltransferase,SpoU_methylase,COG0219,1,1,published,17:04.8,17:04.8,curated`
	s := os.Args[1]
	v := SplitLine(s, ",")
	for i := 0; i < len(v); i++ {
		fmt.Printf("i = %d, substr: %s\n", i, string(v[i]))
	}
}
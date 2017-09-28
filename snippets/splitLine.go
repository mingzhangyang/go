package main

import(
	"fmt"
	//"os"
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
		c := string(str[i])
		elem += c

		switch c {
		case "\"":
			if state.open {
				state.open = false
				state.ready = true
			} else {
				state.open = true
				state.sign = "\""
			}
		case sep:
			if state.open == false {
				res = append(res, strings.TrimSpace(elem))
				elem = ""
				continue
			}
			if state.ready {
				state.sign = ""
				state.ready = false
				res = append(res, strings.TrimSpace(elem))
				elem = ""
			}
		}
	}

	res = append(res, strings.TrimSpace(elem))
	return res
}

func main() {
	s := `10,2,tRNA (cytidine(34)-2'-O)-methyltransferase TrmL catalyzes the methyl transfer from S-adenosyl-L-methionine to the ribose at the nucleotide 34 wobble position in the two leucyl isoacceptors tRNA(Leu)(CmAA) and tRNA(Leu)(cmnm5UmAA),tRNA (cytidine(34)-2'-O)-methyltransferase,SpoU_methylase,COG0219,1,1,published,17:04.8,17:04.8,curated`
	//s := os.Args[1]
	v := SplitLine(s, ",")
	for i := 0; i < len(v); i++ {
		fmt.Printf("i = %d, substr: %s\n", i, string(v[i]))
	}
}
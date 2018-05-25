package utilities

import(
	//"fmt"
	//"os"
	"strings"
)

type state struct{
	open bool
	sign string
	ready bool
}

func splitLine(str, sep string) []string {
	elem := ""
	st := state{false, "", false}
	res := []string{}

	for _, r := range str {
		c := string(r)

		switch c {
		case "\"":
			if st.open {
				st.open = false
				st.ready = true
			} else {
				st.open = true
				st.sign = "\""
			}
		case sep:
			if st.open == false {
				res = append(res, strings.TrimSpace(elem))
				elem = ""
				continue
			}
			if st.ready {
				st.sign = ""
				st.ready = false
				res = append(res, strings.TrimSpace(elem))
				elem = ""
			}
		}
		elem += c
	}

	res = append(res, strings.TrimSpace(elem))
	return res
}

// func main() {
// 	s := `10,2,tRNA (cytidine(34)-2'-O)-methyltransferase TrmL catalyzes the methyl transfer from S-adenosyl-L-methionine to the ribose at the nucleotide 34 wobble position in the two leucyl isoacceptors tRNA(Leu)(CmAA) and tRNA(Leu)(cmnm5UmAA),tRNA (cytidine(34)-2'-O)-methyltransferase,SpoU_methylase,COG0219,1,1,published,17:04.8,17:04.8,curated`
// 	//s := os.Args[1]
// 	v := SplitLine(s, ",")
// 	for i := 0; i < len(v); i++ {
// 		fmt.Printf("i = %d, substr: %s\n", i, string(v[i]))
// 	}
// }
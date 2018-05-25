package utilities

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

// func main() {
// 	s := `10,2,tRNA (cytidine(34)-2'-O)-methyltransferase TrmL catalyzes the methyl transfer from S-adenosyl-L-methionine to the ribose at the nucleotide 34 wobble position in the two leucyl isoacceptors tRNA(Leu)(CmAA) and tRNA(Leu)(cmnm5UmAA),tRNA (cytidine(34)-2'-O)-methyltransferase,SpoU_methylase,COG0219,1,1,published,17:04.8,17:04.8,curated`
// 	//s := os.Args[1]
// 	v := SplitLine(s, ",")
// 	for i := 0; i < len(v); i++ {
// 		fmt.Printf("i = %d, substr: %s\n", i, string(v[i]))
// 	}
// }
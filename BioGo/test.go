package main

import "fmt"
import "./sequence"

func main() {
	s, _ := sequence.NewDNA("ATGCCTT")
	fmt.Println(s.Seq())
	fmt.Println(s.GCContent())
	fmt.Println(s.Complementary())
	fmt.Println(s.Reverse())
	fmt.Println(s.Composition())
	a := sequence.Protein{Name: "cc"}
	a.SetSeq("sjfiosflsdfsd")
	fmt.Println(a)
}

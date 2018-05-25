package utilities

import (
	"testing"
)

func TestSplitLineB(t *testing.T) {
	lines := []string{
		`a, b, c, d`,
      `"A", B, C, D`,
      `"A, B", c, d, e`,
      `1, 2, "OK", ,test"Hello World", 567, "","I am fine.", "A, B, C"`,
      `10,2,tRNA (cytidine(34)-2\'-O)-methyltransferase TrmL catalyzes the methyl transfer from S-adenosyl-L-methionine to the ribose at the nucleotide 34 wobble position in the two leucyl isoacceptors tRNA(Leu)(CmAA) and tRNA(Leu)(cmnm5UmAA),tRNA (cytidine(34)-2\'-O)-methyltransferase,SpoU_methylase,COG0219,1,1,published,17:04.8,17:04.8,curated`,
      `10000155,HY190125,03/19/2015 02:30:00 AM,008XX E 65TH ST,0486,BATTERY,DOMESTIC BATTERY SIMPLE,APARTMENT,false,true,0312,003,20,42,08B,1183157,1862099,2015,02/10/2018 03:50:01 PM,41.776797342,-87.604098905,"(41.776797342, -87.604098905)"`,
	}
	res := []int{4, 4, 4, 9, 12, 22}
	for i, line := range lines {
		x := splitLine(line, ',')
		if len(x) != res[i] {
			t.Errorf("wrong number, expect %d, but get %d\n", res[i], len(x))
		}
		t.Logf("%d, %d, %d, %t\n", i, res[i], len(x), (res[i] == len(x)))
	}
}
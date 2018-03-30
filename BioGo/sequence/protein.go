package sequence

import "fmt"

type Protein struct {
	Name string
	seq  string
}

func (p *Protein) Length() int {
	return len(p.seq)
}

func (p *Protein) SetSeq(s string) {
	p.seq = s
}

func (p *Protein) Seq() string {
	return p.seq
}

func (p Protein) String() string {
	return fmt.Sprintf("Name: %s | Length: %d\n%s", p.Name, p.Length(), p.Seq())
}

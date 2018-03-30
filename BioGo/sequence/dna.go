package sequence

import (
	"errors"
	"fmt"
	"strings"
)

type DNA struct {
	Name string
	seq  string
}

const (
	a uint8 = 65
	t uint8 = 84
	c uint8 = 67
	g uint8 = 71
)

func NewDNA(s ...string) (DNA, error) {
	if len(s) > 1 {
		return DNA{
			Name: s[1],
			seq:  s[0],
		}, nil
	} else if len(s) == 1 {
		return DNA{seq: s[0]}, nil
	}
	return DNA{}, errors.New("At least a string is required")
}

func (dna *DNA) Length() int {
	return len(dna.seq)
}

func (dna *DNA) StdSequence() string {
	return strings.ToUpper(string(dna.seq))
}

func (dna *DNA) Seq() string {
	return dna.seq
}

func (dna *DNA) GCContent() (float64, error) {
	seq := dna.StdSequence()
	var n int
	for i := 0; i < len(seq); i++ {
		switch seq[i] {
		case g, c:
			n += 1
		case a, t:
		default:
			return 0, errors.New("Illeagle chars other than 'ATGC' found")
		}
	}
	return float64(n) / float64(dna.Length()), nil
}

func (dna *DNA) Reverse() DNA {
	b := []byte(string(dna.seq))
	h := len(b)
	for i, j := 0, h-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return DNA{seq: string(b)}
}

func (dna *DNA) Complementary() (DNA, error) {
	s := dna.StdSequence()
	b := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case a:
			b[i] = t
		case t:
			b[i] = a
		case g:
			b[i] = c
		case c:
			b[i] = g
		default:
			return DNA{}, errors.New("Illeagle chars other than 'ATGC' found")
		}
	}
	return DNA{seq: string(b)}, nil
}

func (dna *DNA) SetName(s string) {
	dna.Name = s
}

func (dna *DNA) SetSeq(s string) {
	dna.seq = s
}

func (dna *DNA) Composition() (map[string]int, error) {
	n := dna.Length()
	s := dna.StdSequence()
	cp := make(map[string]int)
	for i := 0; i < n; i++ {
		switch s[i] {
		case a:
			cp["A"] += 1
		case t:
			cp["T"] += 1
		case g:
			cp["G"] += 1
		case c:
			cp["C"] += 1
		default:
			return nil, errors.New("Illeage chars other than 'ATGC' found")
		}
	}
	return cp, nil
}

func (dna DNA) String() string {
	return fmt.Sprintf("Name: %s | Length: %d\n%s", dna.Name, dna.Length(), dna.Seq())
}

package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// Team represent a team
type Team struct {
	id int
	scores int
	log [3]int
}

// ChanceDistribution set the probabilities of win and draw
type ChanceDistribution struct {
	win, draw float64
}

// NewCD create an instance of ChanceDistribution
func NewCD(f1, f2 float64) *ChanceDistribution {
	return &ChanceDistribution{f1, f2}
}

func (t *Team) MatchWith (team *Team, n int, chances *ChanceDistribution) {
	r := rand.Float64()
	switch {
	case r < chances.win:
		t.scores += 3
		t.log[n] = 3
		return
	case r < chances.draw:
		t.scores++
		t.log[n] = 1
		team.scores++
		team.log[n] = 1
	default:
		team.scores += 3
		team.log[n] = 3
	}
	// fmt.Println(t, team)
}

func (t *Team) Reset() {
	t.scores = 0
	t.log[0] = 0
	t.log[1] = 0
	t.log[2] = 0
}

func NewTeam(id int) *Team {
	return &Team{
		id: id,
		scores: 0,
		log: [3]int{0, 0, 0},
	}
}

type Teams [4]*Team

func NewTeams() Teams {
	var teams Teams
	for i := 0; i < 4; i++ {
		teams[i] = NewTeam(i)
	}
	return teams
}

func (ts Teams) Reset() {
	for i := 0; i < 4; i++ {
		ts[i].scores = 0
		ts[i].log[0] = 0
		ts[i].log[1] = 0
		ts[i].log[2] = 0
	}
}

func (ts Teams) Rank() string {
	var s = make([]int, 4)
	for i := 0; i < 4; i++ {
		s[i] = ts[i].scores
	}
	sort.Ints(s)
	return fmt.Sprintf("%d-%d-%d-%d", s[3], s[2], s[1], s[0])
}

func Repeat(n int) {
	var teams = NewTeams()
	var res = make(map[string]int)
	cd := NewCD(0.4, 0.6)
	var key string
	for i := 0; i < n; i++ {
		teams[0].MatchWith(teams[1], 0, cd)
		teams[2].MatchWith(teams[3], 0, cd)
		teams[0].MatchWith(teams[2], 1, cd)
		teams[1].MatchWith(teams[3], 1, cd)
		teams[0].MatchWith(teams[3], 2, cd)
		teams[1].MatchWith(teams[2], 2, cd)
		
		key = teams.Rank()
		// fmt.Println(key)
		res[key]++
		teams.Reset()
	}
	fmt.Println(res)
	var count int
	for _, _ = range res {
		count++
	}
	fmt.Println(count)
}

func main() {
	Repeat(10000)
}
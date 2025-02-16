package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

type TeamName string

type Team struct {
	Name    TeamName
	Players []string
}

type League struct {
	Teams []Team
	Wins  map[TeamName]int
}

type Ranker interface {
	Ranking() []TeamName
}

func (l *League) MatchResult(team1 TeamName, team2 TeamName, team1Goals int, team2Goals int) {

	if team1Goals > team2Goals {
		l.Wins[team1]++
	} else if team2Goals > team1Goals {
		l.Wins[team2]++
	}
}

func (l *League) Ranking() []TeamName {

	sort.Slice(l.Teams, func(i, j int) bool {
		return l.Wins[l.Teams[i].Name] > l.Wins[l.Teams[j].Name]
	})

	names := make([]TeamName, len(l.Teams))
	for i, team := range l.Teams {
		names[i] = team.Name
	}
	return names
}

func RankPrinter(ranker Ranker, writer io.Writer) {
	names := ranker.Ranking()
	for i, name := range names {
		io.WriteString(writer, fmt.Sprintf("%d. %s\n", i+1, name))
	}
}

func main() {
	league := League{
		Teams: []Team{
			{Name: "A"},
			{Name: "B"},
			{Name: "C"},
			{Name: "D"},
		},
		Wins: make(map[TeamName]int),
	}

	league.MatchResult(TeamName("A"), TeamName("B"), 1, 2)
	league.MatchResult(TeamName("C"), TeamName("D"), 1, 2)
	league.MatchResult(TeamName("A"), TeamName("C"), 1, 2)
	league.MatchResult(TeamName("C"), TeamName("B"), 1, 2)
	league.MatchResult(TeamName("A"), TeamName("D"), 1, 2)
	league.MatchResult(TeamName("C"), TeamName("A"), 1, 2)

	file, err := os.OpenFile("rank.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	RankPrinter(&league, file)
	fmt.Println(league.Wins)
}

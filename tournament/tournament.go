package tournament

import (
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strings"
)

const formatString string = "%-31v| %2v | %2v | %2v | %2v | %2v\n"

type team struct {
	name              string
	wins, draws, loss int
}

func (team *team) Points() int {
	return (3 * team.wins) + (1 * team.draws)
}

func (team *team) TableRow() string {
	return fmt.Sprintf(formatString, team.name, (team.wins + team.draws + team.loss), team.wins, team.draws, team.loss, team.Points())
}

type tournament struct {
	teams   []team
	indexOf map[string]int
}

func (tour *tournament) addMatch(teamName1, teamName2, result string) {
	team1 := &(tour.teams[tour.indexOf[teamName1]])
	team2 := &(tour.teams[tour.indexOf[teamName2]])

	if result == "draw" {
		team1.draws++
		team2.draws++
	} else if result == "win" {
		team1.wins++
		team2.loss++
	} else {
		team2.wins++
		team1.loss++
	}
}

func (tour *tournament) recreateIndex() {
	var teams = &tour.teams
	sort.Slice(*teams, func(i, j int) bool {
		if (*teams)[i].Points() != (*teams)[j].Points() {
			return (*teams)[i].Points() > (*teams)[j].Points()
		}
		return (*teams)[i].name < (*teams)[j].name
	})

	tour.indexOf = make(map[string]int)
	for index, team := range *teams {
		tour.indexOf[team.name] = index
	}
}

func (tour *tournament) addTeams(textEntries [][]string) {
	tour.indexOf = make(map[string]int)
	var currentIndex int

	for _, entry := range textEntries {
		if _, present := tour.indexOf[entry[0]]; !present {
			tour.indexOf[entry[0]] = currentIndex
			currentIndex++
			tour.teams = append(tour.teams, team{name: entry[0]})
		}

		if _, present := tour.indexOf[entry[1]]; !present {
			tour.indexOf[entry[1]] = currentIndex
			currentIndex++
			tour.teams = append(tour.teams, team{name: entry[1]})
		}
	}
}

func (tour *tournament) WriteOutput(out io.Writer) {
	tour.recreateIndex()
	fmt.Fprintf(out, formatString, "Team", "MP", "W", "D", "L", "P")
	for _, team := range tour.teams {
		fmt.Fprint(out, team.TableRow())
	}
}

func (tour *tournament) Process(textEntries [][]string) {
	tour.addTeams(textEntries)

	for _, entry := range textEntries {
		tour.addMatch(entry[0], entry[1], entry[2])
	}

	tour.recreateIndex()
}

// Tally takes a reader and writer and writes the tournament table generated
// from the reader input to the writer
func Tally(in io.Reader, out io.Writer) error {
	bytesIn, err := ioutil.ReadAll(in)
	if err != nil {
		return nil
	}

	// Input processing
	inputString := string(bytesIn)
	inputStringArray := strings.Split(inputString, "\n")

	var textEntries [][]string

	for i := 0; i < len(inputStringArray); i++ {
		inputStringArray[i] = strings.TrimSpace(inputStringArray[i])

		if len(inputStringArray[i]) == 0 {
			continue
		}

		temp := strings.Split(inputStringArray[i], ";")
		if len(temp) == 3 {
			if temp[2] == "win" || temp[2] == "loss" || temp[2] == "draw" {
				if temp[1] == temp[0] {
					return fmt.Errorf("Error : Both teams are same")
				}
				textEntries = append(textEntries, temp)
			} else {
				return fmt.Errorf("Error : Not win draw loss")
			}
		} else {
			if inputStringArray[i][0] == '#' {
				continue
			}
			return fmt.Errorf("Error : Not a comment")
		}
	}

	// Create table and write
	var tour tournament
	tour.Process(textEntries)
	tour.WriteOutput(out)

	return nil
}

package tournament

import (
	"encoding/csv"
	"fmt"
	"io"
	"sort"
)

type matchResult int

const (
	win  matchResult = 0
	draw matchResult = 1
	loss matchResult = 2
)

type teamResult struct {
	team   string
	wins   int
	draws  int
	losses int
}

func (r *teamResult) played() int {
	return r.wins + r.draws + r.losses
}

func (r *teamResult) points() int {
	return 3*r.wins + r.draws
}

type inputEntry struct {
	teams        [2]string
	matchResults [2]matchResult
}

func readEntries(reader io.Reader) ([]inputEntry, error) {
	csvReader := csv.NewReader(reader)
	csvReader.Comma = ';'
	csvReader.Comment = '#'
	csvReader.FieldsPerRecord = -1

	var entries []inputEntry
	for {
		input, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if len(input) == 3 {
			teams := [2]string{input[0], input[1]}
			var matchResults [2]matchResult
			switch input[2] {
			case "win":
				matchResults = [2]matchResult{win, loss}
			case "loss":
				matchResults = [2]matchResult{loss, win}
			case "draw":
				matchResults = [2]matchResult{draw, draw}
			default:
				return nil, fmt.Errorf("invalid result %q", input[2])
			}
			entries = append(entries, inputEntry{teams: teams, matchResults: matchResults})
		} else {
			return nil, fmt.Errorf("invalid input: %v", input)
		}
	}
	return entries, nil
}

func tallyEntries(entries []inputEntry) []teamResult {
	var teamResultsMap = make(map[string]teamResult)
	for _, entry := range entries {
		for i := 0; i < 2; i++ {
			team := entry.teams[i]
			matchResult := entry.matchResults[i]

			teamResult, exists := teamResultsMap[team]
			if !exists {
				teamResult.team = team
			}

			switch matchResult {
			case win:
				teamResult.wins++
			case draw:
				teamResult.draws++
			case loss:
				teamResult.losses++
			}
			teamResultsMap[team] = teamResult
		}
	}

	teamResults := make([]teamResult, 0, len(teamResultsMap))
	for _, teamResult := range teamResultsMap {
		teamResults = append(teamResults, teamResult)
	}
	sort.Slice(teamResults, func(i, j int) bool {
		r1, r2 := teamResults[i], teamResults[j]
		r1Points, r2Points := r1.points(), r2.points()
		switch {
		case r1Points != r2Points:
			return r1Points > r2Points
		case r1.wins != r2.wins:
			return r1.wins > r2.wins
		default:
			return r1.team < r2.team
		}
	})

	return teamResults
}

func writeResults(writer io.Writer, teamResults []teamResult) error {
	_, err := fmt.Fprintf(writer, "Team                           | MP |  W |  D |  L |  P\n")
	if err != nil {
		return err
	}
	for _, tr := range teamResults {
		_, err = fmt.Fprintf(writer, "%-30s | %2d | %2d | %2d | %2d | %2d\n",
			tr.team, tr.played(), tr.wins, tr.draws, tr.losses, tr.points())
		if err != nil {
			return err
		}
	}

	return nil
}

func Tally(reader io.Reader, writer io.Writer) error {
	entries, err := readEntries(reader)
	if err != nil {
		return err
	}

	err = writeResults(writer, tallyEntries(entries))
	return err
}

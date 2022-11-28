package cmd

import (
	"Go-cup/services"
	"encoding/json"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"log"
	"os"
)

type StandingResponse struct {
	Standings []Standing `json:"standings"`
}

type Standing struct {
	Stage string       `json:"stage"`
	Type  string       `json:"type"`
	Group string       `json:"group"`
	Table []TeamResult `json:"table"`
}

type TeamResult struct {
	Position       int         `json:"position"`
	Team           Team        `json:"team"`
	PlayedGames    int         `json:"playedGames"`
	Form           interface{} `json:"form"`
	Won            int         `json:"won"`
	Draw           int         `json:"draw"`
	Lost           int         `json:"lost"`
	Points         int         `json:"points"`
	GoalsFor       int         `json:"goalsFor"`
	GoalsAgainst   int         `json:"goalsAgainst"`
	GoalDifference int         `json:"goalDifference"`
}

type Team struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
	TlA       string `json:"tla"`
	Crest     string `json:"crest"`
}

const (
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

// standingCmd represents the standing command
var standingCmd = &cobra.Command{
	Use:   "standing",
	Short: "World Cup standing",
	Long:  `You can track World Cup standing details`,
	Run: func(cmd *cobra.Command, args []string) {
		listStandings()
	},
}

func init() {
	rootCmd.AddCommand(standingCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// standingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// standingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func listStandings() {
	result, err := services.Get("https://api.football-data.org/v4/competitions/WC/standings", nil)
	if err != nil {
		return
	}

	var standings StandingResponse
	//fmt.Println(result)
	err = json.Unmarshal([]byte(result.Response), &standings)
	if err != nil {
		log.Fatalln(err)
		return
	}

	for _, standing := range standings.Standings {
		fmt.Println(standing.Group)
		renderTable(standing.Table)
		fmt.Println()
	}
}

func renderTable(teams []TeamResult) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	headers := table.Row{"Position", "Name", "Played games", "Won", "Draw", "Lost", "Points", "Goals for", "Goals against", "Goal difference"}
	t.AppendHeader(headers)
	for _, team := range teams {
		t.AppendRow(table.Row{team.Position, team.Team.Name, team.PlayedGames, team.Won, team.Draw, team.Lost, team.Points, team.GoalsFor, team.GoalsAgainst, team.GoalDifference})
	}
	t.Render()
}

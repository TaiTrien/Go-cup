/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"Go-cup/services"
	"Go-cup/util"
	"encoding/json"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"log"
	"os"
	"reflect"
	"time"
)

type TeamResponse struct {
	Teams []Team `json:"teams"`
}

type MatchResponse struct {
	Matches   []Match   `json:"matches"`
	ResultSet ResultSet `json:"resultSet"`
}

type ResultSet struct {
	Count  int `json:"count"`
	Played int `json:"played"`
	Wins   int `json:"wins"`
	Draws  int `json:"draws"`
	Losses int `json:"losses"`
}

type Match struct {
	Id          int         `json:"id"`
	UtcDate     time.Time   `json:"utcDate"`
	Status      string      `json:"status"`
	Stage       string      `json:"stage"`
	Group       string      `json:"group"`
	LastUpdated time.Time   `json:"lastUpdated"`
	HomeTeam    Team        `json:"homeTeam"`
	AwayTeam    Team        `json:"awayTeam"`
	Result      MatchResult `json:"score"`
}

type MatchResult struct {
	Winner   string `json:"winner"`
	Duration string `json:"duration"`
	FullTime Score  `json:"fullTime"`
	HalfTime Score  `json:"halfTime"`
}

type Score struct {
	Home int `json:"home"`
	Away int `json:"away"`
}

var teamId string

// teamCmd represents the team command
var teamCmd = &cobra.Command{
	Use:   "team",
	Short: "Get World Cup team details",
	Long:  `You can get World Cup team details`,
	Run: func(cmd *cobra.Command, args []string) {
		if teamId != "" {
			teamInfo := fetchTeamInfo(teamId)
			fetchTeamMatches(teamId, teamInfo.Name)
		} else {
			fetchAllTeams()
		}
	},
}

func init() {
	rootCmd.AddCommand(teamCmd)
	teamCmd.Flags().StringVarP(&teamId, "id", "t", "", "use team id get from command 'team'")
}

func fetchAllTeams() {
	result, err := services.Get("https://api.football-data.org/v4/competitions/WC/teams", nil)
	if err != nil {
		return
	}

	var teamRes TeamResponse
	err = json.Unmarshal([]byte(result.Response), &teamRes)
	if err != nil {
		log.Fatalln(err)
		return
	}
	renderTeamTable(teamRes.Teams)
}

func fetchTeamInfo(id string) Team {
	result, err := services.Get(fmt.Sprintf("https://api.football-data.org/v4/teams/%v", id), nil)
	if err != nil {
		log.Fatalln(err)
		return Team{}
	}
	var team Team
	err = json.Unmarshal([]byte(result.Response), &team)
	return team
}

func fetchTeamMatches(id string, teamName string) {
	result, err := services.Get(fmt.Sprintf("https://api.football-data.org/v4/teams/%v/matches", id), nil)
	if err != nil {
		log.Fatalln(err)
		return
	}
	var matchRes MatchResponse
	err = json.Unmarshal([]byte(result.Response), &matchRes)
	fmt.Print("Team: " + teamName + "\t")
	summaryFields := util.ExtractStructFieldNames(ResultSet{})
	for _, field := range summaryFields {
		r := reflect.ValueOf(matchRes.ResultSet)
		fmt.Printf("%s: %v \t", field, reflect.Indirect(r).FieldByName(field))
	}
	fmt.Println()
	renderTeamMatchesTable(matchRes.Matches)
}

func renderTeamTable(teams []Team) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	headers := table.Row{"ID", "Name", "Code"}
	t.AppendHeader(headers)
	for _, team := range teams {
		t.AppendRow(table.Row{team.ID, team.Name, team.TlA})
	}
	t.Render()
}

func renderTeamMatchesTable(matches []Match) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	headers := table.Row{"ID", "STAGE", "GROUP", "HOME", "SCORE", "AWAY", "STATUS", "VN-TIME", "UTC-TIME"}
	t.AppendHeader(headers)
	for _, match := range matches {
		vnTime := util.DateFormat(match.UtcDate.Add(time.Hour * 7))
		utcTime := util.DateFormat(match.UtcDate)
		score := fmt.Sprintf("%v - %v", match.Result.FullTime.Home, match.Result.FullTime.Away)
		t.AppendRow(table.Row{match.Id, match.Stage, match.Group, match.HomeTeam.Name, score, match.AwayTeam.Name, match.Status, vnTime, utcTime})
	}
	t.Render()
}

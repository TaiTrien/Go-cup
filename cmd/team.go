/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"Go-cup/services"
	"encoding/json"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"log"
	"os"
)

type TeamResponse struct {
	Teams []Team `json:"teams"`
}

// teamCmd represents the team command
var teamCmd = &cobra.Command{
	Use:   "team",
	Short: "World cup team",
	Long:  `You can get World Cup team details`,
	Run: func(cmd *cobra.Command, args []string) {
		listAllTeam()
	},
}

func init() {
	rootCmd.AddCommand(teamCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// teamCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// teamCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func listAllTeam() {
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

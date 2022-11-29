/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"Go-cup/services"
	"encoding/json"
	"github.com/spf13/cobra"
)

// matchCmd represents the match command
var matchCmd = &cobra.Command{
	Use:   "match",
	Short: "Get World Cup matches details",
	Long:  `You can get World Cup matches info`,
	Run: func(cmd *cobra.Command, args []string) {
		getNextMatches()
	},
}

func getNextMatches() {
	result, err := services.Get("https://api.football-data.org/v4/competitions/WC/matches?status=SCHEDULED", nil)
	if err != nil {
		return
	}

	var matchRes MatchResponse
	err = json.Unmarshal([]byte(result.Response), &matchRes)

	renderTeamMatchesTable(matchRes.Matches)
}

func init() {
	rootCmd.AddCommand(matchCmd)
}

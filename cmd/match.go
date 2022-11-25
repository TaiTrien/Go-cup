/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"Go-cup/util"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net/http"
)

// matchCmd represents the match command
var matchCmd = &cobra.Command{
	Use:   "match",
	Short: "Use for related matches",
	Long:  `Use to find matches purposes. For example: match --next to find next match`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := util.LoadConfig(".")
		if err != nil {
			log.Fatalln(err)
		}

		client := &http.Client{}
		req, _ := http.NewRequest(http.MethodGet, "http://api.football-data.org/v4/teams/759/matches", nil)
		req.Header.Add("X-Auth-Token", config.ApiKey)

		resp, httpErr := client.Do(req)
		if httpErr != nil {
			log.Fatalln(err)
		}

		defer func(Body io.ReadCloser) {
			err = Body.Close()
			if err != nil {
				log.Fatalln(err)
			}
		}(resp.Body)

		responseBody, err1 := io.ReadAll(resp.Body)
		if err1 != nil {
			log.Fatal(err1)
		}

		fmt.Println(resp.Status)
		fmt.Println(string(responseBody))
	},
}

func init() {
	rootCmd.AddCommand(matchCmd)
	matchCmd.Flags().String("next", "nx", "To find next match")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// matchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// matchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

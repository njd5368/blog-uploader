/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

type Post struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Type        string `json:"type"`
	Image       []byte  `json:"image"`

	Languages    []string `json:"languages"`
	Technologies []string `json:"technologies"`

	File []byte `json:"file"

}

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		dir, err := os.Getwd()
		if err != nil {
			fmt.Print(err)
			return
		}

		configBytes, err := ioutil.ReadFile(dir + "/post.json")
		if err != nil {
			fmt.Print(err)
			return
		}
		
		var post Post
		json.Unmarshal(configBytes, &post)

		markdown, err := ioutil.ReadFile(dir + "/post.md")
		if err != nil {
			fmt.Print(err)
			return
		}

		post.

		hero, err := ioutil.ReadFile(dir + "/hero.jpeg")
		if err != nil {
			fmt.Print(err)
			return
		}



		apiCall := http.NewRequest("POST", CmdConfig.APIURL + "/api/project", )
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uploadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uploadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("Enter Name: ")
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			fmt.Print(err)
			return
		}
		postName := strings.TrimSpace(scanner.Text())

		fmt.Print("Enter Description: ")
		scanner.Scan()
		err = scanner.Err()
		if err != nil {
			fmt.Print(err)
			return
		}
		postDescription := strings.TrimSpace(scanner.Text())

		postDate := time.Now().Format("2006-01-02") //YYYY-MM-DD

		fmt.Print("Enter Type: ")
		scanner.Scan()
		err = scanner.Err()
		if err != nil {
			fmt.Print(err)
			return
		}
		postType := strings.TrimSpace(scanner.Text())

		fmt.Print("Enter Languages (Space Separated): ")
		scanner.Scan()
		err = scanner.Err()
		if err != nil {
			fmt.Print(err)
			return
		}
		postLanguagesInput := scanner.Text()

		var postLanguages []string 
		for _, l := range(strings.Split(postLanguagesInput, " ")) {
			l = strings.TrimSpace(l)
			if len(l) != 0 {
				postLanguages = append(postLanguages, l)
			}
		}
		
		fmt.Print("Enter Technologies (Space Separated): ")
		scanner.Scan()
		err = scanner.Err()
		if err != nil {
			fmt.Print(err)
			return
		}
		postTechnologiesInput := scanner.Text()

		var postTechnologies []string 
		for _, l := range(strings.Split(postTechnologiesInput, " ")) {
			l = strings.TrimSpace(l)
			if len(l) != 0 {
				postTechnologies = append(postTechnologies, l)
			}
		}

		postJSON := struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Date        string `json:"date"`
			Type        string `json:"type"`
		
			Languages    []string `json:"languages"`
			Technologies []string `json:"technologies"`
		}{
			Name: postName,
			Description: postDescription,
			Date: postDate,
			Type: postType,
			Languages: postLanguages,
			Technologies: postTechnologies,
		}

		dir, err := os.Getwd()
		if err != nil {
			fmt.Print(err)
			return
		}
		dir += "/" + strings.ReplaceAll(postName, " ", "-")
		err = os.Mkdir(dir, os.FileMode(0755))
		if err != nil {
			fmt.Print(err)
			return
		}

		postJSONBytes, err := json.MarshalIndent(postJSON, "", "    ")
		if err != nil {
			fmt.Print(err)
			return
		}

		err = os.WriteFile(dir + "/post.json", postJSONBytes, os.FileMode(0644))
		if err != nil {
			fmt.Print(err)
			return
		}

		postMD := []byte("# " + postName)
		err = os.WriteFile(dir + "/post.md", postMD, os.FileMode(0644))
		if err != nil {
			fmt.Print(err)
			return
		}

		fmt.Println("Make sure to add a hero.jpeg to your post. Happy posting!")
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

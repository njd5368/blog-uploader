/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"bytes"
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

	File []byte `json:"file"`

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
		usernamePassword, err := retrieveUsernamePassword()
		if err != nil {
			fmt.Print(err)
			return
		}

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

		markdownFile, err := os.Open(dir + "/post.md")
		if err != nil {
			fmt.Print(err)
			return
		}

		var markdown []byte

		scanner := bufio.NewScanner(markdownFile)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			line := scanner.Bytes()
			if bytes.HasPrefix(line, []byte("![")) {
				s := bytes.IndexByte(line, byte('('))
				e := bytes.IndexByte(line, byte(')'))
				if s != -1 && e != -1 && s < e{
					s += 1
					e -= 1
					imageName := line[s:e]
					image, err := ioutil.ReadFile(dir + "/" + string(imageName))
					if err != nil {
						fmt.Print(err)
						return
					}

					request, err := http.NewRequest("Post", CmdConfig.APIURL + "/api/image", bytes.NewBuffer(image))
					if err != nil {
						fmt.Print(err)
						return
					}
					request.Header.Set("Content-type", "application/octet-stream")
					request.Header.Set("Authorization", "Basic " + usernamePassword)
					response, err := http.DefaultClient.Do(request)
					if err != nil {
						fmt.Print(err)
						return
					}

					location := response.Header.Get("Location")
					if len(location) == 0 {
						fmt.Print("Error uploading an image file.")
						return
					}

					line = bytes.Replace(line, imageName, []byte(location), 1)
				}
			}

			markdown = append(markdown, line...)
			markdown = append(markdown, byte('\n'))
		}

		post.File = markdown

		hero, err := ioutil.ReadFile(dir + "/hero.jpeg")
		if err != nil {
			fmt.Print(err)
			return
		}

		post.Image = hero

		reqBody, err := json.Marshal(post)
		if err != nil {
			fmt.Print(err)
			return
		}

		apiCall, err := http.NewRequest("Post", CmdConfig.APIURL + "/api/project", bytes.NewBuffer(reqBody))
		if err != nil {
			fmt.Print(err)
			return
		}

		apiCall.Header.Set("Content-type", "application/json")
		apiCall.Header.Set("Authorization", "Basic " + usernamePassword)

		response, err := http.DefaultClient.Do(apiCall)
		if err != nil {
			fmt.Print(err)
			return
		}
		if response.StatusCode != http.StatusAccepted {
			fmt.Print("Error sending post request.")
			return
		}
		fmt.Print("Submitted blog post.")
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

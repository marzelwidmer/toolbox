/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package net

import (
	"fmt"
	"net/http"
	"time"
	"github.com/spf13/cobra"
)

var (
	urlPath string
	client = http.Client{
		Timeout: time.Second * 10,
	}
)

func ping(domain string) (int, error) {
	url := "http://" + domain
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return 0, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	resp.Body.Close()
	return resp.StatusCode, nil
}


var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "This ping a remote url.",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		if resp, err := ping(urlPath); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(resp)
		}
	},
}

func init() {
	pingCmd.Flags().StringVarP(&urlPath, "url", "u", "", "The url to ping")
	if err := pingCmd.MarkFlagRequired("url"); err != nil {
		fmt.Println(err)
	}
	NetCmd.AddCommand(pingCmd)
}

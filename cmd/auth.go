/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"os"
	"path/filepath"
	"you-up/controller"

	"github.com/spf13/cobra"
)

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate to your google service",
	Long:  `Authenticate to your google service. To use the application you need to download "client_secrets.json" from the YouTube Data API (v3).`,
	Run: func(cmd *cobra.Command, args []string) {
		filename, _ := cmd.Flags().GetString("filename")
		usrHomeDir := controller.GetHomeDir()
		clientSecretPath := filepath.Join(usrHomeDir, ".credentials", "client_secret.json")
		if _, err := os.Stat(clientSecretPath); err == nil {
			controller.GoogleLogin()
		} else if errors.Is(err, os.ErrNotExist) {
			controller.GoogleFirstLogin(filename)
		}
	},
}

func init() {
	rootCmd.AddCommand(authCmd)

	authCmd.Flags().StringP("filename", "f", "", "Client secret file")
	authCmd.MarkFlagRequired("filename")
}

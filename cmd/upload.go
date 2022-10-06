/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload video to youtube",
	Long:  `This command enable the upload of videos to youtube via your terminal.`,
	Run: func(cmd *cobra.Command, args []string) {
		filename, _ := cmd.Flags().GetString("filename")
		title, _ := cmd.Flags().GetString("title")
		description, _ := cmd.Flags().GetString("description")
		category, _ := cmd.Flags().GetString("category")
		keywords, _ := cmd.Flags().GetString("keywords")
		privacy, _ := cmd.Flags().GetString("privacy")
		fmt.Println(filename, title, description, category, keywords, privacy)
		// controller.UploadVideo()
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)

	//Available flags
	uploadCmd.Flags().StringP("filename", "f", "", "Name of video file to upload (required)")
	uploadCmd.MarkFlagRequired("filename")
	uploadCmd.Flags().StringP("title", "t", "", "Name of video file to upload (required)")
	uploadCmd.MarkFlagRequired("title")
	uploadCmd.Flags().String("description", "", "Video description")
	uploadCmd.Flags().String("category", "", "Video category")
	uploadCmd.Flags().String("keywords", "", "Comma separated list of video keywords")
	uploadCmd.Flags().String("privacy", "unlisted", "Video privacy status")

}

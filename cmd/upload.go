package cmd

import (
	"log"
	"you-up/controller"

	"github.com/spf13/cobra"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload video to youtube",
	Long:  `This command enable the upload of videos to youtube via your terminal.`,
	Run: func(cmd *cobra.Command, args []string) {
		filename, err := cmd.Flags().GetString("filename")
		controller.HandleError(err, "Cannot read the filename")
		title, err := cmd.Flags().GetString("title")
		controller.HandleError(err, "Cannot read the filename")
		description, err := cmd.Flags().GetString("description")
		controller.HandleError(err, "Cannot read the description")
		category, err := cmd.Flags().GetString("category")
		controller.HandleError(err, "Cannot read the category")
		keywords, err := cmd.Flags().GetString("keywords")
		controller.HandleError(err, "Cannot read the keywords")
		privacy, err := cmd.Flags().GetString("privacy")
		controller.HandleError(err, "Cannot read the privacy")
		multiple, err := cmd.Flags().GetBool("multiple")
		controller.HandleError(err, "Cannot read the multiple")

		if multiple {
			controller.BulkUpload(filename)
		} else {
			if title != "" {
				video := controller.NewVideo(filename, title, description, category, keywords, privacy)
				controller.UploadVideo(video)
			} else {
				log.Fatalf("Title is not provided")
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)

	//Available flags
	uploadCmd.Flags().StringP("filename", "f", "", "Name of video file to upload (required)")
	uploadCmd.MarkFlagRequired("filename")
	uploadCmd.Flags().StringP("title", "t", "", "Name of video file to upload (required)")
	uploadCmd.Flags().String("description", "", "Video description")
	uploadCmd.Flags().String("category", "", "Video category")
	uploadCmd.Flags().String("keywords", "", "Comma separated list of video keywords")
	uploadCmd.Flags().String("privacy", "unlisted", "Video privacy status")
	uploadCmd.Flags().BoolP("multiple", "m", false, "Upload several video using a JSON file")

}

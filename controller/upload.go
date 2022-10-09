package controller

import (
	"fmt"
	"log"
	"os"
	"strings"

	"google.golang.org/api/youtube/v3"
)

type Video struct {
	Filename    string
	Title       string
	Description string
	Category    string
	Keywords    string
	Privacy     string
}

func upload(service *youtube.Service, video *Video) {

	upload := &youtube.Video{
		Snippet: &youtube.VideoSnippet{
			Title:       video.Title,
			Description: video.Description,
			CategoryId:  video.Category,
		},
		Status: &youtube.VideoStatus{PrivacyStatus: video.Privacy},
	}

	// The API returns a 400 Bad Request response if tags is an empty string.
	if strings.Trim(video.Keywords, "") != "" {
		upload.Snippet.Tags = strings.Split(video.Keywords, ",")
	}

	parts := []string{"snippet", "status"}
	call := service.Videos.Insert(parts, upload)

	file, err := os.Open(video.Filename)
	defer file.Close()
	if err != nil {
		log.Fatalf("Error opening %v: %v", video.Filename, err)
	}

	response, err := call.Media(file).Do()
	handleError(err, "")
	url := "https://www.youtube.com/watch?v=" + response.Id
	fmt.Printf("Upload successful! Video ID: %v\n", url)
}

func UploadVideo(video *Video) {
	client := authenticate()
	service, err := youtube.New(client)
	handleError(err, "Error creating YouTube client")
	upload(service, video)
	// fmt.Println(video, video.Filename)
}

func NewVideo(filename string, title string, description string, category string, keywords string, privacy string) *Video {
	v := Video{Filename: filename, Title: title, Description: description, Category: category, Keywords: keywords, Privacy: privacy}
	return &v
}

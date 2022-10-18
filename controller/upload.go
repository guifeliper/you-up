package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"google.golang.org/api/youtube/v3"
)

type Video struct {
	Filename    string `json: "filename"`
	Title       string `json: "title"`
	Description string `json: "description"`
	Category    string `json: "category"`
	Keywords    string `json: "keywords"`
	Privacy     string `json: "privacy"`
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
}

func NewVideo(filename string, title string, description string, category string, keywords string, privacy string) *Video {
	v := Video{Filename: filename, Title: title, Description: description, Category: category, Keywords: keywords, Privacy: privacy}
	return &v
}

func BulkUpload(filename string) {
	file, _ := ioutil.ReadFile(filename)

	data := []Video{}

	_ = json.Unmarshal([]byte(file), &data)

	for i := 0; i < len(data); i++ {
		fmt.Println("Uploading " + data[i].Title + "...")
		video := NewVideo(data[i].Filename, data[i].Title, data[i].Description, data[i].Category, data[i].Keywords, data[i].Privacy)
		UploadVideo(video)
	}
}

// Sample Go code for user authorization

package controller

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

const missingClientSecretsMessage = `
Please configure OAuth 2.0
`

var (
	filename    = flag.String("filename", "", "Name of video file to upload")
	title       = flag.String("title", "Test Title", "Video title")
	description = flag.String("description", "Test Description", "Video description")
	category    = flag.String("category", "22", "Video category")
	keywords    = flag.String("keywords", "", "Comma separated list of video keywords")
	privacy     = flag.String("privacy", "unlisted", "Video privacy status")
)

// getClient uses a Context and Config to retrieve a Token
// then generate a Client. It returns the generated Client.
func getClient(ctx context.Context, config *oauth2.Config) *http.Client {
	cacheFile, err := tokenCacheFile()
	if err != nil {
		log.Fatalf("Unable to get path to cached credential file. %v", err)
	}
	tok, err := tokenFromFile(cacheFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(cacheFile, tok)
	}
	return config.Client(ctx, tok)
}

// getTokenFromWeb uses Config to request a Token.
// It returns the retrieved Token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)
	exec.Command("open", authURL).Start()
	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}

// tokenCacheFile generates credential file path/filename.
// It returns the generated credential path/filename.
func tokenCacheFile() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	tokenCacheDir := filepath.Join(usr.HomeDir, ".credentials")
	os.MkdirAll(tokenCacheDir, 0700)
	return filepath.Join(tokenCacheDir,
		url.QueryEscape("youtube-go-quickstart.json")), err
}

// tokenFromFile retrieves a Token from a given file path.
// It returns the retrieved Token and any read error encountered.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	defer f.Close()
	return t, err
}

// saveToken uses a file path to create a file and store the
// token in it.
func saveToken(file string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", file)
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func handleError(err error, message string) {
	if message == "" {
		message = "Error making API call"
	}
	if err != nil {
		log.Fatalf(message+": %v", err.Error())
	}
}

func authenticate() *http.Client {
	ctx := context.Background()

	b, err := ioutil.ReadFile("client_secret.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, youtube.YoutubeUploadScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(ctx, config)

	return client
}

func channelsListByUsername(service *youtube.Service, part string, forUsername string) {
	parts := []string{part}
	call := service.Channels.List(parts)
	call = call.ForUsername(forUsername)
	response, err := call.Do()
	handleError(err, "")
	fmt.Println(fmt.Sprintf("This channel's ID is %s. Its title is '%s', "+
		"and it has %d views.",
		response.Items[0].Id,
		response.Items[0].Snippet.Title,
		response.Items[0].Statistics.ViewCount))
}

func upload(service *youtube.Service) {
	flag.Parse()

	if *filename == "" {
		log.Fatalf("You must provide a filename of a video file to upload")
	}

	upload := &youtube.Video{
		Snippet: &youtube.VideoSnippet{
			Title:       *title,
			Description: *description,
			CategoryId:  *category,
		},
		Status: &youtube.VideoStatus{PrivacyStatus: *privacy},
	}

	// The API returns a 400 Bad Request response if tags is an empty string.
	if strings.Trim(*keywords, "") != "" {
		upload.Snippet.Tags = strings.Split(*keywords, ",")
	}

	var parts []string
	parts = append(parts, "snippet")
	parts = append(parts, "status")
	call := service.Videos.Insert(parts, upload)

	file, err := os.Open(*filename)
	defer file.Close()
	if err != nil {
		log.Fatalf("Error opening %v: %v", *filename, err)
	}

	response, err := call.Media(file).Do()
	handleError(err, "")
	url := "https://www.youtube.com/watch?v=" + response.Id
	fmt.Printf("Upload successful! Video ID: %v\n", url)
}

func GoogleLogin() {
	authenticate()
	fmt.Println("Google user authenticate with success!")
}

func GetChannelByUsername() {
	client := authenticate()
	service, err := youtube.New(client)
	handleError(err, "Error creating YouTube client")
	channelsListByUsername(service, "snippet,contentDetails,statistics", "GoogleDevelopers")
}

func UploadVideo() {
	client := authenticate()
	service, err := youtube.New(client)
	handleError(err, "Error creating YouTube client")
	upload(service)
}

// Sample Go code for user authorization

package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

// getClient uses a Context and Config to retrieve a Token
// then generate a Client. It returns the generated Client.
func getClient(ctx context.Context, config *oauth2.Config) *http.Client {
	cacheFile := tokenCacheFile()
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
func tokenCacheFile() string {
	usrHomeDir := GetHomeDir()
	tokenCacheDir := filepath.Join(usrHomeDir, ".credentials")
	os.MkdirAll(tokenCacheDir, 0700)
	return filepath.Join(tokenCacheDir,
		url.QueryEscape("google.json"))
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

func saveClientSecret(sourceFile string) {
	// Set all variables
	usrHomeDir := GetHomeDir()
	tokenCacheDir := filepath.Join(usrHomeDir, ".credentials")
	destinationFile := tokenCacheDir + "/client_secret.json"

	// Make Directory
	os.MkdirAll(tokenCacheDir, 0700)

	//Read the file
	input, err := ioutil.ReadFile(sourceFile)
	HandleError(err, "Cannot read the file")

	// Save the file
	err = ioutil.WriteFile(destinationFile, input, 0644)
	HandleError(err, "Error creating "+destinationFile)

}

func authenticate() *http.Client {
	ctx := context.Background()
	usrHomeDir := GetHomeDir()
	clientSecretFile := filepath.Join(usrHomeDir, ".credentials", "client_secret.json")
	b, err := ioutil.ReadFile(clientSecretFile)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v, \n\nPlease login with: you-up auth -f <<Google Client secret file>>", err)
	}

	config, err := google.ConfigFromJSON(b, youtube.YoutubeUploadScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(ctx, config)

	return client
}

func GoogleLogin() {
	authenticate()
	fmt.Println("Google user authenticate with success!")
}

func GoogleFirstLogin(sourceFile string) {
	saveClientSecret(sourceFile)
	GoogleLogin()
}

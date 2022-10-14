# youtube-up
A cli tool for fast uploading videos to youtube.

## How to install

```
brew update
brew install guifeliper/you-up/you-up
```

## How to use
### Authentication 
First thing you should authenticate with your Youtube Data API V3 from google. For that follow the steps: 

**Step 1**: 
1. Turn on the YouTube Data API
Use this [wizard](https://console.developers.google.com/start/api?id=youtube). to create or select a project in the Google Developers Console and automatically turn on the API. Click Continue, then Go to credentials.

2. On the Create credentials page, click the Cancel button.

3. At the top of the page, select the OAuth consent screen tab. Select an Email address, enter a Product name if not already set, and click the Save button.

4. Select the Credentials tab, click the Create credentials button and select OAuth client ID.

5. Select the application type Other, enter the name "YouTube Data API Quickstart", and click the Create button.

6. Click OK to dismiss the resulting dialog.

7. Click the file_download (Download JSON) button to the right of the client ID.

8. Move the downloaded file to your working directory and rename it client_secret.json.

**Step 2**: 
Run the following command and the Browser will show up asking you to connect with you Google account, accept the oAuth2 connection. 
```
you-up auth -f "path/client_secret.json"
```
and you are done to upload videos.

<br>

### Upload videos 

```
you-up upload -f "./assets/sample_video.mp4" -t "testing auth"
```

## Dev Environment

```
docker build --target dev . -t go
docker run -it -v ${PWD}:/work go sh
go version
```

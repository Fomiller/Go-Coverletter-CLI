package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"google.golang.org/api/drive/v3"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/docs/v1"
)

var (
	DOC docs.Document
)

// Retrieves a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Requests a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(oauth2.NoContext, authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	defer f.Close()
	if err != nil {
		log.Fatalf("Unable to cache OAuth token: %v", err)
	}
	json.NewEncoder(f).Encode(token)
}

// reads template.json file and returns a BatchUpdateDocumentRequest for updating documents.
func templateFromFile(file string) (*docs.BatchUpdateDocumentRequest, error) {
	// init template stuct
	template := &docs.BatchUpdateDocumentRequest{}
	// read file
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	// decode JSON into BatchUpdateDocumentRequest struct
	err = json.NewDecoder(f).Decode(template)
	return template, err
}

func main() {
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/docs")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := docs.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Docs client: %v", err)
	}

	// <----------------------------------->
	// <----------------------------------->
	// <----------------------------------->
	// My Code
	// var newDoc docs.Document
	newDoc := DOC
	newDoc.Title = "GOOD TO GO"
	fmt.Println(newDoc.Title)
	res, err := srv.Documents.Create(&newDoc).Do()
	if err != nil {
		log.Panicf("THIS IS THE ERROR: %v", err)
	}

	fmt.Printf("\n RESPONSE DocumentId: %v", res.DocumentId)

	tpl, err := templateFromFile("template.json")
	fmt.Println("\n TEMPLATE: \n", tpl)
	batchRes, err := srv.Documents.BatchUpdate("1OXbxsaMG8TkPTZCuj_cKu0A2B_Ft4jfUjtA2FcBNUW0", tpl).Do()
	if err != nil {
		log.Fatalf("BATCH FAIL %v ", err)
	}
	fmt.Printf("SUCCESSFUL BATCH UPDATE: %v \n", batchRes.DocumentId)

	// COPYING FILE FROM DRIVE
	newFile := drive.File{}
	newFile.Name = "NEW TEMPLATE"
	ctx := context.Background()
	driveSrv, err := drive.NewService(ctx)
	if err != nil {
		log.Fatal(err)
	}

	driveRes, err := driveSrv.Files.Copy("1yMx9J4z6cJCVpzp9zXjknkVX6xrtMFufsp_iNv9aZ40", &newFile).Do()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DRIVE RES: ", driveRes)
}
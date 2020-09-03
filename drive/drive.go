package drive

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
)

var driveSrv *drive.Service

func init() {
	b, err := ioutil.ReadFile("./drive/credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, drive.DriveScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	driveSrv, err = drive.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "./drive/token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

// returns a list of files and file Ids that contain the query string
func SearchForFiles(q string) {
	// create query
	// more info about creating queries can be found https://developers.google.com/drive/api/v3/search-files
	query := fmt.Sprintf("name contains '%v'", q)
	// query drive
	fl, err := driveSrv.Files.List().Q(query).Do()
	if err != nil {
		log.Panic("fl: ", err)
	}

	// print out files found
	for i, v := range fl.Files {
		fmt.Printf("%v: %v\n", i, v.Name)
		fmt.Printf("%v: %v\n", i, v.Id)
	}
}

// return template Id from specified templateName
func GetFileId(File string) string {
	// create query where file matches query exactly
	// more info about creating queries can be found https://developers.google.com/drive/api/v3/search-files
	query := fmt.Sprintf("name='%v'", File)
	// search drive for the matching query
	driveRes, err := driveSrv.Files.List().Q(query).Do()
	if err != nil {
		log.Panic("fl: ", err)
	}

	// return the file Id
	return driveRes.Files[0].Id
}

// create a new file from Template, takes in a fileName and a docId in the form of templateId
func NewTemplate(newFileName string, templateId string) string {
	// create a new file struct
	newFile := drive.File{}
	// set newFile name
	newFile.Name = newFileName

	// copy the template using the information stored in newFile
	driveRes, err := driveSrv.Files.Copy(templateId, &newFile).Do()
	if err != nil {
		log.Fatal(err)
	}
	// return the newFile Id
	return driveRes.Id
}

// download a file to output folder
// TODO*** make the output directory and dynamic value; specified by a flag??
func DownloadFile(fileId string, fileName string) {
	// set out put folder name
	path := "output"
	// append file type to file name
	fileName = fmt.Sprintf("%v.pdf", fileName)
	// create call to export file from drive
	fileCall := driveSrv.Files.Export(fileId, "application/pdf")
	// execute download of file call
	res, err := fileCall.Download()
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// read the res.Body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// check if output folder if exists, if not then create the folder
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0644)
	}

	// write the file to specified folder
	err = ioutil.WriteFile(fmt.Sprintf("./%v/%v", path, fileName), body, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

// delete file from drive
func DeleteFile(docId string) error {
	// delete file from drive
	err := driveSrv.Files.Delete(docId).Do()
	// return err
	return err
}

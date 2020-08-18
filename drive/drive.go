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

const (
	TEMPLATE = "1yMx9J4z6cJCVpzp9zXjknkVX6xrtMFufsp_iNv9aZ40"
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

// OLD FUNCTION NEEDS TO BE REMOVED BROKEN APRART
// func CreateTemplateCopy() string {
// 	r, err := driveSrv.Files.List().PageSize(10).
// 		Fields("nextPageToken, files(id, name)").Do()
// 	if err != nil {
// 		log.Fatalf("Unable to retrieve files: %v", err)
// 	}
// 	fmt.Println("Files:")
// 	if len(r.Files) == 0 {
// 		fmt.Println("No files found.")
// 	} else {
// 		for _, i := range r.Files {
// 			fmt.Printf("%s (%s)\n", i.Name, i.Id)
// 		}
// 	}

// 	fmt.Println("\n<---------------------------->\n")
// 	fmt.Println("List Call")

// 	fl, err := driveSrv.Files.List().Q("name='Cover Letter Template'").Do()
// 	if err != nil {
// 		log.Panic("fl: ", err)
// 	}
// 	fmt.Println("This is from Q:", fl.Files[0].Name)

// 	fmt.Println("\n<---------------------------->\n")
// 	// generate list of ids
// 	res, err := driveSrv.Files.GenerateIds().Do()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// loop through ids
// 	if len(res.Ids) == 0 {
// 		fmt.Println("No Files found.")
// 	} else {
// 		for i, v := range res.Ids {
// 			fmt.Printf("ID-%v: %v\n", i, v)
// 		}
// 	}

// 	fmt.Println("\n<---------------------------->\n")

// 	// **** make this a dynamic value ***
// 	copyTitle := "UPDATED TEMPLATE"
// 	newFile := drive.File{}
// 	newFile.Name = copyTitle

// 	driveRes, err := driveSrv.Files.Copy(TEMPLATE, &newFile).Do()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// print document id if successful.
// 	fmt.Println("FILE/DOCUMENT-ID: ", driveRes.Id)
// 	return driveRes.Id
// }

func SearchForFiles(q string) {
	query := fmt.Sprintf("name contains '%v'", q)

	fl, err := driveSrv.Files.List().Q(query).Do()
	if err != nil {
		log.Panic("fl: ", err)
	}
	fmt.Println(fl.Files[0])
	for i, v := range fl.Files {
		fmt.Printf("%v: %v\n", i, v.Name)
	}

}

func NewTemplate(newFileName string) string {
	newFile := drive.File{}
	newFile.Name = newFileName

	driveRes, err := driveSrv.Files.Copy(TEMPLATE, &newFile).Do()
	if err != nil {
		log.Fatal(err)
	}

	return driveRes.Id
}

func DownloadFile() {
	fileCall := driveSrv.Files.Export(TEMPLATE, "application/pdf")
	res, err := fileCall.Download()
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("testing.pdf", body, 0644)
	if err != nil {
		log.Fatal(err)
	}

}

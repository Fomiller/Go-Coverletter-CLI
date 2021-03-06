package drive

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/fomiller/scribe/config"
	"github.com/fomiller/scribe/utils"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
)

var driveSrv *drive.Service

func init() {
	b, err := ioutil.ReadFile(fmt.Sprintf("%v\\src\\github.com\\fomiller\\scribe\\drive\\%v", config.GoPath, config.Scribe.Credentials.Drive))
	if err != nil {
		log.Fatalf(`%v
 
	Please navigate to https://console.cloud.google.com/apis/credentials, to create and download credentials for a 0Auth client ID.

	Save your new credentials as 'credentials.json' inside the scribe/drive folder.
			
	Once credentials are saved to the correct location run 'go install' from the applications root directory.
			
	to test your applicaton run 'scribe' in your terminal

`, err)
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
func getClient(cfig *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := fmt.Sprintf("%v\\src\\github.com\\fomiller\\scribe\\drive\\%v", config.GoPath, "token.json")
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(cfig)
		saveToken(tokFile, tok)
	}
	return cfig.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser, follow the instructions, then type the "+
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
func GetFileId(File string) (string, error) {
	// create query where file matches query exactly
	// more info about creating queries can be found https://developers.google.com/drive/api/v3/search-files
	query := fmt.Sprintf("name='%v'", File)
	// search drive for the matching query
	driveRes, err := driveSrv.Files.List().Q(query).PageSize(1).Do()
	if err != nil {
		return "", err
	}
	if len(driveRes.Files) == 0 {
		log.Fatal("Could not find file, please check your search")
	}
	if len(driveRes.Files) > 1 {
		log.Fatal("Found more then 1 file, please be more specific with your search")
	}
	return driveRes.Files[0].Id, nil
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
func DownloadFile(fileId string, fileName string) {
	// set out put folder name
	var path string
	// check config file for custom path
	if config.Scribe.Download.Path != "" {
		path = config.Scribe.Download.Path
	} else {
		// if download.Path not set use the users $HOME path
		dir, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("Error determining the users $HOME or %%USERPORFILE%% enviornment variable:%v", err)
		}
		path = fmt.Sprintf("%v\\scribe", dir)
	}

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

	if config.Scribe.Download.FolderGeneration == true && strings.Contains(fileName, "-") == true {
		path = fmt.Sprintf("%v\\%v", path, utils.GetFolderName(fileName))
		fmt.Printf("using folderGeneration to download: %v\n", fileName)
	}
	// check if output folder exists, if not then create the folder
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0644)
	}

	// append file type to file name
	fileName = fmt.Sprintf("%v.pdf", fileName)
	// write the file to specified folder
	err = ioutil.WriteFile(fmt.Sprintf("%v\\%v", path, fileName), body, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

// return a slice of fields from a template
func ParseTemplateFields(fileId string) []string {
	// create call to export file from drive
	fileCall := driveSrv.Files.Export(fileId, "text/html")
	// execute download of file call
	res, err := fileCall.Download()
	if err != nil {
		log.Fatal(err)
	}
	// defer closing response body
	defer res.Body.Close()
	// read the res.Body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	// convert body []bytes to a string to be used in regex test
	strBody := string(body)
	// init parsed fields slice to push fields into
	parsedFields := []string{}
	// create regex test
	rgx := regexp.MustCompile(`{{([A-Za-z_]*)}}`)
	// search document for regex matches
	rs := rgx.FindAllStringSubmatch(strBody, -1)
	// push parsed fields to a parsed fields []string
	for _, v := range rs {
		parsedFields = utils.AppendIfMissing(parsedFields, v[1])
	}
	// return parsed fields
	return parsedFields
}

// delete file from drive
func DeleteFile(docId string) error {
	// delete file from drive
	err := driveSrv.Files.Delete(docId).Do()
	// return err
	return err
}

func ListFileNames() []string {
	fileNameList := []string{}
	drvRes, err := driveSrv.Files.List().Do()
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range drvRes.Files {
		fileNameList = append(fileNameList, v.Name)
	}
	return fileNameList
}

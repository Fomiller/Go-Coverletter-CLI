package docs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/docs/v1"
)

var docSrv *docs.Service

func init() {
	b, err := ioutil.ReadFile("./docs/credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/docs")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	docSrv, err = docs.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Docs client: %v", err)
	}
}

// Retrieves a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	tokFile := "./docs/token.json"
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

func TemplateFromFunc(rs []*docs.Request) *docs.BatchUpdateDocumentRequest {
	template := &docs.BatchUpdateDocumentRequest{}
	template.Requests = rs
	return template
}

// OLD FUNCTION NEEDS TO BE REMOVED BROKEN APRART
// func UpdateTemplateFile(templateId string) string {
// 	// **** make this file a dynamic file ****
// 	tpl, err := templateFromFile("./docs/template.json")
// 	batchRes, err := docSrv.Documents.BatchUpdate(templateId, tpl).Do()
// 	if err != nil {
// 		log.Fatalf("BATCH FAIL %v ", err)
// 	}
// 	fmt.Printf("SUCCESSFUL BATCH UPDATE: %v \n", batchRes.DocumentId)
// 	return batchRes.DocumentId
// }

func CreateRequestStruct(m map[string]string) []*docs.Request {
	// create slice to store request objects.
	var rss []*docs.Request
	// create a request object for each key value pair
	for key, val := range m {
		// remove any trailing white space after a comma or  quotation marks in field flags
		key = strings.TrimSpace(key)
		// capitilize and parse templating format
		key = fmt.Sprintf("{{%v}}", strings.ToUpper(key))
		rs := docs.Request{
			ReplaceAllText: &docs.ReplaceAllTextRequest{
				ContainsText: &docs.SubstringMatchCriteria{
					Text:      key,
					MatchCase: true,
				},
				ReplaceText: val,
			},
		}
		rss = append(rss, &rs)
	}
	return rss
}

func NewUpdateTemplateFile(templateId string, rs []*docs.Request) string {
	tpl := TemplateFromFunc(rs)
	batchRes, err := docSrv.Documents.BatchUpdate(templateId, tpl).Do()
	if err != nil {
		log.Fatalf("BATCH FAIL %v ", err)
	}
	return batchRes.DocumentId
}

func ParseTemplateFields(str string) {
	rgx := regexp.MustCompile(`{{([A-Za-z_]*)}}`)
	rs := rgx.FindAllStringSubmatch(str, -1)
	for _, v := range rs {
		fmt.Println(v[1])
	}
}

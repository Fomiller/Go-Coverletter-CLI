package docs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/fomiller/scribe/config"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/docs/v1"
)

var docSrv *docs.Service

func init() {
	b, err := ioutil.ReadFile(fmt.Sprintf("./docs/%v", config.Scribe.Credentials.Docs))
	if err != nil {
		log.Fatalf(`%v

		Please navigate to https://console.cloud.google.com/apis/credentials, to create and download credentials for a 0Auth client ID.

		Save your new credentials as 'credentials.json' inside the scribe/docs folder.

		Once credentials are saved to the correct location run 'go install' from the applications root directory.

		to test your applicaton run 'scribe' in your terminal

	`, err)
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
	fmt.Printf("Go to the following link in your browser, follow the instructions, then type the "+
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

func TemplateFromFunc(rs []*docs.Request) *docs.BatchUpdateDocumentRequest {
	template := &docs.BatchUpdateDocumentRequest{}
	template.Requests = rs
	return template
}

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

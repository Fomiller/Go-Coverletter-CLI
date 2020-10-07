package sheets

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/fomiller/scribe/config"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

var sheetsSrv *sheets.Service

type TempData struct {
	FieldName  interface{}
	FieldValue interface{}
}

func init() {
	b, err := ioutil.ReadFile(fmt.Sprintf("./sheets/%v", config.Scribe.Credentials.Sheets))
	if err != nil {
		log.Fatalf(`%v
 
	Please navigate to https://console.cloud.google.com/apis/credentials, to create and download credentials for a 0Auth client ID.

	Save your new credentials as 'credentials.json' inside the scribe/sheets folder.
			
	Once credentials are saved to the correct location run 'go install' from the applications root directory.
			
	to test your applicaton run 'scribe' in your terminal

`, err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, sheets.SpreadsheetsScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	sheetsSrv, err = sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "./sheets/token.json"
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
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
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

func RunSheets() {

	// Prints the names and majors of students in a sample spreadsheet:
	// https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit
	spreadsheetId := "1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms"
	readRange := "Class Data!A2:E"
	resp, err := sheetsSrv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		fmt.Println("Name, Major:")
		for _, row := range resp.Values {
			// Print columns A and E, which correspond to indices 0 and 4.
			fmt.Printf("%s, %s\n", row[0], row[4])
		}
	}
}

func GetSpreadsheetColumnNames() []string {
	spreadsheetId := "1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms"
	readRange := "Class Data!1:1"
	var columnNames []string
	resp, err := sheetsSrv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrueve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		for _, row := range resp.Values {
			// print each value from the row
			for _, cell := range row {
				// append cell value to column names
				columnNames = append(columnNames, fmt.Sprint(cell))
			}
		}
	}
	return columnNames
}

func SheetsRanges() {
	spreadsheetId := "1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms"
	readRange := "Class Data"
	resp, err := sheetsSrv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(resp.Values))
	fmt.Printf("%v\n", resp.Values[1:])

}

func GetRowData() [][]interface{} {
	spreadsheetId := "1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms"
	readRange := "Class Data"
	resp, err := sheetsSrv.Spreadsheets.Values.Get(spreadsheetId, readRange).MajorDimension("ROWS").Do()
	if err != nil {
		log.Fatal(err)
	}

	return resp.Values
	// CreateTempDataTypes(fieldNames, resp.Values)

}

func FmtSpreadsheetData(fieldNames []string, rows [][]interface{}) [][]TempData {
	newData := TempData{}
	// create slice to hold all data for one field
	spreadsheetData := [][]TempData{}
	rowData := []TempData{}

	// add field names
	for _, fName := range fieldNames {
		newData.FieldName = fName
		rowData = append(rowData, newData)
	}

	// range over all rows
	for _, row := range rows[1:] {
		// print index and value of each item in slice
		// newSlice := rowData
		for i, v := range row {
			// create new slice
			// print index and value of each item in slice
			// fmt.Printf("%v:%v\n", ii, v)
			rowData[i].FieldValue = v
		}
		// create new slice the same length of the data slice
		appendData := make([]TempData, len(rowData))

		// copy over the values for rowData to s
		for i, _ := range rowData {
			appendData[i] = rowData[i]
		}

		// append appendData slice to spreadsheetData
		spreadsheetData = append(spreadsheetData, appendData)
	}
	// fmt.Println(spreadsheetData)
	return spreadsheetData
}

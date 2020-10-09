// Package utils are utilitie functions used in Scribe
package utils

import (
	"fmt"
	"strings"

	"github.com/fomiller/scribe/config"
)

func AppendIfMissing(slice []string, s string) []string {
	for _, str := range slice {
		if str == s {
			return slice
		}
	}
	return append(slice, s)
}

// convert a map[string]Interface to map[string]string
func StrIntfToStrStr(strInterface map[string]interface{}) map[string]string {
	mapStrStr := make(map[string]string)
	for k, v := range strInterface {
		strKey := fmt.Sprintf("%v", k)
		strValue := fmt.Sprintf("%v", v)

		mapStrStr[strKey] = strValue
	}
	return mapStrStr
}

func GetFolderName(fileName string) string {
	folder := strings.Split(fileName, "-")
	if config.Scribe.Download.UsePrefix == false {
		folderName := strings.TrimSpace(folder[1])
		return folderName
	}
	folderName := strings.TrimSpace(folder[0])
	return folderName
}

func RemoveStringsFromFileName(Name string, replaceStr ...string) string {
	for _, v := range replaceStr {
		if strings.Contains(Name, v) {
			Name = strings.ReplaceAll(Name, v, "")
			// remove "  " replace with " "
		}
	}
	// remove double spaces
	Name = strings.ReplaceAll(Name, "  ", " ")
	return Name
}

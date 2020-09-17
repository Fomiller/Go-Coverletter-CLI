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

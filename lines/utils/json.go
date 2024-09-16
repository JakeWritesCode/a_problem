package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func DecodeJSONResponse(response *http.Response, data interface{}) error {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, data)
	if err != nil {
		return err
	}
	return nil
}

func SplitCsvString(csvString string, spaces bool) []string {
	if spaces {
		return strings.Split(csvString, ", ")
	}
	return strings.Split(csvString, ",")
}

package main

import "strings"
import "os"
import "flag"
import "fmt"
import "github.com/pkg/browser"
import "encoding/json"
import "io/ioutil"
import "path/filepath"

var baseURL = flag.String("baseURL", "", "Set Base URL for open-jira")
var ticketNumber = flag.String("ticket", "", "The ticket number you would like to open")

type Configuration struct {
	JIRA_URL string
}

func main() {
	flag.Parse()
	if (strings.TrimRight(*ticketNumber, "\n") == "") {
		fmt.Println("Must use the flag -ticket to determine what ticket in jira to open")
		fmt.Println("Example: open-jira -ticket=ABC-123")
		return
	}

	if strings.TrimRight(*baseURL, "\n") != "" {
		setBaseURL(*baseURL)
	} else {
		*baseURL = getBaseURL()
	}

	browser.OpenURL(*baseURL + "/browse/" + *ticketNumber)
}

func setBaseURL(url string) {
	var configuration Configuration
	configuration.JIRA_URL = url
    exPath := getExecutableDirectory()
	jsonConfig, _ := json.Marshal(configuration)
	path := filepath.Join(exPath, "/config.json")

	ioutil.WriteFile(path, jsonConfig, 0644)
	fmt.Println("Base URL Saved!")
}

func getBaseURL() (url string) {
    exPath := getExecutableDirectory()
	path := filepath.Join(exPath, "/config.json")
	jsonFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	} 

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var configuration Configuration

	json.Unmarshal(byteValue, &configuration)

	return configuration.JIRA_URL
}

func getExecutableDirectory() (path string) {
	ex, exError := os.Executable()

	if exError != nil {
		panic(exError)
	}
    exPath := filepath.Dir(ex)

	return exPath
}


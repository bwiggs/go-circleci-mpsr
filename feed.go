package circleci_mpsr

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Feed struct {
	url      string
	projects []Project
}

// poll fetches the build feed
func (f Feed) poll() ([]Project, error) {
	log.Println("Fetching: ", f.url)
	res, err := http.Get(f.url)
	if err != nil {
		log.Fatal(err)
	}

	// close out the body
	defer res.Body.Close()

	xml, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return parseXml(xml)
}

// FromXml returns the projects xml from circleci.com with the provided
// api token
func parseXml(xmlData []byte) ([]Project, error) {

	projects := []Project{}

	err := xml.Unmarshal(xmlData, &projects)
	if err != nil {
		log.Panic(err)
	}

	projectCount := len(projects)
	if projectCount == 0 {
		fmt.Println("No projects found")
	} else {
		fmt.Printf("Found %v project(s).\n", projectCount)
	}

	// fmt.Printf(": %v", projects.ProjectList)

	return projects, err
}
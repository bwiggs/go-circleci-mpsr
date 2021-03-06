package circleci_mpsr

import (
	"encoding/xml"
)

// Projects is a struct representing a list of projects from the circleci
// response.
type projectsXmlRoot struct {
	XMLName  xml.Name  `xml:"Projects"`
	Projects []Project `xml:"Project"`
}

// Project represents a single project from the project summary feed.
type Project struct {
	XMLName         xml.Name `xml:"Project"`
	Name            string   `xml:"name,attr"`
	Activity        string   `xml:"activity,attr"`
	LastBuildStatus string   `xml:"lastBuildStatus,attr"`
	LastBuildLabel  string   `xml:"lastBuildLabel,attr"`
	LastBuildTime   string   `xml:"lastBuildTime,attr"`
	NextBuildTime   string   `xml:"nextBuildTime,attr"`
	WebUrl          string   `xml:"webUrl,attr"`
}

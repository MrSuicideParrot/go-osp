package messages

import "encoding/xml"

type getVersionRequest string `xml:"get_version"`

type getVersionResponse struct {
	XMLName    xml.Name `xml:"get_version_response"`
	Text       string   `xml:",chardata"`
	StatusText string   `xml:"status_text,attr"`
	Status     string   `xml:"status,attr"`
	Protocol   struct {
		Text    string `xml:",chardata"`
		Version string `xml:"version"`
		Name    string `xml:"name"`
	} `xml:"protocol"`
	Daemon struct {
		Text    string `xml:",chardata"`
		Version string `xml:"version"`
		Name    string `xml:"name"`
	} `xml:"daemon"`
	Scanner struct {
		Text    string `xml:",chardata"`
		Version string `xml:"version"`
		Name    string `xml:"name"`
	} `xml:"scanner"`
}

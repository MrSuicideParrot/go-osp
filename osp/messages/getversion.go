package messages

import "encoding/xml"

type GetVersionRequest struct {
	XMLName    xml.Name `xml:"get_version"`
	Text       string   `xml:",chardata"`
}

type GetVersionResponse struct {
	XMLName    xml.Name `xml:"get_version_response"`
	Text       string   `xml:",chardata"`
	StatusText string   `xml:"status_text,attr"`
	Status     int   `xml:"status,attr"`
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

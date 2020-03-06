package messages

import "encoding/xml"

type getScannerDetailsRequest string `xml:"get_scanner_details"`

type getScannerDetailsResponse struct {
	XMLName       xml.Name `xml:"get_scanner_details_response"`
	Text          string   `xml:",chardata"`
	StatusText    string   `xml:"status_text,attr"`
	Status        string   `xml:"status,attr"`
	Description   string   `xml:"description"`
	ScannerParams struct {
		Text         string `xml:",chardata"`
		ScannerParam []struct {
			Text        string `xml:",chardata"`
			ID          string `xml:"id,attr"`
			Type        string `xml:"type,attr"`
			Name        string `xml:"name"`
			Description string `xml:"description"`
			Default     string `xml:"default"`
		} `xml:"scanner_param"`
	} `xml:"scanner_params"`
}


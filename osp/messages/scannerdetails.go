package messages

import "encoding/xml"

type GetScannerDetailsRequest struct {
	XMLName    xml.Name `xml:"get_scanner_details"`
}

type GetScannerDetailsResponse struct {
	XMLName       xml.Name `xml:"get_scanner_details_response"`
	Text          string   `xml:",chardata"`
	StatusText    string   `xml:"status_text,attr"`
	Status        int   `xml:"status,attr"`
	Description   string   `xml:"description"`
	ScannerParams struct {
		Text         string `xml:",chardata"`
		ScannerParam []struct {
			Text        string `xml:",chardata"`
			ID          string `xml:"id,attr"`
			Type        bool `xml:"type,attr"`
			Name        string `xml:"name"`
			Description string `xml:"description"`
			Default     string `xml:"default"`
		} `xml:"scanner_param"`
	} `xml:"scanner_params"`
}


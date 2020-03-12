package messages

import "encoding/xml"

type DeleteScanRequest struct {
	XMLName xml.Name `xml:"delete_scan"`
	Text    string   `xml:",chardata"`
	ScanID  string   `xml:"scan_id,attr"`
}

type DeleteScanResponse struct {
	XMLName    xml.Name `xml:"delete_scan_response"`
	Text       string   `xml:",chardata"`
	StatusText string   `xml:"status_text,attr"`
	Status     int   `xml:"status,attr"`
}

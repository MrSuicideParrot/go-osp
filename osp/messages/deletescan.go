package messages

import "encoding/xml"

type deleteScanRequest struct {
	XMLName xml.Name `xml:"delete_scan"`
	Text    string   `xml:",chardata"`
	ScanID  string   `xml:"scan_id,attr"`
}

type deleteScanResponse struct {
	XMLName    xml.Name `xml:"delete_scan_response"`
	Text       string   `xml:",chardata"`
	StatusText string   `xml:"status_text,attr"`
	Status     string   `xml:"status,attr"`
}

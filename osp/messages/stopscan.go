package messages

import "encoding/xml"

type StopScanRequest struct {
	XMLName xml.Name `xml:"stop_scan"`
	Text    string   `xml:",chardata"`
	ScanID  string   `xml:"scan_id,attr"`
}

type StopScanResponse struct {
	XMLName    xml.Name `xml:"stop_scan_response"`
	Text       string   `xml:",chardata"`
	StatusText string   `xml:"status_text,attr"`
	Status     int   `xml:"status,attr"`
}


package messages

import "encoding/xml"

type stopScanRequest struct {
	XMLName xml.Name `xml:"stop_scan"`
	Text    string   `xml:",chardata"`
	ScanID  string   `xml:"scan_id,attr"`
}

type stopScanResponse struct {
	XMLName    xml.Name `xml:"stop_scan_response"`
	Text       string   `xml:",chardata"`
	StatusText string   `xml:"status_text,attr"`
	Status     string   `xml:"status,attr"`
}


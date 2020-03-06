package messages

import "encoding/xml"

type startScanRequest struct {
	XMLName       xml.Name `xml:"start_scan"`
	Text          string   `xml:",chardata"`
	Target        string   `xml:"target,attr"`
	Ports         string   `xml:"ports,attr"`
	ScannerParams struct {
		Text       string `xml:",chardata"`
		TargetPort string `xml:"target_port"`
		UseHttps   string `xml:"use_https"`
		Profile    string `xml:"profile"`
	} `xml:"scanner_params"`
}

type startScanResponse struct {
	XMLName    xml.Name `xml:"start_scan_response"`
	Text       string   `xml:",chardata"`
	StatusText string   `xml:"status_text,attr"`
	Status     string   `xml:"status,attr"`
	ID         string   `xml:"id"`
}


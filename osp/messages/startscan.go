package messages

import "encoding/xml"

type VtSingle struct {
	Text    string `xml:",chardata"`
	ID      string `xml:"id,attr"`
	VtValue struct {
		Text string `xml:",chardata"`
		ID   string `xml:"id,attr"`
	} `xml:"vt_value"`
}

type VtGroup struct {
	Text   string `xml:",chardata"`
	Filter string `xml:"filter,attr"`
}

type VtSelection struct {
	Text     string `xml:",chardata"`
	VtSingles []VtSingle `xml:"vt_single"`
	VtGroups []VtGroup `xml:"vt_group"`
}


type StartScanRequest struct {
	XMLName       xml.Name `xml:"start_scan"`
	Text          string   `xml:",chardata"`
	Target        string   `xml:"target,attr"`
	Ports         string   `xml:"ports,attr"`
	ScannerParams []ScannerParam `xml:"scanner_params"`
	VtSelection  VtSelection `xml:"vt_selection"`
}

type StartScanResponse struct {
	XMLName    xml.Name `xml:"start_scan_response"`
	Text       string   `xml:",chardata"`
	StatusText string   `xml:"status_text,attr"`
	Status     int   `xml:"status,attr"`
	ID         string   `xml:"id"`
}


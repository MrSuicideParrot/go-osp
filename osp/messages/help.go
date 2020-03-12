package messages

import "encoding/xml"

type HelpConfig struct {
	XMLName xml.Name `xml:"help"`
	Text    string   `xml:",chardata"`
	Format  string   `xml:"format,attr"`
}

type HelpResponse struct {
	XMLName    xml.Name `xml:"help_response"`
	Text       string   `xml:",chardata"`
	StatusText string   `xml:"status_text,attr"`
	Status     int   `xml:"status,attr"`
	DeleteScan struct {
		Text       string `xml:",chardata"`
		Attributes struct {
			Text   string `xml:",chardata"`
			ScanID string `xml:"scan_id"`
		} `xml:"attributes"`
		Elements    string `xml:"elements"`
		Description string `xml:"description"`
	} `xml:"delete_scan"`
	Help struct {
		Text       string `xml:",chardata"`
		Attributes struct {
			Text   string `xml:",chardata"`
			Format string `xml:"format"`
		} `xml:"attributes"`
		Elements    string `xml:"elements"`
		Description string `xml:"description"`
	} `xml:"help"`
	GetVersion struct {
		Text        string `xml:",chardata"`
		Attributes  string `xml:"attributes"`
		Elements    string `xml:"elements"`
		Description string `xml:"description"`
	} `xml:"get_version"`
	StopScan []struct {
		Text       string `xml:",chardata"`
		Attributes struct {
			Text   string `xml:",chardata"`
			ScanID string `xml:"scan_id"`
		} `xml:"attributes"`
		Elements    string `xml:"elements"`
		Description string `xml:"description"`
	} `xml:"stop_scan"`
	GetScannerDetails struct {
		Text        string `xml:",chardata"`
		Attributes  string `xml:"attributes"`
		Elements    string `xml:"elements"`
		Description string `xml:"description"`
	} `xml:"get_scanner_details"`
	StartScan struct {
		Text       string `xml:",chardata"`
		Attributes struct {
			Text   string `xml:",chardata"`
			ScanID string `xml:"scan_id"`
			Target string `xml:"target"`
			Ports  string `xml:"ports"`
		} `xml:"attributes"`
		Elements struct {
			Text          string `xml:",chardata"`
			ScannerParams struct {
				Text        string `xml:",chardata"`
				Profile     string `xml:"profile"`
				TargetPort  string `xml:"target_port"`
				UseHttps    string `xml:"use_https"`
				W3afTimeout string `xml:"w3af_timeout"`
			} `xml:"scanner_params"`
		} `xml:"elements"`
		Description string `xml:"description"`
	} `xml:"start_scan"`
	GetScans struct {
		Text       string `xml:",chardata"`
		Attributes struct {
			Text    string `xml:",chardata"`
			ScanID  string `xml:"scan_id"`
			Details string `xml:"details"`
		} `xml:"attributes"`
		Elements    string `xml:"elements"`
		Description string `xml:"description"`
	} `xml:"get_scans"`
}

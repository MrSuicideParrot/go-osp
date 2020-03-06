package messages

import "encoding/xml"

type getScansRequest struct {
	XMLName xml.Name `xml:"get_scans"`
	Text    string   `xml:",chardata"`
	ScanID  string   `xml:"scan_id,attr"`
	Details string   `xml:"details,attr"`
}

type getScansResponse struct {
	XMLName    xml.Name `xml:"get_scans_response"`
	Text       string   `xml:",chardata"`
	StatusText string   `xml:"status_text,attr"`
	Status     string   `xml:"status,attr"`
	Scan       struct {
		Text      string `xml:",chardata"`
		ID        string `xml:"id,attr"`
		Target    string `xml:"target,attr"`
		EndTime   string `xml:"end_time,attr"`
		Progress  string `xml:"progress,attr"`
		StartTime string `xml:"start_time,attr"`
		Results   struct {
			Text   string `xml:",chardata"`
			Result struct {
				Text     string `xml:",chardata"`
				Host     string `xml:"host,attr"`
				Severity string `xml:"severity,attr"`
				Port     string `xml:"port,attr"`
				TestID   string `xml:"test_id,attr"`
				Name     string `xml:"name,attr"`
				Type     string `xml:"type,attr"`
			} `xml:"result"`
		} `xml:"results"`
	} `xml:"scan"`
}





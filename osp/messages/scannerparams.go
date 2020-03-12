package messages

type ScannerParam struct {
	Text       string `xml:",chardata"`
	TargetPort string `xml:"target_port"`
	UseHttps   string `xml:"use_https"`
	Profile    string `xml:"profile"`
}
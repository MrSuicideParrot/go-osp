package osp

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/mrsuicideparrot/go-osp/osp/messages"
	"github.com/mrsuicideparrot/go-osp/osp/ports"
	"github.com/mrsuicideparrot/go-osp/osp/profiles"
	"github.com/mrsuicideparrot/go-osp/osp/vtgroups"
	"golang.org/x/net/html/charset"
	"io/ioutil"
	"net"
	"strconv"
	"strings"
	"time"
)

// Client struct with connection details of an OSP server
type Client struct {
	host    string
	port    string
	tlsConf *tls.Config
}

var connectionTimeout time.Duration

func init() {
	connectionTimeout = time.Second * 2
}

// Creates a client struct  with the provided details
func New(host string, port int, clientCertificatePath, clientKeyPath, certificateAuthorityPath string) (*Client, error) {
	cl := Client{}

	cert, err := tls.LoadX509KeyPair(clientCertificatePath, clientKeyPath)
	if err != nil {
		return nil, err
	}

	caRaw, err := ioutil.ReadFile(certificateAuthorityPath)
	if err != nil {
		return nil, err
	}

	rootPool := x509.NewCertPool()
	rootPool.AppendCertsFromPEM([]byte(caRaw))

	cl.tlsConf = &tls.Config{
		Certificates:       []tls.Certificate{cert},
		RootCAs:            rootPool,
		InsecureSkipVerify: true,
	}

	cl.host = host
	cl.port = strconv.Itoa(port)

	resp, err := cl.GetVersion()

	if err != nil {
		return nil, err
	} else if resp.Status != 200 {
		return nil, errors.New("Response message " + strconv.Itoa(resp.Status))
	}

	return &cl, nil
}

func (cl *Client) sendRequest(req interface{}, resp interface{}) error {
	dialer := &net.Dialer{Timeout: connectionTimeout}
	conn, err := tls.DialWithDialer(dialer, "tcp", cl.host+":"+cl.port, cl.tlsConf)
	if err != nil {
		return err
	}
	defer conn.Close()

	reqXML, err := xml.Marshal(req)
	if err != nil {
		return err
	}

	if _, err := conn.Write(reqXML); err != nil {
		return err
	}

	dec := xml.NewDecoder(conn)
	dec.CharsetReader = charset.NewReaderLabel
	dec.Strict = false

	if err := dec.Decode(resp); err != nil {
		return err
	}

	return nil
}

// Retrieve a scan result
func (cl *Client) GetScan(uuid string, details bool) (*messages.GetScansResponse, error) {

	m := &messages.GetScansRequest{
		ScanID:  uuid,
		Details: details,
	}

	resp := &messages.GetScansResponse{}

	if err := cl.sendRequest(m, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// Get openvas version
func (cl *Client) GetVersion() (*messages.GetVersionResponse, error) {

	m := &messages.GetVersionRequest{}
	resp := &messages.GetVersionResponse{}

	if err := cl.sendRequest(&m, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// Remove scan results on server
func (cl *Client) DeleteScan(uuid string) (*messages.DeleteScanResponse, error) {

	m := &messages.DeleteScanRequest{
		ScanID: uuid,
	}

	resp := &messages.DeleteScanResponse{}

	if err := cl.sendRequest(m, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// Start a new openvas scan with the provided details.
func (cl *Client) StartScan(target string, portList ports.PortList, vtGroup []vtgroups.VtGroup, profile profiles.Profile) (*messages.StartScanResponse, error) {

	groups := []messages.VtGroup{}

	for _, i := range vtGroup {
		groups = append(groups, messages.VtGroup{Filter: "family=" + string(i)})
	}

	m := &messages.StartScanRequest{
		Target: target,
		Ports:  portArrayToString(portList),
		ScannerParams: []messages.ScannerParam{
			{
				Profile: string(profile),
			},
		},
		VtSelection: messages.VtSelection{
			VtGroups: groups,
		},
	}

	resp := &messages.StartScanResponse{}

	if err := cl.sendRequest(m, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (cl *Client) StartScanV2(target string, portListTCP ports.PortList, portListUDP ports.PortList, vtGroup []vtgroups.VtGroup, profile profiles.Profile) (*messages.StartScanResponse, error) {

	groups := []messages.VtGroup{}

	for _, i := range vtGroup {
		groups = append(groups, messages.VtGroup{Filter: "family=" + string(i)})
	}

	m := &messages.StartScan{
		ScannerParams: []messages.ScannerParam{
			{
				Profile: string(profile),
			},
		},
		VtSelection: messages.VtSelection{
			VtGroups: groups,
		},
		Targets: struct {
			Text    string            `xml:",chardata"`
			Targets []messages.Target `xml:"target"`
		}(struct {
			Text    string
			Targets []messages.Target
		}{Text: "", Targets: []messages.Target{
			{
				Hosts:     target,
				Ports:     "T:" + portArrayToString(portListTCP) + ",U:" + portArrayToString(portListUDP),
				AliveTest: "1",
				AliveTestMethods: struct {
					Text          string `xml:",chardata"`
					Icmp          string `xml:"icmp"`
					TcpAck        string `xml:"tcp_ack"`
					TcpSyn        string `xml:"tcp_syn"`
					Arp           string `xml:"arp"`
					ConsiderAlive string `xml:"consider_alive"`
				}{ConsiderAlive: "1"},
			}},
		}),
	}

	resp := &messages.StartScanResponse{}

	if err := cl.sendRequest(m, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func portArrayToString(ports []int) string {
	var result string

	if len(ports) >= 1 {
		result = strconv.Itoa(ports[0])
	}

	if len(ports) > 1 {
		for _, p := range ports[1:] {
			result += ", " + strconv.Itoa(p)
		}
	}

	return result
}

func stringToPortArray(ports string) []int {
	pS := strings.Split(ports, ",")

	var pI []int

	for _, i := range pS {
		iI, err := strconv.Atoi(i)
		if err != nil {
			fmt.Errorf(err.Error())
			continue
		}

		pI = append(pI, iI)
	}
	return pI
}

// Stop a scan
func (cl *Client) StopScan(uuid string) (*messages.StopScanResponse, error) {

	m := &messages.StopScanRequest{
		ScanID: uuid,
	}
	resp := &messages.StopScanResponse{}

	if err := cl.sendRequest(m, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (cl *Client) Help() (*messages.HelpResponse, error) {
	m := &messages.HelpConfig{
		Format: "xml",
	}
	resp := &messages.HelpResponse{}

	if err := cl.sendRequest(m, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

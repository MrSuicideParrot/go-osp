package osp

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/mrsuicideparrot/go-osp/osp/messages"
	ports2 "github.com/mrsuicideparrot/go-osp/osp/ports"
	"github.com/mrsuicideparrot/go-osp/osp/profiles"
	"github.com/mrsuicideparrot/go-osp/osp/vtgroups"
	"golang.org/x/net/html/charset"
	"io/ioutil"
	"strconv"
	"strings"
)

type Client struct{
	host string
	port string
	tlsConf *tls.Config
}


func New(host string, port int, clientCertificatePath, clientKeyPath, certificateAuthorityPath string)  (*Client, error){
	cl := Client{}

	cert, err := tls.LoadX509KeyPair(clientCertificatePath, clientKeyPath)
	if err != nil{
		return nil, err
	}

	caRaw, err := ioutil.ReadFile(certificateAuthorityPath)
	if err != nil{
		return nil, err
	}

	rootPool := x509.NewCertPool()
	rootPool.AppendCertsFromPEM([]byte(caRaw))

	cl.tlsConf = &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs: rootPool,
		InsecureSkipVerify:true,
	}

	cl.host = host
	cl.port = strconv.Itoa(port)

	resp, err := cl.GetVersion()

	if err != nil{
		return nil,err
	} else if resp.Status != 200 {
		return nil, errors.New("Response message " + strconv.Itoa(resp.Status))
	}

	return &cl, nil
}

func (cl *Client)sendRequest(req interface{}, resp interface{}) error{
	conn, err := tls.Dial("tcp", cl.host + ":" + cl.port, cl.tlsConf)
	if err != nil{
		return err
	}
	defer conn.Close()

	reqXML, err := xml.Marshal(req)
	if err != nil{
		return err
	}

	if _, err := conn.Write(reqXML); err != nil{
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

func (cl *Client)GetScan(uuid string, details bool) (*messages.GetScansResponse, error){

	m := &messages.GetScansRequest{
		ScanID:  uuid,
		Details: details,
	}

	resp :=  &messages.GetScansResponse{}

	if err := cl.sendRequest(m,resp); err != nil{
		return nil,err
	}

	return resp, nil
}

func (cl *Client)GetVersion()(*messages.GetVersionResponse, error){

	m := &messages.GetVersionRequest{}
	resp :=  &messages.GetVersionResponse{}

	if err := cl.sendRequest(&m, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (cl *Client)DeleteScan(uuid string)(*messages.DeleteScanResponse, error){

	m := &messages.DeleteScanRequest{
		ScanID: uuid,
	}

	resp := &messages.DeleteScanResponse{}

	if err := cl.sendRequest(m, resp); err != nil{
		return nil, err
	}

	return resp, nil
}

func (cl *Client)StartScan(target string, ports []int, vtGroup vtgroups.VtGroup, profile profiles.Profile) (*messages.StartScanResponse, error) {

	m := &messages.StartScanRequest{
		Target: target,
		Ports: portArrayToString(ports2.CommonPorts),
		ScannerParams: []messages.ScannerParam{
			{
				Profile:    string(profile),
			},
		},
		VtSelection: messages.VtSelection{
			VtGroups:  []messages.VtGroup{
				{
					Filter: "family=" + string(vtGroup),
				},
			},
		},
	}

	resp := &messages.StartScanResponse{}

	if err := cl.sendRequest(m, resp); err != nil{
		return nil, err
	}

	return resp, nil
}

func portArrayToString(ports []int) string{
	var result string

	if len(ports) >= 1{
		result = strconv.Itoa(ports[0])
	}

	if len(ports) > 1 {
		for _, p := range ports[1:]{
			result += ", " + strconv.Itoa(p)
		}
	}

	return result
}

func stringToPortArray(ports string) []int{
	pS := strings.Split(ports,",")

	var pI []int

	for _,i := range pS{
		iI, err := strconv.Atoi(i)
		if err != nil{
			fmt.Errorf(err.Error())
			continue
		}

		pI = append(pI, iI)
	}
	return pI
}

func (cl *Client)StopScan(uuid string)(*messages.StopScanResponse, error){

	m := &messages.StopScanRequest{
		ScanID:uuid,
	}
	resp := &messages.StopScanResponse{}

	if err := cl.sendRequest(m, resp); err != nil{
		return nil, err
	}

	return resp, nil
}

func (cl *Client)Help()(*messages.HelpResponse, error){
	m := &messages.HelpConfig{
		Format:  "xml",
	}
	resp := &messages.HelpResponse{}

	if err := cl.sendRequest(m, resp); err != nil{
		return nil, err
	}

	return resp, nil
}
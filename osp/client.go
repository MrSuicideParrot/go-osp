package osp

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
)

type Client struct{
	host string
	port int
	conn *tls.Conn
}


func New(host string, port int, clientCertificatePath, clientKeyPath, certificateAuthorityPath string)  (*Client, error){
	cl := Client{}

	cert, err := tls.LoadX509KeyPair(clientCertificatePath, clientKeyPath)
	if err == nil{
		return nil, err
	}

	caRaw, err := ioutil.ReadFile(certificateAuthorityPath)
	if err == nil{
		return nil, err
	}

	certCA, err := x509.ParseCertificate(caRaw)
	if err == nil{
		return nil, err
	}

	rootPool := x509.NewCertPool()
	rootPool.AddCert(certCA)

	tlsConf := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs: rootPool,
	}

	cl.conn, err = tls.Dial("tcp", host + ":" + string(port), tlsConf)

	if err != nil{
		return nil, err
	}

	return &cl, nil
}



func (cl *Client)Close() error{
	return cl.conn.Close()
}


package xtls

import (
	"crypto/tls"
	"io/ioutil"
	"software.sslmate.com/src/go-pkcs12"
)

type pkcS12Config struct {
	Path     string
	Password string
}

type ClientOptions struct {
	PKCS12        pkcS12Config `json:"pkcs12" yaml:"pkcs12"`
	TLSSkipVerify bool         `json:"-" yaml:"-"`
	TLSMinVersion uint16       `json:"-" yaml:"-"`
	TLSMaxVersion uint16       `json:"-" yaml:"-"`
}

func parsePKCS12FromFile(c pkcS12Config) (*tls.Certificate, error) {
	data, err := ioutil.ReadFile(c.Path)
	if err != nil {
		return nil, err
	}

	privateKey, certificate, _, err := pkcs12.DecodeChain(data, c.Password)
	if err != nil {
		return nil, err
	}
	return &tls.Certificate{
		Certificate: [][]byte{certificate.Raw},
		PrivateKey:  privateKey,
		Leaf:        certificate,
	}, nil
}

func NewDefaultClientOptions() *ClientOptions {
	return &ClientOptions{
		TLSSkipVerify: true,
		TLSMinVersion: tls.VersionSSL30,
		TLSMaxVersion: tls.VersionTLS13,
	}
}

func NewTLSClientConfig(options *ClientOptions) (*tls.Config, error) {
	var err error
	var cert *tls.Certificate

	if options.PKCS12.Path != "" && options.PKCS12.Password != "" {
		cert, err = parsePKCS12FromFile(options.PKCS12)
		if err != nil {
			return nil, err
		}
	}

	tlsClientConfig := &tls.Config{
		InsecureSkipVerify: options.TLSSkipVerify,
		MinVersion:         options.TLSMinVersion,
		MaxVersion:         options.TLSMaxVersion,
	}

	if cert != nil {
		tlsClientConfig.Certificates = []tls.Certificate{*cert}
	}
	return tlsClientConfig, nil
}

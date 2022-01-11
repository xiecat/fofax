package utils

import (
	"crypto/tls"
	"fmt"
	"github.com/xiecat/fofax/internal/printer"
	"net/http"
)

// GetSerialNumber 转换证书
func GetSerialNumber(url string) string {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get(url)
	if err != nil {
		printer.Errorf("%s Request failed,err : %s", url, err)
		return ""
	}
	defer resp.Body.Close()
	cert := resp.TLS
	if cert == nil {
		printer.Errorf("%s Failed to get certificate serial number", url)
		return ""
	}
	certInfo := cert.PeerCertificates
	return fmt.Sprintf(`cert="%s"`, certInfo[0].SerialNumber.String())
}

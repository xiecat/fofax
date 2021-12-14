package utils

import (
	"crypto/tls"
	"fmt"
	"fofax/internal/printer"
	"net/http"
	"os"
)

// GetSerialNumber 转换证书
func GetSerialNumber(url string) string {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get(url)
	defer resp.Body.Close()

	if err != nil {
		printer.Errorf("%s 请求失败,err : %s", url, err)
		os.Exit(1)
	}
	// fmt.Println(resp.TLS.PeerCertificates[0])
	certInfo := resp.TLS.PeerCertificates
	return fmt.Sprintf(`cert="%s"`, certInfo[0].SerialNumber.String())
}

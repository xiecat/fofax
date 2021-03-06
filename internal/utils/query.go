package utils

import (
	"fmt"
	"github.com/xiecat/fofax/internal/printer"
	"net/http"
)

// GetSerialNumber 转换证书
func GetSerialNumber(client *http.Client, url string) string {
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

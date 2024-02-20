/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package database

import (
	"github.com/go-sql-driver/mysql"
	"crypto/tls"
	"io/ioutil"
	"crypto/x509"

	"helianthus/utils"
)



func registryTLSConfig() {
	// Register a custom TLS config
	if utils.GetEnvWithDefault("MYSQL_TLS", "false") == "true" {		
		rootCertPool := x509.NewCertPool()
		// Load the CA certificate
		ca, err := ioutil.ReadFile("/app/ssl/ca.pem")
		if err != nil {
			panic(err)
		}

		if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
			log.Fatal("Failed to append PEM.")
		}
		
		mysql.RegisterTLSConfig("custom", &tls.Config{
			RootCAs:      rootCertPool
		})
		return
	}
}

func init() {
	registryTLSConfig()
}
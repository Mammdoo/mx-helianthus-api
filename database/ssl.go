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
	"log"

	"helianthus/utils"
)



func registryTLSConfig() {
	// Register a custom TLS config
	if utils.GetEnvWithDefault("DB_TLS", "false") != "false" {		
		rootCertPool := x509.NewCertPool()
		// Load the CA certificate
		ca, err := ioutil.ReadFile("/app/ssl/ca.pem")
		if err != nil {
			panic(err)
		}

		if ok := rootCertPool.AppendCertsFromPEM(ca); !ok {
			log.Fatal("Failed to append CA pem.")
		}
	
		mysql.RegisterTLSConfig("custom", &tls.Config{
			RootCAs:      rootCertPool,
		})
		return
	}
}
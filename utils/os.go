/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package utils

import (
	"os"
)

func GetEnvWithDefault(key string, def string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return def
}
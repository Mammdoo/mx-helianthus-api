/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package utils

import (
	"os"
	"strings"
)

func GetEnvWithDefault(key string, def string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return def
}

func StringToSlice(str string, split string) []string {
	var rst []string

	rst = strings.Split(str, split)
	return rst
}

func SliceToString(slice []string, split string) string {
	var rst string

	rst = strings.Join(slice, split)
	return rst
}
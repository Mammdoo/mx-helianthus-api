/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe("0.0.0.0:2048", httpHandler)
}

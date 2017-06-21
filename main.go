/* ASN1 Basic */

package main

import (
	"encoding/asn1"
	"fmt"
	"os"
	"time"
)

func main() {

	t := time.Now()
	fmt.Println("Before marshalling: ", t.String())

	mdata, err := asn1.Marshal(t)

}

func chkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", error.Error())
	}
	os.Exit(1)
}

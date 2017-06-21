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
        ckErr(err)
        fmt.Println("Marshall ok!")
        
        var newtime = new(time.Time)
        _, err1 := asn1.Unmarshal(mdata, newtime)
        ckErr(err1)
        fmt.Println("After marshal/unmarshal: ", newtime.String())
        
}

func ckErr(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
                os.Exit(1)
	}
}

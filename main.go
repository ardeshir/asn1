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
        fmt.Println("Marshalled time ok!")
        
        var newtime = new(time.Time)
        _, err1 := asn1.Unmarshal(mdata, newtime)
        ckErr(err1)
        fmt.Println("After marshal/unmarshal time: ", newtime.String())
       
        s := "hello \u00bc"
        fmt.Println("Before marshalling string: ", s)
        
        mdata1, err := asn1.Marshal(s)
        ckErr(err)
        fmt.Println("Marshalled string ok!")
       
        var newstr string
        _ , err2 := asn1.Unmarshal(mdata1, &newstr)
        ckErr(err2)

        fmt.Println("After marshal/unmarshal new string: ", newstr)        
}

func ckErr(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
                os.Exit(1)
	}
}

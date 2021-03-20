package main

import (
	"encoding/base64"
	"fmt"
)


func main(){
	resp := []byte("\x00" + "your_email_here" + "\x00" + "your_password_here")
	// fmt.Println(string(resp), resp)

	sEnc := base64.StdEncoding.EncodeToString([]byte(resp))
	fmt.Println(sEnc) //this is your encoded key.
}
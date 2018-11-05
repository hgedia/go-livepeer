package main

import (
	"fmt"
	"net/url"
)

func (w *wizard) registerENSSubdomain() {
	fmt.Println("Enter subdomain 'name' in <name>.transcoder.eth :")
	subDomain := w.read()
	fmt.Printf("Subdomain to create wizard %s", subDomain)
	val := url.Values{
		"subDomain": {fmt.Sprintf("%v", subDomain)},
	}
	httpPostWithParams(fmt.Sprintf("http://%v:%v/registerENSsubDomain", w.host, w.httpPort), val)
	return
}

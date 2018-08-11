package main

import (
	"fmt"
	"net/url"
)

func (w *wizard) registerENSSubdomain() {
	fmt.Println("Enter subdomain :")
	subDomain := w.read()
	fmt.Printf("Subdomain to create %s :", subDomain)
	val := url.Values{
		"subDomain": {fmt.Sprintf("%v", subDomain)},
	}
	httpPostWithParams(fmt.Sprintf("http://%v:%v/registerENSsubDomain", w.host, w.httpPort), val)
	return
}

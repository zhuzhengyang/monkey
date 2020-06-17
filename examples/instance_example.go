package main

import (
	"fmt"
	"net/http"
	"reflect"

	"bou.ke/monkey"
)

func main() {
	var d *http.Client
	monkey.PatchInstanceMethod(reflect.TypeOf(d), "Get", func(_ *http.Client, url string) (*http.Response, error) {
		return nil, fmt.Errorf("%s no dialing allowed", url)
	})
	_, err := http.Get("http://google.com")
	fmt.Println(err) // Get http://google.com: no dialing allowed
}

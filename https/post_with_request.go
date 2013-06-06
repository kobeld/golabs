package main

import (
	// "bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	cmd := "srg_create g3 kobeld2.local Group 3 this_is_g3 g3"

	client := &http.Client{}
	reqest, _ := http.NewRequest("POST", "http://localhost:5280/rest/", strings.NewReader(cmd))
	reqest.Host = "kobeld2.local"

	response, err := client.Do(reqest)
	if err != nil {
		fmt.Printf("Err: %+v \n", err)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	bodystr := string(body)
	fmt.Println(bodystr)
}

// $ wget --header="host: kobeld2.local" http://localhost:5280/rest/ --server-response --post-data 'srg_create g3 kobeld2.local Group_3 this_is_g3 g3'

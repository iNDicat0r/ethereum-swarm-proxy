package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// SwarmHandler creates a new subrequest to swarm endpoint and returns the result
func SwarmHandler(swarmHost string, swarmPort string, swarmHash string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fullURI := "http://" + swarmHost + ":" + swarmPort + "/bzz:/" + swarmHash + r.URL.String()
		client := &http.Client{}
		req, err := http.NewRequest(r.Method, fullURI, r.Body)
		if err != nil {
			log.Println(err)
			fmt.Fprint(w, err.Error())
			return
		}
		resp, err := client.Do(req)
		if err != nil {
			log.Println(err)
			fmt.Fprint(w, err.Error())
			return
		}
		defer resp.Body.Close()

		for i, v := range resp.Header {
			w.Header().Set(i, v[0])
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			fmt.Fprint(w, err.Error())
			return
		}
		log.Println(fullURI)
		fmt.Fprint(w, string(body))
	}
}

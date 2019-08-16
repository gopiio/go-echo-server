package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"sort"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var requests []string
		log.Println("Processing request from ", r.Host)
		hostName, _ := os.Hostname()
		address, _ := net.InterfaceAddrs()

		for _, address := range address {
			if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					requests = append(requests, fmt.Sprintf("PodIP:%v", ipnet.IP.String()))
					break
				}
			}
		}
		requests = append(requests, fmt.Sprintf("HostName:%v", hostName))
		for name, headers := range r.Header {
			for _, val := range headers {
				requests = append(requests, fmt.Sprintf("%v:%v", name, val))
			}
		}
		requests = append(requests, fmt.Sprintf("Tag:%v", "test"))
		requests = append(requests, "Whitelist:True")
		sort.Strings(requests)
		fmt.Fprintf(w, strings.Join(requests, "\n"))
	})
	log.Println("Starting server and listening on 8080")
	http.ListenAndServe(":8080", nil)
}

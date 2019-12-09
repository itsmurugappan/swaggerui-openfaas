package main

import (
	"net/url"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/rakyll/statik/fs"
	_ "github.com/itsmurugappan/swaggerui-openfaas/statik"
)

func main() {

	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}	

	http.Handle("/swaggerui/", http.StripPrefix("/swaggerui/", http.FileServer(statikFS)))

	//reverse proxy openfaas request
	http.HandleFunc("/", serveReverseProxy)

	http.ListenAndServe(":8080", nil)

}


func serveReverseProxy(res http.ResponseWriter, req *http.Request) {
	gateway, _ := os.LookupEnv("openfaas_gateway")
	u, _ := url.Parse(gateway)
	u.Path = "/function"	

	// create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(u)

	req.URL.Host = u.Host
	req.URL.Scheme = u.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = u.Host

	proxy.ServeHTTP(res, req)
}
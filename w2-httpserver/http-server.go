package main

import (
	"fmt"
	"io"
	"log"
	"runtime"

	"net/http"
)

func main() {
	defer startServer()

	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/", buzHandler)
}

func startServer() {
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthz(rsp http.ResponseWriter, req *http.Request) {
	rsp.WriteHeader(200)
}

func buzHandler(w http.ResponseWriter, r *http.Request) {
	defer io.WriteString(w, "success")
	setResponseHeader(w, r)
	systemInfo()
	logRequest(r)
}

func systemInfo() {
	fmt.Println("-----------------------------------------------------------")
	fmt.Println(fmt.Sprintf("VERSION:\n  GoArch: %v\n  GoRoot: %v\n  Goos: %v\n  NumCpu: %v\n  NumGoroutine: %v\n",
		runtime.GOARCH, runtime.GOROOT(), runtime.GOOS, runtime.NumCPU(), runtime.NumGoroutine()))
}

func logRequest(r *http.Request) {
	fmt.Println("-----------------------------------------------------------")
	fmt.Println("addr: ", r.RemoteAddr)
	fmt.Println("host: ", r.Host)
}

func setResponseHeader(rsp http.ResponseWriter, req *http.Request) {
	headerRsp := rsp.Header()
	for s, strings := range req.Header {
		headerRsp.Set(s, strings[0])
	}
}

package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type HttpHandler struct{}

func main() {
	handler := HttpHandler{}
	http.ListenAndServe(":3000", handler)

}

func (h HttpHandler) printHttpRequestDetail(req *http.Request) {
	fmt.Printf("::Client address\t: %s\n", req.URL.Hostname())
	fmt.Printf("::Client port\t: %s\n", req.URL.Port())
	fmt.Printf("::Request command\t: %s\n", req.Method)
}

func (h HttpHandler) sendHttpResponseHeader(res http.ResponseWriter) {
	res.WriteHeader(200)
	res.Header().Set("Content-type", "text/html")
}

func simpleCalc(para1 int, para2 int) int {
	return para1 * para2
}

func parameterRetrieval(msg string) []int {
	var result []int
	fields := strings.Split(msg, "&")
	para1, _ := strconv.Atoi(strings.Split(fields[0], "=")[1])
	para2, _ := strconv.Atoi(strings.Split(fields[1], "=")[1])
	result = append(result, para1)
	result = append(result, para2)

	return result
}

func (h HttpHandler) DoGET(res http.ResponseWriter, req *http.Request) {
	fmt.Println("## DoGET() activated.")

	h.printHttpRequestDetail(req)
	h.sendHttpResponseHeader(res)

	if routine := strings.Split(req.URL.RequestURI(), "?"); len(routine) != 1 {
		parameter := parameterRetrieval(routine[1])
		result := simpleCalc(parameter[0], parameter[1])

		res.Write([]byte("<html>"))
		getResponse := fmt.Sprintf("GET request for calculation => %d x %d = %d", parameter[0], parameter[1], result)
		res.Write([]byte(getResponse))
		res.Write([]byte("</html>"))
		fmt.Printf("## GET request for calculation => %d x %d = %d", parameter[0], parameter[1], result)
	} else {
		res.Write([]byte("<html>"))
		res.Write([]byte(fmt.Sprintf("<p>HTTP Request GET for Path: %s</p>", req.URL.RequestURI())))
		res.Write([]byte("</html>"))
		fmt.Printf("## GET request for directory => %s", req.URL.RequestURI())
	}
}

func (h HttpHandler) DoPOST(res http.ResponseWriter, req *http.Request) {
	fmt.Println("## DoPOST() activated.")

	h.printHttpRequestDetail(req)
	h.sendHttpResponseHeader(res)

	postData := req.PostForm
}

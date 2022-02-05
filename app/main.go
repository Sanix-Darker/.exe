package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Code struct {
	Name    string `json:"name"`
	Lang    string `json:"lang"`
	Code    string `json:"code"`
	Version string `json:"version"`
	Stdin   string `json:"stdin"`
}

type ErrResp struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

var pistonUrl = "http://0.0.0.0:2000/api/v2"

func main() {
	fmt.Println("[-] .exe started at :4321")

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/execute", executeHandler)
	mux.HandleFunc("/api/v1/runtimes", runtimesHandler)

	err := http.ListenAndServe(":4321", mux)
	log.Fatal(err)
}

func returnHttpError(w http.ResponseWriter, resp string, status int) {
	var errR ErrResp

	w.WriteHeader(status)
	err := json.Unmarshal([]byte("{\"message\": \""+resp+"\"}"), &errR)
	if err != nil {
		log.Panic(err)
	}

	out, _ := json.Marshal(&errR)
	w.Write(out)
}

func runtimesHandler(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	switch req.Method {
	case "GET":
		req, err := http.NewRequest("GET", pistonUrl+"/runtimes", nil)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Panic(err)
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		w.WriteHeader(resp.StatusCode)
		w.Write(body)
	default:
		returnHttpError(w, "Your payload/method is not good !", http.StatusInternalServerError)
	}
}

func executeHandler(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	switch req.Method {
	case "POST":
		// We try to parse/decode the payload
		d := json.NewDecoder(req.Body)

		code := &Code{}
		err := d.Decode(&code)

		if err != nil {
			returnHttpError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if len(code.Lang) < 2 || len(code.Version) < 2 || len(code.Code) < 5 || len(code.Name) < 2 {
			returnHttpError(w, "Payload incorrect or too short !", http.StatusBadRequest)
			return
		}

		payload := []byte(`{
			"language": "` + code.Lang + `",
			"version": "` + code.Version + `",
			"files": [
				{
					"name": "` + code.Name + `",
					"content": "` + strings.Replace(code.Code, "\"", "\\\"", -1) + `"
				}
			],
			"stdin": "` + code.Stdin + `",
			"encoding": "utf-8",
			"compile_timeout": 10000,
			"run_timeout": 2999
		}`)

		req, err := http.NewRequest("POST", pistonUrl+"/execute", bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Panic(err)
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		w.WriteHeader(resp.StatusCode)
		w.Write(body)
	default:
		returnHttpError(w, "Your payload/route is not good !", http.StatusInternalServerError)
	}
}

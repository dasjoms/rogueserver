package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
)

func main() {
	host := flag.String("host", "http://localhost:8080", "server host")
	username := flag.String("username", "", "account username")
	flag.Parse()

	if *username == "" {
		fmt.Println("-username is required")
		return
	}

	// GET /training/data
	dataURL := fmt.Sprintf("%s/training/data?username=%s", *host, *username)
	resp, err := http.Get(dataURL)
	if err != nil {
		fmt.Println("training data request failed:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("/training/data response:")
	fmt.Println(string(body))

	// GET /training/sessions
	sessURL := fmt.Sprintf("%s/training/sessions?username=%s", *host, *username)
	resp, err = http.Get(sessURL)
	if err != nil {
		fmt.Println("training sessions request failed:", err)
		return
	}
	defer resp.Body.Close()
	body, _ = io.ReadAll(resp.Body)
	fmt.Println("/training/sessions response:")
	fmt.Println(string(body))

	// POST /training/actions
	actionURL := fmt.Sprintf("%s/training/actions?username=%s", *host, *username)
	action := map[string]interface{}{"name": "UseItem", "args": map[string]interface{}{"itemId": 1}}
	actionBody, _ := json.Marshal(action)
	resp, err = http.Post(actionURL, "application/json", bytes.NewReader(actionBody))
	if err != nil {
		fmt.Println("post action failed:", err)
		return
	}
	resp.Body.Close()
	fmt.Println("Queued action")

	// GET /training/actions
	resp, err = http.Get(actionURL)
	if err != nil {
		fmt.Println("fetch actions failed:", err)
		return
	}
	defer resp.Body.Close()
	body, _ = io.ReadAll(resp.Body)
	fmt.Println("/training/actions response:")
	fmt.Println(string(body))
}

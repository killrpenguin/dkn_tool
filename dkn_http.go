package dkn

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type links struct {
	id   int
	name string
	link string
}


func DnRequest(link string) int {
	fmt.Println("DnRequest saying hello!")
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println(err)
	}
	body, err := io.ReadAll(resp.Body)
	log.Print(body)
	defer resp.Body.Close()
	return resp.StatusCode
}

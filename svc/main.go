package main

import (
	"awesomeProject/handler"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":3000", handler.MakeWebHandler())
	if err != nil {
		return
	}
}

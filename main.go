package main

import (
	"fmt"
	"math"
	"net/http"
)

var (
	Port     string
	Endpoint string
)

func main() {
	Port = ":8080"
	Endpoint = "/calc"

	http.HandleFunc(Endpoint, MyMagicHandler)

	fmt.Println("Приложение запускается на порту", Port)
	fmt.Printf("Выполни запрос на http://127.0.0.1%s%s\n", Port, Endpoint)

	if err := http.ListenAndServe(Port, nil); err != nil {
		fmt.Println("ОШИБОЧКА!", err.Error())
	}

}

func MyMagicHandler(w http.ResponseWriter, r *http.Request) {
	var A, B float64
	A = 207071570.0
	B = 763791297.0
	C := math.Sqrt(A * B)

	result := fmt.Sprintf("{'result' : %.3f}", C)

	w.Header().Set("Content-Type", "application/json")

	fmt.Printf("result is %.3f for %s\n", C, r.RemoteAddr)

	_, _ = w.Write([]byte(result))

}

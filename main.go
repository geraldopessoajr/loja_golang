package main

import (
	"net/http"

	"loja_camisetas/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8080", nil)
}

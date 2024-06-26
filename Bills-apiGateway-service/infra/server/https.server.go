package server

import (
	"log"
	"net/http"

	"github.com/skGab/Bills-microservices/Bills-apigateway-service/app/controllers"
)

func RunHttps(adress string, gateway *controllers.GateWayPipe) {
	http.Handle("/gateway/:routingKey", gateway)

	log.Fatal(http.ListenAndServe(adress, nil))
}

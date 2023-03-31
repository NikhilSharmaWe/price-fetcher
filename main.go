package main

import (
	"flag"
)

func main() {
	// WE GENERALLY WRITE A CLIENT FOR TESTING EACH MICROSERVICES
	// client := client.New("http://localhost:3000")
	// price, err := client.FetchPrice(context.Background(), "GG")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%+v\n", price)
	// return
	listenAddr := flag.String("listenaddr", ":3000", "listen address the service is running")
	flag.Parse()

	svc := NewLoggingService(NewMetricService(&priceFetcher{})) // this is an example of decorator pattern in go

	server := NewJSONAPIServer(*listenAddr, svc)

	server.Run()
}

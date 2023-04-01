package main

import (
	"flag"
)

// MICROSERVICES ARE VERY VERY USEFUL
// FOR EXAMPLE, WE FIRST IMPLEMENTED THE JSON API METHOD FOR CALLING THE SERVICE
// BUT WHEN WE IMPLEMENTED THE GRPC WE DID NOT EVEN LAID HAND ON THE LOGIC FOR LOGGING
// , BUSINESS LOGIC OR ANYTHINGY

// GRPC MAKES SENSE BETWEEN SERVERS, NOT SERVER AND BROWSER
func main() {
	// WE GENERALLY WRITE A CLIENT FOR TESTING EACH MICROSERVICES
	// client := client.New("http://localhost:3000")
	// price, err := client.FetchPrice(context.Background(), "GG")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%+v\n", price)
	// return
	var (
		jsonAddr = flag.String("json", ":3000", "listen address of the json transport")
		grpcAddr = flag.String("grpc", ":4000", "listen address of the grpc transport")
		svc      = NewLoggingService(&priceFetcher{})
		// or
		// svc = loggingService{&priceFetcher{}}
		// ctx = context.Background()
	)

	flag.Parse()
	// HERE IS WE ARE ALSO RUNNING A GRPC CLIENT FOR TESTING THE GRPC SERVER, SINCE WE CANNOT ACCESS IT
	// ON A BROWSER LIKE A JSON API. WE NEED A GRPC CLIENT FOR ACCESSING IT. SO, NOW OTHER SERVICES CAN ACESS THE
	// PRICE FETCHER THROUGH ANOTHER CONTAINER RUNNING A CLIENT FOR ACCESSING THIS.

	// grpcClient, err := client.NewGRPCNewClient(*grpcAddr)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// go func() {
	// 	time.Sleep(3 * time.Second)
	// 	resp, err := grpcClient.FetchPrice(ctx, &proto.PriceRequest{Ticker: "BTC"})
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf("%+v\n", resp)
	// }()

	go makeGRPCServerAndRun(*grpcAddr, svc)
	jsonServer := NewJSONAPIServer(*jsonAddr, svc)
	jsonServer.Run()

}

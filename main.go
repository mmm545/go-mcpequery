package main

import (
	"flag"
	"os"
)

func main()  {
	ip := flag.String("p", "", "IP address to query. Example usage: 127.0.0.1:19132")
	queryType := flag.String("t", "raknet", "Type of query. Options: raknet, gs4")
	flag.Parse()

	if *ip == ""{
		println("An IP address must be specified")
		os.Exit(1)
	}

	switch *queryType{
	case "raknet":
		if err := raknetPing(*ip); err != nil{
			panic(err)
		}

	case "gs4":
		if err := gamespyQuery(*ip); err != nil{
			panic(err)
		}
	default:
		println("Invalid query type. Options: raknet, gs4")
		os.Exit(2)
	}

}

func raknetPing(ip string) error{

}

func gamespyQuery(ip string) error{

}

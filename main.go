package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

var magic = []byte{0x00, 0xFF, 0xFF, 0x00, 0xFE, 0xFE, 0xFE, 0xFE, 0xFD, 0xFD, 0xFD, 0xFD, 0x12, 0x34, 0x56, 0x78}

func main()  {
	ip := flag.String("p", "", "IP address to query. Example usage: 127.0.0.1:19132")
	queryType := flag.String("t", "raknet", "Type of query. Options: raknet, gs4")
	flag.Parse()

	if *ip == ""{
		println("An IP address must be specified")
		os.Exit(1)
	}

	var data map[string]string
	var err error

	switch *queryType{
	case "raknet":
		err, data = raknetPing(*ip)

	case "gs4":
		err, data = gamespyQuery(*ip)

	default:
		println("Invalid query type. Options: raknet, gs4")
		os.Exit(2)
	}

	if err != nil{
		panic(err)
	}

	for k, v := range data{
		fmt.Printf("%v: %v\n", k, v)
	}

}

func raknetPing(ip string) (error, map[string]string){
	addr, err := net.ResolveUDPAddr("udp", ip)
	if err != nil{
		return err, nil
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil{
		return err, nil
	}

	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, byte(0x01))
	binary.Write(&buf, binary.BigEndian, time.Now().Unix())
	binary.Write(&buf, binary.BigEndian, magic)
	binary.Write(&buf, binary.BigEndian, []byte{0, 0, 0, 0, 0, 0, 0, 0}) // there should be a client guid here
	if _, err = conn.Write(buf.Bytes()); err != nil{
		return err, nil
	}

	b := make([]byte, 1024)
	if _, err := conn.Read(b); err != nil{
		return err, nil
	}

	data := strings.Split(string(b[35:]), ";")
	return nil, map[string]string{
		"Game name": data[0],
		"MOTD": data[1],
		"Protocol version": data[2],
		"Game version": data[3],
		"Number of online players": data[4],
		"Max number of players": data[5],
		"Server GUID": data[6],
		"Default world name": data[7],
		"Game mode": data[8],
	}
}

func gamespyQuery(ip string) (error, map[string]string){
	return nil, map[string]string{}
}

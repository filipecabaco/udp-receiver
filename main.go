// MouseCam project main.go
package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
)

func cam(wg *sync.WaitGroup) {
	addr := net.UDPAddr{
		Port: 1234,
		IP:   net.ParseIP("127.0.0.1"),
	}
	fmt.Println(addr)
	sock, err := net.ListenUDP("udp", &addr)
	handleError(err)
	buf := make([]byte, 1500)
	for {
		l, err := sock.Read(buf)
		handleError(err)
		fmt.Println("Reading " + strconv.Itoa(l) + " bytes")
		fmt.Println(buf[0:l])
	}
	wg.Done()
	return
}
func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
func main() {
	var wg sync.WaitGroup
	fmt.Println("Start Cam Routine")

	wg.Add(1)
	go cam(&wg)
	wg.Wait()
	fmt.Println("Done")
}

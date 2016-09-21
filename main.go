package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

/* A Simple function to verify error */
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func main() {

	port := flag.Int("port", 40008, "an int")
	flag.Parse()
	portstr := ":" + fmt.Sprintf("%d", *port)
	fmt.Println("Port: ", portstr)

	/* Lets prepare a address at any address at port 10001*/
	ServerAddr, err := net.ResolveUDPAddr("udp", portstr)
	CheckError(err)

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer ServerConn.Close()

	buf := make([]byte, 1024)

	for {
		// n, addr, err := ServerConn.ReadFromUDP(buf)
		n, _, err := ServerConn.ReadFromUDP(buf)

		// fmt.Println("Received ", string(buf[0:n]), " from ", addr)
		fmt.Printf("\n%v %x %v %x\n", buf[0:5], buf[5:6], buf[6:n-1], buf[n-1:n])
		fmt.Printf("%x %x %x %x\n", buf[0:5], buf[5:6], buf[6:n-1], buf[n-1:n])

		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}

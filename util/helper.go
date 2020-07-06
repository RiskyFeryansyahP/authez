package util

import (
	"log"
	"net"
)

func GetOutboundIP() net.IP {
	// Dial protocol UDP (udp is unlike TCP, udp does not have handshake nor connection)
	// Dial using udp only for get up the local ip address, address can change whatever (exixted or not existed)
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Printf("dial error %+v \n", err)
	}
	defer conn.Close()

	localAddress := conn.LocalAddr().(*net.UDPAddr)

	return localAddress.IP
}
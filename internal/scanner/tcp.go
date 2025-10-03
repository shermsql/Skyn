package scanner

import (
	"fmt"
	"net"
	"sort"
	"time"
)

type Status struct {
	Port    int
	Service string
	Open    bool
}

var Ports = map[int]string {
	20: "FTP Data (File Transfer Protocol Data)",
	21: "FTP (File Transfer Protocol)",
	22: "SSH (Secure Shell)",
	23: "Telnet (Telecommunication Network)",
	25: "SMTP (Simple Mail Transfer Protocol)",
	53: "DNS (Domain Name System)",
	80: "HTTP (Hypertext Transfer Protocol)",
	110: "POP3 (Post Office Protocol 3)",
	139: "NetBIOS (Network Basic Input/Output System)",
	143: "IMAP (Internet Message Access Protocol)",
	443: "HTTPS (Hyper Text Transfer Protocol Secure)",
	445: "Microsoft DS (Microsoft Directory Service)",
	3389: "RDP (Remote Desktop Protocol)",
}

func Scan(host string) []Status {
	var result []Status
	ports := make([]int, 0, len(Ports))

	for port := range Ports {
		ports = append(ports, port)
	}
	sort.Ints(ports)

	for _, port := range ports {
		address := net.JoinHostPort(host, fmt.Sprintf("%d", port))
		connection, err := net.DialTimeout("tcp", address, 1* time.Second)
		status := Status{ Port: port, Service: Ports[port], Open: err == nil }
		result = append(result, status)

		if err == nil {
			connection.Close()
		}
	}
	return result
}
package internal

import (
	"fmt"
	"net"
	"time"
)

func ScanPort(host string, port int, timeout time.Duration) error {
	target := fmt.Sprintf("%s:%d", host, port)

	conn, err := net.DialTimeout("tcp", target, timeout)
	if err != nil {
		return err
	}

	defer conn.Close()
	return nil
}
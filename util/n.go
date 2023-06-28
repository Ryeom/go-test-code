package util

import (
	"fmt"
	"net"
	"time"
)

func CheckPort(host, port string)  {
	timeout := time.Second

	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if err != nil {
		fmt.Printf("방화벽 체크 실패: %v\n", err)
		return
	}
	defer conn.Close()

	fmt.Printf("방화벽 체크 성공: 호스트 %s의 포트 %s에 접근 가능\n", host, port)

}

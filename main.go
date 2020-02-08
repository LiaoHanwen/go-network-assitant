package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("choose to use TCP ('1') / UDP ('2'): ")

	for true {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		// trim \r & \n
		text = strings.TrimRight(strings.TrimRight(text, "\n"), "\r")

		if strings.Compare(text, "1") == 0 || strings.Compare(text, "tcp") == 0 || strings.Compare(text, "TCP") == 0 {
			tcp()
			break
		} else if strings.Compare(text, "2") == 0 || strings.Compare(text, "udp") == 0 || strings.Compare(text, "UDP") == 0 {

			break
		} else {
			fmt.Print("input '1' or '2': ")
		}
	}

	fmt.Println("Bye! ヾ(°∇°*)")
}

func tcp() {
	fmt.Println("choose to use TCP Server ('1') / TCP Circular Send Client ('2') / TCP Input Send Client ('3'): ")

	for true {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		// trim \r & \n
		text = strings.TrimRight(strings.TrimRight(text, "\n"), "\r")

		if strings.Compare(text, "1") == 0 {
			tcpServer()
			break
		} else if strings.Compare(text, "2") == 0 {
			tcpCircuar()
			break
		} else if strings.Compare(text, "3") == 0 {
			tcpInput()
			break
		} else {
			fmt.Print("input '1', '2' or '3': ")
		}
	}
}

package main

import (
	"al-horezmi/gaming"
	"al-horezmi/utilities"
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {
	//fmt.Println("CONNECTION {\"LOGIN\":\"Max\"}")
	//time.Sleep(1000 * time.Millisecond)
	//if n, err := conn.Write([]byte("CONNECTION {\"LOGIN\":\"Max\"}\n"));
	//	n == 0 || err != nil {
	//	fmt.Println(err)
	//	return
	//} else {
	//	fmt.Println("dfngkjdng", n)
	//}
	//fmt.Print("Ответ:")
	//buff := make([]byte, 1024)
	//n, err := conn.Read(buff)
	//if err !=nil{
	//	return
	//}
	//fmt.Print(string(buff[0:n]))
	//fmt.Println()
	//var field = utilities.NewField(
	//	5,
	//	6,
	//	utilities.Position{0,2},
	//	utilities.Position{5,4},
	//	[][4][2]uint8{{utilities.Position{1, 2}, utilities.Position{1, 3}, utilities.Position{2, 2}, utilities.Position{2, 3}}})
	//str, err := json.Marshal(field)
	//if err != nil {
	//	fmt.Println("error:", err)
	//}
	//os.Stdout.Write(str)
	conn, err := net.Dial("tcp4", os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Write([]byte("DISCONNECT {\"QUIT\":\"\"}"))
	defer conn.Close()
	gamesToPlay, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	ch := make(chan string)
	go func() {
		buff := make([]byte, 1024)
		n, err:= conn.Read(buff)
		if n==0 || err != nil{
			fmt.Println("Error reading login response")
		}else {
			fmt.Println("Received login response")
		}
		buff = bytes.Trim(buff, "\x00")
		ch <- string(buff)
	}()
	_, _ = conn.Write([]byte("CONNECTION {\"LOGIN\":\""+ os.Args[3] +"\"}\n"))
	var response = <-ch
	var message utilities.Message
	err = json.Unmarshal([]byte(response), &message)
	if err != nil {
		fmt.Println("Error:", err)
	}
	switch message.Msg {
	case "LOGIN OK":
		_, _ = fmt.Println("Login successful")
	case "LOGIN FAILED":
		_, _ = fmt.Println("Login failed")
		return
	default:
		fmt.Println("[DEBUG] Cannot resolve login message:" + message.Msg)
	}
	for i := 0; i < gamesToPlay; i++ {
		var player = gaming.Player{
			Conn:     	conn,
			Field:    	utilities.Field{},
			Com:      	gaming.Communicator{},
			StringCh: 	nil,
			LobbyInfo: 	utilities.LobbyInfo{},
		}
		fmt.Println(player.PlayGame())
	}
}

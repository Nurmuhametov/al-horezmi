package gaming

import (
	"al-horezmi/utilities"
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"regexp"
)

type Player struct {
	Conn      net.Conn
	Field     utilities.Field
	Com       Communicator
	StringCh  chan string
	LobbyInfo utilities.LobbyInfo
}

func (p Player) PlayGame() string {
	ch := make(chan string)
	go func() {
		fmt.Println("Waiting for JOINLOBBY response")
		buff := make([]byte, 1024)
		n, err:= p.Conn.Read(buff)
		if n==0 || err != nil{
			fmt.Println("Error joining lobby")
		}
		buff = bytes.Trim(buff, "\x00")
		ch <- string(buff)
	}()
	_, _ = p.Conn.Write([]byte("SOCKET JOINLOBBY {\"id\":\"1\"}\n"))
	var response = <-ch
	fmt.Println("Received JOINLOBBY response")
	var joinResponse utilities.JoinResponse
	_ = json.Unmarshal([]byte(response), &joinResponse)
	if joinResponse.Success {
		p.LobbyInfo = joinResponse.Data
		fmt.Println("Joined lobby successfully")
	} else {
		fmt.Println("Cannot join lobby")
		return ""
	}

	p.StringCh = make(chan string)
	p.Com = Communicator{
		active:                false,
		conn:                  p.Conn,
		DataReceivedListeners: []func(string2 string){p.dataReceived},
	}
	p.Com.Start()
	defer p.Com.Stop()
	return "You" + <-p.StringCh + "in lobby " + p.LobbyInfo.Name
}

func (p Player) makeTurn() {

}

func (p Player) dataReceived(data string) {
	re := regexp.MustCompile("[A-Z ]+[A-Z]|(?:{.+})")
	split := re.FindAllString(data, 2)
	if len(split) == 2{
		fmt.Println(split[0], split[1])
	}
	if len(split) > 0 {
		switch split[0] {
		case "SOCKET STARTGAME": p.startGame(split[1])
		case "SOCKET STEP": p.updateField(split[1]) //end also makes turn
		case "SOCKET ENDGAME": showResult(split[1], p.StringCh)
		}
	}
}

func (p Player) startGame(data string) {
	var sg utilities.StartResponse
	_ = json.Unmarshal([]byte(data), &sg)
	if !sg.Move {
		return
	}
	fmt.Println("Yeah, its my turn!")
	p.Field = utilities.Field{
		Width:            sg.Width,
		Height:           sg.Height,
		Position:         sg.Position,
		OpponentPosition: sg.OpponentPosition,
		Barriers:         sg.Barriers,
	}
	p.makeTurn()
}

func (p Player) updateField(data string) {
	var f utilities.Field
	_ = json.Unmarshal([]byte(data), &f)
	p.Field = f
	p.makeTurn()
}

func showResult(data string, sCh chan string) {
	var f utilities.Field
	_ = json.Unmarshal([]byte(data), &f)
	sCh<- " won "
}
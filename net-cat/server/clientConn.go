package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	. "net-cat/models"
	"time"
)

func ClientConn(conn net.Conn) {
	defer conn.Close()
	conn.Write([]byte(welcomeText))
	conn.Write([]byte(logo))
	name, err := getClientName(conn)
	if err != nil {
		log.Printf("Error while parsing name: %v", err)
		closeConn()
		return
	}

	client := Client{Name: name}

	mut.Lock()
	clients[name] = &conn
	joining <- name
	mut.Unlock()

	conn.Write([]byte(history))

	writeMessage(conn, name)
	sc := bufio.NewScanner(conn)
	for sc.Scan() {
		msg := sc.Text()
		if len(msg) == 0 {
			writeMessage(conn, name)
			continue
		}
		if len(msg) >= 100 {
			conn.Write([]byte(maxMessageLength))
			writeMessage(conn, name)
			continue
		}
		if !IsinAscii(msg) {
			conn.Write([]byte(onlyAscii))
		}
		message := Message{
			Sender:  &client,
			Time:    time.Now().Format(timeFormat),
			Content: msg,
		}
		mut.Lock()
		messages <- message
		mut.Unlock()
		writeMessage(conn, name)
	}
	closeConn()
	mut.Lock()
	lefting <- name
	delete(clients, name)
	mut.Unlock()
}

func closeConn() {
	mut.Lock()
	numConnections--
	mut.Unlock()
}

func writeMessage(conn net.Conn, name string) {
	conn.Write([]byte(fmt.Sprintf("[%v][%v]:", time.Now().Format(timeFormat), name)))
}

func getClientName(conn net.Conn) (string, error) {
	conn.Write([]byte(askName))
	clientName := make([]byte, 2048)
	n, err := conn.Read(clientName)
	if err != nil {
		return "", err
	}
	name := string(clientName[:n-1])
L:
	for {
		switch {
		case n <= 2 || n > 16:
			conn.Write([]byte(nameLengthUsage))
		case !IsinAscii(name):
			conn.Write([]byte(onlyAscii))
		case IsClientExist(name):
			conn.Write([]byte(takenName))
		default:
			break L
		}
		conn.Write([]byte(askName))
		n, err = conn.Read(clientName)
		if err != nil {
			return "", err
		}
		name = string(clientName[:n-1])
	}
	return name, nil
}

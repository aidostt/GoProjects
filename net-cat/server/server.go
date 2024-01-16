package server

import (
	"log"
	"net"
	. "net-cat/models"
	"net-cat/pkg/files"
	"sync"
)

const (
	logoPath              = "logo.txt"
	defaultLogo           = "¯\\_(ツ)_/¯\n"
	welcomeText           = "Welcome to TCP-Chat!\n"
	askName               = "[ENTER YOUR NAME]:"
	timeFormat            = "2006-01-02 15:04:05"
	maxConnectionsMessage = "Unable to connect tcp chat. Connection limit reached"
	onlyAscii             = "only characters from ascii are allowed\n"
	nameLengthUsage       = "length of name should be more that 2 and less than 16\n"
	takenName             = "This name is already taken, please choose another one\n"
	maxMessageLength      = "message cannot be more than 100 characters\n"
)

var (
	clients        map[string]*net.Conn
	messages       chan Message
	lefting        chan string
	joining        chan string
	mut            sync.Mutex
	port           string
	logo           string
	numConnections int
	history        string
)

type Server struct {
	Addr           string
	MaxConnections int
}

func New(addr string, maxConnections int) *Server {
	initVars()
	return &Server{Addr: addr, MaxConnections: maxConnections}
}

func initVars() {
	messages = make(chan Message)
	joining = make(chan string)
	lefting = make(chan string)
	port = "8989"
	numConnections = 0
	var err error
	logo, err = files.Contents(logoPath)
	if err != nil {
		logo = defaultLogo
	}
	clients = make(map[string]*net.Conn)
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", s.Addr)
	if err != nil {
		log.Print("Cannot start the server: %v", err.Error())
		return
	}
	defer listener.Close()
	log.Printf("Listening on the port %v", s.Addr)

	go Broadcast()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Unable to accept connection: %v", err.Error())
			conn.Close()
			continue
		}
		mut.Lock()
		if numConnections >= s.MaxConnections {
			conn.Write([]byte(maxConnectionsMessage))
			conn.Close()
			mut.Unlock()
			continue
		}
		numConnections++
		mut.Unlock()
		go ClientConn(conn)
	}
}

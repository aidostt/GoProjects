## NET-CAT
This project aims to recreate a simplified version of NetCat in a Server-Client architecture using the Go programming language. The server operates in server mode, listening for incoming connections on a specified port, while the clients can connect to the server, join a group chat, and exchange messages.

## Objectives
The project objectives are as follows:

Establish a TCP connection between the server and multiple clients (1 to many).
Require clients to provide a unique username.
Implement control over the maximum number of connections (up to 10).
Allow clients to send messages to the chat.
Prevent the broadcast of empty messages from clients.
Identify messages with a timestamp and the username of the sender, e.g., [2020-01-20 15:48:41][client.name]:[client.message].
When a client joins the chat, provide them with previous chat history.
Inform all clients when a new client joins the chat.
Notify all clients when a client leaves the chat.
Ensure that all clients receive messages sent by other clients.
If a client leaves the chat, do not disconnect the remaining clients.
Use port 8989 as the default if no port is specified. Display a usage message in case of incorrect usage.
Implementation Details
The project is written in the Go programming language.
It employs Go-routines to handle concurrent connections.
Channels or Mutexes are used to synchronize and manage data between the server and clients.
A maximum of 10 connections is supported.
Error handling is implemented for server and client-side operations.
Usage
To run the TCP Chat server, use the following command:

$ go run ./cmd/
Listening on the port :8989
$ go run./cmd/ 2525
Listening on the port :2525
$ go run ./cmd/ 2525 localhost
[USAGE]: ./TCPChat $port
$
$ nc $IP $port
Welcome to TCP-Chat!
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    `.       | `' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     `-'       `--'
[ENTER YOUR NAME]:

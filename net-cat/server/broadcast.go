package server

func Broadcast() {
	for {
		select {
		case client := <-joining:
			announce(client, "\n"+client+" has joined our chat...", "\033[0;32m")
		case client := <-lefting:
			announce(client, "\n"+client+" has left our chat...", "\033[0;31m")
		case msg := <-messages:

			announce(msg.Sender.Name, "\n"+msg.ToString(), "\033[0m")
			history = history + msg.ToString() + "\n"
		}
	}
}

func announce(sender string, content string, color string) {
	mut.Lock()
	for k, v := range clients {
		if k == sender {
			continue
		}
		(*v).Write([]byte(color + content + "\033[0m\n"))
		writeMessage(*v, k)
	}
	mut.Unlock()
}

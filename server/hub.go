package chat

var DefaultHub = NewHub()

type Hub struct {
	Join  chan *Conn
	Conns map[*Conn]bool
	Echo  chan string
}

func NewHub() *Hub {
	return &Hub{
		Join:  make(chan *Conn),
		Conns: make(map[*Conn]bool),
		Echo:  make(chan string)}
}

func (h *Hub) Start() {
	for {
		select {
		case conn := <-h.Join:
			h.Conns[conn] = true
		case msg := <-h.Echo:
			for conn := range h.Conns {
				conn.Send <- msg
			}
		}
	}
}

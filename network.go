package innernet

import (
	"net"
	"sync"
)

var inet = newNetwork()

type network struct {
	sync.Mutex
	listeners map[string]*listener
}

func newNetwork() *network {
	return &network{
		listeners: make(map[string]*listener),
	}
}

func (n *network) OpenListener(addr string) (*listener, error) {
	n.Lock()
	defer n.Unlock()

	if _, ok := n.listeners[addr]; ok {
		return nil, ErrorAlreadyInUse
	}

	l := newListener(addr)
	n.listeners[addr] = l
	return l, nil
}

func (n *network) CloseListener(l *listener) {
	n.Lock()
	defer n.Unlock()

	delete(n.listeners, l.addr)
	close(l.newConns)
}

func (n *network) Dial(addr string) (net.Conn, error) {
	n.Lock()
	defer n.Unlock()

	if l, ok := n.listeners[addr]; ok {
		c1, c2 := net.Pipe()
		l.newConns <- c2
		return c1, nil
	}

	return nil, ErrorNotListening
}

package innernet

import "net"

type listener struct {
	addr     string
	newConns chan net.Conn
}

func newListener(addr string) *listener {
	return &listener{
		addr:     addr,
		newConns: make(chan net.Conn),
	}
}

func (l *listener) Accept() (net.Conn, error) {
	if c, ok := <-l.newConns; ok {
		return c, nil
	}
	return nil, ErrorListenerClosed
}

func (l *listener) Close() error {
	inet.CloseListener(l)
	return nil
}

func (l *listener) Addr() net.Addr { return &addr{l.addr} }

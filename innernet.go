/*
Package 'innernet' (do not confuse with inTernet!) is a tiny package for creating in-memory network connections
inside the application, using the standard interfaces from the 'net' package.
*/
package innernet

import (
	"errors"
	"net"
)

const NetworkName = "innernet"

var (
	ErrorAlreadyInUse   = errors.New("Address already in use")
	ErrorListenerClosed = errors.New("Listener was closed")
	ErrorNotListening   = errors.New("Address not listening")
)

// Listen creates listening socket with address 'addr'. The 'addr' can be any string.
func Listen(addr string) (net.Listener, error) { return inet.OpenListener(addr) }

// Dial connects to the listening socket with address 'addr'.
func Dial(addr string) (net.Conn, error) { return inet.Dial(addr) }

package innernet_test

import (
	"io"
	"io/ioutil"
	"sync"
	"testing"

	"github.com/davidmz/innernet"
)

const testAddr = "abc"

func TestListenerCreation(t *testing.T) {
	l, err := innernet.Listen(testAddr)
	ok(t, err)
	equals(t, innernet.NetworkName, l.Addr().Network())
	equals(t, testAddr, l.Addr().String())

	_, err = innernet.Listen(testAddr)
	equals(t, innernet.ErrorAlreadyInUse, err)

	ok(t, l.Close())
}

func TestDialAndListen(t *testing.T) {
	_, err := innernet.Dial(testAddr)
	equals(t, innernet.ErrorNotListening, err)

	l, err := innernet.Listen(testAddr)
	ok(t, err)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		conn, err := l.Accept()
		ok(t, err)
		conn.Close()
	}()
	go func() {
		defer wg.Done()
		conn, err := innernet.Dial(testAddr)
		ok(t, err)
		conn.Close()
	}()
	wg.Wait()

	ok(t, l.Close())
	_, err = l.Accept()
	equals(t, innernet.ErrorListenerClosed, err)
}

func TestDataTransfer(t *testing.T) {
	l, err := innernet.Listen(testAddr)
	ok(t, err)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		conn, err := l.Accept()
		ok(t, err)
		_, err = io.WriteString(conn, "test data")
		ok(t, err)
		conn.Close()
	}()
	go func() {
		defer wg.Done()
		conn, err := innernet.Dial(testAddr)
		ok(t, err)
		b, err := ioutil.ReadAll(conn)
		ok(t, err)
		equals(t, "test data", string(b))
		conn.Close()
	}()
	wg.Wait()
}

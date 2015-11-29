package stopper

import (
	"errors"
	"net"
	"net/http"
	"time"
)

type StoppableListener struct {
	*net.TCPListener
	stop chan int
}

func New(l net.Listener) (*StoppableListener, error) {
	tcpL, ok := l.(*net.TCPListener)

	if !ok {
		return nil, errors.New("Cannot wrap TCP Listener")
	}

	retval := &StoppableListener{}
	retval.TCPListener = tcpL
	retval.stop = make(chan int)

	return retval, nil
}

func (sl *StoppableListener) Accept() (net.Conn, error) {
	for {
		sl.SetDeadline(time.Now().Add(time.Second))

		newConn, err := sl.TCPListener.Accept()

		select {
		case <-sl.stop:
			return nil, errors.New("StoppedError")
		default:

		}

		if err != nil {
			netErr, ok := err.(net.Error)

			if ok && netErr.Timeout() && netErr.Temporary() {
				continue
			}
		}

		return newConn, err
	}
}

func (sl *StoppableListener) Stop() {
	close(sl.stop)
}

type StoppableServer struct {
	*http.Server
}

func (srv *StoppableServer) ListenAndServe() error {
	addr := srv.Addr
	if addr == "" {
		addr = ":http"
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return srv.Serve(tcpKeepAliveListener{ln.(*StoppableListener)})
}

type tcpKeepAliveListener struct {
	*StoppableListener
}

func (ln tcpKeepAliveListener) Accept() (c net.Conn, err error) {
	tc, err := ln.AcceptTCP()
	if err != nil {
		return
	}
	tc.SetKeepAlive(true)
	tc.SetKeepAlivePeriod(3 * time.Minute)
	return tc, nil
}

func StoppableListenAndServe(addr string, handler http.Handler) error {
	server := &StoppableServer{&http.Server{Addr: addr, Handler: handler}}
	return server.ListenAndServe()
}

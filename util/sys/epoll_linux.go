package sys

import (
	"golang.org/x/sys/unix"
	"log"
	"net"
	"reflect"
	"sync"
	"syscall"
)

// Linux epoll event receiver
type epoll struct {
	fd    int              // file descriptor
	conns map[int]net.Conn // connections map
	lock  *sync.RWMutex    // read-write lock
}

func Epoll() (*epoll, error) {
	fd, err := unix.EpollCreate1(0)
	if err != nil {
		return nil, err
	}
	return &epoll{
		fd:    fd,
		conns: make(map[int]net.Conn),
		lock:  &sync.RWMutex{},
	}, nil
}

// Wait for events on connections
func (e *epoll) Wait() ([]net.Conn, error) {
	events := make([]unix.EpollEvent, 100)
	n, err := unix.EpollWait(e.fd, events, 100)
	if err != nil {
		return nil, err
	}
	e.lock.RLock()
	defer e.lock.RUnlock()
	var conns []net.Conn
	for i := 0; i < n; i++ {
		conn := e.conns[int(events[i].Fd)]
		conns = append(conns, conn)
	}
	return conns, nil
}

// Receive a new event from queue
func (e *epoll) Add(conn net.Conn) error {
	// Extract file descriptor associated with the connection
	fd := socketFD(conn)
	err := unix.EpollCtl(e.fd, syscall.EPOLL_CTL_ADD, fd, &unix.EpollEvent{Events: unix.POLLIN | unix.POLLHUP, Fd: int32(fd)})
	if err != nil {
		return err
	}
	e.lock.Lock()
	defer e.lock.Unlock()
	e.conns[fd] = conn
	if len(e.conns)%100 == 0 {
		log.Printf("total number of connections: %v", len(e.conns))
	}
	return nil
}

// Remove a processed event from queue
func (e *epoll) Remove(conn net.Conn) error {
	fd := socketFD(conn)
	err := unix.EpollCtl(e.fd, syscall.EPOLL_CTL_DEL, fd, nil)
	if err != nil {
		return err
	}
	e.lock.Lock()
	defer e.lock.Unlock()
	delete(e.conns, fd)
	if len(e.conns)%100 == 0 {
		log.Printf("total number of connections: %v", len(e.conns))
	}
	return nil
}

func socketFD(conn net.Conn) int {
	//tls := reflect.TypeOf(conn.UnderlyingConn()) == reflect.TypeOf(&tls.Conn{})
	// Extract the file descriptor associated with the connection
	//connVal := reflect.Indirect(reflect.ValueOf(conn)).FieldByName("conn").Elem()
	tcpConn := reflect.Indirect(reflect.ValueOf(conn)).FieldByName("conn")
	//if tls {
	//	tcpConn = reflect.Indirect(tcpConn.Elem())
	//}
	fdVal := tcpConn.FieldByName("fd")
	pfdVal := reflect.Indirect(fdVal).FieldByName("pfd")

	return int(pfdVal.FieldByName("Sysfd").Int())
}

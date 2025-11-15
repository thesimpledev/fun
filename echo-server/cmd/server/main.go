package main

import (
	"bufio"
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const (
	readTimeout  = 1 * time.Second
	writeTimeout = 2 * time.Second
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()
	serverType := flag.String("s", "tcp", "server type")
	port := flag.String("p", "7", "server port")
	flag.Parse()

	switch *serverType {
	case "tcp":
		startTCPServer(ctx, *port)
	case "udp":
		startUDPServer(ctx, *port)
	default:
		log.Fatal("illegal server type")
	}
}

func startUDPServer(ctx context.Context, port string) {
	addr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		log.Fatal(err)
	}
	server, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()
	go func() {
		<-ctx.Done()
		server.Close()
	}()

	buffPool := sync.Pool{
		New: func() any {
			buf := make([]byte, 1024)
			return &buf
		},
	}
	for {
		bufPtr := buffPool.Get().(*[]byte)
		buf := *bufPtr
		n, clientAddr, err := server.ReadFromUDP(buf)
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				buffPool.Put(bufPtr)
				fmt.Println("UDP Server shutting down")
				return
			}
			fmt.Println(err)
			buffPool.Put(bufPtr)
			continue
		}

		data := make([]byte, n)
		copy(data, buf[:n])

		buffPool.Put(bufPtr)
		handleUDPPacket(clientAddr, data, server)
	}
}

func handleUDPPacket(addr *net.UDPAddr, data []byte, conn *net.UDPConn) {
	fmt.Printf("%s - %s", addr.AddrPort(), string(data))
	conn.WriteToUDP(data, addr)
}

func startTCPServer(ctx context.Context, port string) {
	addr, err := net.ResolveTCPAddr("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	server, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()

	var wg sync.WaitGroup
	var mu sync.Mutex
	activeConns := make(map[net.Conn]struct{})

	go func() {
		<-ctx.Done()
		server.Close()

		mu.Lock()
		for conn := range activeConns {
			conn.SetReadDeadline(time.Now().Add(readTimeout))
			conn.SetWriteDeadline(time.Now().Add(writeTimeout))
		}
		mu.Unlock()
	}()

	for {
		select {
		case <-ctx.Done():
			wg.Wait()
			return
		default:
			conn, err := server.Accept()
			if err != nil {
				fmt.Println(err)
				continue
			}

			mu.Lock()
			activeConns[conn] = struct{}{}
			mu.Unlock()
			wg.Go(func() {
				handleTCPConnection(conn, ctx, &mu, activeConns)
			})
		}
	}
}

func handleTCPConnection(conn net.Conn, ctx context.Context, mu *sync.Mutex, activeConns map[net.Conn]struct{}) {
	defer func() {
		conn.Close()
		mu.Lock()
		delete(activeConns, conn)
		mu.Unlock()
	}()
	fmt.Printf("Accepted connection from: %s\n", conn.RemoteAddr())

	scanner := bufio.NewReader(conn)
	for {

		b, err := scanner.ReadByte()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return
			}
			fmt.Println(err)
			continue
		}
		select {
		case <-ctx.Done():
			return
		default:
		}
		conn.Write([]byte{b})
	}
}

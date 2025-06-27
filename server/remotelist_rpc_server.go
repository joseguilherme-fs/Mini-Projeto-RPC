package main

import (
	"fmt"
	"ifpb/remotelist/pkg"
	"net"
	"net/rpc"
)

func main() {
	list := remotelist.NewRemoteList()
	rpcs := rpc.NewServer()
	rpcs.Register(list)

	l, err := net.Listen("tcp", ":5000")
	if err != nil {
		fmt.Println("Erro ao escutar:", err)
		return
	}
	defer l.Close()

	fmt.Println("Servidor RPC ouvindo na porta 5000...")
	for {
		conn, err := l.Accept()
		if err == nil {
			go rpcs.ServeConn(conn)
		}
	}
}

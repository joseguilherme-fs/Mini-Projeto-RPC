package main

import (
	"fmt"
	remotelist "ifpb/remotelist/pkg"
	"net/rpc"
	"time"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:5000")
	if err != nil {
		fmt.Println("Cliente B: Erro ao conectar:", err)
		return
	}
	defer client.Close()

	fmt.Println("--- Cliente B INICIOU ---")

	listaX := "lista-X"
	listaY := "lista-Y"

	for i := 0; i < 5; i++ {

		valAdicionar := 200 + i
		var ok bool
		client.Call("RemoteList.Append", remotelist.AppendArgs{ListID: listaY, Value: valAdicionar}, &ok)
		fmt.Printf("Cliente B: Adicionado '%d' na '%s'\n", valAdicionar, listaY)

		var valRemovido int
		err := client.Call("RemoteList.Remove", remotelist.RemoveArgs{ListID: listaX}, &valRemovido)
		if err == nil {
			fmt.Printf("Cliente B: Removido '%d' da '%s'\n", valRemovido, listaX)
		} else {
			fmt.Printf("Cliente B: Tentou remover da '%s', mas estava vazia.\n", listaX)
		}

		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("--- Cliente B TERMINOU ---")
}

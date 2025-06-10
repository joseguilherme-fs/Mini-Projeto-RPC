package main

import (
	"fmt"
	"ifpb/remotelist/pkg"
	"net/rpc"
	"time"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:5000")
	if err != nil {
		fmt.Println("Cliente A: Erro ao conectar:", err)
		return
	}
	defer client.Close()

	fmt.Println("--- Cliente A INICIOU ---")

	listaX := "lista-X"
	listaY := "lista-Y"


	for i := 0; i < 5; i++ {
		// 1. Adiciona na lista X
		valAdicionar := 100 + i
		var ok bool
		client.Call("RemoteList.Append", remotelist.AppendArgs{ListID: listaX, Value: valAdicionar}, &ok)
		fmt.Printf("Cliente A: Adicionado '%d' na '%s'\n", valAdicionar, listaX)


		var valRemovido int
		err := client.Call("RemoteList.Remove", remotelist.RemoveArgs{ListID: listaY}, &valRemovido)
		if err == nil {
			fmt.Printf("Cliente A: Removido '%d' da '%s'\n", valRemovido, listaY)
		} else {
			fmt.Printf("Cliente A: Tentou remover da '%s', mas estava vazia.\n", listaY)
		}
		
		time.Sleep(80 * time.Millisecond) 
	}
	fmt.Println("--- Cliente A TERMINOU ---")
}
package main

import (
	"fmt"
	"ifpb/remotelist/pkg"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:5000")
	if err != nil {
		fmt.Println("Erro ao conectar:", err)
		return
	}

	listID := "minhaLista"

	valores := []int{10, 20, 30, 40, 50}
	for _, v := range valores {
		args := remotelist.AppendArgs{ListID: listID, Value: v}
		var ok bool
		client.Call("RemoteList.Append", args, &ok)
	}

	getArgs := remotelist.GetArgs{ListID: listID, Index: 2}
	var resultado int
	err = client.Call("RemoteList.Get", getArgs, &resultado)
	if err == nil {
		fmt.Println("Valor na posição 2:", resultado)
	}

	var tamanho int
	client.Call("RemoteList.Size", remotelist.SizeArgs{ListID: listID}, &tamanho)
	fmt.Println("Tamanho atual:", tamanho)

	for i := 0; i < 2; i++ {
		var valorRemovido int
		err = client.Call("RemoteList.Remove", remotelist.RemoveArgs{ListID: listID}, &valorRemovido)
		if err == nil {
			fmt.Println("Removido:", valorRemovido)
		}
	}

	client.Call("RemoteList.Size", remotelist.SizeArgs{ListID: listID}, &tamanho)
	fmt.Println("Tamanho atual:", tamanho)

}

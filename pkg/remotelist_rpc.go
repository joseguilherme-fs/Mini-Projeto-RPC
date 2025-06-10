package remotelist

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"
)

type RemoteList struct {
	mu     sync.Mutex
	Listas map[string][]int
	Log    *os.File
}

type AppendArgs struct {
	ListID string
	Value  int
}

type GetArgs struct {
	ListID string
	Index  int
}

type SizeArgs struct {
	ListID string
}

type RemoveArgs struct {
	ListID string
}

func NewRemoteList() *RemoteList {
	rl := &RemoteList{Listas: make(map[string][]int)}

	var err error
	rl.Log, err = os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	fmt.Println("Restaurando estado")
	rl.loadStateFromDisk()

	go rl.snapshotter()

	return rl
}

func (rl *RemoteList) loadStateFromDisk() {
	snapshotFile, err := os.Open("snapshot.json")
	if err == nil {
		json.NewDecoder(snapshotFile).Decode(&rl.Listas)
		snapshotFile.Close()
		fmt.Println("Snapshot carregado com sucesso!")
	} else {
		fmt.Println("Nenhum snapshot encontrado, recuperando apenas pelo log.")
	}

	rl.replayLog()
}

func (rl *RemoteList) snapshotter() {

	ticker := time.NewTicker(15 * time.Second)

	for {

		<-ticker.C

		fmt.Println("Criando um snapshot")
		rl.createSnapshot()
	}
}

func (rl *RemoteList) createSnapshot() {

	rl.mu.Lock()

	defer rl.mu.Unlock()

	data, err := json.MarshalIndent(rl.Listas, "", "  ")
	if err != nil {
		fmt.Println("Erro ao criar snapshot:", err)
		return
	}

	err = os.WriteFile("snapshot.json", data, 0666)
	if err != nil {
		fmt.Println("Erro ao criar snapshot (escrita):", err)
		return
	}

	rl.Log.Close()
	rl.Log, err = os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}

	fmt.Println("Snapshot criado e log limpo com sucesso!")
}

func (rl *RemoteList) Append(args AppendArgs, reply *bool) error {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	rl.Listas[args.ListID] = append(rl.Listas[args.ListID], args.Value)
	_, err := rl.Log.WriteString(fmt.Sprintf("append %s %d\n", args.ListID, args.Value))
	*reply = (err == nil)
	return err
}

func (rl *RemoteList) Remove(args RemoveArgs, reply *int) error {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	list := rl.Listas[args.ListID]
	if len(list) == 0 {
		return errors.New("lista vazia")
	}
	*reply = list[len(list)-1]
	rl.Listas[args.ListID] = list[:len(list)-1]
	_, err := rl.Log.WriteString(fmt.Sprintf("remove %s\n", args.ListID))
	return err
}

func (rl *RemoteList) Get(args GetArgs, reply *int) error {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	list := rl.Listas[args.ListID]
	if args.Index < 0 || args.Index >= len(list) {
		return errors.New("índice inválido")
	}
	*reply = list[args.Index]
	return nil
}

func (rl *RemoteList) Size(args SizeArgs, reply *int) error {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	*reply = len(rl.Listas[args.ListID])
	return nil
}

func (rl *RemoteList) replayLog() {
	file, err := os.Open("log.txt")
	if err != nil {
		return
	}
	defer file.Close()

	var op, listID string
	var value int
	for {
		_, err := fmt.Fscanf(file, "%s %s", &op, &listID)
		if err != nil {
			break
		}

		if op == "append" {
			fmt.Fscanf(file, "%d", &value)
			rl.Listas[listID] = append(rl.Listas[listID], value)
		} else if op == "remove" {
			l := rl.Listas[listID]
			if len(l) > 0 {
				rl.Listas[listID] = l[:len(l)-1]
			}
		}
	}
}

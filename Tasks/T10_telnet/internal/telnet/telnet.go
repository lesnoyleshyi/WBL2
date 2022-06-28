package telnet

import (
	"WBL2/Tasks/T10_telnet/internal/key"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net"
	"os"
	"sync"
)

type Telnet struct {
	key  *key.Key
	conn net.Conn
	wg   *sync.WaitGroup
}

func New() *Telnet {
	return &Telnet{
		key: key.New(),
		wg:  &sync.WaitGroup{},
	}
}

func (t *Telnet) Start() {
	cmd := &cobra.Command{}

	t.key.SetKeys(cmd)

	if err := cmd.Execute(); err != nil || len(os.Args) < 3 {
		log.Fatal(fmt.Errorf("required argument missing: %v", err))
	}

	t.key.Host = os.Args[len(os.Args)-2]
	t.key.Port = os.Args[len(os.Args)-1]

	t.Connect()
}

func (t *Telnet) Connect() {
	var err error
	t.conn, err = net.DialTimeout("tcp", t.key.Host+":"+t.key.Port, t.key.Timeout)
	if err != nil {
		fmt.Println(t.key.Timeout)
		log.Fatal(fmt.Errorf("error when connect via TCP: %w", err))
	}

	t.wg.Add(1)
	go t.Send()
	t.wg.Add(1)
	go t.Receive()

	t.wg.Wait()
}

func (t *Telnet) Send() {
	defer t.wg.Done()
	_, err := io.Copy(t.conn, os.Stdin)
	if err != nil {
		t.wg.Done()
		log.Fatal(fmt.Errorf("error when send message: %v", err))
	}
	t.conn.Close()
	log.Fatal("EOF")
}

func (t *Telnet) Receive() {
	defer t.wg.Done()
	_, err := io.Copy(os.Stdout, t.conn)
	if err != nil {
		t.wg.Done()
		log.Fatal(fmt.Errorf("error when receive message: %v", err))
	}
	t.conn.Close()
	log.Fatal("Connection closed")
}

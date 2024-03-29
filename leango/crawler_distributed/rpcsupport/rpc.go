package rpcsupport

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func SaverRpc(host string, service interface{}) error {
	_ = rpc.Register(service)
	listener, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}
	log.Printf("Listening on %s", host)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error:%v", err)
		}

		go jsonrpc.ServeConn(conn)
	}

	return nil
}

func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	return jsonrpc.NewClient(conn), nil
}

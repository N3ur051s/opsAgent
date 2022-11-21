package api

import (
	"fmt"
	"net"

	. "opsAgent/conf"
)

func getAddressPort() (string, error) {
	address := Conf.Server.Addr
	return fmt.Sprintf("%v:%v", address, Conf.Server.Port), nil
}

func getListener() (net.Listener, error) {
	address, err := getAddressPort()
	if err != nil {
		return nil, err
	}
	return net.Listen("tcp", address)
}

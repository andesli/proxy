package controllers

import (
//	"flag"
	"net"
	"net/http"
	"time"
)

type Config struct {
	ConnectTimeout   time.Duration
	ReadWriteTimeout time.Duration
}

func TimeoutDialer(config *Config) func(net, addr string) (c net.Conn, err error) {
	return func(netw, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(netw, addr, config.ConnectTimeout)
		if err != nil {
			return nil, err
		}
		conn.SetDeadline(time.Now().Add(config.ReadWriteTimeout))
		return conn, nil
	}
}

//通过设置DisableKeepAlives为true ，避免客户端与服务器因为共享连接导致偶然超时问题
func NewTimeoutClient(conf *Config) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
		Dial: TimeoutDialer(conf),
		DisableKeepAlives: true,
		},
	}
}



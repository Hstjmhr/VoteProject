package remote

import (
	"fmt"
	"framework/game"
	"github.com/nats-io/nats.go"
	"msqp/logs"
)

type NatsClient struct {
	serverId string
	conn     *nats.Conn
	readChan chan []byte
}

func NewNatsClient(serverId string, readChan chan []byte) *NatsClient {
	return &NatsClient{
		serverId: serverId,
		readChan: readChan,
	}
}

func (c *NatsClient) Run() error {
	var err error
	c.conn, err = nats.Connect(game.Conf.ServersConf.Nats.Url)
	if err != nil {
		logs.Error("connect nats server fail,err:%v", err)
		return err
	}
	go c.sub()
	return nil
}

func (c *NatsClient) Close() error {
	if c.conn != nil {
		c.conn.Close()
	}
	return nil
}

func (c *NatsClient) sub() {
	_, err := c.conn.Subscribe(c.serverId, func(msg *nats.Msg) {
		//收到其他nats client传递过来的消息
		c.readChan <- msg.Data
	})
	if err != nil {
		logs.Error("nats sub err:%v", err)
	}
}

func (c *NatsClient) SendMsg(dst string, data []byte) error {
	fmt.Println("发送的nats消息", dst, string(data))
	if c.conn != nil {
		return c.conn.Publish(dst, data)
	}
	return nil
}

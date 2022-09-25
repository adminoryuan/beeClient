package main

import (
	"log"
	"net"
	"time"

	"github.com/adminoryuan/beeclient/common"
	"google.golang.org/protobuf/proto"
)

type Upload struct {
	bclo BeeCol
}

func (u *Upload) startUpload(c *Config) {

	con, err := net.Dial("tcp", c.Server.Addr)

	if err != nil {
		panic(err)
	}
	u.bclo = BeeCol{}
	log.Fatalln("tcp conn..")
	go u.upload(con, c.Tran.Interval)
}
func (u *Upload) upload(con net.Conn, interval int) {

	go func() {
		for {
			body := u.bclo.Collection()
			bytes, err := proto.Marshal(&body)
			if err != nil {
				continue
			}
			message := common.Message{
				MessageId: 1,
				Body:      bytes,
			}

			uplpadByte, err := proto.Marshal(&message)

			if err != nil {
				continue
			}

			log.Fatalln("send..")

			con.Write(uplpadByte)

			time.Sleep(time.Duration(interval) * time.Second)
		}
	}()
}

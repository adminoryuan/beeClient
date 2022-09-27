package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/adminoryuan/beeclient/common"
	"google.golang.org/protobuf/proto"
)

type Upload struct {
	bclo BeeCol
	Node string
	cnf  Config
}

func (u *Upload) startUpload(c *Config) {
	u.cnf = *c
	con, err := net.Dial("tcp", fmt.Sprintf("%s:%s", c.Server.Addr, c.Server.Port))

	if err != nil {

		panic(err)
	}
	u.bclo = BeeCol{}
	log.Print("tcp conn..")
	u.fristUpload(&con)
	//u.upload(con, *c)

}
func (u *Upload) fristUpload(con *net.Conn) {
	//首次上传
	c := ComputeInfo{}
	info, addr := c.getComputerinfo()
	body, err := proto.Marshal(&info)
	if err != nil {
		panic(err)
	}
	u.Node = addr
	newbody, err :=
		proto.Marshal(&common.Message{
			MessageId: 0,
			Apikey:    u.cnf.Server.Apikey,
			Body:      body,
			Node:      u.Node,
		})
	if err != nil {
		panic(err)
	}
	(*con).Write(decoder(newbody))

	log.Print(info)
}
func (u *Upload) upload(con net.Conn, cfg Config) {

	go func() {
		for {
			fmt.Println("start ..")
			body := u.bclo.Collection()

			bytes, err := proto.Marshal(&body)
			if err != nil {
				log.Print("marshal collection error..")

				continue
			}
			message := common.Message{
				MessageId: 1,
				Apikey:    cfg.Server.Apikey,
				Node:      u.Node,
				Body:      bytes,
			}

			uplpadByte, err := proto.Marshal(&message)

			if err != nil {
				log.Print("marshal collection message error..")
				continue
			}

			_, err = con.Write(decoder(uplpadByte))
			if err != nil {
				fmt.Println(err)
				panic(err)
			}
			log.Print("send ..", body)

			time.Sleep(time.Duration(cfg.Tran.Interval) * time.Second)
		}
	}()
}
func IntToBytes(n int) []byte {
	data := int32(n)
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes()
}
func decoder(meg []byte) []byte {
	body := make([]byte, 0)
	body = append(body, 0x00)
	body = append(body, IntToBytes(len(meg))...)
	body = append(body, meg...)

	return body
}

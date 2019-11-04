package echo

import (
	"encoding/binary"
	"errors"
	"io"
	"net"
	//. "variable"
	"github.com/gansidui/gotcp"
)

/*
func NewWallPacket(buff[] byte,hasLengthField bool,tyc int32) *EchoPacket {
	p := &EchoPacket{}
	if hasLengthField {
		p.buff = buf
	} else {
		p.buff = make([]byte, 4+len(buff))
		binary.BigEndian.PutUint16(p.buff[0:2],uint16(len(buff) + 2))
		binary.BigEndian.PutUint16(p.buff[2:4],uint16(tyc))
		copy(p.buff[4:],buff)
	}
}
*/
type EchoPacket struct {
	buff []byte
}

func (this *EchoPacket) Serialize() []byte {
	return this.buff
}

func (this *EchoPacket) GetLength() uint16 {
	return binary.LittleEndian.Uint16(this.buff[0:2])
}

func (this *EchoPacket) GetType() uint16 {
	return binary.LittleEndian.Uint16(this.buff[2:4])
}

func (this *EchoPacket) GetBody() []byte {
	return this.buff[4:]
}

func NewEchoPacket(buff []byte, hasLengthField bool, tyc int32) *EchoPacket {
	p := &EchoPacket{}

	if hasLengthField {
		p.buff = buff

	} else {
		p.buff = make([]byte, 4+len(buff))
		binary.LittleEndian.PutUint16(p.buff[0:2], uint16(len(buff)+2))
		binary.LittleEndian.PutUint16(p.buff[2:4], uint16(tyc))

		//typec := []byte{int16(tyc)}
		//copy(p.buff[2:4], typec)
		copy(p.buff[4:], buff)

	}

	return p
}

type EchoProtocol struct {
}

const maxMsgLength uint16 = 8

func (this *EchoProtocol) ReadPacket(conn *net.TCPConn) (gotcp.Packet, error) {
	println("weqwewqewqee")
	var (
		lengthBytes []byte = make([]byte, 2)
		length      uint16

		typeBytes []byte = make([]byte, 2)
		typec     uint16
	)

	// read length
	if _, err := io.ReadFull(conn, lengthBytes); err != nil {
		return nil, err
	}
	if length = binary.LittleEndian.Uint16(lengthBytes); length > (1024 * maxMsgLength) {
		return nil, errors.New("the size of packet is larger than the limit")
	}

	// read type
	if _, err := io.ReadFull(conn, typeBytes); err != nil {
		return nil, err
	}
	if typec = binary.LittleEndian.Uint16(typeBytes); typec > 65535 {
		return nil, errors.New("the command type is bigger")
	}

	if typec != 35 && length > (1024*maxMsgLength) {
		//log.Printf("ERROR, 消息太长了，断线:%d",typec,length)
		println(length, "aaaaaaaaaa")
		return nil, errors.New("the size of packet is larger than the limit")
	}

	buff := make([]byte, 2+length)
	copy(buff[0:2], lengthBytes)
	copy(buff[2:4], typeBytes)

	// read body ( buff = lengthBytes + body )
	if _, err := io.ReadFull(conn, buff[4:]); err != nil {
		return nil, err
	}

	return NewEchoPacket(buff, true, 0), nil
}

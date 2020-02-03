package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/olefen/horse"
	"github.com/olefen/note/config"
	"github.com/olefen/note/log"
	"github.com/olefen/note/test/protobuf"
	"io"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

type configMsg struct {
	config.Config `yaml:"-"`
	Host          string `json:"host" yaml:"host"`
	Port          string `json:"port" yaml:"port"` // port
	LeftInterval  int    // 左区间
	RightInterval int    // 右区间
	IsProduct     bool
	HWPort        string
}

var (
	LogConn  *log.Logger
	LogHW    *log.Logger
	LogCount *log.Logger
	c        = &configMsg{
		Host:          "47.56.161.105",
		Port:          "8989",
		LeftInterval:  0,
		RightInterval: 99,
		IsProduct:     false,
		HWPort:        "10014",
	}
	conns []*net.TCPConn
)

func main() {
	var (
		err   error
		count int
	)
	if err = config.LoadConfiguration("project.yaml", c, c); err != nil {
		panic(err)
	}
	if err = c.Save(nil); err != nil {
		panic(err)
	}
	LogCount = log.NewLogFile("count.log")
	if c.IsProduct {
		LogConn = log.NewLogFile("looby.log")
		LogHW = log.NewLogFile("hw.log")
	} else {
		LogConn = log.Log
		LogHW = log.Log
	}
	//var (
	//	command string
	//)
	for i := c.LeftInterval; i <= c.RightInterval; i++ {
		time.Sleep(time.Second * 1)
		count++
		LogCount.Infoln("count: ", count)

		go func(_i int) {
			var name = "yace11"
			str := strconv.Itoa(_i)
			if len(str) == 1 {
				str = "00" + str
			} else if len(str) == 2 {
				str = "0" + str
			}
			name += str
			if err = connect(name, c.Host+":"+c.Port); err != nil {
				LogConn.Errorf("[connect] user: %s, err: %v", name, err)
			}
		}(i)

	}
	//go func() {
	//	input := bufio.NewScanner(os.Stdin)
	//	for input.Scan() {
	//		command = input.Text()
	//		if command == "close" {
	//			for _, v := range conns {
	//				if v != nil {
	//					v.Close()
	//				}
	//			}
	//		}
	//	}
	//}()
	//if command == "close" {
	//	time.Sleep(time.Second * 5)
	//	log.Warnln("programing ready close")
	//	os.Exit(0)
	//}
	time.Sleep(time.Second * 10)
	chSig := make(chan os.Signal)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Signal: ", <-chSig)
	log.Infoln("count: ", count)
}

var e = &EchoProtocol{}

func connect(name, addr string) (err error) {
	var (
		tcpAddress *net.TCPAddr
		conn       *net.TCPConn
		data       []byte

		h        interface{}
		iActorID int32
		token    string
		iToken   string
	)
	if tcpAddress, err = net.ResolveTCPAddr("tcp4", addr); err != nil {
		LogConn.Errorln(err)
		return
	}

	if conn, err = net.DialTCP("tcp", nil, tcpAddress); err != nil {
		return
	}
	conns = append(conns, conn)
	req1 := &protobuf.SQLoginCheck{
		Account:          proto.String(name),
		Password:         proto.String("123456"),
		UniqueIdentifier: proto.String(name),
	}

	if data, err = proto.Marshal(req1); err != nil {
		return
	}
	if _, err = conn.Write(NewEchoPacket(data, false, int32(protobuf.MsgID_MSGID_LOGIN_CHECK)).Serialize()); err != nil {
		LogConn.Errorln("[login-check-password] err:", err, "name: ", name)
		return
	} else {
		if h, err = e.ReadPacket(conn); err != nil {
			log.Errorln("[readPacket-check-password] err:", err, ">>>>>>>>>>>>> name: ", name)
			return
		}
		var (
			d = &protobuf.CSMsgBody{}
		)
		_ = proto.Unmarshal(h.(*EchoPacket).GetBody(), d)
		token = *d.GetSALoginCheck().Token
		// 	openId = *d.GetSALoginCheck().Openid

	}

	req := &protobuf.CGameLoginReq{
		SzloginType:  proto.String("account"),
		SzloginID:    proto.String(name),
		SzloginToken: proto.String(token),
		Ip:           proto.String(conn.RemoteAddr().String()),
	}
	data = []byte("")
	if data, err = proto.Marshal(req); err != nil {
		return
	}
	if _, err = conn.Write(NewEchoPacket(data, false, int32(protobuf.MsgID_MSGID_LOGIN_REQ)).Serialize()); err != nil {
		LogConn.Errorln("[login] err:", err, "name: ", name)
		return
	} else {
		if h, err = e.ReadPacket(conn); err != nil {
			log.Errorln("[readPacket] err:", err, ">>>>>>>>>>>>> name: ", name)
			return
		} else {
			resp := &protobuf.CSMsgBody{}
			_ = proto.Unmarshal(h.(*EchoPacket).GetBody(), resp)
			iActorID = *resp.GetStCGameLoginRes().IActorID
			iToken = *resp.GetStCGameLoginRes().SzSvrToken
			if iActorID != -1 && *resp.GetStCGameLoginRes().IResultId == 10000 {
				// LogConn.Infoln("login success >>>>>>>>>>>>>>>> name: ", name, ", iActorId: ", iActorID)
			} else {
				LogConn.Errorln("[readPacket-login-1] err: ", err)
				err = errors.New("登入失败>>>> 1")
				return
			}
		}
	}

	req3 := &protobuf.CSLoginLogicReq{
		IActorId: proto.Int32(iActorID),
		SzToken:  proto.String(iToken),
	}
	data = []byte("")
	data, _ = proto.Marshal(req3)

	if _, err = conn.Write(NewEchoPacket(data, false, int32(protobuf.MsgID_MSGID_LOGINLOGIC_REQ)).Serialize()); err != nil {
		LogConn.Errorln("[login] err:", err, "name: ", name)
		return
	} else {
		if h, err = e.ReadPacket(conn); err != nil {
			log.Errorln("[readPacket] err:", err, ">>>>>>>>>>>>> name: ", name)
			return
		}
		var (
			d = &protobuf.CSMsgBody{}
		)
		_ = proto.Unmarshal(h.(*EchoPacket).GetBody(), d)
		if iActorID != -1 && *d.GetStLoginLogicRes().IEno == 10000 {
			LogConn.Infoln("609 login success >>>>>>>>>>>>>>>> name: ", name, ", iActorId: ", iActorID)
		} else {
			LogConn.Errorln("[readPacket-login-609] err: ", err)
			err = errors.New("登入失败>>>>>609")
			return
		}
	}

	go func() {
		for {
			time.Sleep(time.Second * 8)
			ping := &protobuf.FLCSPing{
				IActorID:     proto.Int32(iActorID),
				FTimeStamp:   proto.Float32(float32(time.Now().Unix())),
				DwPingCount:  proto.Uint32(1000000),
				Ip:           proto.String(conn.RemoteAddr().String()),
				DwServerTick: proto.Uint32(1000000),
			}

			body, _ := proto.Marshal(ping)

			if _, _err := conn.Write(NewEchoPacket(body, false, int32(protobuf.MsgID_MSGID_PING)).Serialize()); _err != nil {
				LogConn.Errorln("[Write] err: ", _err, ">>>>>>>>>>>>>>", name)
				return
			} else {
				LogConn.Infoln("ping success >>>>>>>>>>>>> name", name)
			}
		}
	}()

	// 捕鱼
	go func() {
		if err = joinRoom(c.Host+":"+c.HWPort, iActorID, iToken, name); err != nil {
			LogHW.Errorln("[join-room-fail] err:", err, ", name: ", name)
			return
		}
	}()

	for {
		var (
			d  = &protobuf.CSMsgBody{}
			d2 = &protobuf.FLCSPing{}
		)
		if p2, _err := e.ReadPacket(conn); _err == nil {
			_ = proto.Unmarshal(p2.(*EchoPacket).GetBody(), d)
			_ = proto.Unmarshal(p2.(*EchoPacket).GetBody(), d2)
			LogConn.Infoln("[read] data: ", d, ", name: ", name)
			if d2.FTimeStamp != nil && *d2.FTimeStamp > 0 {
				LogConn.Infoln("[read-ping] data: ", d2, ", name: ", name)
			}

		} else {
			LogConn.Errorln("[read] err: ", _err, ">>>>>>>>>>>>>>>", name)
			break
		}

	}
	return
}

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

func (this EchoProtocol) ReadPacket(c interface{}) (horse.Packet, error) {
	switch conn := c.(type) {
	case *net.TCPConn:
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
	return nil, nil
}

func joinRoom(addr string, iActorID int32, iToken string, name string) (err error) {

	var (
		conn       *net.TCPConn
		tcpAddress *net.TCPAddr
	)
	if tcpAddress, err = net.ResolveTCPAddr("tcp4", addr); err != nil {
		LogHW.Errorln(err)
		return
	}

	if conn, err = net.DialTCP("tcp", nil, tcpAddress); err != nil {
		return
	}

	catchFishReq := &protobuf.CatchFish_JoinRoomReq{
		IPlayerID:   proto.Int32(iActorID),
		SzToken:     proto.String(iToken),
		IRoomTypeId: proto.Int32(400002),
	}
	_d, _ := proto.Marshal(catchFishReq)
	if _, _err := conn.Write(NewEchoPacket(_d, false, int32(protobuf.MsgID_MSGID_CATCHFISH_JOIN_ROOM_REQ)).Serialize()); err != nil {
		LogHW.Errorln("[Write-joinRoom] err: ", _err, ">>>>>>>>>>>>>>", name)
		return
	}
	// LogConn.Infoln("[join-room-success], name: ", name)
	catchFishRes := &protobuf.CatchFish_JoinRoomRes{}
	if _h, _err := e.ReadPacket(conn); _err != nil {
		err = _err
		LogHW.Errorln("[readPacket] err:", _err, ">>>>>>>>>>>>> name: ", name)
		return
	} else {
		_ = proto.Unmarshal(_h.(*EchoPacket).GetBody(), catchFishRes)
		//if catchFishRes.GetIResultId() != 10000 {
		//	LogConn.Warnln("[join-room] msg: ", catchFishRes)
		//	err = errors.New("join room fail")
		//	return
		//} else {
		//	LogConn.Warnln("[join-room] msg: ", catchFishRes)
		//}
	}

	go func() {
		for {
			time.Sleep(time.Second * 8)
			ping := &protobuf.FLCSPing{
				IActorID:     proto.Int32(iActorID),
				FTimeStamp:   proto.Float32(float32(time.Now().Unix())),
				DwPingCount:  proto.Uint32(999999),
				Ip:           proto.String(conn.RemoteAddr().String()),
				DwServerTick: proto.Uint32(99999),
			}

			body, _ := proto.Marshal(ping)

			if _, _err := conn.Write(NewEchoPacket(body, false, int32(protobuf.MsgID_MSGID_PING)).Serialize()); _err != nil {
				LogHW.Errorln("[Write] err: ", _err, ">>>>>>>>>>>>>>", name)
				return
			} else {
				LogHW.Infoln("ping success >>>>>>>>>>>>> name", name)
			}
		}
	}()

	go func() {
		for true {
			time.Sleep(time.Millisecond * 500)
			req := &protobuf.CatchFish_PlayerFireReq{
				WeaponID:     proto.Int32(0),
				Angle:        proto.Float32(45),
				LockFishID:   proto.Int64(-1),
				SpawnTicks:   proto.Int64(time.Now().Unix() - int64(time.Second)),
				GunPosX:      proto.Float32(10),
				GunPosY:      proto.Float32(20),
				BulletPointX: proto.Float32(15),
				BulletPointY: proto.Float32(21),
			}
			_d, _ := proto.Marshal(req)

			if _, err = conn.Write(NewEchoPacket(_d, false, int32(protobuf.MsgID_MSGID_CATCHFISH_PLAYER_FIRE_REQ)).Serialize()); err != nil {
				LogHW.Errorln("[fire-fish] err: ", err)
				return
			}
		}
	}()

	for {
		var (
			d2 = &protobuf.FLCSPing{}
			d3 = &protobuf.CatchFish_PlayerFireRes{}
		)
		if p2, _err := e.ReadPacket(conn); _err == nil {
			_ = proto.Unmarshal(p2.(*EchoPacket).GetBody(), d2)
			_ = proto.Unmarshal(p2.(*EchoPacket).GetBody(), d3)
			if d2.FTimeStamp != nil && *d2.FTimeStamp > 0 {
				LogHW.Infoln("[read-ping] data: ", d2, ", name: ", name)
			}
			if d3.GetWeaponID() > 0 {
				LogHW.Infoln("[read-fire-fish] data: ", d3, ", name: ", name)
			}
		} else {
			LogHW.Errorln("[read] err: ", _err, ">>>>>>>>>>>>>>>", name)
			break
		}

	}

	return
}

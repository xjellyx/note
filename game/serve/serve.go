package main

import (
	"context"
	"github.com/spf13/viper"
	"github.com/topfreegames/pitaya/acceptor"
	"github.com/topfreegames/pitaya/config"
	"github.com/topfreegames/pitaya/groups"
	"strings"

	"fmt"
	"github.com/topfreegames/pitaya"
	"github.com/topfreegames/pitaya/component"
	"github.com/topfreegames/pitaya/logger"
	"github.com/topfreegames/pitaya/serialize/json"
	"github.com/topfreegames/pitaya/timer"
	"time"
)

// Hall
type Hall struct {
	component.Base
	timer *timer.Timer
}

// UserMessage represents a message that user sent
type UserMessage struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

// NewHall
func NewHall() *Hall {
	return &Hall{}
}

//
func (h *Hall) AfterInit() {
	h.timer = pitaya.NewTimer(time.Minute, func() {
		count, err := pitaya.GroupCountMembers(context.Background(), "hall")
		logger.Log.Debugf("userCount: Time=> %s, Count=> %d, Error=> %q", time.Now().String(), count, err)
	})
}

func (h *Hall) Message(ctx context.Context, msg *UserMessage) {
	err := pitaya.GroupBroadcast(ctx, "lobby", "hall", "onMessage", msg)
	if err != nil {
		fmt.Println("error broadcasting message", err)
	}
}

func configApp() *viper.Viper {
	conf := viper.New()
	conf.SetEnvPrefix("chat") // allows using env vars in the CHAT_PITAYA_ format
	conf.SetDefault("pitaya.buffer.handler.localprocess", 15)
	conf.Set("pitaya.heartbeat.interval", "15s")
	conf.Set("pitaya.buffer.agent.messages", 32)
	conf.Set("pitaya.handler.messages.compression", false)
	return conf
}

func main() {
	defer pitaya.Shutdown()
	s := json.NewSerializer()
	conf := configApp()
	pitaya.SetSerializer(s)

	gsi := groups.NewMemoryGroupService(config.NewConfig(conf))
	pitaya.InitGroups(gsi)

	err := pitaya.GroupCreate(context.Background(), "hall")
	if err != nil {
		panic(err)
	}

	h := NewHall()
	pitaya.Register(h,
		component.WithName("hall"),
		component.WithNameFunc(strings.ToLower),
	)
	ws := acceptor.NewWSAcceptor(":9630")
	pitaya.AddAcceptor(ws)
	pitaya.Configure(true, "hall", pitaya.Cluster, map[string]string{}, conf)
	pitaya.Start()
}

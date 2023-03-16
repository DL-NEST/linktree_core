package benchmark

import (
	"context"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-redis/redis/v8"
	"linktree_core/modules/emqx"
	"sync"
	"testing"
	"time"
)

func mqttCreate(id string) mqtt.Client {
	opts := mqtt.NewClientOptions().AddBroker("mqtt://127.0.0.1:1883").SetClientID(id)

	opts.SetKeepAlive(60 * time.Second)
	// 设置消息回调处理函数
	opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("TOPIC: %s\n", msg.Topic())
		fmt.Printf("MSG: %s\n", msg.Payload())
	})
	opts.SetPingTimeout(1 * time.Second)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return c
}

// 测试发送数据emqx
func TestAddMsg(t *testing.T) {
	var cList []mqtt.Client
	// 创建10000个客户端
	for i := 0; i < 10000; i++ {
		cList = append(cList, mqttCreate(fmt.Sprintf("test/%d", i)))
	}

	time.Sleep(2 * time.Second)

	var wg sync.WaitGroup
	wg.Add(len(cList))
	for index, _ := range cList {
		index2 := index
		go func() {
			for i := 0; i < 1; i++ {
				cList[index2].Publish("test/state", 0, false, fmt.Sprintf("ss:%d", 3)).Wait()
				//cList[index].Publish("test/ctrl", 0, false, fmt.Sprintf("ss:%d", 3)).Wait()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

//func TestRPCServiceBenchmark(t *testing.T) {
//	report, err := runner.Run(
//		"HookProvider.OnMessagePublish",
//		"localhost:9981",
//		runner.WithProtoFile("exhook.proto", []string{}),
//		runner.WithDataFromJSON("{}"),
//		runner.WithInsecure(true),
//	)
//
//	if err != nil {
//		fmt.Println(err.Error())
//		os.Exit(1)
//	}
//
//	printer := printer.ReportPrinter{
//		Out:    os.Stdout,
//		Report: report,
//	}
//
//	printer.Print("pretty")
//}

func BenchmarkAdc(b *testing.B) {
	_ = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	for i := 0; i < b.N; i++ {
		msg := emqx.MsgType{
			Time:  324,
			Topic: "wad",
			Msg:   "string(in.Message.Payload)",
		}
		_, _ = json.Marshal(msg)
		//RdbAuth.RPush(context.Background(), "logCache", res)
	}
}

func BenchmarkAddRedis(b *testing.B) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	for i := 0; i < b.N; i++ {
		msg := emqx.MsgType{
			Time:  324,
			Topic: "wad",
			Msg:   "string(in.Message.Payload)",
		}
		res, _ := json.Marshal(msg)
		rdb.RPush(context.Background(), "logCache", res)
	}
}

func BenchmarkAddRediss(b *testing.B) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	for i := 0; i < b.N; i++ {
		rdb.RPush(context.Background(), "logCache", "res")
	}
}

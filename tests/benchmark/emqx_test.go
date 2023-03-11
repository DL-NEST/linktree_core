package benchmark

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
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
			for i := 0; i < 2; i++ {
				cList[index2].Publish("test/state", 0, false, fmt.Sprintf("ss:%d", 3)).Wait()
				//cList[index].Publish("test/ctrl", 0, false, fmt.Sprintf("ss:%d", 3)).Wait()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

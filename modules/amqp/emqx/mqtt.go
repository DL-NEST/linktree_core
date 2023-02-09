package emqx

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/viper"
	"linktree_core/utils/dlog"
	"time"
)

/*
	1.定义一个监听主题的列表  **
	2.做监听绑定（对监听主题的伪实体化）
	3.动态设置监听处理中间件（把函数当参数传）
*/

// TODO 每创建一个设备就

type ClientOpts struct {
	addr     string
	username string
	password string
	clientId string
}

// MqState MqttState Mqtt 一个返回对象
//type MqState struct {
//	// 主题监听队列
//	LinkList   map[string]Sub
//
//}

type Sub struct {
	name string
	qos  int
}

var MqttClient mqtt.Client

func LinkMqttBroker() {
	// 创建mqtt客户端
	MqttClient = CreateMqttClient(ClientOpts{
		addr:     viper.GetString("broker.addr"),
		username: viper.GetString("broker.username"),
		password: viper.GetString("broker.password"),
		clientId: viper.GetString("broker.client"),
	})
	// 在协程里面监听主题
	go Subscribe(MqttClient, "test", 0)
}

// CreateMqttClient 连接emqx服务器
func CreateMqttClient(clientOpts ClientOpts) mqtt.Client {
	connectAddress := fmt.Sprintf("tcp://%s", clientOpts.addr)
	clientId := clientOpts.clientId
	//fmt.Println("connect address: ", connectAddress)
	opts := mqtt.NewClientOptions()
	opts.AddBroker(connectAddress)
	opts.SetUsername(clientOpts.username)
	opts.SetPassword(clientOpts.password)
	opts.SetClientID(clientId)
	opts.SetKeepAlive(60)
	client := mqtt.NewClient(opts)
	// 如果连接失败，则终止程序
	err := client.Connect()
	if err.WaitTimeout(3*time.Second) && err.Error() != nil {
		dlog.Log.Fatal(err.Error())
	}
	// 没有错误返回连接句柄
	return client
}

// Publish 发布信息
func Publish(client mqtt.Client, topic string, qos int) {
	msgCount := 0
	for {
		payload := fmt.Sprintf("message: %d!", msgCount)
		if token := client.Publish(topic, byte(qos), false, payload); token.Wait() && token.Error() != nil {
		} else {

		}
		msgCount++
		time.Sleep(time.Second * 1)
	}
}

// Subscribe 订阅主题
func Subscribe(mqttClient mqtt.Client, topic string, qos int) {
	//mqttClient.LinkList[topic] = Sub{
	//	name: topic,
	//	qos:  qos,
	//}
	dlog.Log.Debug("SC")
	mqttClient.Subscribe(topic, byte(qos), MessageHandler)
}

func MessageHandler(client mqtt.Client, msg mqtt.Message) {
	// 添加进全局队列
	fmt.Printf("主题 <%s> : `%s`\n", msg.Topic(), string(msg.Payload()))
}

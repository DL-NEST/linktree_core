package emqx

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"linktree_core/global"
	pb "linktree_core/modules/emqx/proto"
	"log"
	"net"
)

// server is used to implement emqx_exhook_v1.s *server
type server struct {
	pb.UnimplementedHookProviderServer
}

func (s *server) OnProviderLoaded(ctx context.Context, in *pb.ProviderLoadedRequest) (*pb.LoadedResponse, error) {
	fmt.Println("hook 创建")
	hooks := []*pb.HookSpec{
		//&pb.HookSpec{Name: "client.connected"},
		//&pb.HookSpec{Name: "client.disconnected"},
		//&pb.HookSpec{Name: "client.subscribe"},
		&pb.HookSpec{Name: "message.publish"},
		//&pb.HookSpec{Name: "message.delivered"},
		//&pb.HookSpec{Name: "message.acked"},
		//&pb.HookSpec{Name: "message.dropped"},
	}
	return &pb.LoadedResponse{Hooks: hooks}, nil
}

func (s *server) OnProviderUnloaded(ctx context.Context, in *pb.ProviderUnloadedRequest) (*pb.EmptySuccess, error) {
	fmt.Println("hook 注销")
	return &pb.EmptySuccess{}, nil
}

func (s *server) OnMessagePublish(ctx context.Context, in *pb.MessagePublishRequest) (*pb.ValuedResponse, error) {

	msg := MsgType{
		Time:  in.Message.Timestamp,
		Topic: in.Message.Topic,
		Msg:   string(in.Message.Payload),
	}
	res, _ := json.Marshal(msg)
	// 这是一个耗时操作，丢到协程中
	// TODO 可以使用ants协程池减少协程的开销
	go global.RdGroup.MqMsg.RPush(context.Background(), "logCache", res)
	return &pb.ValuedResponse{}, nil
}

type MsgType struct {
	Time  uint64
	Topic string
	Msg   string
}

func CreateExHook() {
	lis, err := net.Listen("tcp", ":9981")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHookProviderServer(s, &server{})
	global.GLOG.Infof("Started gRPC server on ::9981")
	go func() {
		err1 := s.Serve(lis)
		if err1 != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
}

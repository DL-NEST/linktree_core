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
		Topic: in.Message.Topic,
		Msg:   string(in.Message.Payload),
	}
	res, _ := json.Marshal(msg)
	global.RdGroup.MqMsg.RPush(ctx, "logCache", res)
	//	//fmt.Printf("%d\n", index)
	//	//in.Message.Payload = []byte("hardcode payload by exhook-svr-go :)")
	//	//reply := &pb.ValuedResponse{}
	//	//reply.Type = pb.ValuedResponse_STOP_AND_RETURN
	//	//reply.Value = &pb.ValuedResponse_Message{Message: in.Message}
	return &pb.ValuedResponse{}, nil
}

type MsgType struct {
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
	global.GLOG.Infof("\tStarted gRPC server on ::9981")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

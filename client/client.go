package main

import (
	"context"
	"log"
	"os"
	"time"

	//"time"
	"google.golang.org/grpc"
	pb "github.com/send"
)

const address = "localhost:5050"

func main() {
	// 创建rpc连接
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock()) // 127.0.0.0
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()
	// 创建gprc client对象
	client := pb.NewSenderClient(conn)
	// 发送rpc请求
	name, err := os.Hostname()
	if err != nil {
		log.Fatalf("get hostname err: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20 * time.Second)
	defer cancel()
	resp, err := client.SendData(ctx, &pb.SendRequest{Hostname: name})
	if err != nil {
		log.Fatalf("client.Send failed: %v", err)
	}
	// 输出结果
	log.Printf("cpu usage from %s,\n cpuID:%d, pcore:%d, lcore:%d, occupy:%g, mhz:%g, cachesize:%d, succ", resp.GetReceiver(), resp.GetCpuid(), resp.GetPcore(), resp.GetLcore(), resp.GetOccupancy(), resp.GetMhz(), resp.GetCacheSize())
}
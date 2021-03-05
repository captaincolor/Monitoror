package main

import (
	"context"
	"log"
	"os"
	//"time"
	"google.golang.org/grpc"
	pb "github.com/send"
)

const RPCPORT = "5050"

func main() {
	// 创建rpc连接
	conn, err := grpc.Dial(":"+RPCPORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	// 创建gprc client对象
	client := pb.NewSenderClient(conn)
	// 发送rpc请求

	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()

	name, err := os.Hostname()
	if err != nil {
		log.Fatalf("get hostname err: %v", err)
	}

	resp, err := client.SendData(context.Background(), &pb.SendRequest{Hostname: name})
	if err != nil {
		log.Fatalf("client.Send err: %v", err)
	}
	// 输出结果
	log.Printf("cpu usage from %s,\n cpuID:%d, pcore:%d, lcore:%d, occupy:%g, mhz:%g, cachesize:%d, succ", resp.GetReceiver(), resp.GetCpuid(), resp.GetPcore(), resp.GetLcore(), resp.GetOccupancy(), resp.GetMhz(), resp.GetCacheSize())
}
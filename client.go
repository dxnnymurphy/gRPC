package main

import (
	"context"
	"log"
	"time"

	"github.com/dxnnymurphy/gRPC/pb"

	"google.golang.org/grpc"
	pbtime "google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	addr := "localhost:9999"
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewAnomalyDetectionClient(conn)
	req := pb.AnomalyRequest{
		Metrics: dummyData(),
	}

	resp, err := client.Predict(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(resp.response)
}

func dummyData() *pb.Metric {
	t := time.Date(2020, 5, 22, 14, 13, 11, 0, time.UTC)
	m := pb.Metric{
		Time:    Timestamp(t),
		topic:   "",
		refresh: False,
		stop:    False,
	}
	return m
}

//will remove
func Timestamp(t time.Time) *pbtime.Timestamp {
	return &pb.time.Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.Nanosecond()),
	}
}

package main

import (
	"context"
	"encoding/csv"
	"flag"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"os"
	"socialanticlub/internal/pb/user_service"
	"time"
)

const (
	filePathFlag = "user-file"
)

func main() {
	ctx := context.Background()

	filePtr := flag.String(filePathFlag, "./tmp/users.csv", "path to user file")
	if filePtr == nil {
		log.Fatal("no file path specified")
	}

	file, err := os.Open(*filePtr)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	defer file.Close()
	fileReader := csv.NewReader(file)

	var (
		eg, egCtx = errgroup.WithContext(ctx)
	)
	eg.SetLimit(10)

	conn, gErr := grpc.DialContext(ctx, ":5000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if gErr != nil {
		log.Fatalf("grpc.DialContext: %v", gErr)
	}
	defer conn.Close()

	for {
		row, fErr := fileReader.Read()
		if fErr != nil {
			if fErr == io.EOF {
				break
			}
			log.Fatalf("error reading csv file: %v", fErr)
		}

		request, pErr := ParseToRequest(row)
		if pErr != nil {
			log.Printf("error parsing request: %v", pErr)
			continue
		}

		eg.Go(func() error {
			var rpcErr error
			for range 30 {
				reqCtx, cancel := context.WithTimeout(egCtx, 10*time.Second)
				defer cancel()
				res := new(user_service.RegResponse)
				rpcErr = conn.Invoke(reqCtx, "social_anti_club.UserService/Register", request, res)
				if rpcErr != nil {
					log.Printf("error posting request: %v", rpcErr)
					time.Sleep(100 * time.Millisecond)
					continue
				}

				if res.GetStatus() != user_service.RegResponse_Success {
					log.Printf("error request: %v", res.GetStatus())
					time.Sleep(100 * time.Millisecond)
					continue
				}

				log.Printf("success request: %v", res.GetUserID())

				break
			}
			if rpcErr != nil {
				log.Fatalln("something went wrong")
			}
			return rpcErr
		})
	}

	_ = eg.Wait()
}

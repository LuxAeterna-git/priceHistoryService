package main

import (
	"context"
	"fmt"
	"github.com/LuxAeterna-git/priceHistoryService/internal/repository"
	"github.com/LuxAeterna-git/priceHistoryService/internal/service"
	"log"
)

func main() {
	mng, err := repository.NewMongo(context.Background(), "localhost", "27017", "root", "qwerty", "files", "UsersFiles", "")
	if err != nil {
		log.Fatal("Error while creating MongoDB controller: ", err)
	}
	s := service.NewService(mng)
	fmt.Println(s)
	s.GetRawData()
}

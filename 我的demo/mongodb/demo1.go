package main

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)


func main(){
	// 建立连接

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://111.229.213.40:27017"))




}
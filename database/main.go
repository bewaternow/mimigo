package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

//	mongoDB 数据库连接单例
var MongoDB *mongo.Client

func Mongo() {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URL"))

	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	MongoDB = client

	if os.Getenv("DB_MIGRATE") == "true" {
		migrate()
	}

}

func migrate() {
	log.Println("数据迁移文件写在这里，你也可以理解为数据库脚本")
}

package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var (
	//	mongoDB 数据库连接单例
	MongoDB *mongo.Client
	//	collection 单例，在 collectionMaps.go 中的方法初始化，因为 mongo 是非关系型的，任何字段都可以存储，所以我想在 maps 中创建单例，
	//	服务直接调用示例，不接触名字，就降低了误操作的可能性
	SupportUser                *mongo.Collection
	SupportPersonalAccessToken *mongo.Collection
)

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
	InitCollections()

	if os.Getenv("DB_MIGRATE") == "true" {
		migrate()
	}

}

func migrate() {
	log.Println("数据迁移文件写在这里，你也可以理解为数据库脚本")
}

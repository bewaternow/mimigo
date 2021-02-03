# Flamingo （基于 Gin 和 mongo-driver 搭建的高性能极简 API 框架）



## 声明
本框架以令牌方式授权，借用了jwt的令牌发授方法，但实际上验证已经自己重写。我参考了 singo 等框架，结合自己的使用场景做了适配。  
如果本项目你发现有任何BUG 、使用疑问或者优化意见，欢迎提交PullRequest

### 特别的优化说明：

1、我指定 mongodb 为默认数据库。  
2、序列化我直接放在了模型中。  
3、jwt 我基本是依靠数据库中的记录来验证（其实已经可以不用 jwt，自己来发令牌了）。  
4、允许注销登录状态，即删除相应的令牌记录。  
5、加入了 OSS 相关的示例。  
6、修改了 i18n 的版本 bug，谢谢 Hel1antHu5！  
7、把令牌与（ip 和 userAgent）进行绑定，尽量确保一个令牌只能在一个设备上使用。

## 建议使用集合的时候要直接用变量来使用，集合名称在 database/collections_Maps.go 文件中定义。

虽然 mongo 是 NoSql ，还是建议把 collection 结构写在 database/collections 下的文件中。
所以，我们约定一下：让 database/collections 下的结构体名称和 database 下的 collections_Maps 里的名称相互对应。

我有一个建议，每次插入新的数据，务必使用 collections 中的模型结构体来创建，否则最后统计字段的数量将非常困难。

## Go Mod
本项目使用Go Mod管理依赖。

```
go mod init Flamingo
export GOPROXY=http://mirrors.aliyun.com/goproxy/
go run main.go // 自动安装
```

## 运行
```
go run main.go
```

项目运行后启动在567端口（可以修改，参考gin文档)

## 更新日志
### 【Ver 0.1.2】 2021年1月30日 17:06 ，本次更新内容有：
1、将每个 collection 做成单例，放在 database 中的 collectionMaps 文件中，目的是为了减少开发者对 mongo 的误操作，很多人英文单词都会拼错的。在服务中我们可以直接调用单例，我做了示例。  
2、将序列化从服务中抽离，目的是让服务更加专注于业务。我们可以把序列化方法写在 collections 的模型中，在控制器中调用。  
3、将所有的返回错误码都定义为 const 类型，存放于 serializer 中的 code 中。  
4、修改了 i18n 的 zh-cn.yaml 文件中的一些文案，使其看起来更加友好。如有需要，请自己手动添加。  
5、修改了路由的一些命名。

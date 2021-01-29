# Flamingo

## 声明

### 我参考了 singo，结合自己的使用场景做了适配：

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

## 接下来我要做的

[] OSS token 的控制

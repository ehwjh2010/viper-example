# viper-example

## 各module的作用

## server
```对外服务层。
controller 提供http接口和前端交互
第三方事件的监听listener, 主要是消费消息
grpc 提供rpc接口和后端交互
```

## config
```配置文件和常量配置```

## docs
```swagger自动生成的文档```

## internal
```核心数据和业务层
dao: 数据库操作相关
proxy:对外部http,cache等第三方调用的proxy
service: 核心的业务逻辑层
middleware 中间件
model 数据库模型
```

## storage
```项目生成临时文件```

## script
```脚本```

## lib
```三方包```

## test
```测试用例层, 主要服务于测试同学.   (开发同学写测试用列写在core或server层)```

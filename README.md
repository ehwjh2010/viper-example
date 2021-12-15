# cobra-example

# 说明

## 各module的作用

[//]: # (### client)

[//]: # (给作为微服务的consumer提供的统一接口,包含dto)

### internal
核心数据和业务层
* dao: 数据库操作相关
* dataobject: data object,和数据模型相对应
* proxy:对外部http,cache等第三方调用的proxy
* service: 核心的业务逻辑层

### server
对外服务层。
* provider 对外作为提供的领域服务, 实现client中定义的接口
* controller 提供http接口和前端交互, 包含web层的相关配置
* restful  基于http接口提供的服务
* 第三方事件的监听Listener, 定时任务的入口等
  此module依赖于client、core,  DTO和DO仅在此moudle转换

### test
测试用例层, 主要服务于测试同学.   (开发同学写测试用列写在core或server层)



## 配置与启动

MODE:prod线上 test测试 dev本地 （默认为dev）


```
mvn clean package -P [MODE] -Dmaven.test.skip=true
```

```
nohup java -jar --spring.profiles.active=[MODE] target/cf-web-[MODE].jar  &
```



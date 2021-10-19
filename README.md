# 关于

go-tracing-demo是一个使用集成多个apm trace的demo，目前在不断完善

仅供学习go apm相关知识



tracing demo 功能：

- elastic apm



支持模块：

- gin
- gorm



## elastic apm

参考文档：https://www.elastic.co/guide/en/apm/agent/go/current/configuration.html



程序需要配置环境变量

```
ELASTIC_APM_SERVER_URL=xxx

ELASTIC_APM_SERVICE_NAME=xxx
```

依赖：

elastic-apm-server

elasticsearch

kibana
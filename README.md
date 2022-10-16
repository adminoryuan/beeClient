# beeClient
- beeClient是一个采集服务器性能并上报服务的一个客户端。
- 已经定义好protobuf通信格式
- # 如何使用
- 下载可执行文件
- 修改config.yml
```java
  server:
    addr: 127.0.0.1 # 服务端地址
    port: 9091 # 端口
    apikey: "asdkjaslkdjlaksjdlkasjdk" # apikey
    node: "127.0.0.1:8888" # 节点名称
 tran:
  interval: 5 #ms 上报间隔
  
  
  

```

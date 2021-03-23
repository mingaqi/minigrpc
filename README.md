# minigrpc

## etcd 
./etcd 启动后
可以看到一些信息

- name=default 节点名称默认default
- data dir = default.etcd 保存日志和快照的目录，默认为当前工作目录default.etcd/目录下。
- 在http://localhost:2380和集群中其他节点通信。
- 在http://localhost:2379提供HTTP API服务，供客户端交互
- heartbeat为100ms，该参数的作用是leader多久发送一次心跳到
- followers，默认值是100ms。
- election为1000ms，该参数的作用是重新投票的超时时间，如果follow在该+ 时间间隔没有收到心跳包，会触发重新投票，默认为1000ms。
- snapshot count为10000，该参数的作用是指定有多少事务被提交时，触发+ 截取快照保存到磁盘。
- 集群和每个节点都会生成一个uuid。
- 启动的时候会运行raft，选举出leader。
- 上面的方法只是简单的启动一个etcd服务，但要长期运行的话，还是做成一个服务好一些。下面将以systemd为例，介绍如何建立一个etcd服务。

```shell
nohup ./etcd &

./etcdctl put key value
./etcdctl get key
./etcdctl del key

## 监听某个key
./etcdctl watch key
 
## 锁   
# 只有当正常退出且释放锁后，lock命令的退出码是0，否则这个锁会一直被占用直到过期（默认60秒）
./etcdctl lock mutex

```

## protobuf
[官方文档](https://www.grpc.io/docs/languages/go/quickstart/)
```shell 
go get -u google.golang.org/grpc
```

1. 下载安装protoc
    这是protobuf编译器,将.proto文件,转义为protobuf原生数据结构, 是通用编辑器
   ```
    https://github.com/protocolbuffers/protobuf/releases
    or
    brew install protobuf
    ```
    

2. 安装**protpc-gen-go**, 在执行protoc时会调用这个插件, 用于填充，序列化和检索request以及 response消息类型的代码
   ```
    go get -u github.com/golang/protobuf/protoc-gen-go
    ```
   

3. 安装**protoc-gen-go-grpc**, 生成的客户端和服务器代码
   ```
    go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
    ```
   


## provider

1. 编写proto文件
[proto3语法](https://developers.google.com/protocol-buffers/docs/proto3)
```protobuf
syntax="proto3";

option go_package="pb/";

message ProdRequest{
  int32 prod_id=1; // 传入商品ID
}
message ProdResponse{
  int32 prod_stock=3; //商品库存
}
```

2. 通过proto中间文件生成go代码
   
    ```shell
    cd provider/pb
    protoc --go_out=.. --go-grpc_out=.. Prod.pb
    ```
3. 编写ProdService实现grpc.pb.go中的业务接口 , 编写服务端启动代码
    ```go
      func main() {
            lis, err := net.Listen("tcp", ":5051")
            if err != nil {
                log.Fatalf("failed to listen: %v", err)
            }
            s := grpc.NewServer()
            pb.RegisterProdServiceServer(s,&services.ProdService{})
            if err := s.Serve(lis); err != nil {
                log.Fatalf("failed to serve: %v", err)
        }
    }
    ```

## consumer

```go
func main() {
    s,err:=grpc.Dial(":5051",grpc.WithInsecure())
    if err != nil {
        log.Fatal(err)
    }
    defer s.Close()

    client:=pb.NewProdServiceClient(s)
    prodResponse,err:=client.GetProdStock(context.Background(),&pb.ProdRequest{
        ProdId: 4,
    })
    log.Println(prodResponse.GetProdStock())
}
```

## 服务注册和发现

### 服务端代码

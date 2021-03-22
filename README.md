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

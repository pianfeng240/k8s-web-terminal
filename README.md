## web-terminal

参考：

https://github.com/du2016/web-terminal-in-go

https://github.com/mthenw/frontail/blob/master/web/index.html

基于go1.11重写，使用module特性

#### 可执行文件下载

linux: http://res-xb.oss-cn-beijing.aliyuncs.com/github/k8s-web-terminal.tar.gz


#### 启动项目

./k8s-web-terminal 

支持参数：

-kubeconfig=xxx  // kubeconfig绝对路径，默认加载$HOME/.kube/config文件

-context=xxx // 访问的集群context name，默认kubernetes-admin@kubernetes

-port=9600 // 监听端口号，默认9600


#### web 访问shell

http://127.0.0.1:9600/static/terminal.html?namespace=[namespace]&pod=[pod name]

#### web 查看容器log
http://127.0.0.1:9600/static/logs.html?namespace=[namespace]&pod=[pod name]paths=[logpath]

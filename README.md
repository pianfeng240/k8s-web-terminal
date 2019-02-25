## web-terminal

参考 https://github.com/du2016/web-terminal-in-go

基于go1.11重写，使用module特性

#### 启动项目

./bin/terminal-mac 

支持参数：

-kubeconfig=xxx  // kubeconfig绝对路径，不指定默认加载$HOME/.kube/config文件

-context=kubernetes-admin@kubernetes // 访问的集群context

-port=9600 // 监听端口号，默认9600


#### 访问web shell

http://127.0.0.1:9600/static/terminal.html?namespace=[namespace]&pod=[pod name]


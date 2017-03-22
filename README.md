# monitor

`适用环境centos`
	
拉取项目代码

	$ cd $GOPATH/src && git clone git@github.com:bileji/monitor.git

安装依赖并编译

	$ cd monitor && glide install && make

修改配置文件

	$ vim /etc/monitor.toml
	
以daemon方式启动monitor
	
	$moitor -d

此时monitor可以有两种身份
	
+ manager

		$ monitor server init && monitor server token

+ node
	
		$ monitor join --addr host:port --token c31ffe43a4e31a2be0900ae15ae83170

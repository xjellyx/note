package mynotes

/*Fedora 下docker入门笔记
一，安装docker：

	卸载docker：sudo apt-get remove docker

	安装依赖包：sudo dnf -y install dnf-plugins-core

	更新dnf：sudo dnf update

	安装docker：sudo dnf install docker -ce

	检查是否安装成功：sudo docker run hello-world

	建立docker用户组：sudo groupadd docker

	非root用户加入用户组：sudo gpasswd -a ${USER} docker

	重启：sudo service docker restart

	如果普通用户执行docker命令，如果提示get …… dial unix /var/run/docker.sock权限不够，则修改/var/run/docker.sock权限 :
	sudo chmod a+rw /var/run/docker.sock

二：使用镜像操作
	获取镜像：docker pull [选项] [Docker Registry 地址[:端口号]/]仓库名[:标签]
	docker pull fedora:29

	运行docker：docker run -it --rm fedora:29 bash
-it:交互的终端模式，--rm：避免浪费空间  bash:交互式 Shell
	退出：exit
	列举镜像：docker image ls
	查看镜像空间：docker system df
	删除镜像：docker image  rm [name]
	构建镜像:docker build [选项] <上下文路径/URL/->

三：docker 常用操作指令
	sudo docker image ls 产开所有镜像

四：docker创建mysql到启动mysql
	sudo docker pull mysql //拉下mysql镜像
	sudo docker image -a //查看 镜像
	// 运行mysql
	sudo docker run -p 3306:3306 --name mymysql -e MYSQL_ROOT_PASSWORD=123456 -d mysql
	// 进入mysql容器
	sudo docker exec -it mysql bash
	// 进入mysql
	mysql -uroot -p
*/

### 相关文档
* [docker image 和 container 的区别](https://www.cnblogs.com/bethal/p/5942369.html)
* [docker 17.09 中文社区文档](https://docs-cn.docker.octowhale.com/)
* [docker 官方文档部分翻译版本](http://www.dockerinfo.net/document)

### 命令

#### Gitlab Runner
* [Gitlab CI&CD 在前端项目自动化构建部署中的实践](https://blog.csdn.net/java060515/article/details/84065083)
* [Gitlab Runner Executor SSH](https://docs.gitlab.com/ee/ci/ssh_keys/README.html#ssh-keys-when-using-the-docker-executor)

#### [Containers](https://docs.docker.com/get-started/part2/)
* `docker ps -all` _查看运行的容器_
* `docker --version` _查看版本_
* `docker info` 和 `docker version` _查看 docker 详细安装信息_
* `docker build -t friendlyhello .` _使用当前文件中的 Dockerfile 创建镜像_
* `docker run -p 4000:80 friendlyhello` _运行 friendlyhello 将端口从 4000 映射到 80_
* `docker run -p -d 4000:80 friendlyhello` _同上但是采用后台模式_
* `docker container ls` _查看活动状态的镜像_
* `docker container ls -a` _查看所有的容器_
* `docker container ls -q` _查看容器的 IDs_
* `docker container stop <Container NAME or ID>` _停止指定的 Container_
* `docker container kill <Container NAME or ID>` _强制关闭指定的 Container_
* `docker container rm <Container NAME or ID>` _本地删除制动 Container_
* `docker container rm $(docker container ls -a -q)` _删除本地所有的 Container_
* `docker images`
* `docker image ls -a` _查看本地镜像_
* `docker image rm <image id>/<Container NAME:TAG>` _通过 ID 或者 NAME:TAG 的方式删除本地的镜像_
* `docker image rm $(docker image ls -a -q)` _删除本地所有镜像_
* `docker login` _登录_
* `docker tag <image id> username/repository:tag` _标记本地镜像，将其归入某一仓库_
* `docker push username/repository` _上传本地所有版本（tag）镜像到仓库_
* `docker push username/repository:tag` _上传本地指定 tag 镜像到仓库_
* `docker run username/repository:tag` _从远程服务器获取镜像并运行_

#### [Services](https://docs.docker.com/get-started/part3/)
* `docker stack ls` _查看 stacks 或者 apps_
* `docker stack deploy -c <composefile> <appname>` _运行 Compose File_
* `docker service ls` _查看 app 关联的运行中的 services_
* `docker service ps <service>` _查看 app 关联的 tasks_
* `docker inspect <task or container>` _检查 task 或者 container_
* `docker stack rm <appname>` _关闭一个应用_
* `docker swarm join` _从 manager 中关闭一个 swarm_

#### [Swarms](https://docs.docker.com/get-started/part4/)
* `docker-machine create --driver virtualbox myvm1` _创建一个 VM (Mac, Win7, Linux)_
* `docker-machine create -d hyperv --hyperv-virtual-switch "myswitch" myvm1` _Win10_
* `docker-machine env myvm1` _查看 VM 的基本信息_
* `docker-machine ssh myvm1 "docker node ls"` _查看 swarm 中的 nodes 列表_
* `docker-machine ssh myvm1 "docker node inspect <node ID>"` _检查 node 并返回其具体信息_
* `docker-machine ssh myvm1 "docker swarm init --advertise-addr <myvm1 ip>"`
* `docker-machine ssh myvm1 "docker swarm join-token -q worker"` _查看 join 的 token_
* `docker-machine ssh myvm1` _通过 SSH 连接 VM (SSh 登录)，键入 “exit” 退出_
* `docker node ls` _查看 swarm 中的 nodes 列表_
* `docker-machine ssh myvm2 "docker swarm leave"` _Worker 离开 Swarm_
* `docker-machine ssh myvm1 "docker swarm leave -f"` _Master 离开并结束 Swarm_
* `docker-machine start myvm1` _启动当前未运行的 VM_
* `docker-machine env myvm1` _显示 VM 的环境变量_
* `eval $(docker-machine env myvm1)` _Mac命令 连接到指定 VM 的 shell_
* `eval $(docker-machine env -u)` _Disconnect shell from VMs, use native docker_
* `docker stack deploy -c <file> <app>`  _Deploy an app; command shell must be set to talk to manager (myvm1), uses local Compose file_
* `docker-machine scp docker-compose.yml myvm1:~` _Copy file to node's home dir (only required if you use ssh to connect to manager and deploy the app)_
* `docker-machine ssh myvm1 "docker stack deploy -c <file> <app>"`   _Deploy an app using ssh (you must have first copied the Compose file to myvm1)_
* `docker-machine stop $(docker-machine ls -q)` _Stop all running VMs_
* `docker-machine rm $(docker-machine ls -q)` _Delete all VMs and their disk images_


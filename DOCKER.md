### 相关文档
* [docker image 和 container 的区别](https://www.cnblogs.com/bethal/p/5942369.html)
* [docker 17.09 中文社区文档](https://docs-cn.docker.octowhale.com/)
* [docker 官方文档部分翻译版本](http://www.dockerinfo.net/document)

### 命令

#### Gitlab Runner
* [Gitlab CI&CD 在前端项目自动化构建部署中的实践](https://blog.csdn.net/java060515/article/details/84065083)
* [Gitlab Runner Executor SSH](https://docs.gitlab.com/ee/ci/ssh_keys/README.html#ssh-keys-when-using-the-docker-executor)

### Images
* `docker --version` _查看版本_
* `docker info` 和 `docker version` _查看 docker 详细安装信息_
* `docker login` _登录_
* `docker images`
* `docker image ls -a` _查看本地镜像_
* `docker image rm <image id>/<Container NAME:TAG>` _通过 ID 或者 NAME:TAG 的方式删除本地的镜像_
* `docker image rm $(docker image ls -a -q)` _删除本地所有镜像_
* `docker rmi [OPTIONS]` _删除本地一个或者多个镜像_
* `docker tag <image id> username/repository:tag` _标记本地镜像，将其归入某一仓库_
* `docker push username/repository` _上传本地所有版本（tag）镜像到仓库_
* `docker push username/repository:tag` _上传本地指定 tag 镜像到仓库_
* `docker run username/repository:tag` _从远程服务器获取镜像并运行_
* `docker save -o <filename> <username/repository>` _导出镜像到本地文件_
  ``` bash
    REPOSITORY           TAG                 IMAGE ID            CREATED             SIZE
    9ke123/sinatra       devel               b93319b24b46        21 minutes ago      191MB
    9ke123/sinatra       v1                  b93319b24b46        21 minutes ago      191MB
    9ke123/sinatra       v2                  b93319b24b46        21 minutes ago      191MB

    docker save -o 9ke123_sinatra.tar 9ke123/sinatra # 将全部的版本导出
    docker save -o 9ke123_sinatra_v1.tar 9ke123/sinatras:v1 # 将指定版本的镜像导出
  ```
* `docker image load --input <filename> 或者 docker image load < <filename>` _从本地文件导入到镜像库中_
* `docker run -it <Container NAME or ID> /bin/bash` _使用新创建的镜像来启动一个容器_
* `docker inspect --format '{{ .State.Pid }}' 4ddf4638572d` _当前正在运行的 Docker 容器的进程号（PID_


#### [Containers](https://docs.docker.com/get-started/part2/)
* `docker ps -all` _查看运行的容器_
* `docker ps -a` _查看终止状态的容器_
* `docker build -t friendlyhello .` _使用当前文件中的 Dockerfile 创建镜像_
* `docker run -p 4000:80 friendlyhello` _运行 friendlyhello 将端口从 4000 映射到 80_
* `docker run -p -d 4000:80 friendlyhello` _同上但是采用后台以守护态（Daemonized）形式运行_
* `docker container ls` _查看活动状态的镜像_
* `docker container ls -a` _查看所有的容器_
* `docker container ls -q` _查看容器的 IDs_
* `docker container [OPTIONS] <Container NAME or ID>` _操作指定的 Container_
  ``` bash
    start   # 启动
    restart # 重新启动
    stop    # 停止
    kill    # 强制关闭
    attach  # 进入
    rm      # 本地删除一个处于终止状态的容易，如果删除一个运行中的容器，需要添加 -f 参数
  ```
* `docker container rm $(docker container ls -a -q)` _删除本地所有的 Container_
* `docker rm [OPTIONS]` _删除一个或者多个 Containers_
* `docker exec -it <cintainer name> /bin/bash 或 <command>` _进入 container 并进入控制台或者执行其他命令_
* `cat ubuntu.tar | docker import - test/ubuntu:v1.0` _从容器快照文件中导入为镜像_
* `docker import http://example.com/exampleimage.tgz example/imagerepo` _通过指定 URL 或者某个目录来导入_

  > 用户既可以使用 docker load 来导入镜像存储文件到本地镜像库，也可以使用 docker import 来导入一个容器快照到本地镜像库。这两者的区别在于容器快照文件将丢弃所有的历史记录和元数据信息（即仅保存容器当时的快照状态），而镜像存储文件将保存完整记录，体积也要大。此外，从容器快照文件导入时可以重新指定标签等元数据信息


#### [Volumnes]()


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


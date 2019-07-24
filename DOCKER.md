### 命令

#### Gitlab Runner
* [Gitlab CI&CD 在前端项目自动化构建部署中的实践](https://blog.csdn.net/java060515/article/details/84065083)

#### [Containers](https://docs.docker.com/get-started/part2/)
* `docker ps -all` _查看运行的容器_
* `docker --version` _查看版本_
* `docker info` 和 `docker version` _查看 docker 详细安装信息_
* `docker build -t friendlyhello .` _使用当前文件中的 Dockerfile 创建镜像_
* `docker run -p 4000:80 friendlyhello` _运行 friendlyhello 将端口从 4000 映射到 80_
* `docker run -p -d 4000:80 friendlyhello` _同上但是采用后台模式_
* `docker container ls` _查看活动状态的镜像_
* `docker container ls -a` _查看所有的容器_
* `docker container stop <Container NAME or ID>` _停止指定的 Container_
* `docker container kill <Container NAME or ID>` _强制关闭指定的 Container_
* `docker container rm <Container NAME or ID>` _本地删除制动 Container_
* `docker contailer rm $(docker container ls -a -q)` _删除本地所有的 Container_
* `docker image ls -a` _查看本地镜像_
* `docker image rm <image id>/<Container NAME:TAG>` _通过 ID 或者 NAME:TAG 的方式删除本地的镜像_
* `docker image rm $(docker image ls -a -q)` _删除本地所有镜像_
* `docker login` _登录_
* `docker tag <image> username/repository:tag` _标记本地镜像，将其归入某一仓库_
* `docker push username/repository:tag` _上传本地镜像到仓库_
* `docker run username/repository:tag` _从远程服务器获取镜像并运行_

#### [Services](https://docs.docker.com/get-started/part3/)


### 扩展
* [docker image 和 container 的区别](https://www.cnblogs.com/bethal/p/5942369.html)


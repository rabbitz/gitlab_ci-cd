> Dockerfile 分为四部分：基础镜像信息，维护者信息，镜像操作指令，容器启动时执行指令

### 相关文档
* [Dockerfile 介绍](http://www.dockerinfo.net/dockerfile%e4%bb%8b%e7%bb%8d)

### 文件参数说明

- FROM </br>

  格式为 `FROM <image>` 或 `FROM <image>:<tag>` </br>
  第一条指令必须为 `FROM` 指令，如果在同一个Dockerfile中创建多个镜像时，可以使用多个 `FROM` 指令（每个镜像一次)

- MAINTAINER </br>

  格式为 `MAINTAINER <name>` 指定维护者信息

- RUN </br>

  - `RUN <command>` 在 shell 终端中运行命令，即 `/bin/sh -c`
  - `RUN ["executable", "param1", "param2"]` 使用 `exec` 执行, 即 `RUN ["/bin/bash", "-c", "echo hello"]`
  - 每条 `RUN` 指令将在当前镜像基础上执行指定命令，并提交为新的镜像。当命令较长时可以使用 `\` 来换行

- CMD </br>

  支持三种格式:

  - `CMD ["executable","param1","param2"]` 使用 `exec` 执行【**推荐方式**】
  - `CMD command param1 param2` 在 `/bin/sh` 中执行，提供给需要交互的应用
  - `CMD ["param1","param2"]` 提供给 `ENTRYPOINT` 的默认参数

  指定启动容器时执行的命令，每个 Dockerfile 只能有一条 CMD 命令。如果指定了多条命令，只有最后一条会被执行

  如果用户启动容器时候指定了运行的命令，则会覆盖掉 CMD 指定的命令
















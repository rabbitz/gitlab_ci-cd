1. Docker镜像列表中的none:none是什么?</br>



    >  在构建过Docker镜像的电脑上查看本地镜像列表，有可能看到下图红框中的镜像，在列表中展示为<none>:<none>：</br>
    > [Prune unused Docker objects](https://docs.docker.com/config/pruning/)

    ```bash
    REPOSITORY                    TAG                 IMAGE ID            CREATED              SIZE
    bolingcavalry/eureka-server   0.0.1-SNAPSHOT      be262f101e2c        About a minute ago   683MB
    bolingcavalry/eureka-server   latest              be262f101e2c        About a minute ago   683MB
    <none>                        <none>              90b736eb388e        9 minutes ago        683MB
    ```
    这种镜像在Docker官方文档中被称作dangling images，指的是没有标签并且没有被容器使用的镜像。



    ```bash
    docker image prune
    ```



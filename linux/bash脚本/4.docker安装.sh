#!/bin/bash
VERSION_STRING=18.09.1
# 1.卸载旧版本docker-engine
sudo yum erase docker-common-2:1.12.6-68.gitec8512b.el7.centos.x86_64 -y
sudo yum remove docker-ce docker-ce-cli containerd.io -y
sudo yum -y remove docker \
                docker-client \
                docker-client-latest \
                docker-common \
                docker-latest \
                docker-latest-logrotate \
                docker-logrotate \
                docker-engine
# 2.设置存储库
sudo yum install -y yum-utils device-mapper-persistent-data lvm2
sudo yum-config-manager -y \
    --add-repo \
    https://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
# 3.安装 Docker 引擎
sudo yum install docker-ce docker-ce-cli containerd.io -y
# 查看repo版本  yum list docker-ce --showduplicates | sort -r
# 4.安装特定版本的 Docker Engine
sudo yum - install docker-ce-${VERSION_STRING} docker-ce-cli-${VERSION_STRING} containerd.io docker-compose-plugin
# 5.卸载docker-compose
sudo rm -rf /usr/local/lib/docker/cli-plugins/docker-compose||rm -rf /usr/local/lib/docker-compose
# 6.安装docker-compose
sudo curl -L https://get.daocloud.io/docker/compose/releases/download/1.29.2/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose
# 7.添加x权限
sudo chmod g+x /usr/local/bin/docker-compose
# 开机启动docker和docker-compose
if [ -z "$(ls /etc/systemd/system/multi-user.target.wants/|grep docker)" ];then
    echo "开机启动docker"
    systemctl enable docker
fi
echo "添加docker组"
sudo groupadd docker
echo "添加当前用户到docker组"
sudo usermod -a -G docker $USER
echo "更新docker组"
newgrp docker
echo "/etc/docker目录添加w组权限"
sudo chmod g+w /etc/docker
sudo cat>/etc/docker/daemon.json <<EOF
{
  "registry-mirrors": [
      "http://hub-mirror.c.163.com",
      "https://kls8vom3.mirror.aliyuncs.com",
      "http://hub-mirror.c.163.com",
      "https://3laho3y3.mirror.aliyuncs.com",
      "http://f1361db2.m.daocloud.io",
      "https://mirror.ccs.tencentyun.com",
      "https://docker.mirrors.ustc.edu.cn",
      "https://registry.docker-cn.com"
      ]
}
EOF
echo "启动docker和docker-compse"
sudo systemctl daemon-reload
sudo systemctl start docker
echo "配置Bridge网卡"
docker network create --subnet=172.21.235.0/24  ink_data
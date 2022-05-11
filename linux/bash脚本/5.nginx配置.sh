#!/bin/bash

# 判断url地址是否有效   存在则下载
exists_url() {
    wget --spider -q -o /dev/null --tries=1 -T 5 $1
    if [ $? -eq 0 ]
    then
        wget $1
    else
        echo "$1 is fail."
        exit 1
    fi
}

# 1.删除旧版本nginx
sudo yum erase nginx*
sudo find / -name "nginx"|xargs -i sudo rm -rf {}
# 2.下载依赖包
sudo yum install -y gcc-c++ \
               pcre pcre-devel \
               zlib zlib-devel \
               openssl \
               openssl-devel
exists_nginxDir=$(sudo find / -name "nginx-1.9.9.tar.gz"|head -n 1)
# 3.下载nginx1.9
if [ "$exists_nginxDir" ];then
    cd $(dirname $exists_nginxDir)
else
    exists_url "http://nginx.org/download/nginx-1.9.9.tar.gz"
fi
# 4.解压 进入目录
tar -zxf nginx-1.9.9.tar.gz && cd nginx-1.9.9
# 5.编译安装
sudo ./configure --prefix=/etc/nginx/ 
sudo make && sudo make install
echo "nginx配置中..."
# 6.添加软连接到bin
sudo ln -sf /etc/nginx/sbin/nginx /usr/local/sbin
# 7.注册service
sudo touch -f /usr/lib/systemd/system/nginx.service
if [ ! -w "/usr/lib/systemd/system/nginx.service" ];then
    sudo chmod g+w /usr/lib/systemd/system/nginx.service
fi
sudo cat >/usr/lib/systemd/system/nginx.service<<EOF
[Unit]
Description=NGINX server
After=network.target

[Service]
Type=forking
ExecStartPre=/etc/nginx/sbin/nginx -t
ExecStart=/etc/nginx/sbin/nginx
ExecReload=/etc/nginx/sbin/nginx -s reload
PrivateTmp=True

[Install]
WantedBy=multi-user.target
EOF
# 9.开机启动nginx
if [ -z "$(ls /etc/systemd/system/multi-user.target.wants/|grep nginx)" ];then
    echo "开机启动nginx"
    systemctl enable nginx
fi
# 增加自定义conf配置文件
sudo mkdir -p /etc/nginx/confs
if [ -z "$(cat /etc/nginx/conf/nginx.conf|grep -n "include /etc/nginx/confs/\*.conf;"|cut -f1 -d:)" ];then
    num=$(cat /etc/nginx/conf/nginx.conf|grep -nm 1 server|cut -f1 -d:)
    sudo sed -e 's/#gzip\s*on;/gzip on;/' \
            -e "${num} i include\ /etc/nginx/confs/*.conf;" \
            -ibak \
            /etc/nginx/conf/nginx.conf
fi
# 10.启动nginx
sudo systemctl daemon-reload
sudo systemctl start nginx
# 11.开放服务
if [ ! "$(sudo firewall-cmd --list-ports|grep 80)" ];then
    echo "开放80端口"
    sudo firewall-cmd --permanent --zone=public --add-port=80/tcp
    sudo firewall-cmd --reload
fi
cd ~
sudo find / -name nginx-1.9.9|xargs -i sudo rm -rf {}
if [ $(curl -m 5 -s -w %{http_code} localhost:80 -o /dev/null/span) -eq 200 ];then
echo "nginx搭建成功！！！"
fi
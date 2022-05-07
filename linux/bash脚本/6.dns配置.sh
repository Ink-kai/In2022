#!/bin/bash

# 1.卸载named
if [ "$(yum list installed|grep bind)" ];then
    echo "删除旧版本named"
    sudo yum erase -y bind
    whereis named|xargs -i sudo rm -rf {}
fi
# 2.安装bind并添加confs配置
sudo yum install -y bind dind-devel bind-utils
if [ -a "/etc/named.conf" ];then
    sudo usermod -a -G named $USER
    sudo chmod g+w /etc/named
    sudo sed -e 's/listen-on port 53 { 127.0.0.1; };/listen-on port 53 { any; };/' \
             -e 's/allow-query\s*{ localhost; };/allow-query     { any; };/' \
             -i.bak /etc/named.conf
    if [ -z "$(cat /etc/named.conf|grep 'include /etc/named/confs/*.zones;')" ];then
        sudo sed -e '$a include "/etc/named/confs/named.inkCloud.zones";' -i.bak /etc/named.conf
    fi
    sudo mkdir -p /etc/named/confs;
    echo "firewall开放named服务"
    sudo firewall-cmd --permanent --zone=public --add-port=53/tcp
    sudo firewall-cmd --permanent --zone=public --add-port=53/udp
    sudo firewall-cmd --reload
    echo "开机启动DNS服务"
    sudo systemctl enable named
    sudo systemctl start named
fi
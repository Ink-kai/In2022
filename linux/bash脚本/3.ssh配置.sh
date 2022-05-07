#!/bin/bash
config=/etc/ssh/sshd_config
echo "备份sshd_config"
sed -e 's/#Port 22/Port 22/' \
    -e 's/#PermitRootLogin yes/PermitRootLogin no/' \
    -e 's/#PasswordAuthentication yes/PasswordAuthentication yes/' \
    -i.bak ${config}`
echo "开放ssh服务端口"
if [ -z "$(firewall-cmd --list-service|grep ssh)" ];then
    echo "开放ftp服务端口"
    firewall-cmd --zone=public --permanent --add-service=ssh
    echo "重启防火墙"
    firewall-cmd --reload
fi

#!/bin/bash

##########################################################################
#Author:                     Ink
#Email:                      ink_kai@163.com
#Date:                       2022-05-11
#FileName:                   GNOME定制.sh
#Description:                 The test script
#Copyright (C):             2022 All rights reserved
##########################################################################

# sudo yum -y upgrade
# sudo yum -y install grub2-efi fwupdate
echo "请重启电脑，删除老版本内核"
# yum -y remove kernel
# echo "请再次重启电脑"
# 查看可安装组列表
# yum grouplist
# echo "安装桌面"
# sudo yum -y groupinstall 'GNOME Desktop' 'Graphical Administration Tools'
# echo "当前系统启动模式：" systemctl get-default
# 设置默认启动为图形化界面：
# systemctl set-default graphical.target
# 设置默认启动为命令行界面：
# systemctl set-default multi-user.target
# 命令行界面热切换到图形界面：
# init 5
# 图形界面热切换到命令行界面：
# init 3

# 安装xrdp和VNC服务端（实现rdp远程）
# sudo yum remove xrdp && sudo yum remove xrdp -y
# sudo yum install -y tigervnc-server
# sudo systemctl --now enable vncserver
# sudo wget -O /etc/yum.repos.d/epel.repo http://mirrors.aliyun.com/repo/epel-7.repo
# sudo yum install -y xrdp
# systemctl enable --now xrdp
# 可修改远程端口：/etc/xrdp/xrdp.ini
# sed -i.bak 's/port=3389/port=6666/' /etc/xrdp/xrdp.ini
# if [ ! "$(sudo firewall-cmd --list-ports|grep 6666)" ];then
#     echo "开放3389端口"
#     sudo firewall-cmd --permanent --zone=public --add-port=6666/tcp
#     sudo firewall-cmd --reload
# fi
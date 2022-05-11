#!/bin/bash
ftpDir=/home/ink/ftpdata
if [ -z "$(yum list installed|grep vsftpd)" ];then
    echo "安装vsftpd"
    yum -y install vsftpd
fi
#ftp_home_dir和allow_ftpd_full_access必须为on 才能使vsftpd 具有访问ftp根目录，以及文件传输等权限
echo "开启ftp_home_dir和allow_ftpd_full_access"
setsebool -P tftp_home_dir 1&&setsebool -P allow_ftpd_full_access 1
if [ -n "${ftpDir}" ];then
    echo "创建$ftpDir"
    mkdir -p $ftpDir
    sudo chown "$USER".ftp $ftpDir && sudo chmod 755 $ftpDir
fi
echo "开始配置vsftpd.conf"
exists_sshd_config=$(ls /etc/vsftpd|grep 'vsftpd\(.*\)bak')
if [ -z "$exists_sshd_config" ];then
    cp /etc/vsftpd/vsftpd.conf /etc/vsftpd/vsftpd.confbak
fi
echo -e "#自定义配置\n
# 关闭反向解析
reverse_lookup_enable=NO\n
#匿名主目录
anon_root=/home/ink/ftpdata/\n
#匿名登录FTP服务器
anonymous_enable=yes\n
#匿名允许上传
anon_upload_enable=YES\n
# 匿名用户允许创建目录
anon_mkdir_write_enable=yes\n
# 匿名用户允许下载
anon_world_readable_only=no\n
#匿名用户不允许删除、下载
anon_other_write_enable=no\n
# 匿名用户上传文件的umask
anon_umask=002\n
# 是否允许本地用户(即linux系统中的用户帐号)登录FTP服务器，默认设置为YES允许
# 本地用户登录后会进入用户主目录，而匿名用户登录后进入匿名用户的下载目录/var/ftp/pub
# 若只允许匿名用户访问，前面加上#注释掉即可阻止本地用户访问FTP服务器
local_enable=YES\n
# 是否允许本地用户对FTP服务器文件具有写权限，默认设置为YES允许
write_enable=YES\n
# 掩码，本地用户默认掩码为077
# 你可以设置本地用户的文件掩码为缺省022，也可根据个人喜好将其设置为其他值
local_umask=002\n
# 是否设定FTP服务器将启用FTP数据端口的连接请求
# ftp-data数据传输，21为连接控制端口
connect_from_port_20=YES\n
#启用记录上传下载日志（默认）
xferlog_enable=YES\n
#使用wu-ftp日志格式（默认）
xferlog_std_format=YES\n
#可自动生成（默认）
xferlog_file=/var/log/xferlog\n
#使用vsftpd日志格式，默认不启用
dual_log_enable=YES\n
#可自动生成（默认）
vsftpd_log_file=/var/log/vsftpd.log\n
# 是否允许监听。
# 如果设置为YES，则vsftpd将以独立模式运行，由vsftpd自己监听和处理IPv4端口的连接请求
listen=YES\n
# 设定是否支持IPV6。如要同时监听IPv4和IPv6端口，
# 则必须运行两套vsftpd，采用两套配置文件
# 同时确保其中有一个监听选项是被注释掉的
#listen_ipv6=YES\n
#是否将所有用户限制在主目录,YES为启用 NO禁用.(该项默认值是NO,即在安装vsftpd后不做配置的话，ftp用户是可以向上切换到要目录之外的)
chroot_local_user=yes\n
# 所有的用户都将拥有chroot权限
allow_writeable_chroot=yes\n
# vsftpd:192.168.57.1:DENY 和vsftpd:192.168.57.9:DENY
# 表明限制IP为192.168.57.1/192.168.57.9主机访问IP为192.168.57.2的FTP服务器
# 此时FTP服务器虽可以PING通，但无法连接
tcp_wrappers=YES\n
pam_service_name=vsftpd
# 最大并发连接数
# 如果为0的话，默认不限制
max_clients=5\n
# 每个IP同时发起的最大连接数
# 如果为0的话，则默认不限制数量
max_per_ip=5\n" >/etc/vsftpd/vsftpd.conf
echo "重启vsftpd服务"
service vsftpd restart
if [ -z "$(systemctl list-units|grep vsftpd)" ];then
    echo "开机启动vsftpd"
    systemctl enable vsftpd
fi
if [ ! "$(sudo firewall-cmd --list-service|grep ftp)" ];then
    echo "开放ftp服务端口"
    firewall-cmd --zone=public --permanent --add-service=ftp
    echo "重启防火墙"
    firewall-cmd --reload
fi
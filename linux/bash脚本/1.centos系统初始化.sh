#!/bin/bash
#eth_exists_114num=$(cat ifcfg-ens33|grep -E -n 'DNS([0-9]+)=(\"?)114.114.114.114("$)?'|cut -f1 -d:)
eth_device_name=$(nmcli connection show|awk 'NR==2 {print $NF}')
eth_exists_114num=$(cat ifcfg-ens33|grep -E -n 'DNS([0-9]+)=(\"?)114.114.114.114("$)?'|cut -f1 -d:)
if [ -z "$eth_exists_114" ];then
    if [ ! -w "$/etc/sysconfig/network-scripts/ifcfg-${eth_device_name}" ];then
        sudo chmod g+w /etc/sysconfig/network-scripts/ifcfg-${eth_device_name}
    fi
    echo "添加114DNS到网卡..."
    sudo sed -i.bak "${eth_exists_114num}i DNS2=114.114.114.114" /etc/sysconfig/network-scripts/ifcfg-${eth_device_name}
    sudo service network restart
fi
echo "centos7更换清华源..."
sed -e 's|^mirrorlist=|#mirrorlist=|g' \
    -e 's|^#baseurl=http://mirror.centos.org|baseurl=https://mirrors.tuna.tsinghua.edu.cn|g' \
    -i.bak \
    /etc/yum.repos.d/CentOS-*.repo
echo "清除所有缓存，重新建立元数据"
yum clean all&& yum makecache

exists_user=$(cat /etc/passwd|grep "ink")
exists_sudoers=$(cat /etc/sudoers|grep "ink")
if [ -z "$exists_user" ] && [ -z "$exists_sudoers" ];then
    echo "添加用户到sudoers"
    sed -i{,bak} "/root.*ALL=(ALL).*ALL/a\ink ALL=(ALL) ALL" /etc/sudoers
fi
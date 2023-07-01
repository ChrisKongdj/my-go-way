# linux 配置 k8s节点

### 防火墙的开启、关闭、禁用命令

设置开机启用防火墙：systemctl enable firewalld.service
设置开机禁用防火墙：systemctl disable firewalld.service
启动防火墙：systemctl start firewalld
关闭防火墙：systemctl stop firewalld
检查防火墙状态：systemctl status firewalld

## 禁用swap分区

vi /etc/fstab

i #进入insert 插入模式
:wq #保存退出



### 关闭selinux

临时关闭：输入命令`setenforce 0`。但重启系统后还会开启。
永久关闭：输入命令`vi /etc/selinux/config`，将SELINUX=enforcing改为SELINUX=disabled，然后保存退出。





## 安装组件

yum install -y socat
yum install -y conntrack



## 升级内核

```sh
#导入ELRepo仓库的公钥
rpm --import https://www.elrepo.org/RPM-GPG-KEY-elrepo.org
#为yum安装ELRepo仓库
rpm -Uvh http://www.elrepo.org/elrepo-release-7.0-3.el7.elrepo.noarch.rpm
#查看可用版本
yum --disablerepo="*" --enablerepo="elrepo-kernel" list available
#安装最新内核
yum --enablerepo=elrepo-kernel install kernel-ml


# 查看当前可用内核
awk -F\' '$1=="menuentry " {print i++ " : " $2}' /etc/grub2.cfg

# 0为上一条指令查看记录序号
grub2-set-default 0 

# 重启
reboot

# 查看内核版本
uname -r 

```

## hosts设置

`vi /etc/hosts` 即可设置hosts。K8S集群一般包含多台计算机，hosts的内容应该包含集群内的所有机器

```sh
#vi /etc/hosts

192.168.1.21 audi.server.home
192.168.1.22 aito.server.home
192.168.1.23 aion.server.home
192.168.1.24 bentley.server.home
192.168.1.25 benz.server.home
192.168.1.26 bmw.server.home
192.168.1.27 buck.server.home
192.168.1.28 byd.server.home

```



## 出现containerd unknown service runtime.v1alpha2.ImageService处理

```sh
vi /etc/containerd/config.toml

systemd_cgroup = true
snapshotter = "native"

systemctl restart containerd
```


# Systemd 配置、安装和启动

- [Systemd 配置、安装和启动](#systemd-配置安装和启动)
	- [1. 前置操作](#前置操作)
	- [2. 创建 zenith systemd unit 模板文件](#创建-zenith-systemd-unit-模板文件)
	- [3. 复制 systemd unit 模板文件到 sysmted 配置目录](#复制-systemd-unit-模板文件到-sysmted-配置目录)
	- [4. 启动 systemd 服务](#启动-systemd-服务)

## 1. 前置操作

1. 创建需要的目录

```bash
sudo mkdir -p /data/zenith /opt/zenith/bin /etc/zenith /var/log/zenith
```

2. 编译构建 `zenith` 二进制文件

```bash
make build # 编译源码生成 zenith 二进制文件
```

3. 将 `zenith` 可执行文件安装在 `bin` 目录下

```bash
sudo cp _output/platforms/linux/amd64/zenith /opt/zenith/bin # 安装二进制文件
```

4. 安装 `zenith` 配置文件

```bash
sed 's/.\/_output/\/etc\/zenith/g' configs/zenith.yaml > zenith.sed.yaml # 替换 CA 文件路径
sudo cp zenith.sed.yaml /etc/zenith/ # 安装配置文件
```

5. 安装 CA 文件

```bash
make ca # 创建 CA 文件
sudo cp -a _output/cert/ /etc/zenith/ # 将 CA 文件复制到 zenith 配置文件目录
```

## 2. 创建 zenith systemd unit 模板文件

执行如下 shell 脚本生成 `zenith.service.template`

```bash
cat > zenith.service.template <<EOF
[Unit]
Description=APIServer for blog platform.
Documentation=https://github.com/LiangNing7/zenith/blob/master/init/README.md

[Service]
WorkingDirectory=/data/zenith
ExecStartPre=/usr/bin/mkdir -p /data/zenith
ExecStartPre=/usr/bin/mkdir -p /var/log/zenith
ExecStart=/opt/zenith/bin/zenith --config=/etc/zenith/zenith.yaml
Restart=always
RestartSec=5
StartLimitInterval=0

[Install]
WantedBy=multi-user.target
EOF
```

## 3. 复制 systemd unit 模板文件到 sysmted 配置目录

```bash
sudo cp zenith.service.template /etc/systemd/system/zenith.service
```

## 4. 启动 systemd 服务

```bash
sudo systemctl daemon-reload && systemctl enable zenith && systemctl restart zenith
```

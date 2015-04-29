[lightweight docker containers with buildroot](https://blog.docker.com/2013/06/create-light-weight-docker-containers-buildroot/)

[buildroot](http://buildroot.uclibc.org/)

[docker consul](https://github.com/progrium/docker-consul)


### notes
	
	vagrant up
	vagrant ssh node1

	fleetctl submit /etc/systemd/system/myapp@.service
	fleetctl start myapp@{1..3}

	docker run -t -i ubuntu
	apt-get install curl dnsutils
	dig myapp.service.consul
	curl node3.node.dc1.consul:49154/demo


#### buildroot

	curl http://buildroot.uclibc.org/downloads/buildroot-2013.05.tar.bz2 | tar jx
	cd buildroot-2013.05/
	make menuconfig
	# select x86_64
	make

	cd output/images
	mkdir extra extra/etc extra/sbin extra/lib extra/lib64
	touch extra/etc/resolv.conf
	touch extra/sbin/init
	cp /lib/x86_64-linux-gnu/libpthread.so.0 /lib/x86_64-linux-gnu/libc.so.6 extra/lib
	cp /lib64/ld-linux-x86-64.so.2 extra/lib64


	cp rootfs.tar fixup.tar
	tar rvf fixup.tar -C extra .
	
	docker import - dietfs < fixup.tar

#### systemd

	journalctl -f -u hello.service
	systemctl status|start|... hello.service


  - path: /etc/consul/config.json
    permissions: 0644
    owner: root
    content: |
        {
            "data_dir": "/opt/consul",
            "client_addr": "0.0.0.0",
            "ports": {
                "dns": 53
            },
            "recursors": ["8.8.8.8"],
            "disable_update_check": true
        }

    - name: nameservers.network
      content: |
        [Network]
        DNS=$public_ipv4
        #DNS=8.8.8.8

    - name: consul-server.service
      command: start
      content: |
        [Unit]
        Description=Consul Server Agent
        Requires=docker.service
        After=docker.service

        [Service]
        ExecStartPre=/usr/bin/mkdir -p /opt/bin
        ExecStartPre=/usr/bin/wget -q https://dl.bintray.com/mitchellh/consul/0.5.0_linux_amd64.zip -O /tmp/consul.zip
        ExecStartPre=/usr/bin/unzip -o /tmp/consul.zip -d /opt/bin/
        ExecStartPre=/usr/bin/chmod +x /opt/bin/consul
        ExecStartPre=/usr/bin/rm /tmp/consul.zip
        ExecStart=/bin/bash -c '/opt/bin/consul agent -server -config-dir=/etc/consul/ -advertise $public_ipv4 -bootstrap-expect 3 $(/etc/systems/scripts/consul-join-args)'
        ExecStartPost=/usr/bin/etcdctl set consul.io/nodes/%m $public_ipv4
        ExecReload=/bin/kill -HUP $MAINPID
        ExecStop=/usr/bin/etcdctl rm consul.io/nodes/%m
        Restart=on-failure
        RestartSec=20s

        [Install]
        WantedBy=multi-user.target


#### vagrant


	if [ ! `which docker` ]; then
	  wget -qO- https://get.docker.com/ | sh
	  usermod -aG docker vagrant
	  BRIDGE=$(ifconfig | grep -A 1 docker | tail -n 1 | awk -F: '{print $2}' | awk '{print $1}')
	  echo "DOCKER_OPTS='--dns $BRIDGE --dns 8.8.8.8 --dns-search service.consul'" >> /etc/default/docker
	  service docker restart
	fi




	$provision = <<SCRIPT
	# docker ps -q | xargs docker kill
	# rm -rf /mnt

	BRIDGE=$(ifconfig | grep -A 1 docker | tail -n 1 | awk '{print $2}')
	ETH1=$(ifconfig | grep -A 1 eth1 | tail -n 1 | awk '{print $2}')

	echo nameserver $BRIDGE > /etc/resolv.conf
	echo nameserver $ETH1 >> /etc/resolv.conf

	mkdir -p /etc/systemd/resolved.conf.d/
	echo [Resolve] > /etc/systemd/resolved.conf.d/consul.conf
	echo DNS=$BRIDGE >> /etc/systemd/resolved.conf.d/consul.conf
	echo FallbackDNS=8.8.8.8 >> /etc/systemd/resolved.conf.d/consul.conf

	systemctl restart systemd-resolved

	#-h `hostname` -v /mnt:/data -p $ETH1:8300:8300 -p $ETH1:8301:8301 -p $ETH1:8301:8301/udp -p $ETH1:8302:8302 -p $ETH1:8302:8302/udp -p $ETH1:8400:8400 -p $ETH1:8500:8500 -p $BRIDGE:53:53 -p $BRIDGE:53:53/udp benschw/consul agent -server -config-dir=/config -advertise $ETH1 -data-dir /data -bootstrap-expect 3 -join 172.20.20.10


	# EXTRA=""
	# if [ `hostname` != "node1" ]; then
	#   EXTRA="-join 172.20.20.10"
	# fi

	# docker run -d -h `hostname` -v /mnt:/data \
	#     -p $ETH1:8300:8300 \
	#     -p $ETH1:8301:8301 \
	#     -p $ETH1:8301:8301/udp \
	#     -p $ETH1:8302:8302 \
	#     -p $ETH1:8302:8302/udp \
	#     -p $ETH1:8400:8400 \
	#     -p $ETH1:8500:8500 \
	#     -p $BRIDGE:53:53 \
	#     -p $BRIDGE:53:53/udp \
	#     benschw/consul agent -server -config-dir=/config -advertise $ETH1 -data-dir /data -bootstrap-expect 3 $EXTRA

	SCRIPT


[lightweight docker containers with buildroot](https://blog.docker.com/2013/06/create-light-weight-docker-containers-buildroot/)

[buildroot](http://buildroot.uclibc.org/)

[docker consul](https://github.com/progrium/docker-consul)


### notes
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


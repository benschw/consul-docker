[lightweight docker containers with buildroot](https://blog.docker.com/2013/06/create-light-weight-docker-containers-buildroot/)

[buildroot](http://buildroot.uclibc.org/)

[docker consul](https://github.com/progrium/docker-consul)


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









	if [ ! `which docker` ]; then
	  wget -qO- https://get.docker.com/ | sh
	  usermod -aG docker vagrant
	  BRIDGE=$(ifconfig | grep -A 1 docker | tail -n 1 | awk -F: '{print $2}' | awk '{print $1}')
	  echo "DOCKER_OPTS='--dns $BRIDGE --dns 8.8.8.8 --dns-search service.consul'" >> /etc/default/docker
	  service docker restart
	fi

# -*- mode: ruby -*-
# vi: set ft=ruby :

require 'fileutils'

VAGRANTFILE_API_VERSION = "2"

CLOUD_CONFIG_PATH = File.join(File.dirname(__FILE__), "user-data")

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


Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.ssh.insert_key = false

  # config.vm.box_url = "https://cloud-images.ubuntu.com/vagrant/trusty/current/trusty-server-cloudimg-amd64-vagrant-disk1.box"
  # config.vm.box = "trusty64"

  config.vm.box = "coreos-alpha"
  config.vm.box_version = ">= 308.0.1"
  config.vm.box_url = "http://alpha.release.core-os.net/amd64-usr/current/coreos_production_vagrant.json"

  config.vm.define "node1" do |node1|
      node1.vm.provision "shell", inline: $provision
      if File.exist?(CLOUD_CONFIG_PATH)
        node1.vm.provision :file, :source => "#{CLOUD_CONFIG_PATH}", :destination => "/tmp/vagrantfile-user-data"
        node1.vm.provision :shell, :inline => "mv /tmp/vagrantfile-user-data /var/lib/coreos-vagrant/", :privileged => true
      end
      node1.vm.hostname = "node1"
      node1.vm.network "private_network", ip: "172.20.20.10"
  end

  config.vm.define "node2" do |node2|
      if File.exist?(CLOUD_CONFIG_PATH)
        node2.vm.provision :file, :source => "#{CLOUD_CONFIG_PATH}", :destination => "/tmp/vagrantfile-user-data"
        node2.vm.provision :shell, :inline => "mv /tmp/vagrantfile-user-data /var/lib/coreos-vagrant/", :privileged => true
      end
      node2.vm.provision "shell", inline: $provision
      node2.vm.hostname = "node2"
      node2.vm.network "private_network", ip: "172.20.20.11"
  end

  config.vm.define "node3" do |node3|
      if File.exist?(CLOUD_CONFIG_PATH)
        node3.vm.provision :file, :source => "#{CLOUD_CONFIG_PATH}", :destination => "/tmp/vagrantfile-user-data"
        node3.vm.provision :shell, :inline => "mv /tmp/vagrantfile-user-data /var/lib/coreos-vagrant/", :privileged => true
      end
      node3.vm.provision "shell", inline: $provision
      node3.vm.hostname = "node3"
      node3.vm.network "private_network", ip: "172.20.20.12"
  end

  config.vm.provider :virtualbox do |vb|
    vb.customize ["modifyvm", :id, "--memory", "512"]
  end

end

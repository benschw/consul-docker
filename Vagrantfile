# -*- mode: ruby -*-
# vi: set ft=ruby :

require 'fileutils'

VAGRANTFILE_API_VERSION = "2"

CLOUD_CONFIG_PATH = File.join(File.dirname(__FILE__), "user-data")


Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.ssh.insert_key = false

  # config.vm.box_url = "https://cloud-images.ubuntu.com/vagrant/trusty/current/trusty-server-cloudimg-amd64-vagrant-disk1.box"
  # config.vm.box = "trusty64"

  config.vm.box = "coreos-alpha"
  config.vm.box_version = ">= 308.0.1"
  config.vm.box_url = "http://alpha.release.core-os.net/amd64-usr/current/coreos_production_vagrant.json"

  config.vm.define "node1" do |node1|
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
      node2.vm.hostname = "node2"
      node2.vm.network "private_network", ip: "172.20.20.11"
  end

  config.vm.define "node3" do |node3|
      if File.exist?(CLOUD_CONFIG_PATH)
        node3.vm.provision :file, :source => "#{CLOUD_CONFIG_PATH}", :destination => "/tmp/vagrantfile-user-data"
        node3.vm.provision :shell, :inline => "mv /tmp/vagrantfile-user-data /var/lib/coreos-vagrant/", :privileged => true
      end
      node3.vm.hostname = "node3"
      node3.vm.network "private_network", ip: "172.20.20.12"
  end

  config.vm.provider :virtualbox do |vb|
    vb.customize ["modifyvm", :id, "--memory", "512"]
  end

end

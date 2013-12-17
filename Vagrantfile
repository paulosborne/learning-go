# -*- mode: ruby -*-
# vi: set ft=ruby :
 
Vagrant.configure("2") do |config|
 
  config.vm.box = "golang-dev"
  config.vm.box_url = "http://files.vagrantup.com/precise64.box"
 
  config.vm.network "forwarded_port", guest: 80, host: 8080
 
  config.vm.synced_folder "./", "/home/vagrant/workspace"
 
  config.vm.provider :virtualbox do |vb|
    vb.gui = false
    vb.customize ["modifyvm", :id, "--memory", "2048"]
    vb.customize ["modifyvm", :id, "--cpus", "4"]
  end
 
  config.vm.provision :shell, :path => "vagrant-bootstrap.sh"
 
end

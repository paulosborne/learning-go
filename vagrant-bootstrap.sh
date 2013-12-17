#!/bin/bash
 
echo "Updating system"
apt-get update > /vagrant/vagrantup.log
apt-get install -y wget git build-essential mercurial bzr >> /vagrant/vagrantup.log
 
echo "Getting go source"
wget https://go.googlecode.com/files/go1.2.src.tar.gz >> /vagrant/vagrantup.log
tar -zxf go1.2.src.tar.gz >> /vagrant/vagrantup.log
 
echo "Building go compiler"
cd go/src >> /vagrant/vagrantup.log
./all.bash >> /vagrant/vagrantup.log

echo "Telling the system about the go compiler"
echo "" >> /home/vagrant/.profile
echo "export GOROOT=\$HOME/go" >> /home/vagrant/.profile
echo "export PATH=\$PATH:\$GOROOT/bin" >> /home/vagrant/.profile
echo "export GOPATH=/home/vagrant/workspace" >> /home/vagrant/.profile
echo "export PATH=\$PATH:\$GOPATH/bin" >> /home/vagrant/.profile

echo "Downloading Redis..."
wget http://download.redis.io/redis-stable.tar.gz >> /vagrant/vagrant.log
tar xvzf redis-stable.tar.gz
cd redis-stable
echo "Compiling Redis from source..."
make

echo "Removing non-required files"
rm -f go1.2.src.tar.gz
rm -f redis-stable.tar.gz
 
echo ""
echo "You are all set up, now run 'vagrant ssh' and start using your golang build environment"

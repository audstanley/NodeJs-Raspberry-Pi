#!/bin/bash
#Shell script created by @audstanley
#Special Thanks to Willi Thiel @ https://ni-c.github.io/heimcontrol.js/get-started.html
#This will work specifically for a RaspberryPi 2 ARM7.
#Willi's code installed a old version of NodeJS
#First, Installation of git.
sudo apt-get -y install git;
#Next, Creates directory for downloads, and downloads node 5.5.0
cd ~/ && mkdir temp && cd temp && wget https://nodejs.org/dist/latest-v5.x/node-v5.5.0-linux-armv7l.tar.gz;
tar -xzf node-v5.5.0-linux-armv7l.tar.gz;
#Remove the tar after extracing it.
sudo rm node-v5.5.0-linux-armv7l.tar.gz;
#This next line will copy Node over to the appropriate folder.
sudo mv node-v5.5.0-linux-armv7l/ /opt/nodejs/;
#This line will remove the nodeJs we downloaded.
sudo rm -R node-v5.5.0-linux-armv7l.tar.gz/* && sudo rmdir node-v5.5.0-linux-armv7l.tar.gz/;
#Create symlinks to node && npm
sudo ln -s /opt/nodejs/bin/node /usr/bin/node;
sudo ln -s /opt/nodejs/bin/node /usr/sbin/node;
sudo ln -s /opt/nodejs/bin/node /sbin/node;
sudo ln -s /opt/nodejs/bin/node /usr/local/bin/node;
sudo ln -s /opt/nodejs/bin/npm /usr/bin/npm;
sudo ln -s /opt/nodejs/bin/npm /usr/sbin/npm;
sudo ln -s /opt/nodejs/bin/npm /sbin/npm;
sudo ln -s /opt/nodejs/bin/npm /usr/local/bin/npm;
#Installation of MongoDB by downloading binaries, and compiling.
git clone git://github.com/RickP/mongopi.git;
#Download the dependencies to install mongoDB on R-Pi2
sudo apt-get install -y scons build-essential libboost-filesystem-dev libboost-program-options-dev libboost-system-dev libboost-thread-dev libboost-all-dev;
cd mongopi;
#SCons will build the binaries appropriate to the Pi.
scons;
#Here is the install of MongoDB wih SCons.
sudo scons --prefix=/opt/mongo install;
#SCons -c cleans up unwanted files.
scons -c;
#Creation of a path variable into the envoirnment
sudo su;
echo "PATH=$PATH:/opt/mongo/bin/" >> /etc/enviornment;
echo "export PATH" >> /etc/enviornment;
su pi;
#simlink mongo for shell access:
sudo ln -s /opt/mongo/bin/mongo /usr/bin/mongo;
sudo ln -s /opt/mongo/bin/mongoinfo /usr/bin/mongoinfo;
#add a new mongodb user
sudo useradd mongodb;
sudo mkdir /var/lib/mongodb;
sudo chown mongodb:mongodb /var/lib/mongodb;
sudo mkdir /etc/mongodb/;
sudo sh -c 'echo "dbpath=/var/lib/mongodb" > /etc/mongodb/mongodb.conf';
cd /etc/init.d;
#grab the mongodb.sh from a github gist
sudo wget -O mongodb https://gist.github.com/ni-c/fd4df404bda6e87fb718/raw/36d45897cd943fbd6d071c096eb4b71b37d0fcbb/mongodb.sh;
#make the shell script exicutable.
sudo chmod +x mongodb;
#Update rc.d for the mongodb.sh to run on startup
sudo update-rc.d mongodb defaults;
#finally, start mongoDB after all the installations are finished.
sudo service mongodb start;
#This whole process should take six hours on the raspberry pi 2.  Go get a coffee, or do this before going to bed.

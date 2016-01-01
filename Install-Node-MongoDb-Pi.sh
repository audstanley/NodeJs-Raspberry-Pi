#!/bin/bash
#Shell script created by @audstanley
#Special Thanks to Willi Thiel @ https://ni-c.github.io/heimcontrol.js/get-started.html
#This will work specifically for a RaspberryPi 2 ARM7.
#Willi's code installed a old version of NodeJS
#First, Creates directory for downloads, and downloads node 4.2.4
cd ~/ && mkdir temp && cd temp && wget https://nodejs.org/dist/latest/node-v4.2.4-linux-armv7l.tar.gz;
tar -xzf node-v4.2.4-linux-armv7l.tar.gz;
cd node-v4.0.0-linux-armv6l;
./configure;
make;
sudo make install;
#This next line will copy Node over to the appropriate folder.
sudo cp -R * /usr/local/;
cd ..;
#Installation of git.
sudo apt-get -y install git;
#Installation of MongoDB by downloading binaries, and compiling.
git clone https://github.com/skrabban/mongo-nonx86;
#Download the dependencies to install mongoDB on R-Pi2
sudo apt-get install -y scons build-essential libboost-filesystem-dev libboost-program-options-dev libboost-system-dev libboost-thread-dev libboost-all-dev;
cd mongo-nonx86;
#SCons will build the binaries appropriate to the Pi.
scons;
#Here is the install of MongoDB
sudo scons --prefix=/opt/mongo install;
#Creation of a path variable into the envoirnment
echo "PATH=$PATH:/opt/mongo/bin/" >> /etc/enviornment;
echo "export PATH" >> /etc/enviornment;
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
#This whole process should take many hours.  Go get a coffee, or do this before going to bed.

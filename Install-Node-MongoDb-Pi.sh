#!/bin/bash
#creates directory for downloads, and downloads node 4.2.4
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
sudo apt-get install -y scons build-essential libboost-filesystem-dev libboost-program-options-dev libboost-system-dev libboost-thread-dev libboost-all-dev;
cd mongo-nonx86;
scons;
sudo scons --prefix=/opt/mongo install;
#This whole process should take many hours.  Go get a coffee, or do this before going to bed. 
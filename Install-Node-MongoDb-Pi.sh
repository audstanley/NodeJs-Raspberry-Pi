#!/bin/bash
cd ~/ && mkdir temp && cd temp && wget http://node-arm.herokuapp.com/node_latest_armhf.deb;
sudo dpkg -i node_latest_armhf.deb;
sudo apt-get install git;
git clone https://github.com/skrabban/mongo-nonx86;
sudo apt-get install scons build-essential libboost-filesystem-dev libboost-program-options-dev libboost-system-dev libboost-thread-dev libboost-all-dev;
cd mongo-nonx86;
scons;
sudo scons --prefix=/opt/mongo install;
#!/bin/bash
# curl -G https://nodejs.org/dist/latest-v7.x/ | awk '{print $2}' | grep -P 'href=\"node-v7\.\d{1,}\.\d{1,}-linux-armv7l\.tar\.gz' | sed 's/href="//' | sed 's/<\/a>//' | sed 's/">.*//'
LINKTONODE=$(curl -G https://nodejs.org/dist/latest-v7.x/ | awk '{print $2}' | grep -P 'href=\"node-v7\.\d{1,}\.\d{1,}-linux-armv7l\.tar\.gz' | sed 's/href="//' | sed 's/<\/a>//' | sed 's/">.*//')
NODEFOLDER=$(echo $LINKTONODE | sed 's/.tar.gz/\//')
#Shell script created by @audstanley
#This will work specifically for a RaspberryPi 2/3 ARM7.
#Next, Creates directory for downloads, and downloads node 7.x
cd ~/ && mkdir tempNode && cd tempNode && wget https://nodejs.org/dist/latest-v7.x/$LINKTONODE
tar -xzf $LINKTONODE;
#Remove the tar after extracing it.
sudo rm $LINKTONODE;
#This next line will copy Node over to the appropriate folder.
sudo mv $NODEFOLDER /opt/nodejs/;
#This line will remove the nodeJs we downloaded.
sudo rm -R $LINKTONODE/* && sudo rmdir $LINKTONODE/;
#Create symlinks to node && npm
sudo ln -s /opt/nodejs/bin/node /usr/bin/node;
sudo ln -s /opt/nodejs/bin/node /usr/sbin/node;
sudo ln -s /opt/nodejs/bin/node /sbin/node;
sudo ln -s /opt/nodejs/bin/node /usr/local/bin/node;
sudo ln -s /opt/nodejs/bin/npm /usr/bin/npm;
sudo ln -s /opt/nodejs/bin/npm /usr/sbin/npm;
sudo ln -s /opt/nodejs/bin/npm /sbin/npm;
sudo ln -s /opt/nodejs/bin/npm /usr/local/bin/npm;
cd ../.. && rm -R tempNode/* && rmdir tempNode/

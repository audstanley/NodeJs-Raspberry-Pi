#!/bin/bash
# written by Richard Stanley (audstanley);
PICHIP=$(uname -m);
if [ "$EUID" -ne 0 ]
        then echo "You need to install as root by using sudo ./Install-Node.sh";
        exit
else LINKTONODE=$(curl -sG https://nodejs.org/dist/latest-v9.x/ | awk '{print $2}' | grep -P 'href=\"node-v9\.\d{1,}\.\d{1,}-linux-'$PICHIP'\.tar\.gz' | sed 's/href="//' | sed 's/<\/a>//' | sed 's/">.*//');
# curl -G https://nodejs.org/dist/latest-v9.x/ | awk '{print $2}' | grep -P 'href=\"node-v9\.\d{1,}\.\d{1,}-linux-armv9l\.tar\.gz' | sed 's/href="//' | sed 's/<\/a>//' | sed 's/">.*//'
NODEFOLDER=$(echo $LINKTONODE | sed 's/.tar.gz/\//');
#Next, Creates directory for downloads, and downloads node 8.x
cd ~/ && mkdir tempNode && cd tempNode && wget https://nodejs.org/dist/latest-v9.x/$LINKTONODE;
tar -xzf $LINKTONODE;
#Remove the tar after extracing it.
rm $LINKTONODE;
#remove older version of node:
rm -R -f /opt/nodejs/;
#remove symlinks
rm /usr/bin/node /usr/sbin/node /sbin/node /sbin/node /usr/local/bin/node /usr/bin/npm /usr/sbin/npm /sbin/npm /usr/local/bin/npm 2> /dev/null;
#This next line will copy Node over to the appropriate folder.
mv ~/tempNode/$NODEFOLDER /opt/nodejs/;
#This line will remove the nodeJs tar we downloaded.
rm -R -f ~/tempNode/$LINKTONODE/;
#Create symlinks to node && npm
sudo ln -s /opt/nodejs/bin/node /usr/bin/node; sudo ln -s /opt/nodejs/bin/node /usr/sbin/node; 
sudo ln -s /opt/nodejs/bin/node /sbin/node; sudo ln -s /opt/nodejs/bin/node /usr/local/bin/node; 
sudo ln -s /opt/nodejs/bin/npm /usr/bin/npm; 
sudo ln -s /opt/nodejs/bin/npm /usr/sbin/npm; sudo ln -s /opt/nodejs/bin/npm /sbin/npm; 
sudo ln -s /opt/nodejs/bin/npm /usr/local/bin/npm; 
rm -R -f /root/tempNode/;
#su pi;
#cd ~/ && rm -R NodeJs-Raspberry-Pi-Arm9/;
fi

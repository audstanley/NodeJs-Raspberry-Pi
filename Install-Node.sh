#!/bin/bash
# written by Richard Stanley (audstanley);
PICHIP=$(uname -m);
if [ "$PICHIP" == "aarch64" ]
then
        PICHIP="arm64";
fi
if [ "$EUID" -ne 0 ] 
then 
        echo "You need to install as root by using sudo ./Install-Node.sh";
        exit
else 
        cd /bin/;
        wget "https://raw.githubusercontent.com/audstanley/NodeJs-Raspberry-Pi/master/build/node-install-"$PICHIP;
        mv node-install-$PICHIP node-install
        chmod +x node-install;
        /bin/node-install -a;

fi

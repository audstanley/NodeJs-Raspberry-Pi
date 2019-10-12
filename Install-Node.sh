#!/bin/bash
# written by Richard Stanley (audstanley);
PICHIP=$(uname -m);
MAJOR_NODEJS_VERSION=v12
if [ "$EUID" -ne 0 ]
        then echo "You need to install as root by using sudo ./Install-Node.sh";
        exit
else cd /bin/;
rm node-install;
wget https://raw.githubusercontent.com/audstanley/NodeJs-Raspberry-Pi/master/node-install;
chmod +x node-install;
if [ $PICHIP = "armv6l" ]; then /bin/node-install -v 11; fi; # because there is no ARM6 build for version 12.
else node-install -a;

fi

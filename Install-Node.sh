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
        wget "https://raw.githubusercontent.com/audstanley/NodeJs-Raspberry-Pi/master/build/sha256-node-install-"$PICHIP".checksum";
        if [[ $(cat sha256-node-install-$PICHIP.checksum) = $(sha256sum node-install-$PICHIP | awk '{print $1}') ]]; then
                rm "sha256-node-install-"$PICHIP".checksum"
                mv node-install-$PICHIP node-install
                chmod +x node-install;
                /bin/node-install -a;
        else
                echo "The checksums for the downloaded binary do not match. Please open an issue on GitHub"
        fi

fi

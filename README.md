# Node-MongoDb-Pi
This is a simple shell script that will install NodeJs v7.x.  The "Latest version" of Node in the raspberry pi repo is some really old version like 1.4.0 or something. Too old, and this script will compile the latest version of 7.x<p>
You can install everything on your pi by going into you Raspberry Pi zero/1/2/3's terminal, and running: <p>
```sh
sudo apt-get install git;
git clone https://github.com/audstanley/NodeJs-Raspberry-Pi-Arm7;
cd NodeJs-Raspberry-Pi-Arm7;
chmod +x Install-Node.sh;
sudo ./Install-Node.sh;
echo "done installing node";
```
<p> 

Please note that this is for the raspberry pi zero/1/2/3 ARM6 or ARM7.

As new integer versions of NodeJs come out, I will update the script.

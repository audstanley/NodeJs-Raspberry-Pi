# Node-MongoDb-Pi
This is a simple shell script that will install NodeJs v7.x.  The "Latest version" of Node in the raspberry pi repo is some really old version like 1.4.0 or something. Too old, and this script will compile the latest version of 7.x<p>
The script will then compile the MongoDB binary, and the process will take about four to five hours.
You can install everything on your pi by:<p>
going into you Raspberry Pi 2 terminal, and running: <p>
 <code>sudo apt-get install git; </code><p>
git clone https://github.com/audstanley/NodeJs-Raspberry-Pi-Arm7; <p>
cd NodeJs-Raspberry-Pi-Arm7; <p>
sudo bash Install-Node-MongoDb-Pi-sh; </code><p> 

Please note that this is SPECIFICALLY for the raspberry pi 2/3 which run on an ARM7.

As new integer versions of NodeJs come out, I will update the script.

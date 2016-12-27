# NodeJs-Raspberry-Pi
This is a simple shell script that will install NodeJs v7.x.  The "Latest version" of NodeJs version of 7.x in the raspberry pi on every version of raspberry pis (zero|1|2|3) <p>
You can install everything on your pi by going into you Raspberry Pi zero|1|2|3's terminal, and running: <p>
```sh
sudo apt-get install git;
git clone https://github.com/audstanley/NodeJs-Raspberry-Pi;
cd NodeJs-Raspberry-Pi;
chmod +x Install-Node.sh;
sudo ./Install-Node.sh;
cd .. && rm -R -f NodeJs-Raspberry-Pi/;
node -v;
```
<p> 

Please note that this is for the raspberry pi zero|1|2|3 ARM6 or ARM7.

As new integer versions of NodeJs come out, I will update the script.

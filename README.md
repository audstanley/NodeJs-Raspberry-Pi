# Node-MongoDb-Pi
This is a simple shell script that will install NodeJs v5.3.0, git (if you don't have git), and MongoDB.  The "Latest version" of Node in the raspberry pi repo is some really old version like 1.4.0 or something. Too old, and this script will compile v5.3.0.<p>
The script will then compile the MongoDB binary, and the process will take about four to five hours.
Thank you Willi Thiel @ https://ni-c.github.io/heimcontrol.js/get-started.html
 All I did was place this in a shell script for easy downloading and install via a bash script, and write some extra code that will clean up the downloaded files. I also updated the version of node, as the version on his site is not optimized for the Raspberry Pi 2.  The cool thing about this script is that you can just walk away from your pi for four hours, and everything will be ready to go when you get back.  I will try to keep this script up-to-date as new versions of Node are released for the ARM7.<p> 
You can install everything on your pi by:<p>
going into you Raspberry Pi 2 terminal, and running: <p><code>
 $git clone git://github.com/audstanley/Node-MongoDb-Pi; <p>
 $cd Node-MongoDb-Pi; <p>
 $sudo bash Install-Node-MongoDb-Pi-sh <p> </code>
After the install you can control mongodb anywhere on the system with:<p>
$sudo service mongodb start|stop|restart
<p>
Please note that this is SPECIFICALLY for the raspberry pi 2, as NodeJS has a compiles ARM7 version available.  If you have an older Raspberry Pi, just change eveywhere in the script that says, "arm7" wih "arm6"
 

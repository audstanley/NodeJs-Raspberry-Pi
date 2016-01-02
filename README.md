# Node-MongoDb-Pi
This is a simple shell script that will install NodeJs v5.3.0, git (if you don't have git), and MongoDB.  The "Latest version" of Node in the raspberry pi repo is version 0.something.somthing. Too old, and this script will compile v4.2.4.<p>
The script then compiles the MongoDB binary, and the process will take about a couple of hours.
Thank you Willi Thiel @ https://ni-c.github.io/heimcontrol.js/get-started.html
 All I did was place this in a shell script for easy downloading and install via a bash script.<p>
 You can run this Bash script within the folder you extract the script with the command: <p>
 $sudo bash Install-Node-MongoDb-Pi-sh
 <p>
After the install you can control mongodb with:<p>
$sudo service mongodb start|stop|restart
<p>
Please note that this is SPECIFICALLY for the raspberry pi 2, as NodeJS has a compiles ARM7 version available.  If you have an older Raspberry Pi, just change eveywhere in the script that says, "arm7" wih "arm6"
 

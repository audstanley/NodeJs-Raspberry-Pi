# NodeJs-Raspberry-Pi
This is a simple shell script that will install NodeJs v9.x.  The "Latest version" of NodeJs version of 9.x in the raspberry pi on every version of raspberry pis (zero|1|2|3) This also happens to work on a Beaglebone and other ARM linux computers.<p>
You can install everything on your pi by going into you Raspberry Pi zero|1|2|3's terminal, and running: <p>

```sh
sudo su;
wget -O - https://raw.githubusercontent.com/audstanley/NodeJs-Raspberry-Pi/master/Install-Node.sh | bash;
node -v;

```
<p>

Please note that this should work for EVERY raspberry pi.

As new integer versions of NodeJs come out, I will update the script.  I even wrote some NodeJs code that checks for integer updates every hour, and will send me a push notification to my phone, so this script should be up to date for a LONG time.  Here is that code:
[checkNodeJsForLatestDistro](https://github.com/audstanley/checkNodeJsForLatestDistro)

<a href='https://ko-fi.com/A687KA8' target='_blank'><img height='36' style='border:0px;height:36px;' src='https://az743702.vo.msecnd.net/cdn/kofi4.png?v=f' border='0' alt='Buy Me a Coffee at ko-fi.com' /></a>

or stop by at my blog: [audstanley.com](http://www.audstanley.com)

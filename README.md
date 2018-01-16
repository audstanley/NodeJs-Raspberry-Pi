# NodeJs-Raspberry-Pi
This is a simple shell script that will install NodeJs v9.x.  The "Latest version" of NodeJs version of 9.x in the raspberry pi on every version of raspberry pis (zero|1|2|3) This also happens to work on a Beaglebone and other ARM6 or ARM7 linux computers.<p>
You can install everything on your pi by going into you Raspberry Pi zero|1|2|3's terminal, and running: <p>

```sh
sudo su;
wget -O - https://raw.githubusercontent.com/audstanley/NodeJs-Raspberry-Pi/master/Install-Node.sh | bash;
exit;
node -v;

```
<p>

Please note that this should work for EVERY raspberry pi.

As new integer versions of NodeJs come out, I will update the script.  I even wrote some NodeJs code that checks for integer updates every hour, and will send me a push notification to my phone, so this script should be up to date for a LONG time.  Here is that code:
[checkNodeJsForLatestDistro](https://github.com/audstanley/checkNodeJsForLatestDistro)

<a href='https://ko-fi.com/A687KA8' target='_blank'><img height='36' style='border:0px;height:36px;' src='https://az743702.vo.msecnd.net/cdn/kofi4.png?v=f' border='0' alt='Buy Me a Coffee at ko-fi.com' /></a>

or stop by at my blog: [audstanley.com](http://www.audstanley.com)


## Thank you so much:

[Phil](https://Ko-fi.com/home/coffeeshop?txid=ea3fc9e8-1e81-4198-a555-a595e3eeae76&mode=public&img=ogsomeoneboughtme)

[Shane](https://Ko-fi.com/home/coffeeshop?txid=c0356500-6d0d-452f-a93e-2974f8987e26&mode=public&img=ogsomeoneboughtme)

[Dan](https://Ko-fi.com/home/coffeeshop?txid=d486abde-c02d-454f-a602-025672a835c8&mode=public&img=ogsomeoneboughtme)


<center> <img src="Node-Pi-ASCII.png"> </center>

# NodeJs-Raspberry-Pi

## Latest NodeJS Install:
This is a simple shell script that will install NodeJs v11.x.  The "Latest version" of NodeJs version of 11.x in the raspberry pi on every version of raspberry pis (zero|1|2|3) This also happens to work on a Beaglebone, Nvidia TX2, and other ARM6, ARM7, and ARM64 linux computers.<p>

>First, you will need to install the latest version of NodeJs on your pi by going into you Raspberry Pi zero|1|2|3's terminal, and running: <p>

```sh
wget -O - https://raw.githubusercontent.com/audstanley/NodeJs-Raspberry-Pi/master/Install-Node.sh | sudo bash
node -v
# if you get an error that node did not install, it is because currently there is no ARM6l build for node 12
# (this only really effects the pi zero), but the node-install utility is there, so you can just sudo node-install -v 11
# for the time being
exit
```
<p>

## Specific Version of NodeJs:
Once you have run the initial installer, you can install different versions at any time by opening a new terminal and typing:
```sh
sudo node-install -v 4.9.1; # if you want to be specific
```
or
```sh
sudo node-install -v 4;
# then you will get prompted with which 
# specific version of 4 you wish to install
```
or
```sh
sudo node-install -v 5;
# then you will get prompted with which 
# specific version of 5 you wish to install
```
or
```sh
sudo node-install -v 6;
# then you will get prompted with which 
# specific version of 6 you wish to install
```
or
```sh
sudo node-install -v 7;
# then you will get prompted with which 
# specific version of 7 you wish to install
```
or
```sh
sudo node-install -v 8;
# then you will get prompted with which 
# specific version of 8 you wish to install
```
or
```sh
sudo node-install -v 9;
# then you will get prompted with which 
# specific version of 9 you wish to install
```
or
```sh
sudo node-install -v 10;
# then you will get prompted with which 
# specific version of 10 you wish to install
```

or
```sh
sudo node-install -v 11;
# then you will get prompted with which 
# specific version of 11 you wish to install
```

or
```sh
sudo node-install -v 12;
# then you will get prompted with which 
# specific version of 12 you wish to install
# there is currently no build for the ARM6l, so you will have to
# sudo node-install -v 11; # until they have an ARM6 build (only effects pi zeros, and pi1 models)
```

Please note that this will work for **EVERY raspberry pi**.

If you have installed any *global* npm modules, and change your systems version of nodejs with **node-install** then **node-install** will reinstall your npm modules for you.

<a href='https://ko-fi.com/A687KA8' target='_blank'><img height='36' style='border:0px;height:36px;' src='https://az743702.vo.msecnd.net/cdn/kofi4.png?v=f' border='0' alt='Buy Me a Coffee at ko-fi.com' /></a>

or stop by at my blog: [audstanley.com](http://www.audstanley.com)


## Always keep NodeJs up to date:

If you want to keep your Raspberry Pi up to date with the Latest Version of Node at all times, edit  your **/etc/crontab** file:

```bash
sudo nano /etc/crontab
```

Add this line to the bottom of your crontab:

```cron
0 22 1,7,14,21 * * root node-install -a;
```
This will update your raspberry pi with the latest version of NodeJs on days 1st, 7st, 14th, and 21st of every month at 10:00pm.

## Thank you so much:

[Hansie](https://Ko-fi.com/home/coffeeshop?txid=d40b4253-995a-4bd7-9eea-e56bf463ddbe&mode=public&img=ogsomeoneboughtme)

[Casper](https://Ko-fi.com/home/coffeeshop?txid=8504b181-2138-45e8-9a3a-66414828b024&mode=public&img=ogsomeoneboughtme)

[Ladvien](https://Ko-fi.com/home/coffeeshop?txid=025faeb3-d715-4d0c-8cf9-ce0f7bf4e1bf&mode=public&img=ogsomeoneboughtme)

[Phil](https://Ko-fi.com/home/coffeeshop?txid=ea3fc9e8-1e81-4198-a555-a595e3eeae76&mode=public&img=ogsomeoneboughtme)

[Shane](https://Ko-fi.com/home/coffeeshop?txid=c0356500-6d0d-452f-a93e-2974f8987e26&mode=public&img=ogsomeoneboughtme)

[Dan](https://Ko-fi.com/home/coffeeshop?txid=d486abde-c02d-454f-a602-025672a835c8&mode=public&img=ogsomeoneboughtme)

[Bob](https://Ko-fi.com/home/coffeeshop?txid=98c46da8-45bb-478d-b44e-df9257b87edb&mode=public&img=ogsomeoneboughtme)

[Anon](https://Ko-fi.com/home/coffeeshop?txid=e119b3ef-4171-4beb-a2dc-97d67bf09122&mode=public&img=ogsomeoneboughtme)

## Thank you ![](https://cdn4.iconfinder.com/data/icons/bug-fix/512/qa-quality-assurance-bug-15-128.png) (bug catchers)

[derpmarine168](https://github.com/derpmarine168)
  
[Kryten0807](https://github.com/Kryten0807)
  
## Traffic:

Thank you for checking out the project, and using the node-install utility.  This github project (as of 2019), reaches about an average of 10,000 people every year.

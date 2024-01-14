# Raspberry Pi Server

This guide is for a Raspberry Pi 1 Model B, Revision 2.0

## DietPi

- Download [DietPi](https://dietpi.com/) image.  Used DietPi_RPi-ARMv6-Bookworm.
- Follow [DietPi install instructions](https://dietpi.com/docs/install/) to flash and configure the Pi for the first time.
  - Configure host name as `berry`, because.
- `apt install -y vim less git avahi-daemon openssh-client rsync`
  - avahi-daemon makes the hostname discoverable
  - openssh-client is required for scp
- Maybe install `perf`?
  - `apt-cache search linux-tools` to find latest version
  - `apt install -y linux-perf-5.10`
  - Unfortunately `perf top` doesn't work on Pi 1 B
- dietpi user isn't needed
  - `passwd -d dietpi`
- Reduce swap file?  It is initially 1769 MiB, and this is too much for a 4GB SD card.
  - `/boot/dietpi/func/dietpi-set_swapfile 512`
- Set timezone
  - `sudo timedatectl set-timezone Australia/Perth`

## Realtek 8811CU WiFi

To get a Realtek 8811CU WiFi dongle to work, you'll first need to do the previous steps with another network adaptor, then:

- Log in as root
- `apt install -y bc raspberrypi-kernel-headers build-essential dkms`
- `git clone https://github.com/morrownr/8821cu-20210916.git`
- `cd 8821cu-20210916`
- `chmod +x dkms-install.sh`
- `./dkms-install.sh`  (Building the driver took some hours)
- `modprobe 8821cu`
- Open `dietpi-config` and configure SSIDs

Try `ssh root@berry.local`

## Connecting to LEDs

[Pin layout](https://www.raspberry-pi-geek.com/howto/GPIO-Pinout-Rasp-Pi-1-Rev1-and-Rev2)

- Wire ground
- If using SPI:
  - Wire data to GPIO10
  - Edit `/boot/config.txt` and replace `#dtparam=spi=off` with `dtparam=spi=on`
- If using PWM: (don't, because it doesn't work as a non-root user)
  - Wire data to GPIO18
  - Edit `/etc/modprobe.d/snd-blacklist.conf` and add `blacklist snd_bcm2835`
  - Edit `/boot/config.txt` and replace `dtparam=audio=off` with `dtparam=audio=on`

## Service user

Add user for service
- Log in as root
- `useradd -m xyzmas`
- `usermod -aG gpio,spi xyzmas` to allow user to use SPI
- `chsh -s /bin/bash xyzmas`
- `su - xyzmas`
- `mkdir -m700 ~/.ssh`
- `vi ~/.ssh/authorized_keys`
- Paste you ssh key there, `ZZ`
- `exit`, `exit`
- You can now `ssh xyzmas@berry.local`

Does it need an SSH key?
- `dropbearkey -t rsa -f ~/.ssh/id_rsa`
- `dropbearkey -y -f ~/.ssh/id_rsa | grep "^ssh-rsa " > ~/.ssh/id_rsa.pub`

## Service

First time, or when the service changes:
- From this directory
- `rsync xyzmas.service root@berry.local:/lib/systemd/system`
- `ssh root@berry.local "systemctl daemon-reload && systemctl enable xyzmas.service"`

Grant user access to restart itself:
- From this directory
- `rsync xyzmas.sudoers root@berry.local:/etc/sudoers.d/xyzmas`

Build and deploy server changes:
- From this directory
- `make deploy && make restart`

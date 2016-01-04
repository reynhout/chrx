# chrx

Install Linux onto your Chromebook. Dual-boot alongside ChromeOS for maximum flexibility.


| | |
| ------------ | ---------- |
|**works on**|Compatible Chromebook models. See [chromebooks](#chromebooks).|
|**installs**|Several Linux distributions. See [operating systems](#operating systems) and [recommendations](#recommendations).|



**[chrx.org](https://chrx.org/)**
<br />
**[github.com/reynhout/chrx](https://github.com/reynhout/chrx)**

## status

**Version 2.1** Updated for GalliumOS and Ubuntu 15.10. See [changelog](#changelog).


<a name="usage"></a>
## usage

Installing Linux via **chrx** onto a new (or freshly recovered) Chromebook
is a two-phase process:

- The first phase reserves space on your SSD or other storage device for
the new operating system, **and then reboots**.
- The second phase installs your chosen distribution, and configures the
new system according to your selected options.

If you reinstall later, or switch to a another distribution, **chrx** will
skip directly to phase two.


### step-by-step

1. Enable [Developer Mode](http://www.chromium.org/chromium-os/developer-information-for-chrome-os-devices) (process is model-specific; for Acer C720, press `ESC+F3(Refresh)+Power`), then reboot
1. Load ChromeOS by pressing `CTRL+D` at the white "OS verification is OFF" screen
1. Configure your Wi-Fi network, if necessary
1. Switch to Virtual Terminal (VT) 2 by pressing `CTRL+ALT+F2(top row right arrow)`
1. Log in as user `chronos` (no password) to enter `chronos@localhost` shell
1. Run **chrx**: `curl -Os https://chrx.org/go && sh go` (see [options](#options))
1. Follow on-screen instructions to prepare your Chromebook for installation
1. Reboot, then repeat steps 2-7 to install and configure your new system


<a name="options"></a>
### options

**chrx** can accept several command-line options:

```
Usage: chrx [ option ... ]

Options
   -d DISTRIBUTION OS-specific distribution to install [lubuntu]
                   (galliumos, ubuntu, lubuntu, xubuntu, kubuntu, edubuntu)
   -e ENVIRONMENT  distribution-specific environment [desktop]
                   (desktop, minimal, standard, server)
   -r RELEASE      distribution release number or name []
                   (lts, latest, dev, 15.04, 15.10, vivid, wily, etc)
   -a ARCH         processor architecture (i386, amd64) [amd64]
   -t TARGETDISK   target disk (/dev/mmcblk1, /dev/sdb, etc) []
   -p PACKAGE      additional packages, quote or repeat for multiple []
                   (kodi, minecraft, steam, etc)
   -H HOSTNAME     hostname for new system [chrx]
   -U USERNAME     username of first created user [chrx]
   -L LOCALE       locale for new system [en_US.UTF-8]
   -Z TIMEZONE     timezone for new system, Eggert convention [America/New_York]
                   (America/San_Francisco, Europe/Amsterdam, Etc/UTC, etc)
   -n              disable success/failure notifications
   -s              skip all customization, install stock OS only
   -y              run non-interactively, take defaults and do not confirm
   -v              increase output verbosity
   -h              show this help
  
  Default values are shown in brackets, e.g.: [default].
  
  If TARGETDISK is not specified, chrx will select the internal SSD.

```

<a name="packages"></a>
### packages

**chrx** can install additional software packages after installing
your new operating system, using the `-p PACKAGE` option.

You can install any package in the Ubuntu repositories via this
method, plus a few non-Ubuntu packages for which **chrx** has
special handling, and some aliases for convenience:

- `minecraft` installs [Minecraft](https://minecraft.net/)
- `steam` installs [Steam](http://store.steampowered.com/about/)
- `kodi` installs [Kodi Media Center](http://kodi.tv/about/)
- `admin-misc` is an alias for `"ssh tmux rsync vim"`
- `dev-misc` is an alias for `"arduino geany geany-plugins ruby"`

To install multiple packages from the **chrx** command line, you
can repeat the `-p PACKAGE` option as many times as you need, or
you can quote the argument, e.g.: `-p "gimp blender inkscape"`.


### examples

[Lubuntu](http://lubuntu.net/) Desktop (latest):

`curl -Os https://chrx.org/go && sh go`

[GalliumOS](https://galliumos.org/) Desktop (latest), verbosely:

`curl -Os https://chrx.org/go && sh go -d galliumos -v`

[GalliumOS](https://galliumos.org/) Desktop (latest), plus packages:

`curl -Os https://chrx.org/go && sh go -d galliumos -p "minecraft steam kodi"`

[Ubuntu](https://ubuntu.com/) Standard (good for servers), version 15.04, system name `hal`, first user `dave`, including some administrative tools:

`curl -Os https://chrx.org/go && sh go -d ubuntu -e standard -r 15.04 -H hal -U dave -p admin-misc`


### advanced usage

You may choose to host or cache these installation files yourself.
There are many good reasons to do so, especially if you'll be doing
a large number of installations. However, setup can be somewhat more
complicated, and instructions are outside the scope of this README.

To point **chrx** at your cache, just set the `CHRX_WEB_ROOT`
environment variable before running the `go` script, like this:

```
export CHRX_WEB_ROOT="http://myserver/chrx"
curl -O $CHRX_WEB_ROOT/go && sh go
```

<a name="compatibility"></a>
## compatibility


<a name="chromebooks"></a>
### chromebooks
status            |chromebook                           |hwid   |notes
:----------------:|-------------------------------------|-------|------------------
:white_check_mark:|Acer C720                            |PEPPY  |Best-tested
:white_check_mark:|Acer C720P                           |PEPPY  |
:white_check_mark:|Acer Chromebook 11 C740              |PAINE  |
:white_check_mark:|Acer Chromebook 15 C910              |YUNA   |
:white_check_mark:|Acer Chromebook 15 CB5-571           |YUNA   |
:white_check_mark:|Dell Chromebook 11                   |WOLF   |
:white_check_mark:|Dell Chromebook 13                   |LULU   |Requires [firmware feature update](https://johnlewis.ie/custom-chromebook-firmware/rom-download/)
:white_check_mark:|Google Pixel (2013)                  |LINK   |
:white_check_mark:|Google Pixel (2015)                  |SAMUS  |
:white_check_mark:|HP Chromebook 14                     |FALCO  |
:white_check_mark:|Toshiba CB30/CB35 Chromebook (2014)  |LEON   |
:white_check_mark:|Toshiba CB30/CB35 Chromebook 2 (2015)|GANDOF |Requires [firmware feature update](https://johnlewis.ie/custom-chromebook-firmware/rom-download/)
:white_check_mark:|Other Haswell Chromebooks            |       |Expected to work, but untested
:white_check_mark:|Other Broadwell Chromebooks          |       |Expected to work, some models will require [firmware feature update](https://johnlewis.ie/custom-chromebook-firmware/rom-download/)
:question:        |Other Sandy/Ivy Bridge Chromebooks   |       |Requires SeaBIOS with Legacy Boot capability
:question:        |Other Intel Chromebooks              |       |Requires SeaBIOS with Legacy Boot capability
:x:               |Bay Trail Chromebooks                |       |Requires compatible firmware, presently unavailable
:x:               |ARM Chromebooks                      |       |ARM support is very unlikely
:x:               |HP Chromebook 11                     |SPRING, SKATE|(previously erroneously marked possible)


<a name="operating systems"></a>
### operating systems

status| OS  | distribution | notes
:----:| --- | ------------ | -----
:white_check_mark:|Linux|[GalliumOS](https://galliumos.org/)|Derived from Xubuntu, developed specifically for compatibility and optimized performance on Chromebook hardware.
:white_check_mark:|Linux|[Lubuntu](http://lubuntu.net/)|A light-weight variant of Ubuntu, with the LXDE desktop environment.
:white_check_mark:|Linux|[Xubuntu](http://xubuntu.org/)|A light-weight variant of Ubuntu, with the Xfce desktop environment.
:white_check_mark:|Linux|[Kubuntu](http://kubuntu.org/)|Ubuntu with the KDE desktop environment.
:white_check_mark:|Linux|[Edubuntu](http://edubuntu.org/)|Full Ubuntu plus application bundles used in educational settings.
:white_check_mark:|Linux|[Ubuntu](https://ubuntu.com/)|The standard full Ubuntu distro.
:x:|FreeBSD||Work in progress!


<a name="recommendations"></a>
### recommendations

Chromebooks perform best with lighter-weight operating systems and desktop environments, and they often require updated kernel drivers to support their new and tightly integrated hardware.

Selecting a distribution which meets these needs is therefore an important part of Linux-on-Chromebook happiness. While any updated distro will work for ordinary tasks, there are a few that stand out:

- **GalliumOS** is optimized specifically for Chromebooks. It scores well on all metrics, looks great, and installs quickly. Some memory-hungry applications (e.g. Steam games) perform *best* on GalliumOS thanks to careful optimizations. GalliumOS is in BETA but is available for testing now.
- **Lubuntu** also scores and performs well. It installs slightly smaller and uses significantly less RAM than other distros. Lubuntu is currently the default distro installed by **chrx**.
- **Xubuntu** is another good choice. It's a bit heavier-weight than Lubuntu, but still performs very well.
- I would not choose standard, full, Ubuntu for a Chromebook. It is perfectly usable, bit it's heavier and suffers in performance, without offering any important benefits. Memory use starts higher and increases much more quickly as you use the desktop apps (not reflected in measurements below). If your Chromebook model has 4GB of RAM, the performance differences are reduced but not eliminated.


#### sample measurements

distribution&sup1; | disk space&sup2; | RAM use&sup3; | install time&#8308; | version | recommended? |
--------- | ----- | ----- | ------- | -------- |:---:|
GalliumOS | 2.9GB | 311MB | 9 mins  | 1.0 BETA | :white_check_mark: |
Lubuntu   | 2.7GB | 227MB | 18 mins | 15.10    | :white_check_mark: |
Xubuntu   | 3.0GB | 360MB | 22 mins | 15.04    | :white_check_mark: |
Ubuntu    | 3.5GB | 440MB | 28 mins | 15.04    | :x: |
Kubuntu   | 4.2GB | 613MB |         | 15.10    | :x: |


1. All distributions were installed with the `desktop` environment option.
1. Disk space can be reduced by removing unwanted packages. The number shown reflects the default install for the desktop environment.
1. RAM use is measured after graphical login, connecting to Wi-Fi, and opening one window of the default Terminal program to run `/usr/bin/free` after a couple minutes for the system to stabilize. The number shown is an average of several tests, and variance is very low (2-3%).
1. Installation time will vary greatly depending on your Internet connection, but the ratios should be representative.


## test suite

"Working" is defined as:

- system boots cleanly and quickly
- installation remnants are cleaned up
- swap and compressed RAM are enabled
- proper drivers are properly loaded
- trackpad works (standard & australian)
- trackpad settings are usable
- audio works, including after sleep/wake
- wireless works, including after sleep/wake
- function keys for backlight are functional
- function keys volume control are functional
- microphone input works
- webcam input works
- power management sleeps system when lid closed
- power management wakes system when lid opened
- no user configuration is required for basic use

This list might evolve. Your input is welcome!

## evolution

### chrx past

**chrx** began as an updated and enhanced version of
[ChrUbuntu](http://chromeos-cr48.blogspot.fr/2013/10/chrubuntu-for-new-chromebooks-now-with.html),
and still retains some original ChrUbuntu code (especially the
disk partitioning bits).


### chrx present

**chrx** has been used to install Linux on thousands of Chromebooks.
Users and discussion can be found on [/r/chrubuntu](https://www.reddit.com/r/chrubuntu).


### chrx future

I'd like to test on a wider variety of hardware, and to install other
Linux distributions.

Support for FreeBSD is coming. See http://blog.grem.de/pages/c720.html if you can't wait.


## alternatives

**chrx** is a command-line installer which requires requires no physical
media or other preparation to install. It allows you to dual-boot, so you
can choose Linux or ChromeOS each time you turn on your Chromebook. This
is a flexible setup, well-suited for many users, but of course not all.

Consider these alternatives:

- Hugh Greenberg's [Distroshare](https://www.distroshare.com/) has nicely
updated ISOs (for Ubuntu and many other Linux distros!), which can be
installed from USB/SD flash RAM. This method completely removes ChromeOS
from your Chromebook, and devotes your entire SSD to Linux.
  - See also John Lewis's [Alternate Firmware](https://johnlewis.ie/custom-chromebook-firmware/rom-download/) options for Chromebooks that do not support SeaBIOS Legacy Boot
with stock firmware.
- [Crouton](https://github.com/dnschneid/crouton) allows you to run ChromeOS
and Linux simultaneously, instead of dual-booting like **chrx** or ChrUbuntu.
This arrangement has a few drawbacks, but if you spend most of your time in
ChromeOS and your Linux needs are limited, it should serve well.
- The original [ChrUbuntu](http://chromeos-cr48.blogspot.fr/2013/10/chrubuntu-for-new-chromebooks-now-with.html) has been tested on a wide variety of hardware. Unfortunately, it is now significantly outdated, and fails to install Ubuntu
15.04 and newer properly. 14.10 and older should install successfully via
ChrUbuntu, but they will require significant additional configuration to work
well. If your Chromebook is older and unsupported by **chrx**, give ChrUbuntu
a try.


## notes on security and privacy

Running code from the net is always an act that requires careful thought.
**chrx** can be run directly from the net, and by default will download
additional code via the same mechanisms. Any of these downloads could be
misdirected or compromised. Downloading over an unsecured network (e.g.
public Wi-Fi) raises the likelihood of such malfeasance, but it can never
be fully eliminated.

If these are concerns of yours, you can mitigate your risks by auditing
all of the code involved, comparing checksums of downloaded packages, and
hosting local caches (see [advanced usage](#advanced usage)).

Also, **chrx** "pings home" on every install to report success or failure.
This ping includes **no personal information**, only data that might be
useful for investigating failures.

Log entries created by these pings look like this:

```
17.x.x.x - - [01/Jun/2015:07:37:00 +0000] "GET /end_ok HTTP/1.1" 200 0 "-"
  "chrx/1.1.1 hw=PEPPY_C6A-V7C-A2C sw=linux,ubuntu,latest,-,15.04,amd64,ubuntu-minimal" "-"
```

`hw` is a hardware ID that corresponds to your model of Chromebook
(**not** a serial number).

`sw` combines a few of the command-line settings (or defaults) that you
used to run **chrx**.

If this level of information sharing makes you uncomfortable, the behaviour
can be disabled with the `-n` switch.


## meta, obligatory

**chrx** is pronounced "marshmallow".


## thanks

To Jay Lee for [ChrUbuntu](http://chromeos-cr48.blogspot.fr/2013/10/chrubuntu-for-new-chromebooks-now-with.html), to [/r/chrubuntu](https://www.reddit.com/r/chrubuntu) for assembling links to tons of helpful resources, and to the dozens of people who found answers and solved problems before I even started looking.


<a name="changelog"></a>
## changelog

- **1.0** (20141223)
- **1.1** (20150504): add support for Ubuntu 15.04
  - **1.1.1** (20150508): add "-r RELEASE" option; validate some input
  - **1.1.2** (20151005): update Ubuntu "trusty" to 14.04.3; add recognized HWIDs (PEPTO, LINK, SAMUS, LEON, PAINE, YUNA, SPRING, SKATE, FALCO, WOLF); always verify chrx.org certificates
- **2.0** (20151025): add GalliumOS support; add support for Ubuntu 15.10; add detection and installation prognosis for all known ChromeOS devices; add "-d DISTRIBUTION" and "-e ENVIRONMENT" options; remove "-m METAPACKAGE" option; remove "-i IMAGE" option, make RELEASE smarter; work around `systemd` conflict; refactor code into functions to facilitate multiple distros and future operating systems
  - **2.0.1** (20151113): update core image pathname for GalliumOS
  - **2.0.2** (20151118): update some HWIDs
  - **2.0.3** (20151119): bugfix: issue #4, parted and partprobe removed from ChromeOS
  - **2.0.4** (20151120): bugfix: issue #5, "-r RELEASE" handling failing for some values of RELEASE
  - **2.0.5** (20151212): add first user to important groups; use generic coreimage for GalliumOS
  - **2.0.6** (20151214): bugfix: issue #7, add GalliumOS hwspecific pkgs properly
  - **2.0.7** (20151214): update detection for all known ChromeOS devices; improve prognosis descriptions
  - **2.0.8** (20160102): add CHRX_NO_REBOOT env var for use with https://github.com/MattDevo/scripts
- **2.1** (20150102): add support for "-p PACKAGE" to install additional package(s)

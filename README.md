# chrx

Install Linux onto your Chromebook. Dual-boot alongside ChromeOS for maximum flexibility.


| | |
| ------------ | ---------- |
|**works on**|Most Chromebook models. See [chromebooks](#chromebooks).|
|**installs**|Several Linux distributions. See [operating systems](#operating systems) and [recommendations](#recommendations).|



**[chrx.org](https://chrx.org/)**
<br />
**[github.com/reynhout/chrx](https://github.com/reynhout/chrx)**

## status

**Version 2.3.2** Add support for Fedora; Updated for Ubuntu 16.10; Installs GalliumOS by default. See [changelog](#changelog).


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


<a name="step-by-step"></a>
### step-by-step

1. Enable [Developer Mode](http://www.chromium.org/chromium-os/developer-information-for-chrome-os-devices) (process is model-specific; for Acer C720, press `ESC+F3(Refresh)+Power`), then reboot
1. Load ChromeOS by pressing `CTRL+D` at the white "OS verification is OFF" screen
1. Configure your Wi-Fi network if necessary, then log in (Guest account is fine)
1. Open the ChromeOS Terminal by pressing `CTRL+ALT+T`, and enter `shell` at the prompt
1. Update firmware, if necessary (*required* for Bay Trail and Braswell models, *recommended* for Broadwell and Skylake models, *optional* for Haswell models -- see [chromebooks](#chromebooks))
1. Run **chrx**: `cd ; curl -Os https://chrx.org/go && sh go` (see [options](#options))
1. Follow on-screen instructions to prepare your Chromebook for installation
1. Reboot, then repeat steps 2-4 and 6 to install and configure your new system


<a name="options"></a>
### options

**chrx** can accept several command-line options:

```
Usage: chrx [ option ... ]

Options
   -d DISTRIBUTION OS-specific distribution to install [galliumos]
                     galliumos, ubuntu, lubuntu, xubuntu, kubuntu, edubuntu,
                     fedora
   -e ENVIRONMENT  distribution-specific environment [desktop]
                     galliumos: desktop
                     ubuntu etc: desktop, minimal, standard, server
                     fedora: desktop, workstation, kde, xfce, lxde, mate,
                       cinnamon, sugar
   -r RELEASE      distribution release number or name [latest]
                     galliumos: latest, 2.0, xenon, nightly
                     ubuntu etc: latest, lts, dev, 16.04.1, xenial, etc
                     fedora: latest, 23, 24, 25
   -a ARCH         processor architecture (i386, amd64) [amd64]
   -t TARGETDISK   target disk (/dev/mmcblk1, /dev/sdb, etc) []
   -p PACKAGE      additional packages to install, may repeat []
                     kodi, minecraft, steam, etc, see chrx.org for more
                     (not yet supported on fedora)
   -H HOSTNAME     hostname for new system [chrx]
   -U USERNAME     username of first created user [chrx]
   -L LOCALE       locale for new system [en_US.UTF-8]
   -Z TIMEZONE     timezone for new system, Eggert convention [America/New_York]
                     America/San_Francisco, Europe/Amsterdam, Etc/UTC, etc
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
- `chrome` installs Google Chrome
- `admin-misc` is an alias for `"ssh tmux rsync vim"`
- `dev-misc` is an alias for `"arduino geany geany-plugins ruby"`

To install multiple packages from the **chrx** command line, you
can repeat the `-p PACKAGE` option as many times as you need, or
you can quote the argument, e.g.: `-p "gimp blender inkscape"`.


### examples

[GalliumOS](https://galliumos.org/) Desktop (latest), verbosely:

`cd ; curl -Os https://chrx.org/go && sh go -v`

[GalliumOS](https://galliumos.org/) Desktop (latest), plus packages:

`cd ; curl -Os https://chrx.org/go && sh go -p "minecraft steam kodi"`

[Lubuntu](http://lubuntu.net/) Desktop (latest):

`cd ; curl -Os https://chrx.org/go && sh go -d lubuntu`

[Ubuntu](https://ubuntu.com/) Standard, version 16.04, system name `hal`, first user `dave`, including some administrative tools:

`cd ; curl -Os https://chrx.org/go && sh go -d ubuntu -e standard -r 16.04 -H hal -U dave -p admin-misc`


### advanced usage

You may choose to host or cache these installation files yourself.
There are many good reasons to do so, especially if you'll be doing
a large number of installations. However, setup can be somewhat more
complicated, and instructions are outside the scope of this README.

To point **chrx** at your cache, just set the `CHRX_WEB_ROOT`
environment variable before running the `go` script, like this:

```
export CHRX_WEB_ROOT="http://myserver/chrx"
cd ; curl -O $CHRX_WEB_ROOT/go && sh go
```

<a name="compatibility"></a>
## compatibility


<a name="chromebooks"></a>
### chromebooks
status            |CPU family                           |notes
:----------------:|-------------------------------------|------------------
:white_check_mark:|Intel Haswell                        |[Firmware update](https://mrchromebox.tech/#fwscript) available
:white_check_mark:|Intel Broadwell                      |[Firmware update](https://mrchromebox.tech/#fwscript) *recommended*
:white_check_mark:|Intel Skylake                        |[Firmware update](https://mrchromebox.tech/#fwscript) *recommended*
:white_check_mark:|Intel Bay Trail                      |[Firmware update](https://mrchromebox.tech/#fwscript) **required**
:white_check_mark:|Intel Braswell                       |[Firmware update](https://mrchromebox.tech/#fwscript) **required**
:question:        |Intel Sandy/Ivy Bridge               |Requires SeaBIOS with Legacy Boot capability
:question:        |Intel Pineview                       |Requires SeaBIOS with Legacy Boot capability
:x:               |ARM                                  |ARM support is very unlikely

If you do not know the CPU in your device, check here: https://wiki.galliumos.org/Hardware_Compatibility

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
:white_check_mark:|Linux|[Fedora](https://fedoraproject.org/)|New 20161121!
:x:|FreeBSD||Work in progress!


<a name="recommendations"></a>
### recommendations

Chromebooks perform best with lighter-weight operating systems and desktop environments, and they often require updated kernel drivers to support their new and tightly integrated hardware.

Selecting a distribution which meets these needs is therefore an important part of Linux-on-Chromebook happiness. While any updated distro will work for ordinary tasks, there are a few that stand out:

- **GalliumOS** is optimized specifically for Chromebooks. It scores well on all metrics, looks great, and installs quickly. Some memory-hungry applications (e.g. Steam games) perform *best* on GalliumOS thanks to careful optimizations. GalliumOS is the default distro installed by **chrx**.
- **Lubuntu** also scores and performs well. It uses significantly less RAM than other distros.
- **Xubuntu** is another good choice. It's a bit heavier-weight than Lubuntu, but still performs very well.
- **Fedora** comes in several "spins" (desktop environments, selected with `-e ENVIRONMENT`), some of which (lxde) are lightweight, and some of which (desktop (gnome), default) are heavier. A few sample spins have been added to measurements below.
- I would not choose standard, full, Ubuntu for a Chromebook. It is perfectly usable, bit it's heavier and suffers in performance, without offering any important benefits. Memory use starts higher and increases much more quickly as you use the desktop apps (not reflected in measurements below). If your Chromebook model has 4GB of RAM, the performance differences are reduced but not eliminated.


#### sample measurements

distribution&sup1; | disk space&sup2; | RAM use&sup3; | install time&#8308; | recommended? |
------------------ | ----- | ----- | ------- |:---:|
GalliumOS 2.0      | 2.5GB | 291MB | 9 mins  | :white_check_mark: |
GalliumOS 1.0      | 2.8GB | 287MB | 10 mins | :white_check_mark: |
Lubuntu 15.10      | 2.7GB | 227MB | 18 mins | :white_check_mark: |
Lubuntu 16.04      | 3.1GB | 185MB | 19 mins | :white_check_mark: |
Xubuntu 15.04      | 3.0GB | 360MB | 22 mins | :white_check_mark: |
Ubuntu 15.04       | 3.5GB | 440MB | 28 mins | :x: |
Kubuntu 15.10      | 4.2GB | 613MB |         | :x: |
Fedora 24 (lxde)   | 2.9GB | 182MB | 20 mins | :white_check_mark: |
Fedora 24 (cinnamon)| 3.8GB | 384MB | 27 mins | :white_check_mark: |
Fedora 24          | 4.5GB | 647MB | 27 mins | :x: |


1. All distributions were installed with the `desktop` environment option, except where noted.
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
17.x.x.x - - [01/May/2016:07:37:00 +0000] "GET /end_ok HTTP/1.1" 200 0 "-"
  "chrx/2.2.3 hw=PEPPY_C6A-V7C-A2C sw=linux,galliumos-desktop,latest,2.0,amd64" "-"
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
- **2.1** (20160103): add "-p PACKAGE" option to install additional packages
  - **2.1.1** (20160120): update URL for GalliumOS coreimage; make sure util pkgs are added
  - **2.1.2** (20160130): add parsing for "-r nightly" (GalliumOS only, installs nightly build); log chrx command line for debugging; add first user to groups more quietly
- **2.2** (20160304): switch default distribution to GalliumOS
  - **2.2.1** (20160316): bugfix: issue #12, errors installing to external media
  - **2.2.2** (20160420): retry/resume failed image downloads; add new HWIDs
  - **2.2.3** (20160426): do not drop to shell before reboot; do not retry coreimage downloads; update steam install for xenial; update docs for Ubuntu 16.04
  - **2.2.4** (20160505): add Google Chrome to installable packages; add new HWIDs, update others
  - **2.2.5** (20160512): update Ubuntu base/core image URL (thanks arsfeld)
  - **2.2.6** (20160619): hide eMMC partitions properly (thanks gmykhailiuta); improve -r RELEASE handling for GalliumOS; add preliminary handling for running under non-ChromeOS
  - **2.2.7** (20160810): use version-dependent Ubuntu URL to match new Canonical schemes; update Ubuntu "trusty" to 14.04.5
  - **2.2.8** (20161002): add support for new GalliumOS hardware-specific pkgs: braswell, skylake, samus
- **2.3** (20161121): add support for Fedora! thanks @jedigo!
  - **2.3.1** (20161208): Fedora: add `-p` support, add latest to auto-detection, add nonfree codecs (thx @jedigo); GalliumOS: use chrx GRUB config; all: add more hidden mmcblk0 partitions, update GRUB config
  - **2.3.2** (20161222): add first user to groups individually in case selected distro/metapackage/spin does not include all (fixes #30)

<!-- don't forget to update the version number in chrx-install! -->

# chrx

**[chrx.org](https://chrx.org/)**
<br />
**[github.com/reynhout/chrx](https://github.com/reynhout/chrx)**

Use **chrx** to install Linux onto your Chromebook.

Presently, this means Ubuntu (or a variant) onto your Acer C720.

**chrx** installs OS components from the net. By default, package files
are downloaded from vendor sites, and several configuration tweaks are
downloaded from a secure mirror of this repository.

**chrx** is run directly from your Chromebook's `chronos@localhost` shell,
accessed via `CTRL-ALT-(top row right arrow)` and then logging in as
`chronos`, with no password. You must first enable
[Developer Mode](http://www.chromium.org/chromium-os/developer-information-for-chrome-os-devices),
and configure your network from the ChromeOS setup screen.


## status

Version 1.1.1: Updated for Ubuntu 15.04 (20150504)

status| chromebook | unix | notes
:----:| ---------- | ---- | -----
:white_check_mark:|Acer C720|Ubuntu Linux (15.04, 14.10, 14.04, development)|I use `lubuntu-desktop`
:question:|Other Haswell Chromebooks|Ubuntu Linux|These should work but might require other config tweaks
:x:|(any of the above)|Other UNIX or Linux distributions|Hopes are high
:x:|ARM Chromebooks|(any)|ARM support is spotty, these might never work

**Best-tested:** Lubuntu 14.10 and 15.04 on Acer C720

## usage

**Be sure to back up any valuable data in ChromeOS or Linux before
running chrx!**

Run the following command from your `chronos@localhost` shell:

`curl -Os https://chrx.org/go && sh go`
<!-- yes, by all means read the all of the code first! -->

Then follow the instructions from there.

When installing onto the internal SSD of a clean Chromebook, **chrx**
will repartition the drive to allocate space for the new operating system,
and will reboot after this first step. After reboot, run **chrx** again
(with the same command line) to perform the installation.
Subsequent installs will not require repartitioning or rebooting.

**chrx** can accept several command-line arguments, all are optional:

```
Usage: chrx [ option ... ]

Options
   -m METAPACKAGE  OS-specific metapackage to install [ubuntu-desktop]
                   (ubuntu-desktop, ubuntu-minimal, ubuntu-standard,
                   lubuntu-desktop, kubuntu-desktop, xubuntu-desktop,
                   edubuntu-desktop, edubuntu-server)
   -a ARCH         processor architecture (i386, amd64) [amd64]
   -t TARGETDISK   target disk (/dev/mmcblk1, /dev/sdb, etc) []
   -i IMAGE        OS-specific image name (lts, dev, latest) [latest]
   -r RELEASE      OS-specific release number or name []
                   (14.10, 15.04, 15.10, utopic, vivid, wily, etc)
                   RELEASE takes precedence over IMAGE
   -U USERNAME     username of first created user [chrx]
   -H HOSTNAME     hostname for new system [chrx]
   -L LOCALE       locale for new system [en_US.UTF-8]
   -Z TIMEZONE     timezone for new system [America/New_York]
                   (America/San_Francisco, Etc/UTC, etc)
   -n              disable success/failure notifications
   -s              skip all customization, install stock OS only
   -y              run non-interactively, take defaults and do not confirm
   -v              increase output verbosity
   -h              show this help
  
  Default values are shown in brackets, e.g.: [default].
  
  If TARGETDISK is not specified, chrx will select the internal SSD.

```

### examples

Ubuntu 15.04, system name `hal`, first user `dave`:

`curl -Os https://chrx.org/go && sh go -r 15.04 -H hal -U dave`

Lubuntu (latest), verbosely:

`curl -Os https://chrx.org/go && sh go -v -m lubuntu-desktop`

### advanced usage

You may choose to host or cache these installation files yourself.
There are many good reasons to do so, especially if you'll be doing
a large number of installations. However, setup can be somewhat more
complicated, and instructions are outside the scope of this README.

To point **chrx** at your cache, just set the `CHRX_WEB_ROOT`
environment variable before running the `go` script, like this:

```
export CHRX_WEB_ROOT="http://myserver/chrx"
curl -O http://myserver/chrx/go && sh go
```


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

## chrx past

**chrx** began as an updated and enhanced version of
[ChrUbuntu](http://chromeos-cr48.blogspot.fr/2013/10/chrubuntu-for-new-chromebooks-now-with.html),
and still retains some original ChrUbuntu code (especially the
disk partitioning bits).


## chrx present

**chrx** has been used to install Linux on hundreds of Chromebooks.
Users and discussion can be found on [/r/chrubuntu](https://www.reddit.com/r/chrubuntu).


## chrx future

I'd like to test on a wider variety of hardware, and to install other
Linux distributions.

I'd dearly love to install FreeBSD, but my reading of the lists suggests
that driver support is not imminent.


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
- [Crouton](https://github.com/dnschneid/crouton) allows you to run ChromeOS
and Linux simultaneously, instead of dual-booting like **chrx** or ChrUbuntu.
This arrangement has a few drawbacks, but if you spend most of your time in
ChromeOS and your Linux needs are limited, it should serve well.
- The original [ChrUbuntu](http://chromeos-cr48.blogspot.fr/2013/10/chrubuntu-for-new-chromebooks-now-with.html) has been tested on a wide variety of hardware. Unfortunately, it hasn't been updated in a while, and requires quite a bit of additional configuration after installation. No longer recommended for Acer C720s,
but if your hardware is not supported by **chrx**, it's worth a try.


## notes on security and privacy

Running code from the net is always an act that requires careful thought.
**chrx** can be run directly from the net, and by default will download
additional code via the same mechanisms. Any of these downloads could be
misdirected or compromised. Downloading over an unsecured network (e.g.
public Wi-Fi) raises the likelihood of such malfeasance, but it can never
be fully eliminated.

If these are concerns of yours, you can mitigate your risks by auditing
all of the code involved, comparing checksums of downloaded packages, and
hosting local caches of everything. Almost no one actually does this, but
I don't want to contribute to the problem by ignoring it. See "advanced
usage", above.

Also, pre-release versions of **chrx** "ping home" on every install to
report success or failure. This ping includes **no personal information**,
only data that might be useful for investigating failures.

Log entries created by these pings look like this:

```
17.x.x.x - - [01/Jun/2015:07:37:00 +0000] "GET /end_ok HTTP/1.1" 200 0 "-"
  "chrx/1.1.1 hw=PEPPY_C6A-V7C-A2C sw=linux,ubuntu,latest,-,15.04,amd64,ubuntu-minimal" "-"
```

`hw` is a hardware ID that corresponds to your model of Chromebook.

`sw` combines a few of the command-line settings (or defaults) that you
used to run **chrx**.

If this level of information sharing makes you uncomfortable, the behaviour
can be disabled with the `-n` switch.


## meta, obligatory

**chrx** is pronounced "marshmallow".


## thanks

To Jay Lee for [ChrUbuntu](http://chromeos-cr48.blogspot.fr/2013/10/chrubuntu-for-new-chromebooks-now-with.html), to [/r/chrubuntu](https://www.reddit.com/r/chrubuntu) for assembling links to tons of helpful resources, and to the dozens of people who found answers and solved problems before I even started looking.


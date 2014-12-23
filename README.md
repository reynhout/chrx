# chrx

Use **chrx** to install Linux onto your Chromebook.

Presently, this means Ubuntu (or a variant) onto your Acer C720.

**chrx** installs OS components from the net. By default, package files
are downloaded from vendor sites, and several configuration tweaks are
downloaded from a secure mirror of this repository.

**chrx** is run directly from your Chromebook's `chronos@localhost` shell.
You must first enable
[Developer Mode](http://www.chromium.org/chromium-os/developer-information-for-chrome-os-devices),
and configure your network from the ChromeOS setup screen.

## status

Version 0.9.5 beta -- **working** but more testers welcome!

status| chromebook | unix | notes
:----:| ---------- | ---- | -----
:white_check_mark:|Acer C720|Ubuntu Linux (14.10, 14.04, development)|I use `lubuntu-desktop`
:question:|Other Haswell Chromebooks|Ubuntu Linux|These should work but might require other config tweaks
:x:|(any of the above)|Other UNIX or Linux distributions|Hopes are high
:x:|ARM Chromebooks|(any)|ARM support is spotty, these might never work

**Best-tested:** Lubuntu 14.10 on Acer C720

## usage

**Be sure to back up any valuable data in ChromeOS or Linux before
running chrx!**

Run the following command from your `chronos@localhost` shell:

`curl -Oks https://chrx.org/go && sh go`
<!-- yes, by all means read the all of the code first! -->

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
                   (ubuntu-minimal, xubuntu-desktop, etc)
   -a ARCH         processor architecture (i386, amd64) [amd64]
   -t TARGETDISK   target disk (/dev/mmcblk1, /dev/sdb, etc) []
   -i IMAGE        OS-specific image name (lts, dev, latest) [latest]
   -U USERNAME     username of first created user [chrx]
   -H HOSTNAME     hostname for new system [chrx]
   -L LOCALE       locale for new system [en_US.UTF-8]
   -Z TIMEZONE     timezone for new system [America/New_York]
                   (America/San_Francisco, Etc/UTC, etc)
   -n              disable success/failure notifications
   -y              run non-interactively, take defaults and do not confirm
   -v              increase output verbosity
   -h              show this help
  
  Default values are shown in brackets, e.g.: [default].
  
  If TARGETDISK is not specified, chrx will select the internal SSD.

```

### example

To verbosely install a Lubuntu system named "unicorn", with a first user
named "marc", run:

`curl -Oks https://chrx.org/go && sh go -v -m lubuntu-desktop -H unicorn -U marc`

### advanced usage

You may choose to host or cache these installation files yourself.
There are many good reasons to do so, especially if you'll be doing
a large number of installations. However, setup can be somewhat more
complicated, and instructions are outside the scope of this README.

To point **chrx** at your cache, just set the `CHRX_WEB_ROOT`
environment variable before running the `go` script, like this:

`CHRX_WEB_ROOT="http://myserver/chrx" curl -O http://myserver/chrx/go && sh go`


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

**chrx** is being used to mass-install Lubuntu onto dozens of Acer C720s.

The 1.0 release will mark the point when things are working well enough
to begin the first large rollout batches.

## chrx future

I'd like to test on a wider variety of hardware, and to install other
Linux distributions.

I'd dearly love to install FreeBSD, but my reading of the lists suggests
that driver support is not imminent.


## alternatives

**chrx** is a command-line installer, by necessity (mine). If this is not
a necessity that you share, you should take a look at Hugh Greenberg's
[Distroshare](https://www.distroshare.com/) for nicely updated ISOs which
can be installed from USB/SD flash RAM.


## notes on security and privacy

Running code from the net is always an act that requires careful thought.
**chrx** can be run directly from the net, and by default will download
additional code via the same mechanisms. Any of these downloads could be
misdirected or compromised. Downloading from a unsecured network (e.g.
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
17.x.x.x - - [22/Dec/2014:07:37:00 +0000] "GET /end_ok HTTP/1.1" 200 0 "-"
  "chrx/0.9.5 hw=PEPPY_C6A-V7C-A2C sw=linux,ubuntu,latest,amd64,ubuntu-minimal" "-"
```

`hw` is a hardware ID that corresponds to your model of Chromebook.

`sw` combines a few of the command-line settings (or defaults) that you
used to run **chrx**.

If this level of information sharing makes you uncomfortable, the behaviour
can be disabled with the `-n` switch.


## meta, obligatory

**chrx** is pronounced "marshmallow".


## thanks

To Jay Lee for ChrUbuntu, to [/r/chrubuntu](http://reddit.com/r/chrubuntu)
for assembling links to tons of helpful resources, and to the dozens of people
who found answers and solved problems before I even started looking.


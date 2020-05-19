#!/bin/sh
#
# go
#
# Instructions for update
#

cat << EOMSG

  Instructions for running chrx have been updated to work on newer
  versions of ChromeOS (R82 and later).

  These new instructions also work on older versions of ChromeOS,
  and are recommended for all.

  The new chrx command line is:

    curl https://chrx.org/ | sudo tar xzfC - /usr/local && chrx

  Please see https://chrx.org/ for full details.

EOMSG

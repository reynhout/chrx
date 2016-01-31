#!/bin/sh
#
# chrx-go
#
# A wrapper for chrx-install
#

: ${CHRX_WEB_ROOT:="https://chrx.org"}
CHRX_CACHE0_DIR="/var/tmp/chrx"
CHRX_LOG_FILE="${CHRX_CACHE0_DIR}/chrx-install.log"

## silence sudo msg
sudo -v >/dev/null 2>&1

rm -rf ${CHRX_CACHE0_DIR}
mkdir -p ${CHRX_CACHE0_DIR}
cd ${CHRX_CACHE0_DIR}

echo "## $0 $@" > $CHRX_LOG_FILE

## need sudo -E to keep exported environment vars
(
export CHRX_WEB_ROOT CHRX_CACHE0_DIR
curl -Os ${CHRX_WEB_ROOT}/chrx-install && sudo -E bash ./chrx-install "$@"
) 2>&1 | tee -a ${CHRX_LOG_FILE}

## logfile will be copied to chroot by install-chrx.
## can't copy here -- partition is unmounted when we get back.

## get stderr into tee: bash: `|& tee`;  sh: `2>&1 | tee`

#   

   ##################################################################
   ##                                                              ##
   ##        IF YOU ARE SEEING THIS MESSAGE UNEXPECTEDLY...        ##
   ##                                                              ##
   ##  You probably ran the 'curl' command with '-0' (dash-zero)   ##
   ##  instead of '-O' (dash-capital-O).                           ##
   ##                                                              ##
   ##  No harm done, just run the command again with a capital-O,  ##
   ##  and all will be right with the world.                       ##
   ##                                                              ##
   ##  See https://chrx.org/ for full chrx documentation.          ##
   ##                                                              ##
   ##################################################################


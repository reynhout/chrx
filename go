#!/bin/sh
#
# chrx-go
#
# A wrapper for chrx-install
#

: ${CHRX_WEB_ROOT:="https://chrx.org"}
CHRX_CACHE0_DIR="/var/tmp/chrx"

## silence sudo msg
## TODO: use sudo -v instead
sudo true >/dev/null 2>&1

rm -rf ${CHRX_CACHE0_DIR}
mkdir -p ${CHRX_CACHE0_DIR}
cd ${CHRX_CACHE0_DIR}

## need sudo -E to keep exported environment vars
(
export CHRX_WEB_ROOT CHRX_CACHE0_DIR
curl -Oks ${CHRX_WEB_ROOT}/chrx-install && sudo -E bash ./chrx-install $*
) | tee -a ${CHRX_CACHE0_DIR}/chrx-install.log

## logfile will be copied to chroot by install-chrx.
## can't copy here -- partition is unmounted when we get back.

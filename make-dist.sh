#!/bin/bash
set -e

# If you specify a version it will get packed into the chrx-install otherwise we append a string to denote it isn't a release build
# NEED_TEST
# pseudo TEST chrx -h | grep -E 'installer.*version'
: ${CHRX_VERSION:=$(grep -E '^CHRX_VERSION=' ./chrx-install | cut -f 2 -d '='| sed -e 's/"//g')-dev}
chrx_src=$(pwd)
build_dir=$(mktemp -d -t chrx.XXX)
cleanup() { rm -rf $build_dir; }
trap cleanup EXIT
echo $build_dir
cd $build_dir
mkdir bin
cp $chrx_src/chrx* ./bin
cp $chrx_src/dist/chrx* ./bin
# NEED_TEST
# Tweak the CHRX_VERSION in chrx-install
sed -i -e '/CHRX_VERSION=/ s/CHRX_VERSION=.*/CHRX_VERSION='$CHRX_VERSION'/' ./bin/chrx-install
# Debug print updated version
grep -E '^CHRX_VERSION=' ./bin/chrx-install
# Make sure they are executable as tar should preserve this on extract
chmod -R +x ./bin
mkdir -p etc/chrx-files
mv ./bin/chrx-devices ./etc/
rm -rf $chrx_src/dist/etc/etc/
cp -r -t ./etc/chrx-files/ $chrx_src/dist/etc/
cd $chrx_src
# NEED_TEST
tar czf dist.tar.gz -C $build_dir $(ls -A $build_dir)
# Debug print contents for comparison to official chrx.org version
echo "Confirm contents"
tar tf dist.tar.gz

#!/bin/bash

ARCH=$(uname -m)
case $ARCH in
"x86_64")
	ARCHIVE="flow_Linux_x86_64.tar.gz"
	;;
"i386")
	ARCHIVE="flow_Linux_i386.tar.gz"
	;;
"arm64")
	ARCHIVE="flow_Linux_arm64.tar.gz"
	;;
*)
	echo "Unsupported architecture $ARCH. Exiting."
	exit 1
	;;
esac
wget "https://github.com/TristanSch1/flow/releases/latest/download/$ARCHIVE" -O "/tmp/$ARCHIVE"

tar -xzf "/tmp/$ARCHIVE" -C /tmp/flow

mv /tmp/flow/flow /usr/local/bin/

rm "/tmp/$ARCHIVE"
rm -rf "/tmp/flow"

echo "Flow installation complete."

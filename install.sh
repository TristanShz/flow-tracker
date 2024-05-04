#!/bin/bash

# Déterminer l'architecture du système
ARCH=$(uname -m)
case $ARCH in
"x86_64")
	ARCHIVE_URL="https://github.com/TristanSch1/flow/releases/latest/download/flow_Linux_x86_64.tar.gz"
	;;
"i386")
	ARCHIVE_URL="https://github.com/TristanSch1/flow/releases/latest/download/flow_Linux_i386.tar.gz"
	;;
"arm64")
	ARCHIVE_URL="https://github.com/TristanSch1/flow/releases/latest/download/flow_Linux_arm64.tar.gz"
	;;
*)
	echo "Architecture non prise en charge: $ARCH"
	exit 1
	;;
esac

wget $ARCHIVE_URL

tar -xzf flow_Linux_*.tar.gz

mv flow /usr/local/bin/

rm flow_Linux_*.tar.gz

echo "Flow installation complete."

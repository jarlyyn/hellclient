#!/bin/bash

#settings
binname="hellclient"
buildername="build-linux.sh"
mclconvertorname="mclconvertor"

winbinname="hellclient.exe"
winbuildername="build-win.sh"
winpcredll="./lib/pcre/libpcre.dll"
winmclconvertorname="mclconvertor.exe"


macbinname="hellclient-mac"
macbuildername="build-mac.sh"


if [ -z "$1" ]
then 
    basename=$(basename "$0")
    echo "Usage $basename [targetpah]"
    exit 0
fi
if [ -e "$1" ]
then 
    echo "Target path $1 exists."
    exit 0
fi
path=$(dirname "$0")
cd $path

echo "Publish to $1."
echo "Building linux"
bash ./$buildername
echo "Building win"
bash ./$winbuildername
# echo "Building mac"
# bash ./$macbuildername
echo "Creating folder $1."
mkdir $1
mkdir $1/appdata
cp ../../appdata/readme.md $1/appdata/readme.md
echo "Copying bin file."
mkdir $1/bin
cp -rpf ../../bin/$binname $1/bin/$binname
echo "upx linux"
upx -9 $1/bin/$binname
cp -rpf ../../bin/$winbinname $1/bin/$winbinname
echo "upx win"
upx -9 $1/bin/$winbinname
# cp -rpf ../../bin/$macbinname $1/bin/$macbinname
# echo "upx mac"
# upx -9 $1/bin/$macbinname
cp -rpf ../../bin/$mclconvertorname $1/bin/$mclconvertorname
upx -9 $1/bin/$mclconvertorname
cp -rpf ../../bin/$winmclconvertorname $1/bin/$winmclconvertorname
upx -9 $1/bin/$winmclconvertorname
# cp -rpf ../../bin/$macmclconvertorname $1/bin/$macmclconvertorname
cp -rpf ../../$winpcredll $1/bin/libpcre-1.dll
echo "Copying system files."
cp -rpf ../../system $1/system
echo "Copying resources files."
cp -rpf ../../resources $1/resources
echo "Copying config skeleton files."
cp -rpf ../../system/configskeleton $1/config

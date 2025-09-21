#!/bin/bash
hostname=$USER
kernel=$(uname -r)
# ambil nama distru dan grep ke namanya aja
distro_name=$(grep "^NAME=" /etc/os-release | cut -d= -f2 | tr -d '"')
uptime=$(uptime -p)
shell=$(basename $SHELL)

echo "host    : $hostname on $distro_name"
echo "kernel  : $kernel"
echo "uptime  : $uptime"
echo "shell   : $shell"

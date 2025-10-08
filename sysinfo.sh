#!/bin/bash

# - echo -> mengembalikan niali

# fungsi utuk mendapatkan info theme berdasarkan DE yg di pakani
get_theme() {
  # gunakan pgrep untuk deteksi prosses Destop , untuk mengambil jenis theme nya
  # KDE plasma
  if pgrep -x plasmashell >/dev/null 2>&1; then
    theme=$(grep "widgetStyle=" ~/.config/kdeglobals | cut -d= -f2)
    echo "${theme:-none}"
    return
  fi
  #Gnome
  if pgrep -x gnome-shell >/dev/null 2>&1; then
    theme=$(gsettings get org.gnome.destop.interface gtk-theme 2>/dev/null | tr -d "'")
    echo "${theme:-none}"
    return
  fi
  # xfce4
  if pgrep -x xfce4-session >/dev/null 2>&1; then
    theme=$(xfconf-query -c xsettings -p /Net/ThemeName 2>/dev/null)
    echo "${theme:-none}"
    return
  fi
  echo ""
}

# "VM -> deteksi apakah system berjalan di atas hardward asli atau virtual mechin ( vm)"
check_vm() {
  # cek dgn systemd-detect-virt ( jik ada)
  if command -v systemd-detect-virt >/dev/null 2>&1; then
    virtual=$(systemd-detect-virt)
    if [[ "$virtual" != "none" ]]; then
      echo "$virtual"
      return
    fi
  fi

  # jika yg di atas tida ada (systemd-detect-virt) , gunakan /sys/class/dmi/id/
  if [[ -r /sys/class/dmi/id/product_name ]]; then
    product=$(</sys/class/dmi/id/product_name)
    vendor=$(</sys/class/dmi/id/sys_vendor)

    case "$product $vendor" in
    *VirtualBox* | *VMware* | *QEMU* | *KVM* | *Microsoft* | *Parallels*)
      echo "$product ($vendor)"
      return
      ;;
    esac
  fi

  # jika tidak terdeteksi || bare-metal
  # echo "None"
  echo ""
}

get_icon_theme() {
  # KDE plasma
  if pgrep -x plasmashell >/dev/null 2>&1; then
    awk -F='/^\[Icon\]/{f=1;next} f&&$1~/Theme/{print $2; exit}' ~/.config/kdeglobals
  #Gnome
  elif pgrep -x gnome-shell >/dev/null 2>&1; then
    gsettings get org.gnome.destop.interface icon-theme 2>/dev/null | tr -d "'"
  # xfce4
  elif pgrep -x xfce4-session >/dev/null 2>&1; then
    xfconf-query -c xsettings -p /Net/IconThemeName
  else
    grep "gtk-icon-theme-name" ~/.config/gtk-3.0/settings.ini | cut -d= -f2
  fi
}

get_gpu_info() {
  local GPU=""

  if command -v lspci >/dev/null 2>&1; then
    GPU=$(lspci | grep -i 'vga\|3d\|2d' | cut -d: -f3- | sed 's/^ //')
  elif command -v glxinfo >/dev/null 2>&1; then
    GPU=$(glxinfo | grep "OpenGL renderer" | sed 's/OpenGL renderer string //')
  # nvdia
  elif command -v nvidia-smi >/dev/null 2>&1; then
    GPU=$(nvidia-smi --query-gpu=name --format=csv,noheader | head -n1)
  fi
  # kalau kosong kembalikan None
  # if [[ -z "GPU" ]]; then
  # echo "None"
  # else
  # hapus newline jadi 1 baris
  echo "$GPU" | tr '\n' ',' | sed 's/,$//'
  # fi
}

hostname=$USER
kernel=$(uname -r)
# ambil nama distru dan grep ke namanya aja
distro_name=$(grep "^NAME=" /etc/os-release | cut -d= -f2 | tr -d '"')
uptime=$(uptime -p)
shell=$(basename $SHELL)
resolution=$(xrand | grep '*' | awk '{print $1}')
de=$($XDG_CURRENT_DESTOP || $XDG_CURRENT_DESTOP)
vm=$(check_vm)
theme=$(get_theme)
icons=$(get_icon_theme)
terminal=$TERM
cpu=$(lscpu | awk -F: '/Model name/ {gsub(/^[ \t]+/,"",$2); print $2}') # Harware (vendor/model)
gpu=$(get_gpu_info)
memory=$(awk '/MemTotal/ {printf "%.2f GB\n", $2/1024/1024}' /proc/meminfo) # ouput : ..,..GB

# informasi dasar yg pasti ada tidak perlu di cek ( None ( kembalian function))
# saya manfaatkan -z utntuk mengecek string kosong
echo "host        : $hostname on $distro_name"
echo "kernel      : $kernel"
echo "uptime      : $uptime"
echo "shell       : $shell"
if ![[ -z "$resolution" ]]; then
  echo "Resolution  : $Resolution"
fi
if ! [[ -z "$de" ]]; then
  echo "DE          : $de"
fi
if ! [[ -z "$vm" ]]; then # cek jika vm kosong ( -z)
  echo "WM          : $vm"
fi
if ! [[ -z "$then" ]]; then
  echo "Theme       : $theme"
fi
if ![[ -z "$icons" ]]; then
  echo "icons       : $icons"
fi
echo "Terminal    : $terminal"
if ! [[ -z "$gpu" ]]; then
  echo "GPU         : $gpu"
fi
echo "CPU         : $cpu"
echo "Memory      : $memory"

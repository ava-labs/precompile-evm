#!/bin/bash
# Pulls latest pre-built node binary from GitHub

#stop on errors
set -e

#helper function to check for presence of required commands, and install if missing
check_reqs_deb () {
  if ! command -v curl &> /dev/null
  then
      echo "curl could not be found, will install..."
      apt-get install curl -y
  fi
}
check_reqs_rhel () {
  if ! command -v curl &> /dev/null
  then
      echo "curl could not be found, will install..."
      dnf install curl -y
  fi
  if ! command -v wget &> /dev/null
  then
      echo "wget could not be found, will install..."
      dnf install wget -y
  fi
  if ! command -v dig &> /dev/null
  then
      echo "dig could not be found, will install..."
      dnf install bind-utils -y
  fi
  if ! command -v semanage &> /dev/null
  then
      echo "semanage could not be found, will install..."
      dnf install policycoreutils-python-utils -y
  fi
  if ! command -v restorecon &> /dev/null
  then
      echo "restorecon could not be found, will install..."
      dnf install policycoreutils -y
  fi
}
# Helper function to get OS Type
getOsType () {
  which yum 1>/dev/null 2>&1 && { echo "RHEL"; return; }
  which zypper 1>/dev/null 2>&1 && { echo "openSUSE"; return; }
  which apt-get 1>/dev/null 2>&1 && { echo "Debian"; return; }
}

# Installing necessary system dependencies for AvalacheGO
osType=$(getOsType)
if [ "$osType" = "Debian" ]; then
  check_reqs_deb
elif [ "$osType" = "RHEL" ]; then
  check_reqs_rhel
else
  #sorry, don't know you.
  echo "Unsupported linux flavour/distribution: $osType"
  echo "Exiting."
  exit
fi

# Define architecture
foundArch="$(uname -m)" 

if [ "$foundArch" = "aarch64" ]; then
  getArch="arm64"                               #we're running on arm arch (probably RasPi)
  echo "Found arm64 architecture..."
elif [ "$foundArch" = "x86_64" ]; then
  getArch="amd64"                               #we're running on intel/amd
  echo "Found amd64 architecture..."
elif [ "$foundArch" = "arm64" ]; then
  getArch="arm64"                               #we're running on intel/amd
  echo "Found arm64 architecture..."
else
  #sorry, don't know you.
  echo "Unsupported architecture: $foundArch!"
  echo "Exiting."
  exit
fi

# Import environment variables
source ./versions.sh

# Download AvalancheGo binary
curl -LJ -o avalanchego.tar.gz "https://github.com/ava-labs/avalanchego/releases/download/$AVALANCHEGO_VERSION/avalanchego-linux-$getArch-$AVALANCHEGO_VERSION.tar.gz"
tar -xzf avalanchego.tar.gz --wildcards '*/avalanchego' --strip-components=1 -C /avalanchego
rm avalanchego.tar.gz

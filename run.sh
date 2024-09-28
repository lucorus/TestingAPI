#!/bin/bash

check_go_installation() {
  if command -v go &> /dev/null; then
    return 0
  else
    return 1
  fi
}

install_go() {
  distro=$(uname -a | cut -d' ' -f1 | tr '[:upper:]' '[:lower:]')
  if [ "$distro" == "debian" ] || [ "$distro" == "ubuntu" ]; then
    sudo apt-get update
    sudo apt-get install golang-go
  elif [ "$distro" == "arch" ]; then
    sudo pacman -S go
  fi
}

if ! check_go_installation; then
  echo "Golang is not installed"
  echo "Running installation..."
  install_go
  echo "Success!"
fi

echo "Running project..."
go run main.go

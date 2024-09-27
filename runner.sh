#!/bin/bash

check_go_installation() {
  if command -v go &> /dev/null; then
    return 0
  else
    return 1
  fi
}

main() {
  if check_go_installation; then
    echo "Go is already installed."
    go run main.go
  else
    echo "Go is not installed."
  fi
}

main

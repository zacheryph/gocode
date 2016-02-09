#!/usr/bin/env bash

_build() {
  container_path=/go/src/rsvp

  docker run --rm \
    -v ${PWD}:$container_path \
    -w $container_path \
    golang:1.5.3-alpine \
      go build -v
}

case "${1:-default}" in
  default)
    _build
    ;;
  build)
    _build
    ;;
  vet)
    go vet
    ;;
  *)
    echo "Unknown command: $1"
    echo "Available: build, vet"
esac

#!/bin/bash
cd "$(dirname "$0")"
set -eu

rm -f machine-fuzz
CC=hfuzz-clang CXX=hfuzz-clang++ CFLAGS=-finstrument-functions CXXFLAGS=-finstrument-functions CGO_CFLAGS=-finstrument-functions CGO_CXXFLAGS=-finstrument-functions go build

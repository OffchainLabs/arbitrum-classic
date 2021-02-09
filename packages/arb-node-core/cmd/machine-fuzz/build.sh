#!/bin/bash
cd "$(dirname "$0")"
set -eu

rm -f machine-fuzz
CC=hfuzz-clang CXX=hfuzz-clang++ go build

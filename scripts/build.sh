#!/bin/bash
cd "$(dirname "$0")/.."
docker build $@ -t thavlik/mirbase:latest .
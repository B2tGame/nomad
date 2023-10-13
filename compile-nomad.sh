#!/bin/bash
set -uex
sudo service nomad stop
make clean
make bootstrap
make release
sudo rm /usr/local/bin/nomad
sudo cp pkg/linux_amd64/nomad /usr/local/bin/nomad
sudo service nomad start
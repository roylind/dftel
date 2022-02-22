#!/bin/sh
sudo systemctl disable dftel
sudo rm /usr/bin/dftel
sudo rm /etc/systemd/system/dftel.service

#!/bin/sh 
set -e
# Copy binary
sudo cp ./dftel /usr/bin/dftel
sudo chown root:root /usr/bin/dftel
sudo chmod 755 /usr/bin/dftel
# Install service
sudo cp ./dftel.service /etc/systemd/system/dftel.service
sudo chown root:root /etc/systemd/system/dftel.service
sudo chmod 644 /etc/systemd/system/dftel.service
sudo systemctl enable dftel

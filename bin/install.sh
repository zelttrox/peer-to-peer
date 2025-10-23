#!/bin/bash

# Build go package
cd .. && go build -o peer .

# Move compiled file to bin
sudo mv peer /bin

# Move config files to .config
mkdir ~/.config/peer
cp -r config/. ~/.config/peer/

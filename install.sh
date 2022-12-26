#!/usr/bin/env bash

# Check if the urls.json file exists in the cmd/addurl directory
if [ ! -f "cmd/addurl/urls.json" ]; then
  # Create an empty urls.json file in the cmd/addurl directory
  touch cmd/addurl/urls.json
  echo "[]" > cmd/addurl/urls.json
fi

# Change to the addurl directory
cd cmd/addurl

# Build the binary file
go build -o addurl

# Make the script executable
chmod +x install.sh

# Move the binary file to /usr/local/bin
sudo mv addurl /usr/local/bin/

# Make the binary file executable
sudo chmod +x /usr/local/bin/addurl

# Print success message
echo "addurl installed successfully
#!/bin/zsh
sed -i '/export https\?_proxy/s/^.*e/e/' ~/.zshrc
echo "http proxy recovery"
source ~/.zshrc

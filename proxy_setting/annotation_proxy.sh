#!/bin/zsh
sed -i '/export https\?_proxy/s/^/#/' ~/.zshrc
echo "http proxy annotationned"
source ~/.zshrc 

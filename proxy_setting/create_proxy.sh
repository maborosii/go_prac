#!/bin/zsh

owner_host=`cat /etc/resolv.conf | awk '$0 !~ /^#/ {print $2}'`
# echo "${owner_host}"
is_exists_proxy=`cat ~/.zshrc | grep "https\?_proxy"`
# is_annotationed=`echo ${is_exists_proxyj} | grep "^#"`
# echo "${is_exists_proxy}"
if [ ! -z "${is_exists_proxy}" ];then
  sed -i "\@https\?_proxy@s@://.*:@://${owner_host}:@" ~/.zshrc
  echo "modify success"
else
  echo "export http_proxy=\"http://${owner_host}:10809\"" >> ~/.zshrc
  echo "export https_proxy=\"https://${owner_host}:10809\"" >> ~/.zshrc
  echo "create success"
fi
source ~/.zshrc

#!/bin/bash

#获取主机名
Host_Name=`docker stack ps micro-service |grep "appr-law-api-gateway" |grep Run|awk '{print $4}'`
log_file=gateway.log
> ${log_file}
local_time=`date "+%Y-%m-%d"`
at_time=`date "+%Y-%m-%d %H:%M"`
#获取日志
for i in $Host_Name
do 
	Container_Name=(`ssh $i docker ps |grep "appr-law-api-gateway" |awk '{print $1}'`)
	for j in $Container_Name
	do
		ssh $i "docker exec -t $j /bin/sh -c  'cd /logs && grep pcssologin gateway.${local_time}*.log |grep ${local_time}|grep  userCode'" >> ${log_file}
		ssh $i "docker exec -t $j /bin/sh -c  'cd /logs && grep pcssologin gateway.log |grep ${local_time}|grep  userCode'" >> ${log_file}
		#ssh $i "docker exec -t $j /bin/sh -c  'cd /logs && grep pcssologin gateway*log |grep ${local_time}|grep  userCode'" >> ${log_file}
	done
done 


#获取数量
echo -e  "截止至${at_time},当天访问量为："
cat ${log_file}|sort|uniq|wc -l|awk '{print $1}'
echo -e "截止至${at_time},当天访问用户数为："
cat gateway.log |awk -F "=" '{print $2}'|sort|uniq|wc -l
echo -e "\n"

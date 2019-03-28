"#!/bin/bash"

echo "########################################################"
echo "####################### Update #########################"
echo "########################################################"

#UDT=$(yum -y update)
#echo "$UDT"

echo ""
echo ""

echo "########################################################"
echo "##################### COMPLETE !! ######################"
echo "########################################################"

echo ""
echo ""

echo "########################################################"
echo "#################### ssh security ######################"
echo "########################################################"

ssh1=$(cat /etc/ssh/sshd_config | grep Protocol)
echo "$ssh1"


ssh2=$(cat /etc/ssh/sshd_config | grep RhostsRSA)
echo "$ssh2"


chg1=$(sed -i "60s/^#//" /etc/ssh/sshd_config)$(sed -i "60s/IgnoreRhosts no/IgnoreRhosts yes/" /etc/ssh/sshd_config)
ssh3=$(cat /etc/ssh/sshd_config | grep Ignore)
echo "$chg1$ssh3"
echo ""

ssh4=$(cat /etc/ssh/sshd_config | grep PermitEmptyPasswords)
chg2=$(sed -i "63s/^#//" /etc/ssh/sshd_config)$(sed -i "63s/PermitEmptyPasswords yes/PermitEmptyPassword no/" /etc/ssh/sshd_config)
echo "$chg2$ssh4"
echo ""

ssh5=$(cat /etc/ssh/sshd_config | grep PermitRootLogin)
chg3=$(sed -i "38s/^#//" /etc/ssh/sshd_config)$(sed -i "38s/PermitRootLogin yes/PermitRootLogin no/" /etc/ssh/sshd_config)
echo "$ssh5$chg3"
echo ""

echo ""
echo ""
echo "########################################################"
echo "############### ssh security complete!! ################"
echo "########################################################"

#install_ntp=$(yum install ntp)
#echo "$install_ntp"

#time_synchronization
server_time=$(sed -i "25s/^/server 192.168.8.3/" /etc/ntp.conf)
echo "$server_time"
echo ""

#Daemon_start
Daemon=$(service ntpd start) #daemon ntpd start

echo ""
echo "################# Time and ntpd start ##################"
echo ""

mk_usb=$(touch /etc/modprobe.d/disable-usb)
chg_usb=$(sed -i "1s/^/install usb-storage" /etc/modprobe.d/disable-fireware)
check_usb=$(cat /etc/modprobe.d/disable-usb)

echo "$mk_usb$chg_usb$check_usb"

echo "############### usb #################" 
mk_fireware=$(touch /etc/modprobe.d/disable-fireware)
chg_fireware=$(sed -i "1s/^/install fireware ohci/" /etc/modprobe.d/disable-fireware)
check_fireware=$(cat /etc/modprobe.d/disable-usb)

echo "############### fireware ###################"
echo "$mk_fireware$chg_fireware$check_fireware"

echo ""
echo ""


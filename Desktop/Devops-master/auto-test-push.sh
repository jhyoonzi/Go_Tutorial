#!/bin/bash


echo ""
echo ""
echo "#######################################################################################"
echo "#################################### run test #########################################"
echo "#######################################################################################"
echo ""
echo ""

python3 -m pytest --cov=.

if [  "$?" = "0" ]; then
    echo ""
    echo ""
    echo "###################################################################################"
    echo "################################# git push ########################################"
    echo "###################################################################################"
    echo ""
    echo ""
    
    message="$1"
    if [ $1 = '' ]; then
        message="auto commit and push"
    fi 

    git add .
    git commit -am "$message"
    git push 

else
    echo ""
    echo ""
    echo "################################### test fail ####################################"
    echo ""
    echo ""

else
    echo ""
    echo ""


    exit 1
fi


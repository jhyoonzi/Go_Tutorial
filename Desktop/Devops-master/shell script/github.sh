echo "#/bin/bash"


add=$(git add *)
commit=$(git commit -m "jhyoon")
push=$(git push origin master)

$all="$add,$commit,$push"
echo "$all"


echo "##########################"
echo "#########complete#########"
echo "##########################"

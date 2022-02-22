#!/bin/sh
I=`dpkg -s golang | grep "Status" ` #проверяем состояние пакета (dpkg) и ищем в выводе его статус (grep)
if [ -n "$I" ] #проверяем что нашли строку со статусом (что строка не пуста)
then
    go build dftel.go
else
   echo "Install Go Lang (sudo apt install golang)"
fi

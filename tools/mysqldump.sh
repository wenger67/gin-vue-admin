#!/bin/bash
v_date=`date "+%Y%m%d%H%M%S"`
path=`pwd`
echo $path
echo $v_date
/usr/bin/mysqldump -uroot -p qmPlus > $path/../server/db/dump/$v_date.sql
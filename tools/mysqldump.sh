#!/bin/bash
v_date=`date "+%Y%m%d%_H%M%S"`
path=`pwd`
/usr/bin/mysqldump -uroot -p qmPlus > $path/../server/db/dump/$v_date.sql
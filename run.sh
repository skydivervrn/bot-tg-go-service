#!/bin/bash

systemctl stop cmdrunner
rm /bin/cmdrunner
cp cmdrunner /bin/cmdrunner
cp cmdrunner.service /lib/systemd/system/cmdrunner.service
systemctl enable cmdrunner.service
systemctl start cmdrunner
systemctl status cmdrunner

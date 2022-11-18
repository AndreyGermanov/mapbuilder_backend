#!/bin/bash
git push origin master
ssh root@test.germanov.dev "cd /root/mapbuilder_backend/;git pull origin master;./build.sh"

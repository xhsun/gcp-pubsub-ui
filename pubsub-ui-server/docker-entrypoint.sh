#!/bin/bash
  
# turn on bash's job control
set -m
  
# Start the primary process and put it in the background
/server/pubsubui_server &
  
# Start the helper process
/usr/local/bin/envoy -c /etc/envoy/envoy.yaml -l trace --log-path /tmp/envoy_info.log

# now we bring the primary process back into the foreground
# and leave it there
fg %1
#!/bin/bash
cleanup() { pkill python3; }
# Don't want previous `pkill -P $$` which would kill parent session if sourced?
# trap cleanup EXIT
export CHRX_WEB_ROOT="http://localhost:8000/"
python3 -m http.server &
sleep 3 # Lets the webserver start
curl $CHRX_WEB_ROOT/dist.tar.gz | sudo tar xzfC - /usr/local 
# Only necessary if tar doesn't preserve permissions
# && sudo chmod +x /usr/local/bin/chrx*
chrx -h
chrx -h | grep -E 'installer.*version'
cleanup

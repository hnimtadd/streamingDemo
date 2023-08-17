#!/bin/bash

#MONGODB1=`ping -c 1 video-analytics-mongodb | head -1  | cut -d "(" -f 2 | cut -d ")" -f 1`

MONGODB1=video-analytics-mongodb-hls-streaming
UserName=$MONGO_INITDB_ROOT_USERNAME
Password=$MONGO_INITDB_ROOT_PASSWORD
Database=$MONGO_INITDB_DATABASE

echo "********MONGODB1 = " ${MONGODB1}
echo "********UserName = " ${UserName}
echo "********Password = " ${Password}
echo "********Database = " ${Database}

echo "********Waiting for startup..********"
until curl http://${MONGODB1}:27017/serverStatus\?text\=1 2>&1 | grep uptime | head -1; do
  printf '.'
done

echo curl http://${MONGODB1}:27017/serverStatus\?text\=1 2>&1 | grep uptime | head -1
echo "********Started..********"

echo SETUP.sh time now: `date +"%T" `
mongosh --host ${MONGODB1} --port 27017 <<EOF
use video_analytics;
var cfg = {
    "_id": "rs0",
    "protocolVersion": 1,
    "version": 1,
    "members": [
        {
            "_id": 0,
            "host": "${MONGODB1}:27017",
            "priority": 1
        }
    ],settings: {chainingAllowed: true}
};
rs.initiate(cfg, { force: true });
rs.status();

EOF


mongosh --host ${MONGODB1} --port 27017  <<EOF

use admin;
db.createUser( { user: 'vinai_User', pwd: 'vinai_Password', roles: [ { role: 'readWrite', db: 'admin' }]});
use video_analytics;
db.createUser( { user: 'vinai_User', pwd: 'vinai_Password', roles: [ { role: 'readWrite', db: 'video_analytics' }]});
EOF

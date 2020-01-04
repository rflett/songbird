#!/bin/bash
docker rm songbird
docker run -d --name songbird --net=host -e API_URL=$API_URL -e CONSUMER_KEY=$CONSUMER_KEY -e CONSUMER_SECRET=$CONSUMER_SECRET -e ACCESS_TOKEN=$ACCESS_TOKEN -e ACCESS_SECRET=$ACCESS_SECRET songbird
echo "Tailing logs.."
docker logs -f songbird

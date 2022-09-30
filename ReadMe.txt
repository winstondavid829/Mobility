////////////////////////////////////////////// Development - 35.228.251.106 /////////////////////////////////////////////////////////
gcloud endpoints services deploy openapi-entitlementregistry-dev.yaml
sudo docker network create --driver bridge esp_net_entertainment
docker build . -t go-dock-entertainment-dev:latest
docker tag go-dock-entertainment-dev:latest gcr.io/iconic-ruler-363702/go-dock-entertainment-dev:latest
docker push gcr.io/iconic-ruler-363702/go-dock-entertainment-dev:latest
sudo docker pull gcr.io/iconic-ruler-363702/go-dock-entertainment-dev:latest
sudo docker run --detach --restart always --name=esp-entertainment --net=esp_net_entertainment gcr.io/iconic-ruler-363702/go-dock-entertainment-dev:latest


 #Creating Server Blocks in nginx
 For this user script 'create_gcpendpoint.sh' under /automate_nginx folder, the command line parameter is cloud end point url
sudo docker run --name=espssl-5011-entertainment \
     --detach \
     --publish=5011:443 \
     --net=esp_net_entertainment \
     -v /etc/letsencrypt/live/entertainment-dev.endpoints.iconic-ruler-363702.cloud.goog/fullchain.pem:/etc/nginx/ssl/nginx.crt \
     -v /etc/letsencrypt/live/entertainment-dev.endpoints.iconic-ruler-363702.cloud.goog/privkey.pem:/etc/nginx/ssl/nginx.key \
     --link=esp-entertainment:esp-entertainment \
     gcr.io/endpoints-release/endpoints-runtime:1 \
     --service=entertainment-dev.endpoints.iconic-ruler-363702.cloud.goog \
     --rollout_strategy=managed \
     --backend=esp-entertainment:8080 \
     --ssl_port=443
======================================================================================================================================================
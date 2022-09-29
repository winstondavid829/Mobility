/////////////////////////////// Production Start -  35.228.150.245 ////////////////////////////////////////////////////////////////////
gcloud endpoints services deploy openapi-entertainment.yaml
sudo docker network create --driver bridge esp_net_entertain
docker build . -t go-dock-entertain-production:latest
docker tag go-dock-entertain-production:latest gcr.io/iconic-ruler-363702/go-dock-entertain-production:latest
docker push gcr.io/iconic-ruler-363702/go-dock-entertain-production:latest
sudo docker pull gcr.io/iconic-ruler-363702/go-dock-entertain-production:latest

sudo docker run --detach --restart always --name=esp-entertain --net=esp_net_entertain --log-opt max-size=10m --log-opt max-file=5 gcr.io/iconic-ruler-363702/go-dock-entertain-production:latest
sudo docker run --name=espssl-5007-entertain \
     --detach \
     --publish=5007:443 \
     --net=esp_net_entertain \
     -v /etc/letsencrypt/live/entertain.endpoints.iconic-ruler-363702.cloud.goog/fullchain.pem:/etc/nginx/ssl/nginx.crt \
     -v /etc/letsencrypt/live/entertain.endpoints.iconic-ruler-363702.cloud.goog/privkey.pem:/etc/nginx/ssl/nginx.key \
     --link=esp-entertain:esp-entertain \
     gcr.io/endpoints-release/endpoints-runtime:1 \
     --service=entertain.endpoints.iconic-ruler-363702.cloud.goog \
     --rollout_strategy=managed \
     --backend=esp-entertain:8080 \
     --ssl_port=443
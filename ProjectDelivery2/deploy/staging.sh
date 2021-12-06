
APPNAME=food-delivery
# copy qua bash chay
docker load -i food-delivery.tar
docker rm -f food-delivery

docker run  -d --name food-delivery \
--network my-net \
-e VIRTUAL_HOST=159.223.81.47



APP_NAME=food-delivery \

echo "Downloading packages ..." \
go mod download \
echo "Compiling ..."\
CGO_ENABLED=0 GOOS=linux  go build -a -installsuffix cgo -o app \

echo "Docker building ..."\
docker build -t food-delivery  -f ./Dockerfile . \
echo "Docker saving ..." \
docker save -o food-delivery.tar food-delivery:latest
#
#echo "Deploying ..."
#

scp  ./food-delivery.tar root@159.223.81.47:~
ssh  root@159.223.81.47 'bash -s' < ./deploy/staging.sh

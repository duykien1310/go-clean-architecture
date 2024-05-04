TAG=$(date '+%Y%d%m.%H%M%S')

ImageName="auth_service"
echo "Building image $ImageName:$TAG"
docker build . -t $ImageName:$TAG
docker tag $ImageName:$TAG $ImageName:latest

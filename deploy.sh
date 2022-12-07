#/bin/bash

source .env

DOCKER_NAME=$DOCKER_REGISTRY/hasura_auth:graphql-default

docker login --username $DOCKERHUB_USERNAME --password $DOCKERHUB_PASSWORD
docker build . -t $DOCKER_NAME
docker push $DOCKER_NAME

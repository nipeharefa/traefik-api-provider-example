#!/bin/bash
docker run \
--name traefik -d -p 8080:80 -p 8081:8080 \
-v $PWD/traefik.yml:/etc/traefik/traefik.yml \
traefik:v2.5
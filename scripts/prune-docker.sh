#!/bin/bash

docker system prune -f
docker volume prune -f
docker network prune -f

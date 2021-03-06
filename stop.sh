#!/bin/bash
echo "Entering virtualenv"
source demo_platform/bin/activate
echo "Stopping all Docker processes and servers"
docker stop pro_tes_flower_1 || true && docker rm pro_tes_flower_1 || true
docker stop pro_tes_protes-worker_1 || true && docker rm pro_tes_protes-worker_1 || true
docker stop pro_tes_rabbitmq_1 || true && docker rm pro_tes_rabbitmq_1 || true
docker stop pro_tes_protes_1 || true && docker rm pro_tes_protes_1 || true
docker stop pro_tes_mongodb_1 || true && docker rm pro_tes_mongodb_1 || true
docker stop drs_mock-drs_1 || true && docker rm drs_mock-drs_1 || true
docker stop drs_mongo_1 || true && docker rm drs_mongo_1 || true
killall -9 funnel
killall -9 wes

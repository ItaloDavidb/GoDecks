#!/bin/bash

echo "Starting the backend..."
cd back || exit 1
make docker-compose
make run &

echo "Waiting for the backend to start..."
sleep 5

echo "Starting the frontend..."
cd ../front || exit 1
npm start

echo "Backend and frontend are running!"

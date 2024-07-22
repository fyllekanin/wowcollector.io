# Development

This document will describe the needed things for developing

## Server

The server is written in GoLang (v1.22)

## Web

The web is using Nuxt with NodeJS (v20.15.1)

## Database

Within the development folder there is a **database-compose.yaml** which will start up both mongoDB and mongo-express
The express UI will be accesible at **http://localhost:8081** with credentails admin/admin

## Running the server

### Build and start with Go
---------
#### Generate swagger documentation
```
./back-end/generate-swagger.sh back-end
```
#### Build the server
```
$(cd back-end && go build) 
```

#### Start the server
```
./back-end/wowcollector.io
```

### Build and start with Docker
---------
#### Build the image (root folder)
```
docker build --tag=wowcollector.io-server ./back-end/
```

#### Start the image
```
docker run -p 8888:8888/tcp \
-e DATABASE_USERNAME=admin \
-e DATABASE_PASSWORD=admin \
-e DATABASE_NAME=wowcollector \
-e DATABASE_HOST=localhost \
-e DATABASE_PORT=27017 \
-e BATTLE_NET_CLIENT_ID=clientId \
-e BATTLE_NET_CLIENT_SECRET=clientSecret \
wowcollector.io-server
```

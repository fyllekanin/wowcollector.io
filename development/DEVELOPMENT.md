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
$(cd back-end && go build -v -o web-app ./cmd/web/)
```

#### Start the server
```
./back-end/web-app
```

### Build and start with Docker
---------
#### Build the image (root folder)
```
docker build -f ./back-end/DockerfileRest --tag=wowcollector.io-server ./back-end/
```

#### Start the image
```
docker run -p 8888:8888/tcp \
-e DATABASE_CONNECTION=mongodb://admin:admin@localhost:27017 \
-e BATTLE_NET_CLIENT_ID=clientId \
-e BATTLE_NET_CLIENT_SECRET=clientSecret \
-e GITHUB_TOKEN=githubPersonalToken \
-e JWT_SECRET_KEY="nptOqeOBZXJ3/T0KKAX6MO5rObFIV4qTOZ7EuALGqOE=" \
wowcollector.io-server
```

## Running the scanner

### Modifications
---------
#### Why
As the scanner is configured to not run everything all the time there is timeouts
before certain scans started.

#### Edit a task to run directly
- go to back-end/cmd/scanner/main.go
- Find the task you want to run
- Edit the time, example below for having realm scanner for EU to run on startup

Before
```
time.AfterFunc(1*time.Hour, ....)
```

After
```
time.AfterFunc(1*time.Second, ....)
```

### Build and run the scanner with Go 
---------
#### Build the scanner (run within the back-end folder)
```
go build -o scanner-app ./cmd/scanner/
```

#### Start it
```
./scanner-app
```

### Build and run the scanner with Docker
---------
#### Build the image (root folder)
```
docker build -f ./back-end/DockerfileScanner --tag=wowcollector.io-scanner ./back-end/
```

#### Start the image
```
docker run -p 8888:8888/tcp \
-e DATABASE_CONNECTION=mongodb://admin:admin@localhost:27017 \
-e BATTLE_NET_CLIENT_ID=clientId \
-e BATTLE_NET_CLIENT_SECRET=clientSecret \
wowcollector.io-scanner
```

### Build and run the migration & seeder with Go 
---------
#### Build the scanner (run within the back-end folder)
```
go build -o migration-seeder-app ./cmd/migration-seeder/
```

#### Start it
```
./migration-seeder-app
```

## Running the web

### Build and start with Node
---------
#### Build the server
```
$(cd front-end && npm install)
```

#### Start the server
```
cd front-end && npm run dev
```

### Build and start with Docker
---------
#### Build the image (root folder)
```
docker build -f ./front-end/DockerfileNuxt --tag=wowcollector.io-ui ./front-end/
```

#### Start the image
```
docker run -p 3000:3000/tcp \
wowcollector.io-ui
```
# Golang Boilerplate

Golang boilerplate based on [Uncle Bob's Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) with slight modification on the DB layer (this project use ORM). The technologies used in this project are:

- [Firebase](firebase.google.com) for authentication
- [Gin](github.com/gin-gonic/gin) for routing
- [Viper](github.com/spf13/viper) for environment settings
- [Gorm](gorm.io/gorm) as ORM
- [SQLite](https://www.sqlite.org) as database
- [PostgreSQL](https://www.postgresql.org) as an alternative to SQLite

## Project Setup

You can rename the project by replacing all "boilerplate" to your new project name in every file.

To start the project, you need to make sure that you have Go (Golang) and Docker installed. Then you can run this commands to install Golang dependency:

```
go get
```

### Integrating Firebase

If you want to use Firebase, you need to provide your service account key and name the file to `firebaseServiceAccountKey.json`. Check out `firebaseServiceAccountKey.json.sample` to make sure you have the correct file and place it in the correct directory.

### Modify Configuration

This app will take configurations from environment variables or `.env` file in root folder. Duplicate the `.env.sample` file and rename it to `.env`. Then, configure the environment variables according to your needs. The default configuration is usable.

### Integrating Database

You can change the database settings by editing the environment variables or `.env` file from the previous step. By default, this app will use SQLite as database stored in `gorm.db` file in root folder. If you want to run a PostgreSQL database in Docker, you can run:

```
bash start_db.sh
```

## Running the Program

You can run the program by simply running these commands:

```
go mod tidy
go run main.go
```

By default, you can access the app from `http://localhost:8080/`.

## API Documentation

The REST API documentation is accessible from `docs/docs.http`. If you want to test the REST API from this file, you can use this [VSCode REST Client Extension](https://marketplace.visualstudio.com/items?itemName=humao.rest-client). Please don't forget to change the Firebase API key in `@firebase_api_key`.

## Future Plan

1. Integrate Redis
2. Websocket
3. Terraform + Docker deployment
4. SonarQube
5. Unit test
6. Prometeus + Grafana

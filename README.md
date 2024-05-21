# BookStore


A simplistic bookstore application built using the Gin Framework.

![terminal](./misc/images/background.png)




### Prerequisites

Before initiating this project, ensure that you have installed all the required dependencies on your machine.
### Software Installation

- [x] [Go](https://go.dev/) - Go is an open source programming language that makes it simple to build secure, scalable systems.
- [x] [Docker](https://www.docker.com/) - Docker helps developers bring their ideas to life by conquering the complexity of app development.
- [x] [PostgreSQL](https://www.postgresql.org/) - The World's Most Advanced Open Source Relational Database
- [x] [AWS CLI](https://aws.amazon.com/cli/) - The AWS Command Line Interface (CLI) is a unified tool to manage AWS services directly from the command line.



For running Postgres locally using Docker, run the following command:

```bash
docker run --name bookstore -p 5432:5432 -e POSTGRES_PASSWORD=******** -d postgres
```

Execute in Postgres DB Shell

```sql
create database bookstore;
```

### Environment Variables

Before initiating the application, ensure that you have configured the required environment variables.
```
- DB_HOST
- DB_USERNAME
- DB_PASSWORD
- DB_NAME
- DB_PORT
- S3_BUCKET
```

We will access AWS resources such as Amazon S3. Therefore, ensure you configure your AWS CLI with both the `AWS Access Key ID` and `Secret Access Key`.

### Application Startup

#### Running App

```bash
go run main.go
```

### Blog

Want to follow step by step, follow this tutorial: [REST API Development with Gin](https://www.jetbrains.com/guide/go/tutorials/bookstore_rest_api/)


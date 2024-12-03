# IAM
> Identity and Access Management

## Folder structure
- `main.go`: This file serves as the entry point of the application. It typically contains the initialization logic for your Gin router and other essential configurations.
- `go.mod` and `go.sum`: These files are used for managing dependencies with Go modules . They list the project’s dependencies and their versions.
- `README.md`: A README file containing essential information about the project, such as its purpose, installation instructions, and usage guidelines.
- `config/`: This directory contains configuration files used to customize the behavior of the application. These files may include settings related to database connections, server ports, logging levels, and other runtime parameters.
- `env/`: The `env` directory typically stores environment-specific configuration files, such as development, staging, and production configurations. Separating configurations by environment helps maintain consistency and simplifies deployment workflows.
- `handlers/`: Handlers are responsible for processing incoming requests and generating appropriate responses. Organize your handlers based on the routes they handle or the resources they manage to maintain clarity and coherence.
- `middleware/`: Middleware functions intercept incoming requests before they reach the handlers, allowing you to perform common tasks such as authentication, logging, and request preprocessing. Separating middleware logic from route handlers promotes modularity and reusability.
- `repository/`: The repository directory contains database access and data retrieval logic. It encapsulates interactions with the database, including CRUD operations, transaction management, and query execution.
- `models/`: Database models represent the structure of your data and define how it is stored and manipulated. Organize your models based on the entities they represent, such as users, products, or orders, to maintain clarity and consistency.
- `utils/`: Utility functions encapsulate common tasks or operations that are used throughout the application. Examples include string manipulation, date formatting, and encryption utilities. Organize your utilities based on functionality to facilitate code discovery and reuse.
- `helpers/`: Helper functions provide reusable snippets of code that assist in specific tasks or workflows. These functions are typically used to simplify complex operations or encapsulate common patterns. Keep your helper functions well-documented and modular to promote code readability and maintainability.

## Running app

### Pre

1. Prepare `.env` file data as below

    ```shell
    DB_DRIVER=postgres
    DB_HOST=localhost
    DB_PORT=5432
    DATABASE=iam
    DB_USER=postgres
    DB_PASSWORD=postgres
    PORT=5051
    TOKEN_HOUR_LIFESPAN=1
    API_SECRET=changeme
    HMAC_SECRET=changeme
    REDIS_PORT=6379
    REDIS_HOST=localhost
    REDIS_DB=0
    REDIS_PASSWORD=""
    
    KAFKA_HOST='localhost:9092'
    
    MASTER_USERNAME=${MASTER_USERNAME}
    MASTER_PASSWORD=${MASTER_PASSWORD}
    MASTER_EMAIL=
    MASTER_FIRST_NAME=
    MASTER_LAST_NAME=
    ```

2. Start postgres on port `5432`
    ```shell
      docker compose up
    ```
### Build and Run app
- Build
```shell
go build -o iam
```

- Run
```shell
./iam run
```

- Test
```shell
go test ./... 
```

In case have error happened like below:
```shell
kafka/producer.go:10:19: undefined: kak.Producer
kafka/producer.go:15:22: undefined: kak.NewProducer
kafka/producer.go:15:39: undefined: kak.ConfigMap
kafka/producer.go:25:23: undefined: kak.Message
kafka/producer.go:26:23: undefined: kak.TopicPartition
kafka/producer.go:26:75: undefined: kak.PartitionAny
kafka/producer.go:30:32: undefined: kak.Event
```

This is the way to fix this:
1. Set variable by cli below:
   ```shell
   go env -w CGO_ENABLED=1
   ```
2. If compiler have error related to gcc
   ```shell
   // on linux we can install all essentials package for dev
   sudo apt-get install build-essential
   ```

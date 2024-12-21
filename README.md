## API and backend layer for Consunet web application

### Prerrequisites

- Install Go.
- Install air. [(See how)](https://github.com/air-verse/air)

### Execute the project locally

1. Clone the repository

```sh
git clone git@github.com:danielRamosMencia/consunet-api.git
```

```sh
git clone https://github.com/danielRamosMencia/consunet-api.git
```

2. Install packages and dependencies.

```sh
go mod tidy
```

NOTE: if not works, instead, use:

```sh
go mod vendor
```

3. Execute the server.

```sh
air
```

NOTE: if you haven't installed air, use:

```sh
go run main.go
```

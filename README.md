## API y capa backend para aplicación Consunet.

### Prerrequisitos

- Instalar Go.
- Instalar air. [(Ver como)](https://github.com/air-verse/air)

### Ejecutar el proyecto

1. Clonar el repositorio.

```sh
git clone git@github.com:danielRamosMencia/consunet-api.git
```

```sh
git clone https://github.com/danielRamosMencia/consunet-api.git
```

2. Instalar los paquetes y dependencias.

```sh
go mod tidy
```

NOTA: En su defecto, usar:

```sh
go mod vendor
```

3. Levantar servidor.

```sh
air
```

NOTA: Si no ha instalado air puede usar:

```sh
go run main.go
```

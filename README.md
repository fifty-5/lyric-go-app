## ⚡️ Configuración 


Dependiendo el ambiente seleccionado, es necesario ejecutar el docker-compose para cada situación descrita a continuación:

## ⚙️ Instalación

### <strong>Development </strong>

Esta configuración de docker usa el módulo de ([air](https://github.com/cosmtrek/air)), el multistage del Dockerfile se centra en target dev:


```Dockerfile
FROM golang:1.19 AS dev

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /home/app

EXPOSE 3000

CMD ["air"]
```

Solo hay que ejecutar:


```bash
docker compose -f docker-compose.dev.yml up
```

Esto expone el servidor http://localhost:3000.

### <strong>Production</strong>

La configuración de producción utiliza una capa de buildeo de la aplicación para compilar un binario que sería copiado y expuesto a una imagen de linux alpine sencilla, esta imagen se conecta mediante política de networking puente en docker compose que será redireccionada por un server Nginx reverse proxy.


```Dockerfile
# build for prod
FROM golang:alpine3.16 AS build

WORKDIR /home/app

COPY . .

RUN go build -o /home/go.app

# target prod
FROM alpine:3.16 AS prod

WORKDIR /home

COPY --from=build /home/go.app /home/go.app

# !important for github.com/joho/godotenv file .env
COPY .env /home/.env

EXPOSE 3000

CMD ["/home/go.app"]
```

Ejecutamos:

```bash
docker compose up 
```

Esto expone el servidor http://localhost:8000.

```go
package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env")
	}
}

```

## 🤖 Características 

El proyecto está construido con el módulo ([fiber](https://github.com/gofiber/fiber)) que se inspira en Express Js lo cual invitó a utilizar una estructura de carpetas similar a lo que utilizaremos en un proyecto de Backend productivo con Node JS. 

* <strong>controllers</strong>: Este directorio contiene las funcionalidades de cada ruta o recurso expuesto en el servidor, infiere en tener una funcionalidad aislada para cada situación.
* <strong>middlewares</strong>: Directorio donde se almacenan los prehandlers necesarios en la aplicación, como lo puede ser autenticación o autorización.
* <strong>models</strong>: Expone los modelos utilizados por la lógica de programación y que se representan en base de datos.
* <strong>utils</strong>: Su función es guardar lógica funcional reutilizable.

Se utiliza un módulo de carga de archivos .env espeficimente el módulo [godotenv](https://github.com/joho/godotenv) donde en este momento se hardcodean claves de conexión a base de datos.


```go
package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env")
	}
}

```

En un ambiente productivo esto debería de modificarse para no tener elementos hardcodeados en el repositorio, sobre todo si se considera una clave de acceso importante.


También expone la información de base de datos en un contenedor de [PMA](https://hub.docker.com/r/phpmyadmin/phpmyadmin/) en el servidor http://localhost:8001


## 🎯 Pruebas

Para realizar pruebas se puede utilizar curl con los siguientes comandos (importante recalcar que la url dependerá del ambiente seleccionado http://localhost:3000 o http://localhost:8000): 

Registro de nueva entidad usuario:

```bash
curl --location --request POST 'http://localhost:3000/signin' \
--header 'Content-Type: application/json' \
--data-raw '{
   "firstname":"Aldo",
   "lastname":"Trujillo",
   "email":"example@example.com",
   "password":"1234"
}'
```

Lectura de datos con basic auth:

```bash
curl --location --request GET 'http://localhost:3000/lyric?name=Costa A Costa&artist=El de La Guitarra&album=Con los Pies en La Tierra y la Mirada en el Cielo' \
-u "example@example.com:1234"
```

En curl cuando mandamos a Nginx una petición como la anterior se debe eliminar los espacios en la url, por ejemplo: 

"Costa A Costa" a "Costa-A-Costa"

Esto para evitar un badrequest en modo producción, también recomiendo copiar y pegar en postman.

## 👍 
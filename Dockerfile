# Utilizamos una imagen base de Go 1.22 para la compilación
FROM golang:1.23

# Establecemos el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copaimos los archivos del módulo Go y descargamos las dependencais
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copaimos el resto del código fuente de la aplicación
COPY . .

# Compilamos la aplicación a un binario
RUN go build -o ./out/dist ./cmd/api/

# Comando para ejecutar la aplicación
CMD ./out/dist
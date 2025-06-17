FROM golang:1.23.3-alpine

RUN apk add --no-cache git gcc g++ make

WORKDIR /app

# Copia primero los archivos del módulo para aprovechar el caché
COPY go.mod go.sum ./
RUN go mod download

# Luego copia el resto del código
COPY . .

EXPOSE 8087

# Aquí aún compila cada vez que arrancas
CMD ["go", "run", "main.go"]

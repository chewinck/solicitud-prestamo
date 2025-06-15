FROM golang:1.23.3-alpine

WORKDIR /solicitud-prestamo

COPY . .

RUN go mod download

EXPOSE 8087

CMD ["go", "run", "main.go"]


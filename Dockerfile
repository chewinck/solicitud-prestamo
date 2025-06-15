FROM golang:1.22-alpine 

WORKDIR /solicitud-prestamo

COPY . .

RUN go mod download

EXPOSE 8087

CMD ["go", "run", "main.go"]


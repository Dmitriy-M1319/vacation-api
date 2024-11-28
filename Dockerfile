FROM golang:1.23.0

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN apt update && apt install make

COPY . .

ENV HTTP_PORT 8080
ENV GRPC_PORT 12201

EXPOSE ${HTTP_PORT}/tcp
EXPOSE ${GRPC_PORT}/tcp

CMD [ "make" ]
# RUN go run cmd/vacation-api/main.go
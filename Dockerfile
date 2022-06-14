FROM golang:1.18

WORKDIR /usr/src/backend

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

EXPOSE 3051
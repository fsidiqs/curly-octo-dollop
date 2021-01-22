FROM golang:alpine

RUN mkdir /app

ADD . /app
WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 80
ENTRYPOINT /app/main --port 80
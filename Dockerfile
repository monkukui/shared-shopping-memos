FROM golang:alpine AS build-go

WORKDIR /app

COPY ./go.* ./

COPY ./*.go ./

RUN CGO_ENABLED=0 go build -o main main.go 

FROM busybox

WORKDIR /app

COPY --from=build-go /app/main ./main

EXPOSE 8080

CMD ./main

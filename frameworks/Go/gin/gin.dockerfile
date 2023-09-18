FROM docker.io/golang:1.20-bullseye as build-env

WORKDIR /gin
COPY ./gin-std /gin

RUN go mod download

RUN GOAMD64=v3 go build -ldflags="-w -s" -o hello

EXPOSE 8080

CMD ./hello

FROM gcr.io/distroless/base:debug
COPY --from=build-env /gin/hello /hello
ENTRYPOINT ["/hello"]
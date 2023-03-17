FROM golang:alpine3.17 AS build
WORKDIR /build
COPY ./src ./src
WORKDIR /build/src
RUN ls
RUN go get .
RUN go build -o /build/static-hoster

FROM alpine:3.17 AS final
COPY --from=build /build/static-hoster /bin/static-hoster
CMD static-hoster

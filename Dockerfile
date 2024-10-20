FROM golang:1.22.3-alpine as builder


WORKDIR /src
COPY . .

# install updates and build executable
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -a -installsuffix cgo -o app ./cmd

# copy to alpine image
FROM alpine:latest


ENV BOT_TOKEN="..."
# create user other than root and install updated
RUN addgroup -g 101 app && \
    adduser -H -u 101 -G app -s /bin/sh -D app
# place all necessary executables and other files into /app directory
WORKDIR /app/
COPY --from=builder --chown=app:app /src/app .

# run container as new non-root user
USER app

CMD ["/app/app"]
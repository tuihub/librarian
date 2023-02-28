FROM debian:stable-slim

COPY server /app/server

WORKDIR /app

EXPOSE 10000
EXPOSE 10001
VOLUME /data

CMD ["./server", "-conf", "/data/conf", "-data", "/data"]

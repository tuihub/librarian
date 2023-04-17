FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY server /app/server

WORKDIR /app

EXPOSE 10000
EXPOSE 10001
VOLUME /data

CMD ["./server", "-conf", "/data/conf", "-data", "/data"]

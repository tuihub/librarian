FROM scratch

COPY server /app/server

WORKDIR /app

EXPOSE 10000
EXPOSE 10001
VOLUME /data/conf

CMD ["./server", "-conf", "/data/conf"]

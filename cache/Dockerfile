
FROM alpine:3.15

ENV REDIS_VERSION 6.2.6

# Descargar y compilar Redis
RUN apk add --no-cache --virtual .build-deps build-base linux-headers \
    && wget http://download.redis.io/releases/redis-$REDIS_VERSION.tar.gz \
    && tar xzf redis-$REDIS_VERSION.tar.gz \
    && rm redis-$REDIS_VERSION.tar.gz \
    && cd redis-$REDIS_VERSION \
    && make \
    && make install \
    && cd .. \
    && rm -r redis-$REDIS_VERSION \
    && apk del .build-deps


WORKDIR /data

# Exponer el puerto de Redis
EXPOSE 6379

# Iniciar Redis
CMD ["redis-server", "--protected-mode", "no"]

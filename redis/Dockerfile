FROM redis:7-alpine

COPY redis.conf /usr/local/etc/redis/redis.conf
COPY users.acl /etc/redis/users.acl

ENTRYPOINT ["redis-server", "/usr/local/etc/redis/redis.conf"]

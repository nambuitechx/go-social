FROM redis:7

COPY redis.conf /usr/local/etc/redis/redis.conf
COPY users.acl /etc/redis/users.acl

ENTRYPOINT ["redis-server", "/usr/local/etc/redis/redis.conf"]

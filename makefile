POSTGRES_IMAGE = go-social/postgres
REDIS_IMAGE = go-social/redis
PROMETHEUS_IMAGE = go-social/prometheus
KAFKA_IMAGE = go-social/kafka
BACKEND_IMAGE = go-social/backend

build-postgres:
	docker build -t $(POSTGRES_IMAGE) ./postgres

build-redis:
	docker build -t $(REDIS_IMAGE) ./redis

build-prometheus:
	docker build -t $(PROMETHEUS_IMAGE) ./prometheus

build-kafka:
	docker build -t $(KAFKA_IMAGE) ./kafka

build-backend:
	docker build -t ${BACKEND_IMAGE} ./backend


build: build-postgres build-redis build-prometheus build-kafka build-backend

run:
	docker compose up -d -V

up_postgres:
	docker compose up postgres -d -V

up_redis:
	docker compose up redis -d -V

up_prometheus:
	docker compose up prometheus -d -V

up_kafka:
	docker compose up kafka -d -V

up: up_postgres up_redis up_prometheus up_kafka

down:
	docker compose down --volumes

# Đăng nhập vào ECR
login:
	aws ecr get-login-password --region $(REGION) | docker login --username AWS --password-stdin $(ECR_REGISTRY)

# Đăng xuất khỏi ECR
logout:
	docker logout $(ECR_REGISTRY)

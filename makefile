IMAGE_POSTGRES = go-social/postgres
IMAGE_REDIS = go-social/redis
IMAGE_PROMETHEUS = go-social/prometheus
IMAGE_BACKEND = go-social/backend

build-postgres:
	docker build -t $(IMAGE_POSTGRES) ./postgres

build-redis:
	docker build -t $(IMAGE_REDIS) ./redis

build-prometheus:
	docker build -t $(IMAGE_PROMETHEUS) ./prometheus

build-backend:
	docker build -t ${IMAGE_BACKEND} ./backend


build: build-postgres build-redis build-prometheus build-backend

run:
	docker compose up -d -V

up_postgres:
	docker compose up postgres -d -V

up_redis:
	docker compose up redis -d -V

up_prometheus:
	docker compose up prometheus -d -V

up: up_postgres up_redis up_prometheus

down:
	docker compose down

# Đăng nhập vào ECR
login:
	aws ecr get-login-password --region $(REGION) | docker login --username AWS --password-stdin $(ECR_REGISTRY)

# Đăng xuất khỏi ECR
logout:
	docker logout $(ECR_REGISTRY)

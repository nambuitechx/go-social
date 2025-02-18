IMAGE_POSTGRES = go-social/postgres
IMAGE_BACKEND = go-social/backend

build-postgres:
	docker build -t $(IMAGE_POSTGRES) ./postgres

build-backend:
	docker build -t ${IMAGE_BACKEND} ./backend


build: build-postgres build-backend

run:
	docker compose up -d -V

up_postgres:
	docker compose up postgres -d -V

down:
	docker compose down

# Đăng nhập vào ECR
login:
	aws ecr get-login-password --region $(REGION) | docker login --username AWS --password-stdin $(ECR_REGISTRY)

# Đăng xuất khỏi ECR
logout:
	docker logout $(ECR_REGISTRY)

IMAGE_NAME = ilyaksyonov/gotgbot:latest
CONTAINER_NAME = gotgbot
ENV_FILE = /opt/go_tg_bot/.env

DOCKERFILE_PATH = ./build/Dockerfile

build:
	docker build -t $(IMAGE_NAME) -f $(DOCKERFILE_PATH) .

pull:
	docker pull $(IMAGE_NAME)

run:
	docker run -d \
		--privileged \
	  	--name $(CONTAINER_NAME) \
  		--env-file $(ENV_FILE) \
		--ulimit nofile=1024:1024 \
  		--network host \
		$(IMAGE_NAME)

stop:
	docker stop $(CONTAINER_NAME) && docker rm $(CONTAINER_NAME) || echo "Skipping stop..."

clean:
	docker rmi $(IMAGE_NAME)

all: stop pull run

.PHONY: build run pull stop clean all

container_name=dp-zebedee-api-stub
app_name=api-stub
BIND_ADDR=8082

PHONY: build
build:
	go build -o ${app_name}

	env GOOS=linux GOARCH=amd64 go build -o ${app_name}-linux

PHONY: debug
debug: build
	./${app_name}

## Build a docker container to run FTB locally
PHONY: container
container:
	@echo "compiling linux binary"
	env GOOS=linux GOARCH=amd64 go build -o ${app_name}-linux

	@echo "building ${container_name}  container"
	docker build -t ${container_name} -f Dockerfile.stub --build-arg BIND_ADDR=${BIND_ADDR} .

## Run the FTB docker container locally (runs in detached state)
run: container
	@echo "running ${container_name} container"
	docker run -d \
		--name ${container_name} \
		-p 0.0.0.0:${BIND_ADDR}:${BIND_ADDR} \
		${container_name} \

## Stop and remove any existing FTB container and delete the image for a fresh build.
PHONY: clean
clean:
	@echo "stopping ${container_name} container"
	docker stop ${container_name} || true

	@echo "removing ${container_name} container"
	docker rm ${container_name} || true
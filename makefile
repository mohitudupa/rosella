APP = "rosella"
MODULE = "githib.com/mohitudupa/rosella"


dev:
	go run $(MODULE)
build:
	go build -o builds/$(APP) $(MODULE)
clean:
	rm builds/$(APP) || true
	docker image rm $(APP):latest || true
docker:
	@make clean
	@echo "Compiling module: $(MODULE)"
	@if [ $ARCH = "arm64" ] || [ $ARCH = "aarch64" ]; then\
		env GOARCH="arm64" GOOS="linux" go build -o builds/$(APP) $(MODULE);\
	else\
		env GOARCH="amd64" GOOS="linux" go build -o builds/$(APP) $(MODULE);\
	fi
	docker build -t $(APP):latest .
	docker container run -it --rm --name $(APP) -v ./config.json:/application/config.json -v ./sources:/application/sources -p 8080:8080 $(APP):latest


.PHONY: linux
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/managerapi main.go

.PHONY: mac
mac:
	go build -o ./build/managerapi main.go


.PHONY: clean
clean:
	rm -rf ./build

# 编译到 全部平台
.PHONY: build-all
all:
	make clean
	make linux
	make mac





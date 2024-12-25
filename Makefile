.PHONY: build run clean test deps

# 构建可执行文件
build:
	go build -o bin/cursor-fake cursor_fake.go

# 运行程序
run:
	go run cursor_fake.go

# 清理构建产物
clean:
	rm -rf bin/
	go clean

# 运行测试
test:
	go test -v ./...

# 更新依赖
deps:
	go mod tidy
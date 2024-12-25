.PHONY: build run clean test deps build-all release-notes

# 版本信息
LATEST_VERSION := $(shell sed -n 's/## \[\(.*\)\].*/\1/p' CHANGELOG.md | head -n 1)
GIT_COMMIT := $(shell git rev-parse --short HEAD)
BUILD_TIME := $(shell date -u '+%Y-%m-%d %H:%M:%S')
LDFLAGS := -X 'main.Version=$(LATEST_VERSION)' -X 'main.BuildTime=$(BUILD_TIME)' -X 'main.GitCommit=$(GIT_COMMIT)'

# 构建可执行文件
build:
	go build -ldflags "$(LDFLAGS)" -o bin/cursor-fake cursor_fake.go

# 构建所有平台版本
build-all:
	GOOS=darwin GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o bin/cursor-fake-darwin-arm64 cursor_fake.go
	GOOS=darwin GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o bin/cursor-fake-darwin-amd64 cursor_fake.go
	GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o bin/cursor-fake-linux-amd64 cursor_fake.go
	GOOS=windows GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o bin/cursor-fake-windows-amd64.exe cursor_fake.go

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

# 显示版本信息
version:
	@echo "Version: $(LATEST_VERSION)"
	@echo "Git Commit: $(GIT_COMMIT)"
	@echo "Build Time: $(BUILD_TIME)"

# 生成发布说明
release-notes:
	@sed -n "/## \[$(LATEST_VERSION)\]/,/## \[/p" CHANGELOG.md | sed '$$d' > release-notes.md
	@echo "\n---\n" >> release-notes.md
	@sed -n "/## \[$(LATEST_VERSION)\]/,/## \[/p" CHANGELOG_CN.md | sed '$$d' >> release-notes.md
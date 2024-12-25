# Cursor Fake

[English](./README.md) | [中文](./README_zh.md)

这是一个用于重置Cursor IDE设备ID的工具。它可以帮助你在需要时生成新的机器ID和设备ID。

## 功能特点

- 重置Cursor IDE设备标识
- 生成新的macMachineId和machineId
- 创建新的devDeviceId
- 跨平台支持（Windows、macOS、Linux）
- 简单的一键操作
- 自动化的多平台发布

## 工具作用

本工具可以帮助你：
- 重置Cursor IDE的设备识别
- 为系统生成新的机器ID
- 为Cursor创建新的设备ID
- 修改storage.json配置文件

工具将生成：
- 新的macMachineId
- 新的machineId
- 新的devDeviceId

## 快速开始

1. 从[发布页面](../../releases)下载适合你系统的版本：
   - Windows系统：`cursor-fake-windows-amd64.exe`
   - macOS系统（Apple Silicon）：`cursor-fake-darwin-arm64`
   - macOS系统（Intel）：`cursor-fake-darwin-amd64`
   - Linux系统：`cursor-fake-linux-amd64`

2. 运行程序：
   - Windows：双击 `.exe` 文件
   - macOS/Linux：打开终端运行：
     ```bash
     chmod +x cursor-fake-*  # 添加执行权限（仅首次���要）
     ./cursor-fake-*  # 运行程序
     ```

## 使用方法

1. 如果Cursor IDE正在运行，请先关闭
2. 按上述方式运行程序
3. 工具将自动：
   - 定位你的Cursor存储文件
   - 生成新的ID
   - 更新配置
4. 完成后按Enter键退出

## 开发者相关

如果你想从源码构建，需要：
- Go 1.21 或更高版本
- Make（用于构建自动化）

### 从源码构建

1. 克隆仓库
```bash
git clone [仓库地址]
cd cursor-fake
```

2. 构建当前平台版本
```bash
make build
```

或构建所有平台版本：
```bash
make build-all
```

### 可用命令

- `make build` - 构建当前平台版本
- `make build-all` - 构建所有平台版本
- `make run` - 运行程序
- `make test` - 运行测试
- `make clean` - 清理构建产物
- `make deps` - 更新依赖

## 贡献

欢迎提交Pull Request。对于重大变更，请先开issue讨论您想要改变的内容。

## 许可证

MIT License 
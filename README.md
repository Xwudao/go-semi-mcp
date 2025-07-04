# Semi-MCP

本项目是一个基于 [Semi-Design](https://semi.design/) 组件库的说明文档服务，支持通过 MCP 协议查询组件列表及组件用法。

## 功能

- 列出 Semi-Design 组件库中所有组件名称
- 查询指定组件的详细用法说明

## 使用方法

1. 克隆本仓库并安装依赖：

   ```bash
   git clone https://github.com/Xwudao/go-semi-mcp.git
   cd go-semi-mcp
   go mod tidy
   ```

2. 运行服务：

   ```bash
   go build main.go
   ```

   生成的main.exe即可使用stdio（command）方式进行调用。

## 主要依赖

- [mark3labs/mcp-go](https://github.com/mark3labs/mcp-go)
- [Semi-Design](https://semi.design/)

## 目录结构

- `main.go`：服务入口
- `tools/`：MCP 工具定义
- `assets/`：组件数据与文档

## 许可证

MIT License

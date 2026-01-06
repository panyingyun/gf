# gf

一个功能强大的Go语言实现的文件搜索命令行工具，兼具 `find` 和 `grep` 的功能。支持递归搜索指定目录，既可以搜索文件内容，也可以搜索文件名。

## 功能特性

- 🔍 **文件内容搜索** (`-f`): 递归搜索目录，查找文件内容包含指定字符串的文件，并输出文件路径、行号和匹配行内容
- 📁 **文件名搜索** (`-g`): 递归搜索目录，查找文件名包含指定字符串的文件，并输出文件路径
- 🚀 **高性能**: 使用Go语言实现，搜索速度快
- 🛡️ **错误处理**: 自动跳过无法访问的文件和目录，确保搜索过程稳定可靠
- 📊 **版本信息**: 支持查看版本号、构建时间和Git提交信息

## 安装

### 从源码编译

确保已安装 Go 1.22.5 或更高版本：

```bash
git clone https://github.com/panyingyun/gf.git
cd gf
go build -o gf .
```

### 使用 Makefile

```bash
make env  //only run once when first build 

make build
```

## 使用方法

### 基本语法

```bash
gf -g "pattern" <目录>    # 文件内容搜索
gf -f "pattern" <目录>    # 文件名搜索
gf -v                     # 显示版本信息
```

### 文件内容搜索 (`-g`)

递归搜索指定目录，查找文件内容包含指定字符串的文件：

```bash
gf -g "abc" /
```

输出格式：`文件路径:行号:匹配行内容`

示例输出：
```
/home/user/project/main.go:10:func main() {
/home/user/project/main.go:25:    fmt.Println("abc")
/home/user/project/utils.go:5:const abc = "test"
```

### 文件名搜索 (`-f`)

递归搜索指定目录，查找文件名包含指定字符串的文件：

```bash
gf -f "test" /home/user/project
```

输出格式：`文件路径`

示例输出：
```
/home/user/project/test.go
/home/user/project/test_utils.go
/home/user/project/tests/test_helper.go
```

### 查看版本信息

```bash
gf -v
```

输出示例：
```
gf v1.0.0, build at 2024-01-01T12:00:00+0800, commit abc1234
```

## 使用示例

### 搜索包含特定函数的文件

```bash
# 搜索包含 "handleRequest" 函数的所有文件
gf -g "handleRequest" ./src
```

### 搜索配置文件

```bash
# 查找所有包含 "config" 的文件名
gf -f "config" /etc
```

### 搜索特定代码模式

```bash
# 搜索包含 "TODO" 注释的文件
gf -g "TODO" ./project
```

### 搜索日志文件

```bash
# 查找所有日志文件
gf -f "log" /var/log
```

## 项目结构

```
gf/
├── main.go              # 主程序入口
├── go.mod              # Go模块定义
├── Makefile            # 构建脚本
├── .goreleaser.yml     # GoReleaser配置
├── LICENSE             # 许可证文件
├── README.md           # 项目说明文档
└── docs/               # 文档目录
    └── 实现go搜索命令行工具.md
```

## 开发

### 环境要求

- Go 1.22.5 或更高版本

### 本地开发

```bash
# 安装依赖
go mod tidy

# 格式化代码
gofumpt -l -w .

# 运行程序
go run main.go -f "pattern" <目录>
```

### 构建

```bash
# 使用 Makefile 构建
make build

# 或直接使用 go build
go build -o gf .
```

## 错误处理

程序会自动处理以下情况：

- 目录不存在：显示错误信息并退出
- 文件权限不足：跳过无法访问的文件，继续搜索
- 文件读取错误：跳过有问题的文件，继续搜索
- 无效参数：显示使用说明

## 许可证

详见 [LICENSE](LICENSE) 文件。

## 贡献

欢迎提交 Issue 和 Pull Request！

## 作者

panyingyun

## 相关链接

- GitHub: https://github.com/panyingyun/gf


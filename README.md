后续还会更新利用GO编写的命令行工具以及其他CVE的POC（由于有一些工具还没完成就没写在README里）

```markdown
# Go Security Tools

## 项目简介
本项目是一个基于 Go 语言开发的安全工具集合，包含以下功能模块：
- **子域名枚举**：支持主动和被动子域名扫描。
- **端口扫描**：快速扫描目标主机的开放端口。
- **漏洞扫描**：检测目标系统是否存在特定漏洞。
- **YAML 配置解析**：解析 YAML 文件以获取扫描规则。

## 功能模块

### 1. 子域名扫描
- **被动扫描**：使用 `subfinder` 库进行子域名的被动枚举。
- **主动扫描**：通过字典爆破的方式发现子域名，并支持解析子域名的 IP 地址。

### 2. 端口扫描
- 快速扫描目标主机的 0-65535 端口，支持自定义协议（如 TCP）。

### 3. 漏洞扫描
- 检测目标 URL 是否存在远程代码执行（RCE）漏洞。
- 支持自定义命令注入和结果验证。

### 4. YAML 配置解析
- 解析 YAML 文件以提取 HTTP 匹配规则。

## 项目结构
```
├── Internal/
│   ├── Domain/
│   │   └── SubDomain.go       # 子域名扫描模块
│   └── Infomation/            # 配置信息模块
├── comand/                    # 工具命令模块
├── getyaml/
│   └── getyaml.go             # YAML 配置解析模块
├── logger/                    # 日志模块
├── portscan/
│   └── Portscan2.go           # 端口扫描模块
├── rce/
│   └── rce.go                 # 漏洞扫描模块
├── go.mod                     # Go 模块依赖文件
└── README.md                  # 项目说明文件
```

## 安装与运行

### 环境要求
- Go 1.20 或更高版本
- Git

### 安装步骤
1. 克隆项目到本地：
   ```bash
   git clone https://github.com/your-username/your-repo-name.git
   cd your-repo-name
```

2. 安装依赖：
   ```bash
   go mod tidy
   ```

3. 编译项目：
   ```bash
   go build -o security-tools
   ```

### 使用说明

#### 子域名扫描
- **被动扫描**：
  ```bash
  ./security-tools subdomain passive -d example.com
  ```
- **主动扫描**：
  ```bash
  ./security-tools subdomain active -d example.com --wordlist /path/to/wordlist.txt
  ```

#### 端口扫描
```bash
./security-tools portscan -ip 192.168.1.1
```

#### 漏洞扫描
```bash
./security-tools rce -url https://example.com -cmd "whoami" -res "expected-result"
```

#### YAML 配置解析
```bash
./security-tools getyaml -file CVE-2024-38856.yaml
```

## 注意事项
- 请确保在合法授权的情况下使用本工具。
- 主动扫描可能会对目标系统产生影响，请谨慎操作。

## 贡献
欢迎提交 Issue 和 Pull Request 来改进本项目。

```

将 `your-username` 和 `your-repo-name` 替换为你的 GitHub 用户名和仓库名，并根据实际情况补充其他信息。

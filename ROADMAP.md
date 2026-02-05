# CLI Tool Roadmap

本文档描述了 CLI 工具的功能增强计划和开发路线图。

## 当前状态

### 已实现功能 ✅

- **`cli init`**: 初始化新的 CLI 项目
  - 支持单命令和多命令模式
  - 交互式、混合模式和非交互式
  - 自动生成项目结构（main.go, go.mod, README.md, .gitignore）
  - Flag 配置（类型、别名、环境变量、默认值）

- **`cli add`**: 向现有项目添加命令或 flag
  - 添加新命令到多命令 CLI
  - 添加 flag 到命令或全局
  - 混合模式支持

## 功能规划

### Phase 1: 核心项目管理功能 (优先级: 高)

#### 1.1 `cli remove` - 删除命令或 Flag
**目标**: 允许用户删除已存在的命令或 flag

**功能点**:
- `cli remove command <name>` - 删除命令
- `cli remove flag <name> [--from-command <cmd>]` - 删除 flag
- 交互式确认
- 自动更新代码
- 备份原文件（可选）

**用例**:
```bash
cli remove command delete
cli remove flag verbose --from-command list
```

#### 1.2 `cli list` - 列出项目信息
**目标**: 显示项目的命令和 flag 列表

**功能点**:
- `cli list commands` - 列出所有命令
- `cli list flags [--command <cmd>]` - 列出 flag
- `cli list` - 显示项目概览
- 格式化输出（table, json, yaml）

**用例**:
```bash
cli list
cli list commands
cli list flags --command list
cli list --format json
```

#### 1.3 `cli update` - 更新命令或 Flag
**目标**: 修改已存在的命令或 flag 属性

**功能点**:
- `cli update command <name>` - 更新命令
- `cli update flag <name> [--from-command <cmd>]` - 更新 flag
- 交互式编辑
- 支持部分更新（只更新指定属性）

**用例**:
```bash
cli update command list --usage "List all items"
cli update flag port --default 3000
```

#### 1.4 `cli validate` - 验证项目结构
**目标**: 检查项目是否符合 CLI 框架规范

**功能点**:
- 验证 main.go 结构
- 检查命令和 flag 定义
- 验证代码语法
- 提供修复建议

**用例**:
```bash
cli validate
cli validate --fix  # 自动修复问题
```

### Phase 2: 模板系统 (优先级: 高)

#### 2.1 `cli template` - 模板管理
**目标**: 支持自定义模板和模板库

**功能点**:
- `cli template list` - 列出可用模板
- `cli template add <name> <path>` - 添加本地模板
- `cli template remove <name>` - 删除模板
- `cli template use <name>` - 使用模板初始化项目
- 支持从 GitHub/GitLab 加载模板
- 模板变量系统

**用例**:
```bash
cli template list
cli template add my-template ./templates/my-template
cli init --template my-template
cli init --template github:user/repo/template
```

#### 2.2 内置模板
**目标**: 提供常用场景的预定义模板

**模板类型**:
- `basic` - 基础模板（当前默认）
- `api-client` - API 客户端模板
- `file-manager` - 文件管理工具模板
- `server` - 服务器管理模板
- `database` - 数据库工具模板
- `devops` - DevOps 工具模板

### Phase 3: 代码质量工具 (优先级: 中)

#### 3.1 `cli format` - 代码格式化
**目标**: 格式化生成的代码

**功能点**:
- 使用 gofmt 格式化
- 使用 goimports 整理导入
- 代码风格检查

**用例**:
```bash
cli format
cli format --check  # 只检查不修改
```

#### 3.2 `cli lint` - 代码检查
**目标**: 代码质量检查

**功能点**:
- 集成 golangci-lint
- 自定义规则配置
- 自动修复建议

**用例**:
```bash
cli lint
cli lint --fix
```

#### 3.3 `cli test` - 测试生成
**目标**: 自动生成测试文件

**功能点**:
- 为命令生成测试模板
- 为 flag 验证生成测试
- 集成测试模板

**用例**:
```bash
cli test generate
cli test generate --command list
```

### Phase 4: 项目增强功能 (优先级: 中)

#### 4.1 `cli upgrade` - 项目升级
**目标**: 升级项目到新版本的 CLI 框架

**功能点**:
- 检测当前版本
- 升级到最新版本
- 自动迁移代码
- 版本兼容性检查

**用例**:
```bash
cli upgrade
cli upgrade --to v1.2.0
cli upgrade --check  # 检查可用的升级
```

#### 4.2 `cli migrate` - 项目迁移
**目标**: 迁移项目结构或配置

**功能点**:
- 单命令 CLI 迁移到多命令 CLI
- 配置格式迁移
- 代码结构重构

**用例**:
```bash
cli migrate --to multiple
cli migrate --config new-format
```

#### 4.3 `cli config` - 配置管理
**目标**: 管理项目配置

**功能点**:
- `cli config get <key>` - 获取配置
- `cli config set <key> <value>` - 设置配置
- `cli config list` - 列出所有配置
- 配置文件管理（.clirc）

**用例**:
```bash
cli config set output.format json
cli config get template.default
cli config list
```

### Phase 5: 开发工具 (优先级: 中)

#### 5.1 `cli build` - 构建项目
**目标**: 构建 CLI 应用

**功能点**:
- 跨平台构建
- 构建优化
- 版本信息注入
- 构建配置

**用例**:
```bash
cli build
cli build --platform linux,amd64
cli build --output ./bin
```

#### 5.2 `cli run` - 运行项目
**目标**: 运行开发中的 CLI 应用

**功能点**:
- 自动检测并运行
- 参数传递
- 环境变量支持

**用例**:
```bash
cli run
cli run -- --help
```

#### 5.3 `cli watch` - 文件监听
**目标**: 监听文件变化并自动重新构建

**功能点**:
- 文件变化检测
- 自动重新构建
- 自动运行测试

**用例**:
```bash
cli watch
cli watch --build
cli watch --test
```

### Phase 6: 文档生成 (优先级: 中)

#### 6.1 `cli docs` - 文档生成
**目标**: 自动生成项目文档

**功能点**:
- 从代码生成 Markdown 文档
- 生成 API 文档
- 生成使用示例
- 支持多种输出格式

**用例**:
```bash
cli docs generate
cli docs generate --output ./docs
cli docs generate --format markdown
```

#### 6.2 `cli man` - Man Page 生成
**目标**: 生成 Unix man page

**功能点**:
- 自动生成 man page
- 安装到系统
- 多语言支持

**用例**:
```bash
cli man generate
cli man install
```

#### 6.3 `cli completion` - Shell 补全
**目标**: 生成 shell 自动补全脚本

**功能点**:
- 支持 bash, zsh, fish, powershell
- 自动安装
- 动态补全

**用例**:
```bash
cli completion generate bash
cli completion install bash
```

### Phase 7: 高级功能 (优先级: 低)

#### 7.1 `cli plugin` - 插件系统
**目标**: 支持插件扩展

**功能点**:
- `cli plugin list` - 列出插件
- `cli plugin install <name>` - 安装插件
- `cli plugin remove <name>` - 删除插件
- 插件开发 SDK

**用例**:
```bash
cli plugin list
cli plugin install cli-plugin-database
cli plugin remove cli-plugin-database
```

#### 7.2 交互式组件集成
**目标**: 在生成代码时自动集成交互式组件

**功能点**:
- 自动添加交互式组件选项
- 生成使用示例
- 配置向导

**用例**:
```bash
cli init --with-interactive
cli add interactive --type select
```

#### 7.3 Daemon 模式支持
**目标**: 生成 daemon 模式代码模板

**功能点**:
- 生成 daemon 配置
- 生成 systemd 服务文件
- 生成启动/停止脚本

**用例**:
```bash
cli init --with-daemon
cli add daemon
```

#### 7.4 版本管理
**目标**: 版本信息管理

**功能点**:
- `cli version` - 显示版本
- `cli check-update` - 检查更新
- 版本号自动管理

**用例**:
```bash
cli version
cli check-update
cli version bump --type patch
```

### Phase 8: 开发者体验 (优先级: 低)

#### 8.1 项目模板库
**目标**: 社区模板库

**功能点**:
- 模板市场
- 模板搜索
- 模板评分
- 模板分享

#### 8.2 代码片段库
**目标**: 常用代码片段

**功能点**:
- 代码片段管理
- 快速插入
- 片段分享

#### 8.3 调试工具
**目标**: 开发调试支持

**功能点**:
- 调试模式
- 日志输出
- 性能分析
- 错误追踪

## 实施优先级

### 短期 (1-2 个月)
1. `cli remove` - 删除功能
2. `cli list` - 列表功能
3. `cli validate` - 验证功能
4. 基础模板系统

### 中期 (3-6 个月)
1. `cli update` - 更新功能
2. 完整模板系统
3. 代码质量工具（format, lint）
4. 文档生成

### 长期 (6+ 个月)
1. 插件系统
2. 开发工具（build, run, watch）
3. 高级功能（daemon, interactive 集成）
4. 社区功能（模板库、代码片段）

## 技术债务

1. **代码解析改进**: 当前使用文本操作，应改为 AST 解析
2. **错误处理**: 增强错误处理和用户友好的错误消息
3. **测试覆盖**: 增加单元测试和集成测试
4. **性能优化**: 大文件处理优化
5. **国际化**: 多语言支持

## 贡献指南

欢迎贡献新功能！请参考以下步骤：

1. 在 GitHub Issues 中讨论新功能
2. 创建功能分支
3. 实现功能并添加测试
4. 更新文档
5. 提交 Pull Request

## 反馈

如有建议或问题，请通过以下方式反馈：
- GitHub Issues
- 文档反馈
- 社区讨论

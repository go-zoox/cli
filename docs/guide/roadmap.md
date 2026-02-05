# CLI Tool Roadmap

本文档概述了 CLI 工具的未来发展计划。详细内容请参考项目根目录的 [ROADMAP.md](https://github.com/go-zoox/cli/blob/master/ROADMAP.md) 和 [TODO.md](https://github.com/go-zoox/cli/blob/master/TODO.md)。

## 快速概览

### 当前功能 ✅

- `cli init` - 初始化新项目
- `cli add` - 添加命令或 flag
- 混合模式支持（命令行参数 + 交互式询问）

### 即将推出 🚀

#### Phase 1: 核心项目管理（高优先级）

1. **`cli remove`** - 删除命令或 flag
   - 安全删除，带确认
   - 自动代码更新

2. **`cli list`** - 列出项目信息
   - 显示所有命令和 flag
   - 多种输出格式

3. **`cli validate`** - 验证项目结构
   - 检查代码规范
   - 自动修复建议

4. **`cli update`** - 更新命令或 flag
   - 交互式编辑
   - 部分更新支持

#### Phase 2: 模板系统（高优先级）

- 自定义模板支持
- 模板库管理
- 内置常用模板（API 客户端、文件管理、服务器等）

#### Phase 3: 代码质量工具（中优先级）

- `cli format` - 代码格式化
- `cli lint` - 代码检查
- `cli test` - 测试生成

#### Phase 4: 开发工具（中优先级）

- `cli build` - 构建项目
- `cli run` - 运行项目
- `cli watch` - 文件监听

#### Phase 5: 文档生成（中优先级）

- `cli docs` - 自动生成文档
- `cli man` - 生成 man page
- `cli completion` - Shell 补全脚本

## 完整路线图

查看完整的功能规划和实施计划：
- [ROADMAP.md](https://github.com/go-zoox/cli/blob/master/ROADMAP.md) - 详细功能规划
- [TODO.md](https://github.com/go-zoox/cli/blob/master/TODO.md) - 具体任务列表

## 贡献

欢迎贡献新功能！请查看 TODO.md 中的任务列表，选择感兴趣的任务开始贡献。

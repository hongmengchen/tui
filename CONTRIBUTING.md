# 贡献

欢迎 issue 和 PR。这个项目主要按自己的使用场景维护，合并前会优先保证兼容性和安全边界。

## 开发

需要 Go 1.24+：

```bash
gofmt -w .
go vet ./...
go test ./...
go build ./...
```

本地验证时用假数据：不要拿真实的 token、kubeconfig、`.npmrc`、Vault 地址/角色或公司内部主机名做测试。

## 提交 PR

- 一个 PR 只做一件事，写清动机、行为和验证方式
- 行为变化补测试；文档和示例用中性占位值（`git.example.com`、`owner/repo`、`alice`）
- 不提交二进制、真实凭据或内部配置
- 提交信息用祈使句，说明“做了什么”，不用写版本号

## 安全问题

不要开公开 issue。请按 [SECURITY.md](SECURITY.md) 走私密报告。

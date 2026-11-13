# Changelog

## v1.4.2-rc.1 (2026-11-13)

- 升级 Go 工具链到 1.27.4, 修复标准库 CVE (GO-2026-6859, GO-2026-6809 等)
- golang.org/x/net v0.26.0 -> v0.62.0, 修复 HTTP/2 系列 DoS 漏洞 (GO-2026-6617, GO-2026-6612, GO-2026-6611, GO-2026-6603, GO-2026-5026)
- google.golang.org/grpc v1.63.2 -> v1.85.0, 修复 HTTP/2 DATA 帧堆内存耗尽 (GO-2026-6348) 及 xDS RBAC 漏洞 (GO-2026-6061)
- golang.org/x/text v0.16.0 -> v0.44.0, 修复非法输入死循环 (GO-2026-5970)
- 新增 pkg/utils 桩包: 该包在开源版中被裁剪但仍有引用, 仅为通过编译, 不含真实功能
- govulncheck 符号级扫描可达漏洞清零
- 将模块路径, 源码导入和文档链接中的仓库地址统一迁移到 git.ovh.im/yeqing

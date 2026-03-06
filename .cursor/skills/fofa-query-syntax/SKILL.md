---
name: fofa-query-syntax
description: 生成、改写并校验 FOFA 查询语句。用于用户提到 FOFA、资产测绘、语法速查、查询语句编写、语句纠错、条件组合优化时。
---

# FOFA Query Syntax

## 使用范围

- 仅允许使用该文件中存在的字段（`key`）。
- 当用户要求“写 FOFA 语句 / 改写语句 / 检查语法 / 给出等价查询”时，优先应用本 Skill。

## 语法基线

- 基础条件写法：`field="value"`。
- 取反条件写法：`field!="value"`（仅当字段支持 `not_equal`）。
- 条件连接：使用 `&&` / `||` 进行组合。
- 布尔字段按示例使用：如 `is_domain=true`。
- 时间字段值格式：`YYYY-MM-DD`。
- 同源 IP 函数字段按示例调用：`ip_filter(...)`、`ip_exclude(...)`。

## 运算符能力规则（必须遵守）

按约束使用运算符：

- `equal + not_equal + fuzzy_search`（可用 `=`、`!=`、模糊）  
  `port, domain, host, os, server, asn, org, cloud_name, protocol, title, header_hash, js_name, js_md5, cname, cname_domain, icp, cert.subject, cert.issuer, cert.subject.org, cert.subject.cn, cert.issuer.org, cert.issuer.cn, cert.domain, jarm, tls.ja3s`
- `equal + not_equal`（可用 `=`、`!=`）  
  `ip, fid, product, product.version, category, banner, banner_hash, banner_fid, base_protocol, header, body, body_hash, icon_hash, status_code, sdk_hash, country, region, city, cert, tls.version, cert.sn, port_size`
- `equal only`（仅可 `=`）  
  `is_domain, is_ipv6, app, type, is_cloud, is_fraud, is_honeypot, cert.is_equal, cert.is_valid, cert.is_match, cert.is_expired, cert.not_after.after, cert.not_after.before, cert.not_before.after, cert.not_before.before, after, before, after&before, ip_filter(), ip_exclude(), port_size_gt, port_size_lt, ip_ports, ip_country, ip_region, ip_city, ip_after, ip_before`

如果用户给了不被支持的写法（例如对 `equal only` 字段使用 `!=`），必须改写为合法表达式并说明改写原因。

## 字段分组速查（按 category.group）

- `group 0` 资产基础  
  `ip, port, domain, host, os, server, asn, org, is_domain, is_ipv6`
- `group 1` 应用与资产分类  
  `app, fid, product, product.version, category, type, cloud_name, is_cloud, is_fraud, is_honeypot`
- `group 2` 协议与横幅  
  `protocol, banner, banner_hash, banner_fid, base_protocol`
- `group 3` Web 内容与页面特征  
  `title, header, header_hash, body, body_hash, js_name, js_md5, cname, cname_domain, icon_hash, status_code, icp, sdk_hash`
- `group 4` 地理信息  
  `country, region, city`
- `group 5` 证书信息  
  `cert, cert.subject, cert.issuer, cert.subject.org, cert.subject.cn, cert.issuer.org, cert.issuer.cn, cert.domain, cert.is_equal, cert.is_valid, cert.is_match, cert.is_expired, jarm, tls.version, tls.ja3s, cert.sn, cert.not_after.after, cert.not_after.before, cert.not_before.after, cert.not_before.before`
- `group 6` 时间筛选  
  `after, before, after&before`
- `group 7` 同源 IP 与独立 IP 聚合  
  `ip_filter(), ip_exclude(), port_size, port_size_gt, port_size_lt, ip_ports, ip_country, ip_region, ip_city, ip_after, ip_before`

## 权限提醒规则

若命中以下字段，输出时追加版本提示：

- `（个人版及以上）`  
  `banner_hash, banner_fid, header_hash, cert.is_equal, cert.is_valid, cert.is_match, cert.is_expired, after, before`
- `（专业版及以上）`  
  `is_fraud, is_honeypot`
- `（商业版及以上）`  
  `sdk_hash, ip_filter(), ip_exclude(), port_size, port_size_gt, port_size_lt, ip_ports, ip_country, ip_region, ip_city, ip_after, ip_before`

## 助手工作流（生成/改写/校验）

按以下顺序执行，不跳步：

1. 明确目标  
   - 提取用户目标（找什么资产、在什么范围、排除什么）。
2. 映射字段  
   - 从分组速查选择字段，优先语义最直接字段。
3. 选择运算符  
   - 严格按“运算符能力规则”选择 `=` / `!=` / 模糊。
4. 组装语句  
   - 先给 1 条主查询，再给 1-2 条等价改写（更宽/更窄）。
5. 语法自检  
   - 检查未知字段、非法运算符、明显冲突（如同字段互斥值）。
6. 输出说明  
   - 简述每个条件含义；有权限字段时附版本提醒。

## 信息不充分时的处理

- 若用户目标缺少关键维度（目标类型/地理范围/时间范围/排除条件），先补问 1-2 个关键问题再生成。
- 若用户明确要求“直接给查询”，可先给一版可执行语句，并标注默认假设。
- 默认假设优先级：
  - 默认使用 `=` 写法，不擅自使用 `!=`。
  - 默认不加时间范围，除非用户明确要求“最近/历史”。
  - 默认不使用付费字段，除非用户主动要求或场景强依赖。

## 自然语言到字段映射规则

- “独立 IP / 同 IP 多端口 / 旁站”优先映射 `group 7`：`ip_filter()`, `ip_exclude()`, `port_size*`, `ip_ports`。
- “网站标题/正文/响应头/图标”优先映射 `group 3`：`title`, `body`, `header`, `icon_hash` 等。
- “证书主体/签发者/证书有效性/过期”优先映射 `group 5`：`cert.*`。
- “资产更新时间范围”优先映射 `group 6`：`after`, `before`（优先，不建议使用 `after&before`）。
- “国家/省份/城市”优先使用：
  - 一般资产：`country`, `region`, `city`
  - 独立 IP 聚合：`ip_country`, `ip_region`, `ip_city`

## 冲突与风险检测清单

在“语法检查”中至少检查以下项：

- 字段拼写冲突：如 `product_version`（应为 `product.version`）。
- 运算符冲突：`equal only` 字段误用 `!=`。
- 类型冲突：布尔字段写成字符串（`"true"`）。
- 时间冲突：`after` 晚于 `before`。
- 语义冲突：同字段在同一逻辑分支出现互斥条件（如 `is_ipv6=true && is_ipv6=false`）。
- 权限风险：命中付费字段但未提醒版本门槛。

## 输出质量标准

- 至少返回 1 条“主查询”+ 1 条“等价改写”。
- 解释每个条件的筛选意图，不只给语句本身。
- 改写版本应体现“更宽”或“更窄”的差异，而不是仅改写排版。
- 发生纠错时，必须同时给“原写法 -> 修正写法 -> 修正原因”。

## 输出模板

```markdown
主查询：
<query>

等价改写：
1) <query_variant_1>
2) <query_variant_2>

语法检查：
- 字段合法性：通过/不通过（原因）
- 运算符合法性：通过/不通过（原因）
- 条件冲突：无/有（说明）

说明：
- 条件解释...
- 版本权限提醒（如有）...
```

## 代表性示例

- 基础网络资产  
  - `ip="1.1.1.1"`  
  - `port="6379"`  
  - `domain="qq.com"`
- Web 指纹与内容  
  - `title="beijing" && body="网络空间测绘"`  
  - `header_hash="1258854265" && icon_hash="-247388890"`
- 证书场景  
  - `cert.subject="Oracle Corporation" && cert.is_valid=true`  
  - `cert.not_after.before="2025-03-01"`
- 时间场景  
  - `after="2023-01-01" && before="2023-12-01"`
- 同源 IP 场景  
  - `ip_filter(banner="SSH-2.0-OpenSSH_6.7p2") && ip_filter(icon_hash="-1057022626")`  
  - `ip_filter(banner="SSH-2.0-OpenSSH_6.7p2" && asn="3462") && ip_exclude(title="EdgeOS")`

## 纠错规则

- 未知字段：提示“字段不存在，并给最接近候选字段。
- 运算符不合法：改写为支持形式（如 `!=` 改为 `=` 并调整逻辑）。
- 布尔值误写为字符串：改为 `true/false`（不加引号）。
- 时间格式错误：改为 `YYYY-MM-DD`。
- `after&before` 标记为不推荐：优先改写为 `after="..." && before="..."`。

## 纠错输出格式（固定）

```markdown
发现问题：
1) <原写法>

修正建议：
1) <修正写法>

原因：
- <对应 grammar 能力限制或字段定义>
```

## 约束

- 不虚构字段和能力。
- 不假设未声明的函数或语法糖。
- 如果需求超出 FOFA 字段能力，明确指出限制并给可执行近似方案。

## 参考资料

- 完整字段目录、能力矩阵、权限门槛见 [reference.md](reference.md)。

# FOFA Query Syntax Reference

本文件是 `SKILL.md` 的审计参考。

## 图例

- 运算符能力：
  - `ENF` = 支持 `=`、`!=`、fuzzy_search
  - `EN` = 支持 `=`、`!=`
  - `E` = 仅支持 `=`
- 版本权限：
  - `Personal+` = 个人版及以上
  - `Pro+` = 专业版及以上
  - `Business+` = 商业版及以上
  - `-` = 无额外版本标注

## 全字段目录（75）

### group 0 资产基础（10）

| 字段 | 能力 | 权限 |
| --- | --- | --- |
| `ip` | EN | - |
| `port` | ENF | - |
| `domain` | ENF | - |
| `host` | ENF | - |
| `os` | ENF | - |
| `server` | ENF | - |
| `asn` | ENF | - |
| `org` | ENF | - |
| `is_domain` | E | - |
| `is_ipv6` | E | - |

### group 1 应用与分类（10）

| 字段 | 能力 | 权限 |
| --- | --- | --- |
| `app` | E | - |
| `fid` | EN | - |
| `product` | EN | - |
| `product.version` | EN | - |
| `category` | EN | - |
| `type` | E | - |
| `cloud_name` | ENF | - |
| `is_cloud` | E | - |
| `is_fraud` | E | Pro+ |
| `is_honeypot` | E | Pro+ |

### group 2 协议与横幅（5）

| 字段 | 能力 | 权限 |
| --- | --- | --- |
| `protocol` | ENF | - |
| `banner` | EN | - |
| `banner_hash` | EN | Personal+ |
| `banner_fid` | EN | Personal+ |
| `base_protocol` | EN | - |

### group 3 Web 内容与页面特征（13）

| 字段 | 能力 | 权限 |
| --- | --- | --- |
| `title` | ENF | - |
| `header` | EN | - |
| `header_hash` | ENF | Personal+ |
| `body` | EN | - |
| `body_hash` | EN | - |
| `js_name` | ENF | - |
| `js_md5` | ENF | - |
| `cname` | ENF | - |
| `cname_domain` | ENF | - |
| `icon_hash` | EN | - |
| `status_code` | EN | - |
| `icp` | ENF | - |
| `sdk_hash` | EN | Business+ |

### group 4 地理信息（3）

| 字段 | 能力 | 权限 |
| --- | --- | --- |
| `country` | EN | - |
| `region` | EN | - |
| `city` | EN | - |

### group 5 证书与 TLS（20）

| 字段 | 能力 | 权限 |
| --- | --- | --- |
| `cert` | EN | - |
| `cert.subject` | ENF | - |
| `cert.issuer` | ENF | - |
| `cert.subject.org` | ENF | - |
| `cert.subject.cn` | ENF | - |
| `cert.issuer.org` | ENF | - |
| `cert.issuer.cn` | ENF | - |
| `cert.domain` | ENF | - |
| `cert.is_equal` | E | Personal+ |
| `cert.is_valid` | E | Personal+ |
| `cert.is_match` | E | Personal+ |
| `cert.is_expired` | E | Personal+ |
| `jarm` | ENF | - |
| `tls.version` | EN | - |
| `tls.ja3s` | ENF | - |
| `cert.sn` | EN | - |
| `cert.not_after.after` | E | - |
| `cert.not_after.before` | E | - |
| `cert.not_before.after` | E | - |
| `cert.not_before.before` | E | - |

### group 6 时间筛选（3）

| 字段 | 能力 | 权限 |
| --- | --- | --- |
| `after` | E | Personal+ |
| `before` | E | Personal+ |
| `after&before` | E | - |

备注：`after&before` 标记 `notRecommend`，建议改写为 `after="..." && before="..."`。

### group 7 同源 IP 与独立 IP 聚合（11）

| 字段 | 能力 | 权限 |
| --- | --- | --- |
| `ip_filter()` | E | Business+ |
| `ip_exclude()` | E | Business+ |
| `port_size` | EN | Business+ |
| `port_size_gt` | E | Business+ |
| `port_size_lt` | E | Business+ |
| `ip_ports` | E | Business+ |
| `ip_country` | E | Business+ |
| `ip_region` | E | Business+ |
| `ip_city` | E | Business+ |
| `ip_after` | E | Business+ |
| `ip_before` | E | Business+ |

## 常用示例

- `ip="220.181.111.1/24"`
- `app="Microsoft-Exchange"`
- `banner_hash="7330105010150477363"`
- `title="beijing"`
- `cert.subject="Oracle Corporation"`
- `cert.is_valid=true`
- `after="2023-01-01" && before="2023-12-01"`
- `ip_filter(banner="SSH-2.0-OpenSSH_6.7p2") && ip_exclude(title="EdgeOS")`

## 审计检查建议

- 字段名必须与表格完全一致（特别是 `product.version`、`cert.*`、函数字段带括号）。
- `E` 字段禁止使用 `!=`。
- 命中 `Personal+` / `Pro+` / `Business+` 字段时必须显式提示版本门槛。

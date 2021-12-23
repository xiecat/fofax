# fofaX

[![Latest release](https://img.shields.io/github/v/release/xiecat/fofax)](https://github.com/xiecat/fofax/releases/latest)![GitHub Release Date](https://img.shields.io/github/release-date/xiecat/fofax)![GitHub All Releases](https://img.shields.io/github/downloads/xiecat/fofax/total)[![GitHub issues](https://img.shields.io/github/issues/xiecat/fofax)](https://github.com/xiecat/fofax/issues)

## 0x00 简介

fofaX是一款命令行fofa查询工具，定位命令行，简单就是最好的！

## 0x01 下载

点击 [Releases下载链接](https://github.com/xiecat/fofax/releases) ，按照自己的系统来选择相应的发行版本。

## 0x02 配置

### Windows



### macOS/Linux





## 0x03 使用

### 使用Tips

不带任何参数时，除了会输出ASCII Logo，还会随机输出一条使用Tips。

```bash
fofax
```

![tips](./docs/images/tips.png)



### 帮助信息

可以使用 `fofax -h` 显示帮助信息。

```bash
fofax -h
```

![help](./docs/images/help.png)

### FOFA语法规则

使用 `-use` 参数，显示FOFA语法查询规则。

```bash
fofax -use
```

![use](./docs/images/use.png)

### 基础查询

有如下两种方式查询 `app="APACHE-Solr"`，不指定数量默认会输出100个host，并且默认会对数据进行去重。

```bash
fofax -q 'app="APACHE-Solr"'
```

![solr-1](./docs/images/solr-1.png)

```bash
echo 'app="APACHE-Solr"' | fofax
```

![solr-2](./docs/images/solr-2.png)

### 指定查询数量

```bash
echo 'app="APACHE-Solr"' | fofax -fs 5
```

![fs](./docs/images/fs.png)

如上只输出了4条数据，是因为fofax自动对重复的数据进行了去重（不信可以看fofa API的返回数据）。

![fofaapi](./docs/images/fofaapi.png)

### 获取URL

添加 `-ffi` 参数，根据查询语句直接获取对应的URL（[scheme]://[host]:[port]）。

```bash
echo 'app="APACHE-Solr"' | fofax -fs 5 -ffi
```

![url](./docs/images/url.png)

### Debug模式

添加 `-debug` 参数，开启Debug详细模式。

```bash
echo 'app="APACHE-Solr"' | fofax -fs 5 -ffi -debug
```

![debug](./docs/images/debug.png)

### 浏览器中打开

```bash
echo 'app="APACHE-Solr"' | fofax -open
```

![openinbrowser](./docs/images/openinbrowser.gif)

### icon hash查询

两种方式，第一种是直接根据提供icon的URL来查询。

```bash
fofax -iu https://www.baidu.com/favicon.ico -fs 5
```

![iconhash-1](./docs/images/iconhash-1.png)

第二种是根据本地icon文件，来计算hash并查询。

![iconhash-2](./docs/images/iconhash-2.png)





## 0x04 联动使用案例

### fofax && httpx

CVE-2021-43798 Grafana未授权目录遍历。

![fofax&httpx](./docs/images/fofax&httpx.png)

### fofax && nuclei

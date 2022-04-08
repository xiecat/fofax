# fofaX

[![Latest release](https://img.shields.io/github/v/release/xiecat/fofax)](https://github.com/xiecat/fofax/releases/latest)  ![GitHub Release Date](https://img.shields.io/github/release-date/xiecat/fofax)  ![GitHub All Releases](https://img.shields.io/github/downloads/xiecat/fofax/total)  [![GitHub issues](https://img.shields.io/github/issues/xiecat/fofax)](https://github.com/xiecat/fofax/issues)  ![LICENSE](https://img.shields.io/badge/LICENSE-GPL-ff69b4)  


[:ledger: 中文 README](https://github.com/xiecat/fofax/blob/main/README.zh-CN.md)   |   [:pushpin: Releases Download](https://github.com/xiecat/fofax/releases)  |  [:racehorse: 使用文档 Docs](http://fofax.xiecat.fun/)
## 0x00 Introduction

fofax is a fofa query tool written in go, positioned as a command-line tool and characterized by simplicity and speed. The following features are currently available:

- Basic FOFA syntax queries
- Icon Hash local/online calculation query
- Asset filtering
- Opening in browser
- Linking other security tools
- More (waiting for your feedback after using) ......


In addition to this it is possible to customize fx syntax queries, and users can write their own specific fx query rules via a configuration file in yaml format.


## 0x01 Download

Click on the [Release download link](https://github.com/xiecat/fofax/releases) and choose the appropriate release for your system architecture.

## 0x02 Configuration

### macOS/Linux

Unzip the downloaded fofax archive and recommend placing it in the `/usr/local/bin/` directory, so that you can run fofax commands from any directory.

```Console
tar -zxvf ~/Downloads/fofax_v0.1.22_darwin_amd64.tar.gz -C /usr/local/bin/.
```


The first time you run the fofax command a configuration file is automatically generated, located at `~/.config/fofax/fofax.yaml`.

```console
fofax

      ____        ____       _  __
     / __/____   / __/____ _| |/ /
    / /_ / __ \ / /_ / __ `/|   /
   / __// /_/ // __// /_/ //   |
  /_/   \____//_/   \__,_//_/|_|
                                    
                         fofax.xiecat.fun

2021/12/23 21:21:28 [SUCC] create config file /Users/user/.config/fofax/fofax.yaml. please modify and use
```

The next step is to configure this configuration file. Generally you only need to configure `email` and `key`.

```console
vim ~/.config/fofax/fofax.yaml
```

```console
# fofa api email
fofa-email: ******@gmail.com

# fofa api key
fofakey: ***************
```

### Windows

Unzip the archive and run fofax.exe for the first time to generate a fofax.yaml configuration file in the same level directory. Then open this configuration file and fill in `email` and `key`.

## 0x03 How to use

### Usage tips

Without any parameters, in addition to the ASCII logo, a random usage hint will be output.

```console
fofax

      ____        ____       _  __
     / __/____   / __/____ _| |/ /
    / /_ / __ \ / /_ / __ `/|   /
   / __// /_/ // __// /_/ //   |
  /_/   \____//_/   \__,_//_/|_|
                              
                         fofax.xiecat.fun

fofaX is a command line fofa query tool, simple is the best!

Tips:
Comment: Search google-reverse in fx, the -fe parameter must be added to the query when using the extension
Usage: fofax -q 'fx="google-reverse"' -fe
```

### Help information

You can use `fofax -h` to display help information.

```console
fofax -h

fofaX is a command line fofa query tool, simple is the best!

Usage:
  fofax [flags]

Flags:
CONFIGS:
   -email, -fofa-email string  Fofa API Email
   -key, -fofakey string       Fofa API Key
   -p, -proxy string           proxy for http like http://127.0.0.1:8080
   -fofa-url string            Fofa url (default "https://fofa.so")
   -debug                      Debug mode

FILTERS:
   -fs, -fetch-size int          The maximum number of query (default 100)
   -e, -exclude                  Exclude the honeypot.
   -ec, -exclude-country-cn      Exclude CN.
   -ffi, -fetch-fullHost-info    URL fetch, with scheme, hostname, port
   -fto, -fetch-titles-ofDomain  Fetch website title

SINGLE QUERY/ERT/ICON:
   -q, -query string              FoFa query statement
   -uc, -url-cert string          Enter the certificate of the https URL to query
   -iu, -url-to-icon-hash string  Enter the URL of an icon, calculate it and query it
   -if, -icon-file-path string    Calculate the hash of the local icon file, then query it

MULTIPLE QUERY/CERT/ICON:
   -qf, -query-file string           Load files, query multiple statements
   -ucf, -url-cert-file string       Read the URL from the file, calculate the cert and then query it
   -iuf, -icon-hash-url-file string  Retrieve the URL from the file, calculate the icon hash and query it

FX GRAMMER:
   -g, -gen string           Generate fx statement files eg: default_fx.yaml
   -fd, -fxdir string        fxdir directory (default "/Users/user/.config/fofax/fxrules")
   -l, -lists                List of fx statements
   -lt, -list-tags           List fx tags
   -s, -search string        Search for fx statements. Statements are separated by semicolons eg: id=fx-2021-01;query="jupyter Unauth"
   -tree                     Print syntax tree
   -fe, -fofa-ext            Using extended syntax(fx)
   -ss, -show-single string  Display a single fx message

OTHER OPTIONS:
   -config string  fofax configuration file.The file reading order(fofax.yaml,/Users/user/.config/fofax/fofax.yaml,/etc/fofa.yaml) (default "/Users/user/.config/fofax/fofax.yaml")
   -v, -version    Show fofaX version
   -use            Syntax queries
   -open           Open with your browser only support pipline/-q/-uc/-iu/-if
   -no-limit-open  No limit to the number of openings in your browser
```

### FOFA Syntax Rules

Use the `-use` argument to display FOFA syntax query rules.

```console
fofax -use

┌───────────────────────────────────────────┬──────────────────────────────────────────────────┐
│ Query                                     │ Explanation                                      │
├───────────────────────────────────────────┼──────────────────────────────────────────────────┤
│ title="beijing"                           │ 从标题中搜索"北京"                               	  │
│ header="elastic"                          │ 从http头中搜索"elastic"                          	 │
│ body="网络空间测绘"                         │ 从html正文中搜索"网络空间测绘"                       │
│ title="beijing"                           │ 从标题中搜索"北京"                              	  │
│ header="elastic"                          │ 从http头中搜索"elastic"                            │
│ body="网络空间测绘"                         │ 从html正文中搜索"网络空间测绘"                       │
│ fid="kIlUsGZ8pT6AtgKSKD63iw=="            │ 查找相同的网站指纹                               	  │
│ domain="google.com"                       │ 搜索根域名带有google.com的网站                   	  │
│ icp="京ICP证030173号"                      │ 查找备案号为"京ICP证030173号"的网站                  │
│ js_name="js/jquery.js"                    │ 查找网站正文中包含js/jquery.js的资产             	   │
│ js_md5="82ac3f14327a8b7ba49baa208d4eaa15" │ 查找js源码与之匹配的资产                         	   │
└───────────────────────────────────────────┴──────────────────────────────────────────────────┘
```

### Basic queries

There are two ways to query `app="APACHE-Solr"`, which will output 100 hosts by default without specifying the number, and will de-duplicate the data by default.

```console
fofax -q 'app="APACHE-Solr"'

2021/12/23 20:17:32 [SUCC] Fetch Data From FoFa: [100/30830]
54.114.20.168:8443
193.8.4.43:8983
208.37.227.95:8983
3.20.255.140:8983
3.114.85.178:8983
82.142.82.197:8983
159.39.10.212:8983
199.102.27.69:8983
…………
```

```console
echo 'app="APACHE-Solr"' | fofax
2021/12/23 20:17:59 [SUCC] Fetch Data From FoFa: [100/30830]
54.114.20.168:8443
193.8.4.43:8983
208.37.227.95:8983
3.20.255.140:8983
3.114.85.178:8983
82.142.82.197:8983
159.39.10.212:8983
199.102.27.69:8983
…………
```

```console
echo 'app="APACHE-Solr"' | fofax
2021/12/23 20:17:59 [SUCC] Fetch Data From FoFa: [100/30830]
54.114.20.168:8443
193.8.4.43:8983
208.37.227.95:8983
3.20.255.140:8983
3.114.85.178:8983
82.142.82.197:8983
159.39.10.212:8983
199.102.27.69:8983
............
```

### Specify the number of queries

```console
echo 'app="APACHE-Solr"' | fofax -fs 5
2021/12/23 20:19:00 [SUCC] Fetch Data From FoFa: [5/30830]
13.57.71.190:8443
165.22.215.32:8983
184.73.40.143:8443
3.20.255.140:8983
```

As above, only 4 data are output because fofax automatically de-duplicates the data.

### Exclude queries

Add the `-e` argument to exclude honeypots.

```console
echo 'app="APACHE-Solr"' | fofax -fs 10 -e
2021/12/23 22:56:14 [SUCC] Fetch Data From FoFa: [10/30849]
13.126.128.253:80
185.22.235.14:8983
151.248.126.4:8983
20.71.77.183:80
23.102.46.20:443
15.113.170.101:8443
52.58.201.109:80
```

Add the `-ec` parameter to exclude data from China.

```console
echo 'app="APACHE-Solr"' | fofax -fs 10 -ec
2021/12/23 22:56:36 [SUCC] Fetch Data From FoFa: [10/26044]
15.113.170.101:8443
52.58.201.109:80
13.126.128.253:80
185.22.235.14:8983
151.248.126.4:8983
20.71.77.183:80
23.102.46.20:443
```

### Get the URL

Add the `-ffi` parameter to get the corresponding URL directly based on the query statement ([scheme]://[host]:[port]).

```console
echo 'app="APACHE-Solr"' | fofax -fs 5 -ffi
2021/12/23 20:21:03 [SUCC] Fetch Data From FoFa: [5/30830]
https://184.73.40.143:8443
http://120.24.42.244:8983
https://13.57.71.190:8443
http://165.22.215.32:8983
```

### Get Title

```console
echo 'domain="baidu.com" && status_code="200"' | fofax -fs 10 -fto
2021/12/23 20:21:19 [SUCC] Fetch Data From FoFa: [10/1124]
https://home.baidu.com [关于百度]
http://research.baidu.com [Baidu Research]
http://fecs.baidu.com [FECS - Front End Code Style Suite]
http://yuntu.baidu.com [企业图谱]
https://ditu.baidu.com [百度地图]
https://sp2.baidu.com [百度一下，你就知道]
https://tushuo.baidu.com [图说]
https://ocpc.baidu.com [百度 oCPC 开发者中心]
https://naotu.baidu.com [百度脑图 - 便捷的思维工具]
http://usa.baidu.com [Baidu USA]
```

### Debug mode

Add the `-debug` parameter to enable Debug detail mode.

```console
echo 'app="APACHE-Solr"' | fofax -fs 5 -ffi -debug
2021/12/25 21:28:57 [DEBUG] FoFa Size : 5
2021/12/25 21:28:57 [DEBUG] FoFa Query of: app="APACHE-Solr"
2021/12/25 21:28:57 [DEBUG] https://fofa.so/api/v1/search/all?email=**********@gmail.com&key=**************************&qbase64=YXBwPSJBUEFDSEUtU29sciI=&size=5&page=1&fields=protocol,ip,port,host
2021/12/25 21:28:57 [DEBUG] Resp Time: 432/millis
2021/12/25 21:28:57 [SUCC] Fetch Data From FoFa: [5/30942]
http://35.183.115.103
http://3.17.203.145:8983
http://195.201.119.15:49154
https://18.169.23.120
http://174.138.127.51:8983
```

### Open in browser

```console
echo 'app="APACHE-Solr"' | fofax -open
```

<!-- ![openinbrowser](./docs/images/openinbrowser.gif) -->
![openinbrowser.gif](https://s2.loli.net/2021/12/25/2lvs4njYEUNmkLS.gif)

### Calculate icon hash and query

There are two ways to do this, the first is to query directly based on the URL that provided the icon.

```console
fofax -iu https://www.baidu.com/favicon.ico -fs 5

2021/12/23 20:21:59 [SUCC] Fetch Data From FoFa: [5/13284]
47.98.104.77:8088
154.39.217.22:80
xueshu.mrsb.tk:80
154.39.217.2:80
154.39.217.28:80
```

The second one is to calculate hash and query based on local icon file.

```console
wget https://www.baidu.com/favicon.ico
fofax -if favicon.ico -fs 5
2021/12/23 21:25:24 [SUCC] Fetch Data From FoFa: [5/13284]
47.98.104.77:8088
154.39.217.22:80
xueshu.mrsb.tk:80
154.39.217.2:80
154.39.217.28:80
```

### Calculate the certificate and query

```console
fofax -fs 5 -uc https://www.baidu.com/

2021/12/23 21:29:54 [SUCC] Fetch Data From FoFa: [5/361619]
180.97.93.146:443
180.97.93.65:443
112.3.25.49:443
itv.leiqiang8.cn:80
owa2.leiqiang8.cn:80
```

## 0x04 fx syntax query

When using fofa to do information collection or other things, it is possible that the query statement will be very, very long and not easy to remember, so we can't just take a small book to remember the fofa query statement.

This time, you can use fofax's fx function. Currently fofax has dozens of built-in fx syntax query rules, users can directly use the corresponding parameters to query. You can also write your own specific fx syntax query rules via yaml format configuration file.

#### shows the built-in fx query statements

```console
fofax -l

┌───────────────┬────────────────────┬────────────────────────────────┬────────┬─────────────┬───────┐
│ Id            │ Query              │ RuleName                       │ Author │ Tag         │ Type  │
├───────────────┼────────────────────┼────────────────────────────────┼────────┼─────────────┼───────┤
│ fx-2021-1001  │ google-reverse     │ Google反代服务器                 │ fofa   │ google      │ 内置   │
│ fx-2021-1002  │ python-simplehttp  │ Python SimpleHTTP              │ fofa   │ python      │ 内置   │
│ fx-2021-1003  │ data-leak          │ 社工库                          │ fofa   │ fun         │ 内置   │
│ fx-2021-1004  │ hfs-rce            │ 存在命令执行的HFS服务             │ fofa   │ fun         │ 内置   │
│ fx-2021-1005  │ satellite-ftp      │ 一键日卫星FTP？                  │ fofa   │ fun         │ 内置   │
│ fx-2021-1006  │ mk-mining          │ mk路由器全球挖矿感染              │ fofa   │ fun         │ 内置   │
│ fx-2021-1007  │ ss-manager-login   │ ss-Manager 登录                 │ fofa   │ fun         │ 内置   │
│ fx-2021-1008  │ heating-monitor    │ 供暖监控系统                     │ fofa   │ fun         │ 内置   │
│ fx-2021-1009  │ free-proxy         │ 免费代理池                      │ fofa    │ fun         │ 内置   │
│ fx-2021-1010  │ honeypot           │ 蜜罐                           │ fofa    │ fun         │ 内置   │
│ fx-2021-1011  │ hacked-website     │ 被挂黑的站点                     │ fofa   │ fun         │ 内置   │
│ fx-2021-1012  │ jupyter-unauth     │ Jupyter 未授权                  │ xiecat │ unauth      │ 内置   │
│ fx-2021-11001 │ APACHE-ActiveMQ    │ APACHE ActiveMQ                │ fofa   │ log4j2,fofa │ 内置   │
│ fx-2021-11002 │ Apache_OFBiz       │ Apache OFBiz                   │ fofa   │ log4j2,fofa │ 内置   │
│ fx-2021-11003 │ Jenkins            │ Jenkins                        │ fofa   │ log4j2,fofa │ 内置   │
│ fx-2021-11004 │ RabbitMQ           │ RabbitMQ                       │ fofa   │ log4j2,fofa │ 内置   │
│ fx-2021-11005 │ Apache-log4j2-Web  │ Apache log4j2 Web              │ fofa   │ log4j2,fofa │ 内置   │
│ fx-2021-11006 │ Jedis              │ Jedis                          │ fofa   │ log4j2,fofa │ 内置   │
│ fx-2021-11007 │ APACHE-tika        │ APACHE tika                    │ fofa   │ log4j2,fofa │ 内置   │
└───────────────┴────────────────────┴────────────────────────────────┴────────┴─────────────┴───────┘
```

### List the details of the fx statement

```console
fofax -ss fx-2021-1001

fx-2021-1001 fx-2021-1001
┌─────────────┬─────────────────────────────────────────────────────────────────────────────────────────────┐
│ Name        │ Value                                                                                       │
├─────────────┼─────────────────────────────────────────────────────────────────────────────────────────────┤
│ ID          │ fx-2021-1001                                                                                │
│ Query       │ google-reverse                                                                              │
│ RuleName    │ Google反代服务器                                                                              │
│ RuleEnglish │ Google Reverse proxy                                                                        │
│ Author      │ fofa                                                                                        │
│ FofaQuery   │ body="var c = Array.prototype.slice.call(arguments, 1);return function() {var d=c.slice();" │
│ Tag         │ google                                                                                      │
│ Type        │ 内置                                                                                         │
│ Description │ 不用挂代理就可以访问的Google搜索，但搜索记录可能会被记录。                                           │
│ FileDir     │                                                                                             │
└─────────────┴─────────────────────────────────────────────────────────────────────────────────────────────┘
```

### Add the `-fe` argument to query by fx syntax

```console
[~] fofax -q 'fx="google-reverse"' -fe -fs 5
2021/12/23 22:27:02 [SUCC] fx query id:google-reverse
2021/12/23 22:27:03 [SUCC] Fetch Data From FoFa: [5/5834]
54.76.26.205:10000
47.74.3.55:80
47.90.7.161:443
23.83.249.79:443
45.76.10.197:8081
```

### Open directly in the browser

```console
fofax -q 'fx="google-reverse"' -fe -open

2021/12/23 22:22:21 [SUCC] fx query id:google-reverse
2021/12/23 22:22:21 [SUCC] the query body="var c = Array.prototype.slice.call(arguments, 1);return function() {var d=c.slice();" will be opened with a browser
```

### Writing custom fx syntax rules

An example of using fofa to collect information about a target is given below.

#### generates a template

A template file is generated by using `-g` and specifying the path to the generated file name.

```console
fofax -g .config/fofax/fxrules/info-gathering.yaml

2021/12/24 20:09:27 [INFO] Will Write Plugin file: .config/fofax/fxrules/info-gathering.yaml
```

Check this yaml file, its contents are as follows.

```console
id: fx-2021-01
query: 查询的字符串用于fx="jupyter Unauth" eg:(jupyter Unauth)
rule_name: 规则名称 eg:(jupyter 未授权)
rule_english: jupyter unauthorized
description: 规则描述
author: 作者<邮箱>eg:(xiecat)
fofa_query: fofa语句 eg:(body="ipython-main-app" && title="Home Page - Select or create a notebook")"
tag:
- 标签1 eg(unauthorized)
- 标签2
source: 语句来源
```

Follow the above instructions and modify the corresponding content to bring in a new fx syntax rule, regarding the path of this file, please put it in the directory `~/.config/fofax/fxrules/` after it is written.

For easy reproduction, the details are as follows (note that the title is filled with your target name)

```yaml
id: fx-2021-01
query: redteam-info-gathering
rule_name: 红队信息收集
rule_english: redteam-info-gathering
description: 使用fofa针对某个目标进行红队常见的高关注CMS/OA系统的信息收集
author: xiecat
fofa_query: title="Target" && (title="平台" || title="OA" || title="系统" || title="协同" || title="办公" || title="致远" || title="泛微" || title="用友" || title="管理" || title="后台" || title="登录" || title="login" || title="admin") && country="CN"
tag:
- redteam
source: 

```

The following can use this fx query rule, this query can not be said to be inconvenient.

```console
fofax -q 'fx="redteam-info-gathering"' -fe -ffi

2021/12/25 21:31:01 [SUCC] fx query id:redteam-info-gathering
2021/12/25 21:31:01 [SUCC] Fetch Data From FoFa: [27/27]
http://60.205.169.36:9080
https://43.243.13.187
http://806f52.ylhskhgyn.com
https://119.28.47.98:8443
http://124.70.197.255:8088
https://223.72.236.165
http://192.144.212.92:8080
https://114.255.204.149
......
```


## 0x05 Linkage Use Case

> Once the red team information is collected, the collected assets can be handed over to the live probing tool, fingerprinting tool and vulnerability scanning tool for live probing, fingerprinting and vulnerability detection.

### fofax && httpx

CVE-2021-43798 Grafana Unauthorized Directory Traversal.

<!-- ![fofax&httpx](./docs/images/fofax&httpx.png) -->
![fofax_httpx](https://s2.loli.net/2021/12/25/kNx281ne7Ou5p4L.png)

### fofax && nuclei

Pass the data obtained by fofax to nuclei and then use the CVE-2021-43798 Template to vulnerability in batch.

<!-- ![fofax&nuclei](./docs/images/fofax&nuclei.png) -->
![fofax_nuclei](https://s2.loli.net/2021/12/25/YztbnOelLZGQAIJ.png)

### fofax && xray

![fofax_xray](https://github.com/xiecat/fofax-doc/blob/dev/docs/.vuepress/public/fofax&xray.png?raw=true)

### fofax && observer_ward

![fofax_observer_ward](https://github.com/xiecat/fofax-doc/blob/dev/docs/.vuepress/public/fofax&observer_ward.png?raw=true)

### fofax && dismap

![fofax_dismap](https://github.com/xiecat/fofax-doc/blob/dev/docs/.vuepress/public/fofax&dismap.png?raw=true)


## 0x06 Stargazers

[![Stargazers over time](https://starchart.cc/xiecat/fofax.svg)](https://starchart.cc/xiecat/fofax)

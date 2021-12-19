# “有趣”的FOFA查询语句收集计划

## FOFA官方收集

> 参考链接：
>
> https://tp.wjx.top/jq/59761097.aspx
>
> http://tp.wjx.top/m/63242496.aspx
>
> https://tp.wjx.top/m/67114073.aspx
>
> 部分失效的规则被删除。

### 第一期

| 规则名                | 规则内容                                             |
| --------------------- | ---------------------------------------------------- |
| 相亲约会？            | body="/tpl/static/varpop/css/box.css"                |
| 具体未知              | (((((((((domain!="6684.com" && (title="专科" \|\| title="学院" \|\| title="大学") && country="CN" && title="查询") && server!="Apache-Coyote/1.1") && server!="Microsoft-IIS/8.5") && os!="centos") && os!="ubuntu") && server!="YxlinkWAF") && port!="8000") && port!="8001") && port!="8088") && server="Apache"        |
| 这个语法可搜索存在命令执行的HFS服务，使用者多数为抓鸡黑客，可以直接上线能捡到不少有趣的东西~ | body="HttpFileServer v2.3 beta 287"                  |
| 日卫星FTP？           | banner="Cobham SATCOM"                               |
| mk路由器全球挖矿感染  | app="Mikrotik-HttpProxy"&&(body="CoinHive.Anonymous" \|\| body="CRLT.Anonymous" \|\| body=" WMP.Anonymous(") |


### 第二期

| 规则名                | 规则内容                                             |
| --------------------- | ---------------------------------------------------- |
| Shadowsocks-Manager登录界面            | body="indeterminate" && body="MainController" && header="X-Powered-By: Express"                |
| 供暖监控系统        | body="s1v13.htm"   |
| 具体未知       | body="admin" && server="tomcat"         |
| 获取免费代理池 | body="get all proxy from proxy pool"                  |
| 具体未知      | body="miner start"                               |
| 蜜罐  | (header="uc-httpd 1.0.0" && server="JBoss-5.0") \|\| server="Apache,Tomcat,Jboss,weblogic,phpstudy,struts" |
| 未授权访问摄像头 | app="webcamXP"  |
| 被挂黑的站点 | body="hacked by" |
| SSR | body="indeterminate" && body="MainController" && header="X-Powered-By: Express" |


### 第三期

| 规则名                | 规则内容                                             |
| --------------------- | ---------------------------------------------------- |
| ss机场框架 | body="SSPanel-Uim" |
| 查看看别人家搭建的蜜罐捕获攻击情况，红色的为攻击者，黄色的为蜜罐部署位置，看看哪个国家的肉鸡多~ | body="img/mhn_logo.png" && body="world-map" |
| 未授权的摄像头 | title="webcam 7"&&body="Live View" |
| 此语句搜索源代码等文件，可以修改rar部分，搜索自己需要的内容，也可根据文件类型进行筛选。 | title="Index of" && body="rar" |
| 该后台为逆苍穹游戏私服运营管理后台，如果存在/action/api_show.php该文件 则密码为value的一句话木马后门。 | body="网站管理员登陆" && port="8090" |
| Cobalt Strike TeamServer | header="HTTP/1.1 404 Not Found Content-Type: text/plain Date:" \|\| protocol="cobaltstrike" \|\| cert="Serial Number: 146473198" |
| vulscan 扫描框架 | title="vulscan" |
| GitHub域名接管 | title="Site not found · GitHub Pages" && server="cloudflare" |
| 私服GM管理后台 一般都有默认密码，数据库弱口令，或者后门，注入 | title="GM管理后台" \|\| title="传奇后台" \|\| body="GM校验码" |
| Canon网络摄像头 | app="Canon-网络摄像头" |
| 这是一个批量上传的webshell，使用密码"r00ts"即可登录，这些目录中还会包含大量的批量上传的木马文件，可供大家进行研究分析。 | body="x.aspx" && body="asp" |
| 批量上传的一句话后门，某些后门密码比较简单，并且密码都一样。这些服务器都有一定的漏洞，比如IIS上传等，并且以企业测试服务器居多，可能包含企业内部数据，并且可以成为跳板进入企业内部网络。 | body="asp;.jpg" |
| 各种机场登录注册 | body="UA-111801619-3" |
| 未知 | body="选择区服" && body="充值" && body="后台" |
| xray扫描结果 | title="X-Ray Report" \|\| body="Powered by xray" |
| 各种指挥系统的登录后台 | title="指挥" && title="登录" |
| html中显示用户名密码 | body="admin" && body="123456" && title="登录" |
| 搜索各种密码文件 | (body="password.txt" \|\| body="密码.txt") && title="index of" |
| 机场科学上网 | `body='<a href= "/staff">STAFF</a >' && body='<a href="/tos">'` |
| 社工库 | title="社工库" \|\| ((title="社工库" && title="系统") \|\| (title="社工库查询" )) |
| 搜索邮件配置文件 | body="intitle:"index of" squirrelmail/" |

### 第四期

| 规则名                | 规则内容                                             |
| --------------------- | ---------------------------------------------------- |
| 另类查找九安监控  | port="600001" && body="login_chk_usr_pwd"  |
| 路由设备登陆界面 | title="欢迎使用RippleOS" |
| 该语法可直接访问到任意使用python3.6.3、3.6.8、3.7.X搭建的简易服务器，然后可以直接访问甚至下载其服务器所有展示内容（毕设Demo、内部测试Demo、源码测试、电影图片XX、相册...） | (server="SimpleHTTP/0.6 Python/3.6.3" \|\| server="SimpleHTTP/0.6 Python/3.6.8" \|\| server="SimpleHTTP/0.6 Python/3.7.0" \|\| server="SimpleHTTP/0.6 Python/3.7.1" \|\| server="SimpleHTTP/0.6 Python/3.7.2" \|\| server="SimpleHTTP/0.6 Python/3.7.3" \|\| server="SimpleHTTP/0.6 Python/3.7.4" \|\| server="SimpleHTTP/0.6 Python/3.7.5" \|\| server="SimpleHTTP/0.6 Python/3.7.6") && title="Directory listing for"   |
| 用户量最大的xss系统 | body="tmp_downloadhelper_iframe" && body="mr_15"  |
| 直播测试系统 | body="直播测试" |
| 关于情报中心的网站内容 | title="情报中心" |
| weblogic | port="7001" && app="Weblogic_interface_7001" |
| 未授权burp  | titile="Burp Suite" && body="Proxy History" |
| Google反代 | body="var c = Array.prototype.slice.call(arguments, 1);return function() {var d=c.slice();"  |

### 其他

……

| 作者 | 规则名 | 规则内容                      | 简述 | Tag        |
| ---- | ------ | ----------------------------- | ---- | ---------- |
| tardc | CMS/OA信息收集 | title="xxxx" && (title="平台" \|\| title="OA" \|\| title="系统" \|\| title="协同" \|\| title="办公" \|\| title="致远" \|\| title="泛微" \|\| title="用友" \|\| title="管理" \|\| title="后台" \|\| title="登录" \|\| title="login" \|\| title="admin") && country="CN"|针对某个目标，对其所有的CMS/OA系统进行收集|红队、|
|      |  ||||




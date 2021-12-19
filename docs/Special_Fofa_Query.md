# “有趣”的FOFA查询语句收集计划

## FOFA官方收集

> 参考链接：
>
> https://tp.wjx.top/jq/59761097.aspx
>
> http://tp.wjx.top/m/63242496.aspx
>
> https://tp.wjx.top/m/67114073.aspx

### 第一期

| 规则名                | 规则内容                                             |
| --------------------- | ---------------------------------------------------- |
| 相亲约会？            | body="/tpl/static/varpop/css/box.css"                |
| 具体未知              | `(((((((((domain!="6684.com" && (title="专科" || title="学院" || title="大学") && country="CN" && title="查询") && server!="Apache-Coyote/1.1") && server!="Microsoft-IIS/8.5") && os!="centos") && os!="ubuntu") && server!="YxlinkWAF") && port!="8000") && port!="8001") && port!="8088") && server=="Apache"`         |
| 存在命令执行的HFS服务 | body="HttpFileServer v2.3 beta 287"                  |
| 日卫星FTP？           | banner="Cobham SATCOM"                               |
| mk路由器全球挖矿感染  | `app="Mikrotik-HttpProxy"&&(body="CoinHive.Anonymous" || body="CRLT.Anonymous" || body=" WMP.Anonymous(")` |


会员 R4v3n 提交的搜索语句 `body="HttpFileServer v2.3 beta 287 随波汉化版"`

这个语法可搜索存在命令执行的HFS服务，使用者多数为抓鸡黑客，可以直接上线能捡到不少有趣的东西~


会员 Snow_zhuimeng 提交的搜索语句 `app="Mikrotik-HttpProxy"&&(body="CoinHive.Anonymous"||body="CRLT.Anonymous"||body=" WMP.Anonymous(")`

这个语法可以查看mk路由器全球挖矿感染的数量

### 第二期

| 规则名                | 规则内容                                             |
| --------------------- | ---------------------------------------------------- |
| Shadowsocks-Manager登录界面            | body="indeterminate" && body="MainController" && header="X-Powered-By: Express"                |
| 供暖监控系统        | body="s1v13.htm"   |
| 具体未知       | body="admin" && server="tomcat"         |
| 获取免费代理池 | body="get all proxy from proxy pool"                  |
| 具体未知      | body="miner start"                               |
| 蜜罐  | `(header="uc-httpd 1.0.0" && server="JBoss-5.0") || server="Apache,Tomcat,Jboss,weblogic,phpstudy,struts"` |
| 未授权访问摄像头 | app="webcamXP"  |
| 被挂黑的站点 | body="hacked by" |
| SSR | body="indeterminate" &&body="MainController" && header="X-Powered-By: Express" |


会员 kankai 提交的搜索语句 `server="Apache,Tomcat,Jboss,weblogic,phpstudy,struts" || header="uc-httpd 1.0.0" && server="JBoss-5.0"`

这个语法可以搜索一些蜜罐，如此多的server。。。

![微信截图_20200320204838.png](https://nosec.org/avatar/uploads/attach/image/7ff51f29bbf078793d11b1fdc73fd6bc/%E5%BE%AE%E4%BF%A1%E6%88%AA%E5%9B%BE_20200320204838.png)




### 第三期

| 规则名                | 规则内容                                             |
| --------------------- | ---------------------------------------------------- |
| ss机场框架 | body="SSPanel-Uim" |
| 查看看别人家搭建的蜜罐捕获攻击情况，红色的为攻击者，黄色的为蜜罐部署位置，看看哪个国家的肉鸡多~ | body="img/mhn_logo.png" && body="world-map" |
| 未授权的摄像头 | title="webcam 7"&&body="Live View" |
| 此语句搜索源代码等文件，可以修改rar部分，搜索自己需要的内容，也可根据文件类型进行筛选。 | title="Index of" && body="rar" |
| 该后台为逆苍穹游戏私服运营管理后台，如果存在/action/api_show.php该文件 则密码为value的一句话木马后门。 | body="网站管理员登陆" && port="8090" |
| Cobalt Strike TeamServer | `header="HTTP/1.1 404 Not Found Content-Type: text/plain Date:" || protocol="cobaltstrike" || cert="Serial Number: 146473198"` |
| vulscan 扫描框架 | title="vulscan" |
| GitHub域名接管 | title="Site not found · GitHub Pages" && server=="cloudflare" |
| 私服GM管理后台 一般都有默认密码，数据库弱口令，或者后门，注入 | `title="GM管理后台" || title="传奇后台" || body="GM校验码"` |
| Canon网络摄像头 | app="Canon-网络摄像头" |
| 这是一个批量上传的webshell，使用密码"r00ts"即可登录，这些目录中还会包含大量的批量上传的木马文件，可供大家进行研究分析。 | body="x.aspx" && body="asp" |
| 批量上传的一句话后门，某些后门密码比较简单，并且密码都一样。这些服务器都有一定的漏洞，比如IIS上传等，并且以企业测试服务器居多，可能包含企业内部数据，并且可以成为跳板进入企业内部网络。 | body="asp;.jpg" |
| 各种机场登录注册 | body="UA-111801619-3" |
| 未知 | body="选择区服" && body="充值" && body="后台" |
| xray扫描结果 | `title="X-Ray Report" || body="Powered by xray"` |
| 各种指挥系统的登录后台 | title="指挥" && title="登录" |
| html中显示用户名密码 | body="admin" && body="123456" && title="登录" |
| 搜索各种密码文件 | `(body="password.txt" || body="密码.txt") && title="index of"` |
| 机场科学上网 | `body='<a href= "/staff">STAFF</a >' && body='<a href="/tos">'` |
| 社工库 | `title=="社工库" || ((title="社工库" && title="系统") || (title=="社工库查询" ))` |
| 搜索邮件配置文件 | body="intitle:"index of" squirrelmail/" |


第三期点赞投票最多的是会员 evalos 提交的查询语法 `header="HTTP/1.1 404 Not Found Content-Type: text/plain Date:"||protocol="cobaltstrike"||cert="Serial Number: 146473198"`

这个语法可以识别cobalt strike服务器，cobalt strike C/S架构的商业渗透软件，适合多人进行团队协作，可模拟APT做模拟对抗，进行内网渗透，是当前比较热门的一款C2软件，无论是红蓝对抗(HW)，还是应急响应中，快速识别C2服务器是重要的工作之一，红队可以提前抹去特征，蓝队可以将此类服务器的IP作为IOC，并加入黑名单大礼包。

这里主要介绍以下三个特征： 第一个利用的是Cobalt Strike 3.13之前版本默认的空格后门特征，第二个是fofa自带的协议识别（推测是teamserver默认证书关键字），第三个是利用cobalt strike的web的默认证书特征 最近爆出的另外一个特征，如何确认web是否为Cobalt Strike，可以请求/manager/text/list/这个路径，如果可以下载到x86的payload，就是Cobalt Strike的服务器。

最二多的是会员 1au 提交的查询语法 `title=="社工库" || ((title="社工库" && title="系统") ||(title=="社工库查询" ))`

这个语句可以搜索社工库，社工库是黑客与大数据方式进行结合的一种产物，黑客们将泄漏的用户数据整合分析，然后集中归档的一个地方，你懂得~

![图片.png](https://nosec.org/avatar/uploads/attach/image/fcf69274afebd92922c8b8b701a30484/%E5%9B%BE%E7%89%87.png)

![图片.png](https://nosec.org/avatar/uploads/attach/image/8cf89e1301c066f3af16869f9d058522/%E5%9B%BE%E7%89%87.png)

最三是会员 Kiven提交的 `body="asp;.jpg"`

这个语句可以搜索带有"asp;.jpg"字样的页面，一般来说，这些文件都是通过批量上传的一句话后门，某些后门密码比较简单，并且密码都一样。这些服务器都有一定的漏洞，比如IIS上传等，并且以企业测试服务器居多，可能包含企业内部数据，并且可以成为跳板进入企业内部网络。

![图片.png](https://nosec.org/avatar/uploads/attach/image/c3884a3e45a12b15632b40349105deab/%E5%9B%BE%E7%89%87.png)![图片.png](https://nosec.org/avatar/uploads/attach/image/3651a7862a75d023d6a18e40b3b4ccec/%E5%9B%BE%E7%89%87.png)

最四是会员 mark0smith 提交的查询语法 `title="指挥" && title="登录"`

这个语法查询的是各种各样的指挥系统，这也是公共资产，大家看看就好~

![图片.png](https://nosec.org/avatar/uploads/attach/image/14269042d186a6f3482f54a791284c9e/%E5%9B%BE%E7%89%87.png)

第五是会员 angel 提交的查询语法 `title="GM管理后台" || title="传奇后台" || body="GM校验码"`

这个语法可以查到很多游戏管理后台，GM管理后台是一个游戏管理后台，包括权限管理，支付系统，兑换系统，邮件发送，信息推送，游戏数据统计等模块，喜欢玩游戏的朋友们~（滑稽）

![图片.png](https://nosec.org/avatar/uploads/attach/image/aa91c19b5aef180359dd0c5dfbb78159/%E5%9B%BE%E7%89%87.png)

### 第四期

| 规则名                | 规则内容                                             |
| --------------------- | ---------------------------------------------------- |
| 另类查找九安监控  | port="600001" && body="login_chk_usr_pwd"  |
| 路由设备登陆界面 | title="欢迎使用RippleOS" |
| 该语法可直接访问到任意使用python3.6.3、3.6.8、3.7.X搭建的简易服务器，然后可以直接访问甚至下载其服务器所有展示内容（毕设Demo、内部测试Demo、源码测试、电影图片XX、相册...） | `(server="SimpleHTTP/0.6 Python/3.6.3" || server="SimpleHTTP/0.6 Python/3.6.8" || server="SimpleHTTP/0.6 Python/3.7.0" || server="SimpleHTTP/0.6 Python/3.7.1" || server="SimpleHTTP/0.6 Python/3.7.2" || server="SimpleHTTP/0.6 Python/3.7.3" || server="SimpleHTTP/0.6 Python/3.7.4" || server="SimpleHTTP/0.6 Python/3.7.5" || server="SimpleHTTP/0.6 Python/3.7.6") && title="Directory listing for"`   |
| 用户量最大的xss系统 | body="tmp_downloadhelper_iframe" && body="mr_15"  |
| 直播测试系统 | body="直播测试" |
| 关于情报中心的网站内容 | title="情报中心" |
| weblogic | port="7001" && app="Weblogic_interface_7001" |
| 未授权burp  | titile="Burp Suite" && body="Proxy History" |
| Google反代 | body="var c = Array.prototype.slice.call(arguments, 1);return function() {var d=c.slice();"  |

### 其他

……
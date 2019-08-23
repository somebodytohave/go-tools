# go-tools

[![Release](https://img.shields.io/github/release/sun-wenming/go-tools.svg?style=plastic)](https://github.com/sun-wenming/go-tools/releases)
[![GitHub stars](https://img.shields.io/github/stars/sun-wenming/go-tools?style=plastic)](https://github.com/sun-wenming/go-tools/stargazers)
[![GitHub issues](https://img.shields.io/github/issues/sun-wenming/go-tools.svg?style=plastic)](https://github.com/sun-wenming/go-tools/issues)
[![GitHub license](https://img.shields.io/github/license/sun-wenming/go-tools.svg?style=plastic)](https://github.com/sun-wenming/go-tools/blob/master/LICENSE)


> 此项目封装了日常开发用到的工具.如果对您有些帮助，为了使项目更完善<br>
[![GitHub issues](https://img.shields.io/github/issues/sun-wenming/go-tools.svg?style=plastic)](https://github.com/sun-wenming/go-tools/issues) 好的建议或修复<br>
[![GitHub stars](https://img.shields.io/github/stars/sun-wenming/go-tools.svg?style=plastic)](https://github.com/sun-wenming/go-tools/stargazers)支持下.
谢谢咯.

## 目录
- [安装](#安装)
- 功能
    - [mfile 文件](https://github.com/sun-wenming/go-tools/tree/master/mfile) 关于文件是否存在、权限、新建...验证图片后缀、验证图片大小
    - [mgin gin返回封装](https://github.com/sun-wenming/go-tools/tree/master/mgin) 对于Gin框架返回数据,分页,的封装,使开发更加便捷.
    - [mgoogleauth 谷歌二次验证](https://github.com/sun-wenming/go-tools/tree/master/mgoogleauth) 谷歌的二次验证,类似于将军令.
    - [mjwt 生成](https://github.com/sun-wenming/go-tools/tree/master/mjwt) 使用密钥对生成防伪Token以及加密解密Token的内容.
    - [mencrypt aes 加密解密](https://github.com/sun-wenming/go-tools/tree/master/) aes 加密解密.
    - [mlog 日志文件封装](https://github.com/sun-wenming/go-tools/tree/master/mlog) 日志文件封装.
    - [mmiddleware gin中间件](https://github.com/sun-wenming/go-tools/tree/master/mmiddleware) 跨域中间件.
    - [mqrcode 二维码](https://github.com/sun-wenming/go-tools/tree/master/mqrcode) 二维码生成.
    - [mrandom 随机字符串或code](https://github.com/sun-wenming/go-tools/tree/master/mrandom) .
    - [mredis redis 缓存的常见操作](https://github.com/sun-wenming/go-tools/tree/master/mredis) .
    - [mstr 字符串处理](https://github.com/sun-wenming/go-tools/tree/master/mstr) .
    - [mtime 时间处理](https://github.com/sun-wenming/go-tools/tree/master/mtime) 时间戳生成...
    - [mvalid 参数与正则验证](https://github.com/sun-wenming/go-tools/tree/master/mvalid) ...
- [TODO](#TODO)
- [三方开源工具](#三方开源工具)

## 安装

@release [![Release](https://img.shields.io/github/release/sun-wenming/go-tools.svg?style=plastic)](https://github.com/sun-wenming/go-tools/releases) 
```
go get github.com/sun-wenming/go-tools@上方版本号
```

## TODO
- 增加功能Demo
- 新功能建议 [![GitHub issues](https://img.shields.io/github/issues/sun-wenming/go-tools.svg?style=plastic)](https://github.com/sun-wenming/go-tools/issues)


## 三方开源工具
- 类型转换 [cast](https://github.com/spf13/cast)
- 时间工具 [now](https://github.com/jinzhu/now) 
- 定时任务 [cron](https://github.com/robfig/cron)

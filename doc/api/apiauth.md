# 授权接口

## CheckPermissions

检查权限

### 原型
```
CheckPermissions(permissions []string) bool
```

### 描述

检查是否具有给定的权限列表中的全部权限

* permissions 权限列表

### 范例代码

Javascript
```
world.Note(world.CheckPermissions(["http"]))
```

Lua
```
Note(CheckPermissions({"http"}))
```
### 返回值

布尔值

## RequestPermissions

请求权限

### 原型

```
RequestPermissions(permissions []string, reason string, script string)
```

### 描述

请求用户对指定的权限进行授权

* permissions 需要授权的权限列表
* reason 提示用户的理由
* script 回调脚本，留空不回调

### 范例代码

Javascript
```
world.RequestPermissions(["http"],"调用接口","httpauthed")
```

Lua
```
RequestPermissions({"http"},"调用接口","httpauthed")
```

### 返回值

无

## CheckTrustedDomains

检查信任域名

### 原型

```
CheckTrustedDomains(domains []string) bool
```

### 描述

检查是否信任了给定的域名列表中的全部域名

* domains 域名列表

### 范例代码

Javascript
```
world.Note(world.CheckTrustedDomains(["127.0.0.1"]))
```

Lua
```
Note(CheckTrustedDomains({"127.0.0.1"}))
```
### 返回值

布尔值

## RequestTrustDomains

请求信任域名

### 原型
```
RequestTrustDomains(domains []string, reason string, script string)
```

### 描述

请求用户对指定的域名进行信任

* domains 需要信任的域名列表
* reason 提示用户的理由
* script 回调脚本，留空不回调

### 范例代码

Javascript
```
world.RequestTrustDomains(["127.0.0.1"],"调用接口","localhostauthed")
```

Lua
```
RequestTrustDomains({"127.0.0.1"},"调用接口","localhostauthed")
```

### 返回值

无
### nginx-ldap-auth

基于 ngx_http_auth_request_module 的 Nginx 反向代理 LDAP 认证。
Golang 实现

### usage
#### Build from source
```
# git clone https://github.com/shanghai-edu/nginx-ldap-auth
# go mod tidy
# chmod +x control
# ./control build
# ./control pack
```
这里会 pack 出一个 tar 包，拿去部署即可
也可以直接从 [release](https://github.com/shanghai-edu/nginx-ldap-auth/releases) 下载编译好的版本
#### Install
```
# tar -zxvf nginx-ldap-auth-0.1.tar.gz 
# cp nginx.conf /etc/nginx/nginx.conf
# mv cfg.example.json cfg.json // and change it 
# service nginx reload
# ./control start
```

根据自己的实际配置情况修改 cfg.json
```
{
     // ldap config
    "ldap": {
        "addr": "ldap.example.org:389",
        "baseDn": "dc=example,dc=org",
        "bindDn": "cn=manager,dc=example,dc=org",
        "bindPass": "password",
        "authFilter": "(&(uid=%s))",
        "attributes": ["uid", "cn", "mail"],
        "tls":        false,
        "startTLS":   false
    },
    // control config
    "control":{
        "ipAcl":{
            "deny":["127.0.0.1","192.168.0.0/24","192.168.1.0-192.168.1.255"], // 这些 IP 会被直接干掉
            "direct":[] // 直通的 IP 不需要认证
        },
        "timeAcl":{
            "deny":["00:00-08:00","17:00-23:59"], // 这些时间范围的访问会被 deny
            "direct":[]  // 这些时间访问会直通
        },
        "allowUser":["user1"]  // 允许认证通过的用户名，如果允许所有的 LDAP 用户通过认证，那么这里留空。
    },
    "http": {
        "debug":false,
        "ips":["127.0.0.1"], // 调用 API 所信任的 IP 范围
        "listen": "0.0.0.0:8080"
    }
}
```
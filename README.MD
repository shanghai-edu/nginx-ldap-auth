### nginx-ldap-auth

A simple example of LDAP authentication using ngx_http_auth_request_module by golang.

[中文](https://github.com/shanghai-edu/nginx-ldap-auth/blob/master/README_CN.MD)

### usage
#### Build from source
```
# git clone https://github.com/shanghai-edu/nginx-ldap-auth
# go get ./...
# chmod +x control
# ./control build
# ./control pack
```

#### Install
```
# tar -zxvf nginx-ldap-auth-0.1.tar.gz 
# cp nginx.conf /etc/nginx/nginx.conf
# mv cfg.example.json cfg.json // and change it 
# service nginx reload
# ./control start
```
cfg.json example
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
            "deny":["127.0.0.1","192.168.0.0/24","192.168.1.0-192.168.1.255"], // deny ips
            "direct":[] // direct ips
        },
        "timeAcl":{
            "deny":["00:00-08:00","17:00-23:59"], // time deny
            "direct":[]  // time direct
        },
        "allowUser":["user1"]  // user allowed to login, should be [] if all ldap users are allowed
    },
    "http": {
        "debug":false,
        "ips":["127.0.0.1"], // api trust ips
        "listen": "0.0.0.0:8080"
    }
}
```

# nginx

```bash
[root@selfhosted-0001 petal]# ./petal-linux-amd64 -file task/mw/nginx/install.yml 
=== Task: Add NGINX Yum Repository ===
[Add NGINX Yum Repository][node1] Running...
cat <<EOF >/etc/yum.repos.d/nginx.repo
[nginx-stable]
name=nginx stable repo
baseurl=http://nginx.org/packages/centos/$releasever/$basearch/
gpgcheck=1
enabled=1
gpgkey=https://nginx.org/keys/nginx_signing.key
module_hotfixes=true
EOF

[Add NGINX Yum Repository][node1] Success

=== Completed: Add NGINX Yum Repository ===

=== Task: Install NGINX 1.20.2 ===
[Install NGINX 1.20.2][node1] Running...
yum clean all
yum makecache
yum install -y nginx-1.20.2

[Install NGINX 1.20.2][node1] Success
Loaded plugins: fastestmirror
Cleaning repos: base epel extras nginx-stable updates
Cleaning up list of fastest mirrors
Other repos take up 71 M of disk space (use --enablerepo='*' to clean all repos)
Loaded plugins: fastestmirror
Determining fastest mirrors
 * base: mirrors.huaweicloud.com
 * extras: mirrors.huaweicloud.com
 * updates: mirrors.huaweicloud.com
Metadata Cache Created
Loaded plugins: fastestmirror
Loading mirror speeds from cached hostfile
 * base: mirrors.huaweicloud.com
 * extras: mirrors.huaweicloud.com
 * updates: mirrors.huaweicloud.com
Resolving Dependencies
--> Running transaction check
---> Package nginx.x86_64 1:1.20.2-1.el7.ngx will be installed
--> Finished Dependency Resolution

Dependencies Resolved

================================================================================
 Package         Arch            Version                   Repository      Size
================================================================================
Installing:
 nginx           x86_64          1:1.20.2-1.el7.ngx        nginx-stable   790 k

Transaction Summary
================================================================================
Install  1 Package

Total download size: 790 k
Installed size: 2.7 M
Downloading packages:
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  Installing : 1:nginx-1.20.2-1.el7.ngx.x86_64                            1/1
  Verifying  : 1:nginx-1.20.2-1.el7.ngx.x86_64                            1/1

Installed:
  nginx.x86_64 1:1.20.2-1.el7.ngx

Complete!
=== Completed: Install NGINX 1.20.2 ===

=== Task: Check NGINX Version After Installation ===
[Check NGINX Version After Installation][node1] Running...
nginx -v
rpm -qa | grep nginx

[Check NGINX Version After Installation][node1] Success
nginx version: nginx/1.20.2
nginx-1.20.2-1.el7.ngx.x86_64
=== Completed: Check NGINX Version After Installation ===
```

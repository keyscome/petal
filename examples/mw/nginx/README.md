# nginx

## Overview

NGINX 1.20.2 is a high-performance HTTP server and reverse proxy commonly used to serve static content, terminate SSL/TLS, and proxy requests to upstream application servers. This example installs the stable release from the official NGINX yum repository on a single node.

## Architecture

- **Topology**: Single node (`node1`)
- **Ports**: `80` (HTTP), `443` (HTTPS, after manual SSL configuration)
- **Package source**: `http://nginx.org/packages/centos/` (stable channel, GPG-verified)
- **Config file**: `/etc/nginx/nginx.conf` (default post-install location)
- **Repo file**: `/etc/yum.repos.d/nginx.repo`
- **Binary**: `/usr/sbin/nginx`

## Files

| File | Description |
|------|-------------|
| `install.yml` | Adds the official NGINX yum repository, installs NGINX 1.20.2, and verifies the installation |
| `uninstall.yml` | Removes NGINX and backs up the yum repository file |

### `install.yml`

1. **Add NGINX Yum Repository** — Writes a `[nginx-stable]` repo definition to `/etc/yum.repos.d/nginx.repo` using a heredoc. The repo points to the official NGINX package CDN, enables GPG signature verification, and activates `module_hotfixes` for CentOS compatibility.
2. **Install NGINX 1.20.2** — Clears the yum cache, rebuilds the metadata cache from the new repo, then installs the pinned version `nginx-1.20.2`.
3. **Check NGINX Version After Installation** — Runs `nginx -v` to confirm the installed version and `rpm -qa | grep nginx` to list all installed NGINX packages.

### `uninstall.yml`

1. **Ensure TMP_NGINX_DIR Exists** — Creates the backup directory (`$TMP_NGINX_DIR`) if it doesn't already exist.
2. **Check NGINX Version Before Uninstallation** — Prints the NGINX version (if present) and lists installed NGINX RPM packages as a pre-removal audit.
3. **Remove NGINX** — Runs `yum remove -y nginx` to uninstall the package.
4. **Move NGINX Repo File to TMP_NGINX_DIR** — If `/etc/yum.repos.d/nginx.repo` exists, moves it to the backup directory so it can be restored later; skips gracefully if not found.
5. **Verify NGINX Removal** — Checks that the `nginx` binary is gone and no NGINX RPMs remain, reporting `[OK]` or `[WARN]` accordingly.

## Example

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

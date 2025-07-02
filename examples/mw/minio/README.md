# minio

```bash
[root@selfhosted-0001 petal]# ./petal-linux-amd64 -file task/mw/minio/install.yml 
=== Task: Download Minio Package ===
[Download Minio Package][node3] Running...
obsutil cp -r -f obs://selfhosted/pkg/minio/ $TMP_DIR

[Download Minio Package][node1] Running...
obsutil cp -r -f obs://selfhosted/pkg/minio/ $TMP_DIR

[Download Minio Package][node2] Running...
obsutil cp -r -f obs://selfhosted/pkg/minio/ $TMP_DIR

[Download Minio Package][node2] Success
Start at 2025-06-30 09:16:47.771531725 +0000 UTC


Parallel:      5                   Jobs:          5                   
Threshold:     50.00MB             PartSize:      auto                
VerifyLength:  false               VerifyMd5:     false               
CheckpointDir: /root/.obsutil_checkpoint     

Task id: c1c36b8f-2cc0-43c3-a9bd-19e4700418f6
OutputDir: /root/.obsutil_output         
TempFileDir: /root                         

Succeed count:      2         Failed count:       0         .59 880.95MB/s 2/2 15.86MB/15.86MB 220ms
Succeed bytes:      15.86MB   
Metrics [max cost:211 ms, min cost:211 ms, average cost:105.50 ms, average tps:4.52, transferred size:15.86MB]

Task id: c1c36b8f-2cc0-43c3-a9bd-19e4700418f6

[Download Minio Package][node1] Success
Start at 2025-06-30 09:16:47.766009806 +0000 UTC


Parallel:      5                   Jobs:          5                   
Threshold:     50.00MB             PartSize:      auto                
VerifyLength:  false               VerifyMd5:     false               
CheckpointDir: /root/.obsutil_checkpoint     

Task id: b8b92adc-9315-4f7b-ac8a-44afb79eef99
OutputDir: /root/.obsutil_output         
TempFileDir: /root                         

Succeed count:      2         Failed count:       0         4.42 70.16MB/s 2/2 15.86MB/15.86MB 427ms
Succeed bytes:      15.86MB   
Metrics [max cost:414 ms, min cost:414 ms, average cost:207.00 ms, average tps:2.34, transferred size:15.86MB]

Task id: b8b92adc-9315-4f7b-ac8a-44afb79eef99

[Download Minio Package][node3] Success
Start at 2025-06-30 09:16:47.767018564 +0000 UTC


Parallel:      5                   Jobs:          5                   
Threshold:     50.00MB             PartSize:      auto                
VerifyLength:  false               VerifyMd5:     false               
CheckpointDir: /root/.obsutil_checkpoint     

Task id: c7dcbd16-8d84-4d03-87b0-5b70ed646785
OutputDir: /root/.obsutil_output         
TempFileDir: /root                         

Succeed count:      2         Failed count:       0         1.88% tps:0.00 ?/s 1/2 12.98MB/15.86MB ?
Succeed bytes:      15.86MB   
Metrics [max cost:437 ms, min cost:437 ms, average cost:218.50 ms, average tps:2.23, transferred size:15.86MB]

Task id: c7dcbd16-8d84-4d03-87b0-5b70ed646785

=== Completed: Download Minio Package ===

=== Task: Install Minio ===
[Install Minio][node3] Running...
yum install -y $TMP_DIR_MINIO/minio-20210326000041.0.0.x86_64.rpm
rpm -qa | grep minio

[Install Minio][node1] Running...
[Install Minio][node2] Running...
yum install -y $TMP_DIR_MINIO/minio-20210326000041.0.0.x86_64.rpm
rpm -qa | grep minio

yum install -y $TMP_DIR_MINIO/minio-20210326000041.0.0.x86_64.rpm
rpm -qa | grep minio

[Install Minio][node3] Success
Loaded plugins: fastestmirror
Examining /tmp/selfhosted/minio/minio-20210326000041.0.0.x86_64.rpm: minio-20210326000041.0.0-1.x86_64
/tmp/selfhosted/minio/minio-20210326000041.0.0.x86_64.rpm: does not update installed package.
Error: Nothing to do
minio-20210326000041.0.0-1.x86_64

[Install Minio][node2] Success
Loaded plugins: fastestmirror
Examining /tmp/selfhosted/minio/minio-20210326000041.0.0.x86_64.rpm: minio-20210326000041.0.0-1.x86_64
/tmp/selfhosted/minio/minio-20210326000041.0.0.x86_64.rpm: does not update installed package.
Error: Nothing to do
minio-20210326000041.0.0-1.x86_64

[Install Minio][node1] Success
Loaded plugins: fastestmirror
Examining /tmp/selfhosted/minio/minio-20210326000041.0.0.x86_64.rpm: minio-20210326000041.0.0-1.x86_64
/tmp/selfhosted/minio/minio-20210326000041.0.0.x86_64.rpm: does not update installed package.
Error: Nothing to do
minio-20210326000041.0.0-1.x86_64

=== Completed: Install Minio ===

=== Task: Write Minio Start Script ===
[Write Minio Start Script][node3] Running...
[ -d /data/minio ] || mkdir -p /data/minio
cat <<EOF > /usr/local/bin/start-minio.sh
#!/bin/bash
# /usr/local/bin/start-minio.sh

# 设置 MinIO 用户名和密码
export MINIO_ROOT_USER=minioadmin
export MINIO_ROOT_PASSWORD=minioadmin#Crm110

[ -d /data/minio ] || mkdir -p /data/minio
[ -d /data/minio/data1 ] || mkdir -p /data/minio/data1
[ -d /data/minio/data2 ] || mkdir -p /data/minio/data2

# 启动 MinIO 服务并将输出追加到日志文件
/usr/local/bin/minio server --address :9000 \
  http://${NODE1_IP}/data/minio/data1 http://${NODE1_IP}/data/minio/data2 \
  http://${NODE2_IP}/data/minio/data1 http://${NODE2_IP}/data/minio/data2 \
  http://${NODE3_IP}/data/minio/data1 http://${NODE3_IP}/data/minio/data2 \
  2>&1 | tee -a /data/minio/minio.log
EOF
chmod +x /usr/local/bin/start-minio.sh

[Write Minio Start Script][node1] Running...
[ -d /data/minio ] || mkdir -p /data/minio
cat <<EOF > /usr/local/bin/start-minio.sh
#!/bin/bash
# /usr/local/bin/start-minio.sh

# 设置 MinIO 用户名和密码
export MINIO_ROOT_USER=minioadmin
export MINIO_ROOT_PASSWORD=minioadmin#Crm110

[ -d /data/minio ] || mkdir -p /data/minio
[ -d /data/minio/data1 ] || mkdir -p /data/minio/data1
[ -d /data/minio/data2 ] || mkdir -p /data/minio/data2

# 启动 MinIO 服务并将输出追加到日志文件
/usr/local/bin/minio server --address :9000 \
  http://${NODE1_IP}/data/minio/data1 http://${NODE1_IP}/data/minio/data2 \
  http://${NODE2_IP}/data/minio/data1 http://${NODE2_IP}/data/minio/data2 \
  http://${NODE3_IP}/data/minio/data1 http://${NODE3_IP}/data/minio/data2 \
  2>&1 | tee -a /data/minio/minio.log
EOF
chmod +x /usr/local/bin/start-minio.sh

[Write Minio Start Script][node2] Running...
[ -d /data/minio ] || mkdir -p /data/minio
cat <<EOF > /usr/local/bin/start-minio.sh
#!/bin/bash
# /usr/local/bin/start-minio.sh

# 设置 MinIO 用户名和密码
export MINIO_ROOT_USER=minioadmin
export MINIO_ROOT_PASSWORD=minioadmin#Crm110

[ -d /data/minio ] || mkdir -p /data/minio
[ -d /data/minio/data1 ] || mkdir -p /data/minio/data1
[ -d /data/minio/data2 ] || mkdir -p /data/minio/data2

# 启动 MinIO 服务并将输出追加到日志文件
/usr/local/bin/minio server --address :9000 \
  http://${NODE1_IP}/data/minio/data1 http://${NODE1_IP}/data/minio/data2 \
  http://${NODE2_IP}/data/minio/data1 http://${NODE2_IP}/data/minio/data2 \
  http://${NODE3_IP}/data/minio/data1 http://${NODE3_IP}/data/minio/data2 \
  2>&1 | tee -a /data/minio/minio.log
EOF
chmod +x /usr/local/bin/start-minio.sh

[Write Minio Start Script][node3] Success

[Write Minio Start Script][node1] Success

[Write Minio Start Script][node2] Success

=== Completed: Write Minio Start Script ===

=== Task: Minio Systemd Service ===
[Minio Systemd Service][node1] Running...
[Minio Systemd Service][node2] Running...
cat <<EOF > /etc/systemd/system/minio.service
# /etc/systemd/system/minio.service

[Unit]
Description=MinIO distributed object storage server
After=network.target

[Service]
User=root
Group=root
ExecStart=/usr/local/bin/start-minio.sh
WorkingDirectory=/data/minio
Restart=always
LimitNOFILE=65536
Environment="MINIO_BROWSER_REDIRECT_PORT=9000"

[Install]
WantedBy=multi-user.target
EOF
systemctl daemon-reload

cat <<EOF > /etc/systemd/system/minio.service
# /etc/systemd/system/minio.service

[Unit]
Description=MinIO distributed object storage server
After=network.target

[Service]
User=root
Group=root
ExecStart=/usr/local/bin/start-minio.sh
WorkingDirectory=/data/minio
Restart=always
LimitNOFILE=65536
Environment="MINIO_BROWSER_REDIRECT_PORT=9000"

[Install]
WantedBy=multi-user.target
EOF
systemctl daemon-reload

[Minio Systemd Service][node3] Running...
cat <<EOF > /etc/systemd/system/minio.service
# /etc/systemd/system/minio.service

[Unit]
Description=MinIO distributed object storage server
After=network.target

[Service]
User=root
Group=root
ExecStart=/usr/local/bin/start-minio.sh
WorkingDirectory=/data/minio
Restart=always
LimitNOFILE=65536
Environment="MINIO_BROWSER_REDIRECT_PORT=9000"

[Install]
WantedBy=multi-user.target
EOF
systemctl daemon-reload

[Minio Systemd Service][node1] Success

[Minio Systemd Service][node3] Success

[Minio Systemd Service][node2] Success

=== Completed: Minio Systemd Service ===

=== Task: Start Minio Service ===
[Start Minio Service][node3] Running...
systemctl enable minio
systemctl start minio
systemctl status minio

[Start Minio Service][node2] Running...
systemctl enable minio
systemctl start minio
systemctl status minio

[Start Minio Service][node1] Running...
systemctl enable minio
systemctl start minio
systemctl status minio

[Start Minio Service][node1] Success
Created symlink from /etc/systemd/system/multi-user.target.wants/minio.service to /etc/systemd/system/minio.service.
● minio.service - MinIO distributed object storage server
   Loaded: loaded (/etc/systemd/system/minio.service; enabled; vendor preset: disabled)
   Active: active (running) since Mon 2025-06-30 17:16:49 CST; 3ms ago
 Main PID: 5467 (start-minio.sh)
   CGroup: /system.slice/minio.service
           ├─5467 /bin/bash /usr/local/bin/start-minio.sh
           ├─5471 /usr/local/bin/minio server --address :9000 http://10.31.0.93/data/minio/data1 http://10.31.0.93/data/minio/data2 http://10.31.0.246/data/minio/data1 http://10.31.0.246/data/minio/data2 http://10.31.0.68/data/minio/data1 http://10.31.0.68/data/minio/data2
           └─5472 tee -a /data/minio/minio.log

Jun 30 17:16:49 selfhosted-0001 systemd[1]: Started MinIO distributed object storage server.

[Start Minio Service][node3] Success
Created symlink from /etc/systemd/system/multi-user.target.wants/minio.service to /etc/systemd/system/minio.service.
● minio.service - MinIO distributed object storage server
   Loaded: loaded (/etc/systemd/system/minio.service; enabled; vendor preset: disabled)
   Active: active (running) since Mon 2025-06-30 17:16:49 CST; 3ms ago
 Main PID: 24039 (start-minio.sh)
   CGroup: /system.slice/minio.service
           ├─24039 /bin/bash /usr/local/bin/start-minio.sh
           ├─24043 /bin/bash /usr/local/bin/start-minio.sh
           └─24044 tee -a /data/minio/minio.log

Jun 30 17:16:49 selfhosted-0003 systemd[1]: Started MinIO distributed object storage server.

[Start Minio Service][node2] Success
Created symlink from /etc/systemd/system/multi-user.target.wants/minio.service to /etc/systemd/system/minio.service.
● minio.service - MinIO distributed object storage server
   Loaded: loaded (/etc/systemd/system/minio.service; enabled; vendor preset: disabled)
   Active: active (running) since Mon 2025-06-30 17:16:49 CST; 3ms ago
 Main PID: 2924 (start-minio.sh)
   CGroup: /system.slice/minio.service
           ├─2924 /bin/bash /usr/local/bin/start-minio.sh
           └─2927 [mkdir]

Jun 30 17:16:49 selfhosted-0002 systemd[1]: Started MinIO distributed object storage server.

=== Completed: Start Minio Service ===

[root@selfhosted-0001 petal]# 
```

# nacos

```bash
[root@selfhosted-0001 petal]# ./petal-linux-amd64 -file task/mw/nacos/install.yml 
=== Task: Download Nacos Packages ===
[Download Nacos Packages][node3] Running...
obsutil cp -r -f obs://selfhosted/pkg/nacos/ $TMP_DIR
tar zxf $TMP_DIR/nacos/nacos_pg.tgz -C /data
mv /data/nacos_pg /data/nacos

[Download Nacos Packages][node1] Running...
obsutil cp -r -f obs://selfhosted/pkg/nacos/ $TMP_DIR
tar zxf $TMP_DIR/nacos/nacos_pg.tgz -C /data
mv /data/nacos_pg /data/nacos

[Download Nacos Packages][node2] Running...
obsutil cp -r -f obs://selfhosted/pkg/nacos/ $TMP_DIR
tar zxf $TMP_DIR/nacos/nacos_pg.tgz -C /data
mv /data/nacos_pg /data/nacos

[Download Nacos Packages][node2] Success
Start at 2025-07-01 08:56:07.16248275 +0000 UTC


Parallel:      5                   Jobs:          5                   
Threshold:     50.00MB             PartSize:      auto                
VerifyLength:  false               VerifyMd5:     false               
CheckpointDir: /root/.obsutil_checkpoint     

Task id: 65ff3aa6-374f-4b08-87bc-40441ab17808
OutputDir: /root/.obsutil_output         
TempFileDir: /root                         

Succeed count:      10        Failed count:       0         .75 384.81MB/s 8/10 325.74MB/368.49MB 0s
Succeed bytes:      368.49MB  
Metrics [max cost:1072 ms, min cost:6 ms, average cost:341.40 ms, average tps:7.86, transferred size:368.49MB]
[---------------------------------------] 100.00% tps:9.54 390.76MB/s 10/10 368.49MB/368.49MB 1.144s
Task id: 65ff3aa6-374f-4b08-87bc-40441ab17808

[Download Nacos Packages][node1] Success
Start at 2025-07-01 08:56:07.160021784 +0000 UTC


Parallel:      5                   Jobs:          5                   
Threshold:     50.00MB             PartSize:      auto                
VerifyLength:  false               VerifyMd5:     false               
CheckpointDir: /root/.obsutil_checkpoint     

Task id: 5fb994db-675d-4dc9-b643-1b5590c1c8e9
OutputDir: /root/.obsutil_output         
TempFileDir: /root                         

Succeed count:      10        Failed count:       0         .72 438.60MB/s 8/10 333.68MB/368.49MB 0s
Succeed bytes:      368.49MB  
Metrics [max cost:1086 ms, min cost:49 ms, average cost:323.10 ms, average tps:7.85, transferred size:368.49MB]
[---------------------------------------] 100.00% tps:9.53 390.76MB/s 10/10 368.49MB/368.49MB 1.145s
Task id: 5fb994db-675d-4dc9-b643-1b5590c1c8e9

[Download Nacos Packages][node3] Success
Start at 2025-07-01 08:56:07.169027957 +0000 UTC


Parallel:      5                   Jobs:          5                   
Threshold:     50.00MB             PartSize:      auto                
VerifyLength:  false               VerifyMd5:     false               
CheckpointDir: /root/.obsutil_checkpoint     

Task id: 297ff64d-b92d-492e-b5c2-29eda4c51ca4
OutputDir: /root/.obsutil_output         
TempFileDir: /root                         

Succeed count:      10        Failed count:       0         .46 407.49MB/s 7/10 336.15MB/368.49MB 0s
Succeed bytes:      368.49MB  
Metrics [max cost:1122 ms, min cost:5 ms, average cost:348.00 ms, average tps:7.90, transferred size:368.49MB]
[---------------------------------------] 100.00% tps:9.62 394.10MB/s 10/10 368.49MB/368.49MB 1.136s
Task id: 297ff64d-b92d-492e-b5c2-29eda4c51ca4

=== Completed: Download Nacos Packages ===

=== Task: Postgre Nacos Setup ===
[Postgre Nacos Setup][node1] Running...
su - postgres -c "psql -c 'CREATE DATABASE nacos_config;'" || echo "Database exists"
su - postgres -c "psql -c \"CREATE USER nacos WITH PASSWORD 'Heli198213';\"" || echo "User exists"
su - postgres -c "psql -c 'GRANT ALL PRIVILEGES ON DATABASE nacos_config TO nacos;'"

PGPASSWORD=Heli198213 psql -U nacos -h 127.0.0.1 -d nacos_config -f $TMP_DIR_NACOS/nacos-pg.sql

[Postgre Nacos Setup][node1] Success
CREATE DATABASE
CREATE ROLE
GRANT
psql:/tmp/selfhosted/nacos/nacos-pg.sql:20: NOTICE:  table "config_info" does not exist, skipping
DROP TABLE
CREATE TABLE
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
psql:/tmp/selfhosted/nacos/nacos-pg.sql:58: NOTICE:  table "config_info_aggr" does not exist, skipping
DROP TABLE
CREATE TABLE
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
BEGIN
COMMIT
psql:/tmp/selfhosted/nacos/nacos-pg.sql:88: NOTICE:  table "config_info_beta" does not exist, skipping
DROP TABLE
CREATE TABLE
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
BEGIN
COMMIT
psql:/tmp/selfhosted/nacos/nacos-pg.sql:129: NOTICE:  table "config_info_tag" does not exist, skipping
DROP TABLE
CREATE TABLE
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
BEGIN
COMMIT
DROP TABLE
psql:/tmp/selfhosted/nacos/nacos-pg.sql:168: NOTICE:  table "config_tags_relation" does not exist, skipping
CREATE TABLE
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
BEGIN
COMMIT
DROP TABLE
psql:/tmp/selfhosted/nacos/nacos-pg.sql:196: NOTICE:  table "group_capacity" does not exist, skipping
CREATE TABLE
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
BEGIN
COMMIT
psql:/tmp/selfhosted/nacos/nacos-pg.sql:231: NOTICE:  table "his_config_info" does not exist, skipping
DROP TABLE
CREATE TABLE
COMMENT
COMMENT
COMMENT
COMMENT
psql:/tmp/selfhosted/nacos/nacos-pg.sql:258: NOTICE:  table "permissions" does not exist, skipping
DROP TABLE
CREATE TABLE
BEGIN
COMMIT
psql:/tmp/selfhosted/nacos/nacos-pg.sql:275: NOTICE:  table "roles" does not exist, skipping
DROP TABLE
CREATE TABLE
BEGIN
INSERT 0 1
COMMIT
psql:/tmp/selfhosted/nacos/nacos-pg.sql:292: NOTICE:  table "tenant_capacity" does not exist, skipping
DROP TABLE
CREATE TABLE
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
BEGIN
COMMIT
psql:/tmp/selfhosted/nacos/nacos-pg.sql:327: NOTICE:  table "tenant_info" does not exist, skipping
DROP TABLE
CREATE TABLE
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
COMMENT
BEGIN
COMMIT
psql:/tmp/selfhosted/nacos/nacos-pg.sql:358: NOTICE:  table "users" does not exist, skipping
DROP TABLE
CREATE TABLE
BEGIN
INSERT 0 1
COMMIT
CREATE INDEX
ALTER TABLE
CREATE INDEX
ALTER TABLE
CREATE INDEX
ALTER TABLE
CREATE INDEX
ALTER TABLE
CREATE INDEX
CREATE INDEX
ALTER TABLE
CREATE INDEX
ALTER TABLE
CREATE INDEX
CREATE INDEX
CREATE INDEX
ALTER TABLE
CREATE INDEX
CREATE INDEX
CREATE INDEX
ALTER TABLE
CREATE INDEX

=== Completed: Postgre Nacos Setup ===

=== Task: Nacos Config Setup ===
[Nacos Config Setup][node3] Running...
sed -i "s|^db.url.0=.*|db.url.0=jdbc:postgresql://${POSTGRESQL_IP}:5432/nacos_config|" /data/nacos/conf/application.properties

[Nacos Config Setup][node1] Running...
sed -i "s|^db.url.0=.*|db.url.0=jdbc:postgresql://${POSTGRESQL_IP}:5432/nacos_config|" /data/nacos/conf/application.properties

[Nacos Config Setup][node2] Running...
sed -i "s|^db.url.0=.*|db.url.0=jdbc:postgresql://${POSTGRESQL_IP}:5432/nacos_config|" /data/nacos/conf/application.properties

[Nacos Config Setup][node2] Success

[Nacos Config Setup][node3] Success

[Nacos Config Setup][node1] Success

=== Completed: Nacos Config Setup ===

=== Task: Systemd Nacos Service ===
[Systemd Nacos Service][node3] Running...
cat <<EOF > /etc/systemd/system/nacos.service
[Unit]
Description=Nacos Service
After=network.target

# Add a condition to wait for PostgreSQL to be available on node1
[Service]
Type=forking
User=root

# Check if PostGreSQL is ready on node1 before starting
ExecStartPre=/bin/bash -c 'until nc -z node1 5432; do echo "Waiting for PostGreSQL to be up on node1..."; sleep 3; done'
Environment=JAVA_HOME=/usr/lib/jvm/java-1.8.0-openjdk-1.8.0.412.b08-1.el7_9.x86_64
ExecStart=/data/nacos/bin/startup.sh
ExecStop=/data/nacos/bin/shutdown.sh

[Install]
WantedBy=multi-user.target
EOF
systemctl daemon-reload
systemctl enable nacos.service
systemctl start nacos.service
systemctl status nacos.service

[Systemd Nacos Service][node2] Running...
cat <<EOF > /etc/systemd/system/nacos.service
[Unit]
Description=Nacos Service
After=network.target

# Add a condition to wait for PostgreSQL to be available on node1
[Service]
Type=forking
User=root

# Check if PostGreSQL is ready on node1 before starting
ExecStartPre=/bin/bash -c 'until nc -z node1 5432; do echo "Waiting for PostGreSQL to be up on node1..."; sleep 3; done'
Environment=JAVA_HOME=/usr/lib/jvm/java-1.8.0-openjdk-1.8.0.412.b08-1.el7_9.x86_64
ExecStart=/data/nacos/bin/startup.sh
ExecStop=/data/nacos/bin/shutdown.sh

[Install]
WantedBy=multi-user.target
EOF
systemctl daemon-reload
systemctl enable nacos.service
systemctl start nacos.service
systemctl status nacos.service

[Systemd Nacos Service][node1] Running...
cat <<EOF > /etc/systemd/system/nacos.service
[Unit]
Description=Nacos Service
After=network.target

# Add a condition to wait for PostgreSQL to be available on node1
[Service]
Type=forking
User=root

# Check if PostGreSQL is ready on node1 before starting
ExecStartPre=/bin/bash -c 'until nc -z node1 5432; do echo "Waiting for PostGreSQL to be up on node1..."; sleep 3; done'
Environment=JAVA_HOME=/usr/lib/jvm/java-1.8.0-openjdk-1.8.0.412.b08-1.el7_9.x86_64
ExecStart=/data/nacos/bin/startup.sh
ExecStop=/data/nacos/bin/shutdown.sh

[Install]
WantedBy=multi-user.target
EOF
systemctl daemon-reload
systemctl enable nacos.service
systemctl start nacos.service
systemctl status nacos.service

[Systemd Nacos Service][node1] Success
Created symlink from /etc/systemd/system/multi-user.target.wants/nacos.service to /etc/systemd/system/nacos.service.
● nacos.service - Nacos Service
   Loaded: loaded (/etc/systemd/system/nacos.service; enabled; vendor preset: disabled)
   Active: active (running) since Tue 2025-07-01 16:56:28 CST; 3ms ago
  Process: 3856 ExecStart=/data/nacos/bin/startup.sh (code=exited, status=0/SUCCESS)
  Process: 3851 ExecStartPre=/bin/bash -c until nc -z node1 5432; do echo "Waiting for PostGreSQL to be up on node1..."; sleep 3; done (code=exited, status=0/SUCCESS)
 Main PID: 3882 (java)
   CGroup: /system.slice/nacos.service
           └─3882 /usr/lib/jvm/java-1.8.0-openjdk-1.8.0.412.b08-1.el7_9.x86_64/bin/java -Djava.ext.dirs=/usr/lib/jvm/java-1.8.0-openjdk-1.8.0.412.b08-1.el7_9.x86_64/jre/lib/ext:/usr/lib/jvm/java-1.8.0-openjdk-1.8.0.412.b08-1.el7_9.x86_64/lib/ext -server -Xms2g -Xmx2g -Xmn1g -XX:MetaspaceSize=128m -XX:MaxMetaspaceSize=320m -XX:-OmitStackTraceInFastThrow -XX:+HeapDumpOnOutOfMemoryError -XX:HeapDumpPath=/data/nacos/logs/java_heapdump.hprof -XX:-UseLargePages -Dnacos.member.list= -Ddb.driverClassName=org.postgresql.Driver -Xloggc:/data/nacos/logs/nacos_gc.log -verbose:gc -XX:+PrintGCDetails -XX:+PrintGCDateStamps -XX:+PrintGCTimeStamps -XX:+UseGCLogFileRotation -XX:NumberOfGCLogFiles=10 -XX:GCLogFileSize=100M -Dloader.path=/data/nacos/plugins,/data/nacos/plugins/health,/data/nacos/plugins/cmdb,/data/nacos/plugins/selector -Dnacos.home=/data/nacos -jar /data/nacos/target/nacos-server.jar --spring.config.additional-location=file:/data/nacos/conf/ --logging.config=/data/nacos/conf/nacos-logback.xml --server.max-http-header-size=524288 nacos.nacos

Jul 01 16:56:27 selfhosted-0001 systemd[1]: Starting Nacos Service...
Jul 01 16:56:28 selfhosted-0001 startup.sh[3856]: /usr/lib/jvm/java-1.8.0-openjdk-1.8.0.412.b08-1.el7_9.x86_64/bin/java -Djava.ext.dirs=/usr/lib/jvm/java-1.8.0-openjdk-1.8.0.412.b08-1.el7_9.x86_64/jre/lib/ext:/usr/lib/jvm/java-1.8.0-openjdk-1.8.0.412.b08-1.el7_9.x86_64/lib/ext  -server -Xms2g -Xmx2g -Xmn1g -XX:MetaspaceSize=128m -XX:MaxMetaspaceSize=320m -XX:-OmitStackTraceInFastThrow -XX:+HeapDumpOnOutOfMemoryError -XX:HeapDumpPath=/data/nacos/logs/java_heapdump.hprof -XX:-UseLargePages -Dnacos.member.list= -Ddb.driverClassName=org.postgresql.Driver -Xloggc:/data/nacos/logs/nacos_gc.log -verbose:gc -XX:+PrintGCDetails -XX:+PrintGCDateStamps -XX:+PrintGCTimeStamps -XX:+UseGCLogFileRotation -XX:NumberOfGCLogFiles=10 -XX:GCLogFileSize=100M -Dloader.path=/data/nacos/plugins,/data/nacos/plugins/health,/data/nacos/plugins/cmdb,/data/nacos/plugins/selector -Dnacos.home=/data/nacos -jar /data/nacos/target/nacos-server.jar  --spring.config.additional-location=file:/data/nacos/conf/ --logging.config=/data/nacos/conf/nacos-logback.xml --server.max-http-header-size=524288
Jul 01 16:56:28 selfhosted-0001 startup.sh[3856]: nacos is starting with cluster
Jul 01 16:56:28 selfhosted-0001 systemd[1]: Started Nacos Service.
Jul 01 16:56:28 selfhosted-0001 startup.sh[3856]: nacos is starting，you can check the /data/nacos/logs/start.out

[Systemd Nacos Service][node2] Success
Created symlink from /etc/systemd/system/multi-user.target.wants/nacos.service to /etc/systemd/system/nacos.service.
● nacos.service - Nacos Service
   Loaded: loaded (/etc/systemd/system/nacos.service; enabled; vendor preset: disabled)
   Active: active (running) since Tue 2025-07-01 16:56:28 CST; 3ms ago
  Process: 14048 ExecStart=/data/nacos/bin/startup.sh (code=exited, status=0/SUCCESS)
  Process: 14045 ExecStartPre=/bin/bash -c until nc -z node1 5432; do echo "Waiting for PostGreSQL to be up on node1..."; sleep 3; done (code=exited, status=0/SUCCESS)
   CGroup: /system.slice/nacos.service
           └─14074 /usr/lib/jvm/java-1.8.0-openjdk-1.8.0.412.b08-1.el7_9.x86_64/bin/java -Djava.ext.dirs=/usr/lib/jvm/java-1.8.0-openjdk-1.8.0.412.b08-1.el7_9.x86_64/jre/lib/ext:/usr/lib/jvm/java-1.8.0-openjdk-1.8.0.412.b08-1.el7_9.x86_64/lib/ext -server -Xms2g -Xmx2g -Xmn1g -XX:MetaspaceSize=128m -XX:MaxMetaspaceSize=320m -XX:-OmitStackTraceInFastThrow -XX:+HeapDumpOnOutOfMemoryError -XX:HeapDumpPath=/data/nacos/logs/java_heapdump.hprof -XX:-UseLargePages -Dnacos.member.list= -Ddb.driverClassName=org.postgresql.Driver -Xloggc:/data/nacos/logs/nacos_gc.log -verbose:gc -XX:+PrintGCDetails -XX:+PrintGCDateStamps -XX:+PrintGCTimeStamps -XX:+UseGCLogFileRotation -XX:NumberOfGCLogFiles=10 -XX:GCLogFileSize=100M -Dloader.path=/data/nacos/plugins,/data/nacos/plugins/health,/data/nacos/plugins/cmdb,/data/nacos/plugins/selector -Dnacos.home=/data/nacos -jar /data/nacos/target/nacos-server.jar --spring.config.additional-location=file:/data/nacos/conf/ --logging.config=/data/nacos/conf/nacos-logback.xml --server.max-http-header-size=524288 nacos.nacos

Jul 01 16:56:27 selfhosted-0002 systemd[1]: Starting Nacos Service...
Jul 01 16:56:27 selfhosted-0002 bash[14045]: Connection to node1 (10.31.0.93) 5432 port [tcp/postgres] succeeded!
Jul 01 16:56:28 selfhosted-0002 startup.sh[14048]: /usr/lib/jvm/java-1.8.0-openjdk-1.8.0.412.b08-1.el7_9.x86_64/bin/java -Djava.ext.dirs=/usr/lib/jvm/java-1.8.0-openjdk-1.8.0.412.b08-1.el7_9.x86_64/jre/lib/ext:/usr/lib/jvm/java-1.8.0-openjdk-1.8.0.412.b08-1.el7_9.x86_64/lib/ext  -server -Xms2g -Xmx2g -Xmn1g -XX:MetaspaceSize=128m -XX:MaxMetaspaceSize=320m -XX:-OmitStackTraceInFastThrow -XX:+HeapDumpOnOutOfMemoryError -XX:HeapDumpPath=/data/nacos/logs/java_heapdump.hprof -XX:-UseLargePages -Dnacos.member.list= -Ddb.driverClassName=org.postgresql.Driver -Xloggc:/data/nacos/logs/nacos_gc.log -verbose:gc -XX:+PrintGCDetails -XX:+PrintGCDateStamps -XX:+PrintGCTimeStamps -XX:+UseGCLogFileRotation -XX:NumberOfGCLogFiles=10 -XX:GCLogFileSize=100M -Dloader.path=/data/nacos/plugins,/data/nacos/plugins/health,/data/nacos/plugins/cmdb,/data/nacos/plugins/selector -Dnacos.home=/data/nacos -jar /data/nacos/target/nacos-server.jar  --spring.config.additional-location=file:/data/nacos/conf/ --logging.config=/data/nacos/conf/nacos-logback.xml --server.max-http-header-size=524288
Jul 01 16:56:28 selfhosted-0002 startup.sh[14048]: nacos is starting with cluster
Jul 01 16:56:28 selfhosted-0002 startup.sh[14048]: nacos is starting，you can check the /data/nacos/logs/start.out
Jul 01 16:56:28 selfhosted-0002 systemd[1]: Started Nacos Service.

[Systemd Nacos Service][node3] Success
Created symlink from /etc/systemd/system/multi-user.target.wants/nacos.service to /etc/systemd/system/nacos.service.
● nacos.service - Nacos Service
   Loaded: loaded (/etc/systemd/system/nacos.service; enabled; vendor preset: disabled)
   Active: active (running) since Tue 2025-07-01 16:56:28 CST; 3ms ago
  Process: 29187 ExecStart=/data/nacos/bin/startup.sh (code=exited, status=0/SUCCESS)
  Process: 29184 ExecStartPre=/bin/bash -c until nc -z node1 5432; do echo "Waiting for PostGreSQL to be up on node1..."; sleep 3; done (code=exited, status=0/SUCCESS)
   CGroup: /system.slice/nacos.service
           └─29213 /usr/lib/jvm/java-1.8.0-openjdk-1.8.0.412.b08-1.el7_9.x86_64/bin/java -Djava.ext.dirs=/usr/lib/jvm/java-1.8.0-openjdk-1.8.0.412.b08-1.el7_9.x86_64/jre/lib/ext:/usr/lib/jvm/java-1.8.0-openjdk-1.8.0.412.b08-1.el7_9.x86_64/lib/ext -server -Xms2g -Xmx2g -Xmn1g -XX:MetaspaceSize=128m -XX:MaxMetaspaceSize=320m -XX:-OmitStackTraceInFastThrow -XX:+HeapDumpOnOutOfMemoryError -XX:HeapDumpPath=/data/nacos/logs/java_heapdump.hprof -XX:-UseLargePages -Dnacos.member.list= -Ddb.driverClassName=org.postgresql.Driver -Xloggc:/data/nacos/logs/nacos_gc.log -verbose:gc -XX:+PrintGCDetails -XX:+PrintGCDateStamps -XX:+PrintGCTimeStamps -XX:+UseGCLogFileRotation -XX:NumberOfGCLogFiles=10 -XX:GCLogFileSize=100M -Dloader.path=/data/nacos/plugins,/data/nacos/plugins/health,/data/nacos/plugins/cmdb,/data/nacos/plugins/selector -Dnacos.home=/data/nacos -jar /data/nacos/target/nacos-server.jar --spring.config.additional-location=file:/data/nacos/conf/ --logging.config=/data/nacos/conf/nacos-logback.xml --server.max-http-header-size=524288 nacos.nacos

Jul 01 16:56:27 selfhosted-0003 systemd[1]: Starting Nacos Service...
Jul 01 16:56:27 selfhosted-0003 bash[29184]: Connection to node1 (10.31.0.93) 5432 port [tcp/postgres] succeeded!
Jul 01 16:56:28 selfhosted-0003 startup.sh[29187]: /usr/lib/jvm/java-1.8.0-openjdk-1.8.0.412.b08-1.el7_9.x86_64/bin/java -Djava.ext.dirs=/usr/lib/jvm/java-1.8.0-openjdk-1.8.0.412.b08-1.el7_9.x86_64/jre/lib/ext:/usr/lib/jvm/java-1.8.0-openjdk-1.8.0.412.b08-1.el7_9.x86_64/lib/ext  -server -Xms2g -Xmx2g -Xmn1g -XX:MetaspaceSize=128m -XX:MaxMetaspaceSize=320m -XX:-OmitStackTraceInFastThrow -XX:+HeapDumpOnOutOfMemoryError -XX:HeapDumpPath=/data/nacos/logs/java_heapdump.hprof -XX:-UseLargePages -Dnacos.member.list= -Ddb.driverClassName=org.postgresql.Driver -Xloggc:/data/nacos/logs/nacos_gc.log -verbose:gc -XX:+PrintGCDetails -XX:+PrintGCDateStamps -XX:+PrintGCTimeStamps -XX:+UseGCLogFileRotation -XX:NumberOfGCLogFiles=10 -XX:GCLogFileSize=100M -Dloader.path=/data/nacos/plugins,/data/nacos/plugins/health,/data/nacos/plugins/cmdb,/data/nacos/plugins/selector -Dnacos.home=/data/nacos -jar /data/nacos/target/nacos-server.jar  --spring.config.additional-location=file:/data/nacos/conf/ --logging.config=/data/nacos/conf/nacos-logback.xml --server.max-http-header-size=524288
Jul 01 16:56:28 selfhosted-0003 startup.sh[29187]: nacos is starting with cluster
Jul 01 16:56:28 selfhosted-0003 startup.sh[29187]: nacos is starting，you can check the /data/nacos/logs/start.out
Jul 01 16:56:28 selfhosted-0003 systemd[1]: Started Nacos Service.

=== Completed: Systemd Nacos Service ===

[root@selfhosted-0001 petal]#
```

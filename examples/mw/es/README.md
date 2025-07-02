# es

```bash
[root@selfhosted-0001 petal]# ./petal-linux-amd64 -file task/mw/es/install.yml 
=== Task: Configure ES System Limits ===
[Configure ES System Limits][node3] Running...
cat <<EOF >> /etc/security/limits.conf
elasticsearch soft nproc 65535
elasticsearch hard nproc 65535
elasticsearch soft nofile 65535
elasticsearch hard nofile 65535
EOF
cat <<EOF >> /etc/sysctl.conf
vm.max_map_count = 262144
EOF
sysctl -p

[Configure ES System Limits][node1] Running...
cat <<EOF >> /etc/security/limits.conf
elasticsearch soft nproc 65535
elasticsearch hard nproc 65535
elasticsearch soft nofile 65535
elasticsearch hard nofile 65535
EOF
cat <<EOF >> /etc/sysctl.conf
vm.max_map_count = 262144
EOF
sysctl -p

[Configure ES System Limits][node2] Running...
cat <<EOF >> /etc/security/limits.conf
elasticsearch soft nproc 65535
elasticsearch hard nproc 65535
elasticsearch soft nofile 65535
elasticsearch hard nofile 65535
EOF
cat <<EOF >> /etc/sysctl.conf
vm.max_map_count = 262144
EOF
sysctl -p

[Configure ES System Limits][node2] Success
vm.swappiness = 0
net.core.somaxconn = 1024
net.ipv4.tcp_max_tw_buckets = 5000
net.ipv4.tcp_max_syn_backlog = 1024
vm.max_map_count = 262144

[Configure ES System Limits][node3] Success
vm.swappiness = 0
net.core.somaxconn = 1024
net.ipv4.tcp_max_tw_buckets = 5000
net.ipv4.tcp_max_syn_backlog = 1024
vm.max_map_count = 262144

[Configure ES System Limits][node1] Success
vm.swappiness = 0
net.core.somaxconn = 1024
net.ipv4.tcp_max_tw_buckets = 5000
net.ipv4.tcp_max_syn_backlog = 1024
vm.max_map_count = 262144

=== Completed: Configure ES System Limits ===

=== Task: Create User ===
[Create User][node3] Running...
useradd elasticsearch -s /sbin/nologin
mkdir -p /data/es717/data
mkdir -p /data/es717/logs

[Create User][node1] Running...
useradd elasticsearch -s /sbin/nologin
mkdir -p /data/es717/data
mkdir -p /data/es717/logs

[Create User][node2] Running...
useradd elasticsearch -s /sbin/nologin
mkdir -p /data/es717/data
mkdir -p /data/es717/logs

[Create User][node3] Success
Creating mailbox file: File exists

[Create User][node2] Success
Creating mailbox file: File exists

[Create User][node1] Success
Creating mailbox file: File exists

=== Completed: Create User ===

=== Task: Download && Install ES Packages ===
[Download && Install ES Packages][node3] Running...
obsutil cp -r -f obs://selfhosted/pkg/es/ /tmp/selfhosted/
cd /tmp/selfhosted/es/
rpm -ivh elasticsearch-7.17.11-x86_64.rpm
systemctl daemon-reload

[Download && Install ES Packages][node1] Running...
obsutil cp -r -f obs://selfhosted/pkg/es/ /tmp/selfhosted/
cd /tmp/selfhosted/es/
rpm -ivh elasticsearch-7.17.11-x86_64.rpm
systemctl daemon-reload

[Download && Install ES Packages][node2] Running...
obsutil cp -r -f obs://selfhosted/pkg/es/ /tmp/selfhosted/
cd /tmp/selfhosted/es/
rpm -ivh elasticsearch-7.17.11-x86_64.rpm
systemctl daemon-reload

[Download && Install ES Packages][node2] Success
Start at 2025-06-30 09:33:45.08596509 +0000 UTC


Parallel:      5                   Jobs:          5                   
Threshold:     50.00MB             PartSize:      auto                
VerifyLength:  false               VerifyMd5:     false               
CheckpointDir: /root/.obsutil_checkpoint     

Task id: 27c101d2-67cf-4e9c-8623-522dd79f4844
OutputDir: /root/.obsutil_output         
TempFileDir: /root                         

Succeed count:      8         Failed count:       0         3.75 843.15MB/s 7/8 769.39MB/769.39MB 0s
Succeed bytes:      769.39MB  
Metrics [max cost:1942 ms, min cost:60 ms, average cost:646.25 ms, average tps:3.58, transferred size:769.39MB]

Task id: 27c101d2-67cf-4e9c-8623-522dd79f4844
warning: elasticsearch-7.17.11-x86_64.rpm: Header V4 RSA/SHA512 Signature, key ID d88e42b4: NOKEY
Preparing...                          ########################################
Updating / installing...
elasticsearch-0:7.17.11-1             ########################################
### NOT starting on installation, please execute the following statements to configure elasticsearch service to start automatically using systemd
 sudo systemctl daemon-reload
 sudo systemctl enable elasticsearch.service
### You can start elasticsearch service by executing
 sudo systemctl start elasticsearch.service
Created elasticsearch keystore in /etc/elasticsearch/elasticsearch.keystore

[Download && Install ES Packages][node3] Success
Start at 2025-06-30 09:33:45.082590084 +0000 UTC


Parallel:      5                   Jobs:          5                   
Threshold:     50.00MB             PartSize:      auto                
VerifyLength:  false               VerifyMd5:     false               
CheckpointDir: /root/.obsutil_checkpoint     

Task id: 906db76b-98fd-4daa-bb26-aa9c94961f30
OutputDir: /root/.obsutil_output         
TempFileDir: /root                         

Succeed count:      8         Failed count:       0          453.38MB/s 8/8 769.39MB/769.39MB 1.899s
Succeed bytes:      769.39MB  
Metrics [max cost:1890 ms, min cost:6 ms, average cost:797.38 ms, average tps:3.69, transferred size:769.39MB]

Task id: 906db76b-98fd-4daa-bb26-aa9c94961f30
warning: elasticsearch-7.17.11-x86_64.rpm: Header V4 RSA/SHA512 Signature, key ID d88e42b4: NOKEY
Preparing...                          ########################################
Updating / installing...
elasticsearch-0:7.17.11-1             ########################################
### NOT starting on installation, please execute the following statements to configure elasticsearch service to start automatically using systemd
 sudo systemctl daemon-reload
 sudo systemctl enable elasticsearch.service
### You can start elasticsearch service by executing
 sudo systemctl start elasticsearch.service
Created elasticsearch keystore in /etc/elasticsearch/elasticsearch.keystore

[Download && Install ES Packages][node1] Success
Start at 2025-06-30 09:33:45.090018135 +0000 UTC


Parallel:      5                   Jobs:          5                   
Threshold:     50.00MB             PartSize:      auto                
VerifyLength:  false               VerifyMd5:     false               
CheckpointDir: /root/.obsutil_checkpoint     

Task id: a592a3a6-5956-4300-a07d-3320c403bce2
OutputDir: /root/.obsutil_output         
TempFileDir: /root                         

Succeed count:      8         Failed count:       0         3.33 735.89MB/s 5/8 769.39MB/769.39MB 0s
Succeed bytes:      769.39MB  
Metrics [max cost:1524 ms, min cost:22 ms, average cost:653.38 ms, average tps:4.55, transferred size:769.39MB]

Task id: a592a3a6-5956-4300-a07d-3320c403bce2
warning: elasticsearch-7.17.11-x86_64.rpm: Header V4 RSA/SHA512 Signature, key ID d88e42b4: NOKEY
Preparing...                          ########################################
Updating / installing...
elasticsearch-0:7.17.11-1             ########################################
### NOT starting on installation, please execute the following statements to configure elasticsearch service to start automatically using systemd
 sudo systemctl daemon-reload
 sudo systemctl enable elasticsearch.service
### You can start elasticsearch service by executing
 sudo systemctl start elasticsearch.service
Created elasticsearch keystore in /etc/elasticsearch/elasticsearch.keystore

=== Completed: Download && Install ES Packages ===

=== Task: Install JDK17 ===
[Install JDK17][node1] Running...
[Install JDK17][node3] Running...
cd /tmp/selfhosted/es/
tar xf jdk-17.0.2.tar -C /opt/
cat <<EOF >> /etc/profile
export ES_JAVA_HOME=/opt/jdk-17.0.2
export ES_PATH_CONF=/etc/elasticsearch
LANG=en_US.UTF-8
SYSFONT=latarcyrheb-sun16
EOF

[Install JDK17][node2] Running...
cd /tmp/selfhosted/es/
tar xf jdk-17.0.2.tar -C /opt/
cat <<EOF >> /etc/profile
export ES_JAVA_HOME=/opt/jdk-17.0.2
export ES_PATH_CONF=/etc/elasticsearch
LANG=en_US.UTF-8
SYSFONT=latarcyrheb-sun16
EOF

cd /tmp/selfhosted/es/
tar xf jdk-17.0.2.tar -C /opt/
cat <<EOF >> /etc/profile
export ES_JAVA_HOME=/opt/jdk-17.0.2
export ES_PATH_CONF=/etc/elasticsearch
LANG=en_US.UTF-8
SYSFONT=latarcyrheb-sun16
EOF

[Install JDK17][node1] Success

[Install JDK17][node3] Success

[Install JDK17][node2] Success

=== Completed: Install JDK17 ===

=== Task: Update Configuration node1 ===
[Update Configuration node1][node1] Running...
cat <<EOF > /etc/elasticsearch/elasticsearch.yml
cluster.name: es717-hqyg
node.name: node-1-7-hqyg
node.attr.rack: r1-7-hqyg
path.data: /data/es717/data
path.logs: /data/es717/logs
bootstrap.memory_lock: false
bootstrap.system_call_filter: false
network.host: node1
http.port: 9200
transport.host: node1
transport.tcp.port: 9300
discovery.seed_hosts: ["$NODE1_IP:9300","$NODE2_IP:9300","$NODE3_IP:9300"]
action.destructive_requires_name: true
indices.query.bool.max_clause_count: 10240
cluster.initial_master_nodes: ["$NODE1_IP:9300","$NODE2_IP:9300","$NODE3_IP:9300"]
node.master: true
node.data: true
xpack.security.enabled: true
xpack.security.transport.ssl.enabled: true
xpack.security.transport.ssl.key: x-pack/instance/instance.key
xpack.security.transport.ssl.certificate: x-pack/instance/instance.crt
xpack.security.transport.ssl.certificate_authorities: x-pack/ca/ca.crt
xpack.security.transport.ssl.verification_mode: certificate
xpack.security.transport.ssl.client_authentication: required
EOF

[Update Configuration node1][node1] Success

=== Completed: Update Configuration node1 ===

=== Task: Update Configuration node2 ===
[Update Configuration node2][node2] Running...
cat <<EOF > /etc/elasticsearch/elasticsearch.yml
cluster.name: es717-hqyg
node.name: node-2-7-hqyg
node.attr.rack: r2-7-hqyg
path.data: /data/es717/data
path.logs: /data/es717/logs
bootstrap.memory_lock: false
bootstrap.system_call_filter: false
network.host: node2
http.port: 9200
transport.host: node2
transport.tcp.port: 9300
discovery.seed_hosts: ["$NODE1_IP:9300","$NODE2_IP:9300","$NODE3_IP:9300"]
action.destructive_requires_name: true
indices.query.bool.max_clause_count: 10240
cluster.initial_master_nodes: ["$NODE1_IP:9300","$NODE2_IP:9300","$NODE3_IP:9300"]
node.master: true
node.data: true
xpack.security.enabled: true
xpack.security.transport.ssl.enabled: true
xpack.security.transport.ssl.key: x-pack/instance/instance.key
xpack.security.transport.ssl.certificate: x-pack/instance/instance.crt
xpack.security.transport.ssl.certificate_authorities: x-pack/ca/ca.crt
xpack.security.transport.ssl.verification_mode: certificate
xpack.security.transport.ssl.client_authentication: required
EOF

[Update Configuration node2][node2] Success

=== Completed: Update Configuration node2 ===

=== Task: Update Configuration node3 ===
[Update Configuration node3][node3] Running...
cat <<EOF > /etc/elasticsearch/elasticsearch.yml
cluster.name: es717-hqyg
node.name: node-3-7-hqyg
node.attr.rack: r3-7-hqyg
path.data: /data/es717/data
path.logs: /data/es717/logs
bootstrap.memory_lock: false
bootstrap.system_call_filter: false
network.host: node3
http.port: 9200
transport.host: node3
transport.tcp.port: 9300
discovery.seed_hosts: ["$NODE1_IP:9300","$NODE2_IP:9300","$NODE3_IP:9300"]
action.destructive_requires_name: true
indices.query.bool.max_clause_count: 10240
cluster.initial_master_nodes: ["$NODE1_IP:9300","$NODE2_IP:9300","$NODE3_IP:9300"]
node.master: true
node.data: true
xpack.security.enabled: true
xpack.security.transport.ssl.enabled: true
xpack.security.transport.ssl.key: x-pack/instance/instance.key
xpack.security.transport.ssl.certificate: x-pack/instance/instance.crt
xpack.security.transport.ssl.certificate_authorities: x-pack/ca/ca.crt
xpack.security.transport.ssl.verification_mode: certificate
xpack.security.transport.ssl.client_authentication: required
EOF

[Update Configuration node3][node3] Success

=== Completed: Update Configuration node3 ===

=== Task: Install Plugins ===
[Install Plugins][node3] Running...
cd /tmp/selfhosted/es/
tar zxf plugins.tgz -C /usr/share/elasticsearch/plugins/
touch /etc/elasticsearch/char_filter_text.txt

[Install Plugins][node1] Running...
cd /tmp/selfhosted/es/
tar zxf plugins.tgz -C /usr/share/elasticsearch/plugins/
touch /etc/elasticsearch/char_filter_text.txt

[Install Plugins][node2] Running...
cd /tmp/selfhosted/es/
tar zxf plugins.tgz -C /usr/share/elasticsearch/plugins/
touch /etc/elasticsearch/char_filter_text.txt

[Install Plugins][node3] Success

[Install Plugins][node1] Success

[Install Plugins][node2] Success

=== Completed: Install Plugins ===

=== Task: Install xpack Files ===
[Install xpack Files][node3] Running...
[Install xpack Files][node1] Running...
cd /tmp/selfhosted/es/
tar xf x-pack.tar -C /etc/elasticsearch/

[Install xpack Files][node2] Running...
cd /tmp/selfhosted/es/
tar xf x-pack.tar -C /etc/elasticsearch/

cd /tmp/selfhosted/es/
tar xf x-pack.tar -C /etc/elasticsearch/

[Install xpack Files][node1] Success

[Install xpack Files][node2] Success

[Install xpack Files][node3] Success

=== Completed: Install xpack Files ===

=== Task: Update File & Directory Permissions ===
[Update File & Directory Permissions][node3] Running...
chown elasticsearch.elasticsearch -R /etc/elasticsearch
chown elasticsearch.elasticsearch -R /var/lib/elasticsearch
chown elasticsearch.elasticsearch -R /etc/init.d/elasticsearch
chown elasticsearch.elasticsearch -R /var/log/elasticsearch
chown elasticsearch.elasticsearch -R /usr/share/elasticsearch
# chown elasticsearch.elasticsearch -R /var/run/elasticsearch
chown elasticsearch.elasticsearch -R /home/elasticsearch
chown elasticsearch.elasticsearch -R /data/es717/data
chown elasticsearch.elasticsearch -R /data/es717/logs

[Update File & Directory Permissions][node1] Running...
chown elasticsearch.elasticsearch -R /etc/elasticsearch
chown elasticsearch.elasticsearch -R /var/lib/elasticsearch
chown elasticsearch.elasticsearch -R /etc/init.d/elasticsearch
chown elasticsearch.elasticsearch -R /var/log/elasticsearch
chown elasticsearch.elasticsearch -R /usr/share/elasticsearch
# chown elasticsearch.elasticsearch -R /var/run/elasticsearch
chown elasticsearch.elasticsearch -R /home/elasticsearch
chown elasticsearch.elasticsearch -R /data/es717/data
chown elasticsearch.elasticsearch -R /data/es717/logs

[Update File & Directory Permissions][node2] Running...
chown elasticsearch.elasticsearch -R /etc/elasticsearch
chown elasticsearch.elasticsearch -R /var/lib/elasticsearch
chown elasticsearch.elasticsearch -R /etc/init.d/elasticsearch
chown elasticsearch.elasticsearch -R /var/log/elasticsearch
chown elasticsearch.elasticsearch -R /usr/share/elasticsearch
# chown elasticsearch.elasticsearch -R /var/run/elasticsearch
chown elasticsearch.elasticsearch -R /home/elasticsearch
chown elasticsearch.elasticsearch -R /data/es717/data
chown elasticsearch.elasticsearch -R /data/es717/logs

[Update File & Directory Permissions][node3] Success

[Update File & Directory Permissions][node2] Success

[Update File & Directory Permissions][node1] Success

=== Completed: Update File & Directory Permissions ===

=== Task: Start ES ===
[Start ES][node3] Running...
systemctl daemon-reload
systemctl start elasticsearch

[Start ES][node1] Running...
systemctl daemon-reload
systemctl start elasticsearch

[Start ES][node2] Running...
systemctl daemon-reload
systemctl start elasticsearch

[Start ES][node1] Success

[Start ES][node3] Success

[Start ES][node2] Success

=== Completed: Start ES ===

=== Task: Set Password ===
[Set Password][node3] Running...
/usr/share/elasticsearch/bin/elasticsearch-setup-passwords auto -b \
  -E "elasticsearch.host=http://$NODE3_IP:9200" \
  > $PASSWORD_FILE
grep '^PASSWORD elastic =' $PASSWORD_FILE | awk '{print $4}' > $ELASTIC_PASSWORD_FILE
export ELASTIC_PASSWORD=$(cat $ELASTIC_PASSWORD_FILE)
echo "Elastic Password: $ELASTIC_PASSWORD"

[Set Password][node3] Success
Elastic Password: UeS3i07XdI3JrRqVZaR6

=== Completed: Set Password ===

=== Task: Check Cluster ===
[Check Cluster][node3] Running...
export ELASTIC_PASSWORD=$(cat $ELASTIC_PASSWORD_FILE)
echo "Elastic Password: $ELASTIC_PASSWORD"
curl -u "elastic:$ELASTIC_PASSWORD" -XGET http://node1:9200/_cat/indices
curl -u "elastic:$ELASTIC_PASSWORD" -XGET http://node1:9200/_cluster/health?pretty
curl -u "elastic:$ELASTIC_PASSWORD" -XGET http://node1:9200/_cat/nodes

[Check Cluster][node3] Success
Elastic Password: UeS3i07XdI3JrRqVZaR6
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0green open .security-7 hloJomwMR3q2mOnWW6bppQ 1 1 7 0 50.4kb 25.2kb
100    68  100    68    0     0    395      0 --:--:-- --:--:-- --:--:--   397
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   463  100   463    0     0  32564 {    0 --:--:-- --:--:-- --:--:--     0
  "cluster_name" : "es717-hqyg",
  "status" : "green",
  "timed_out" : false,
  "number_of_nodes" : 3,
  "number_of_data_nodes" : 3,
  "active_primary_shards" : 2,
  "active_shards" : 4,
  "relocating_shards" : 0,
  "initializing_shards" : 0,
  "unassigned_shards" : 0,
  "delayed_unassigned_shards" : 0,
  "number_of_pending_tasks" : 0,
  "number_of_in_flight_fetch" : 0,
  "task_max_waiting_in_queue_millis" : 0,
  "active_shards_percent_as_number" : 100.0
}
     0 --:--:-- --:--:-- --:--:-- 33071
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   192  100   192    0     0   5542      0 -10.31.0.93  31 98 34 1.24 0.51 0.29 cdfhilmrstw - node-1-7-hqyg
10.31.0.246 30 88 42 0.92 0.27 0.13 cdfhilmrstw - node-2-7-hqyg
10.31.0.68  34 89 53 1.03 0.29 0.13 cdfhilmrstw * node-3-7-hqyg
-:--:-- --:--:-- --:--:--  5647

=== Completed: Check Cluster ===

=== Task: Install Kibana ===
[Install Kibana][node3] Running...
cd /tmp/selfhosted/es/
rpm -ivh kibana-7.17.11-x86_64.rpm

[Install Kibana][node3] Success
warning: kibana-7.17.11-x86_64.rpm: Header V4 RSA/SHA512 Signature, key ID d88e42b4: NOKEY
Preparing...                          ########################################
Updating / installing...
kibana-7.17.11-1                      ########################################
Creating kibana group... OK
Creating kibana user... OK
Created Kibana keystore in /etc/kibana/kibana.keystore

=== Completed: Install Kibana ===

=== Task: Change Kibana Password ===
[Change Kibana Password][node3] Running...
export ELASTIC_PASSWORD=$(cat $ELASTIC_PASSWORD_FILE)
echo "Elastic Password: $ELASTIC_PASSWORD"
cat <<EOF >/etc/kibana/kibana.yml
server.port: 5601
server.host: "0.0.0.0"
elasticsearch.hosts: ["http://node3:9200"]
kibana.index: ".kibana"
pid.file: /var/run/kibana.pid
logging.dest: stdout
elasticsearch.username: elastic
elasticsearch.password: $ELASTIC_PASSWORD
EOF

[Change Kibana Password][node3] Success
Elastic Password: UeS3i07XdI3JrRqVZaR6

=== Completed: Change Kibana Password ===

=== Task: Start Kibana ===
[Start Kibana][node3] Running...
systemctl daemon-reload
systemctl enable kibana
systemctl start kibana
systemctl status kibana      

[Start Kibana][node3] Success
Created symlink from /etc/systemd/system/multi-user.target.wants/kibana.service to /etc/systemd/system/kibana.service.
● kibana.service - Kibana
   Loaded: loaded (/etc/systemd/system/kibana.service; enabled; vendor preset: disabled)
   Active: active (running) since Mon 2025-06-30 17:34:31 CST; 3ms ago
     Docs: https://www.elastic.co
 Main PID: 26833 (kibana)
   CGroup: /system.slice/kibana.service
           ├─26833 /bin/sh /usr/share/kibana/bin/kibana --logging.dest="/var/log/kibana/kibana.log" --pid.file="/run/kibana/kibana.pid" --deprecation.skip_deprecated_settings[0]="logging.dest"
           ├─26836 /bin/sh /usr/share/kibana/bin/kibana --logging.dest="/var/log/kibana/kibana.log" --pid.file="/run/kibana/kibana.pid" --deprecation.skip_deprecated_settings[0]="logging.dest"
           └─26837 grep -v ^#

Jun 30 17:34:31 selfhosted-0003 systemd[1]: Started Kibana.

=== Completed: Start Kibana ===

=== Task: Post ===
[Post][node3] Running...
cd ~
rm -fr /tmp/selfhosted/es/
rm -f $PASSWORD_FILE

[Post][node1] Running...
cd ~
rm -fr /tmp/selfhosted/es/
rm -f $PASSWORD_FILE

[Post][node2] Running...
cd ~
rm -fr /tmp/selfhosted/es/
rm -f $PASSWORD_FILE

[Post][node1] Success

[Post][node3] Success

[Post][node2] Success

=== Completed: Post ===

[root@selfhosted-0001 petal]#
```

# rmq

```bash
[root@selfhosted-0001 petal]# ./petal-linux-amd64 -file task/mw/rmq/install.yml 
=== Task: Configure System Limits ===
[Configure System Limits][node3] Running...
cat <<EOF >> /etc/security/limits.conf
rabbitmq soft nofile 1024000
rabbitmq hard nofile 1024000
rabbitmq soft nproc 65535
rabbitmq hard nproc 65535
EOF
cat <<EOF >> /etc/sysctl.conf
fs.file-max = 65536
EOF
sysctl -p

[Configure System Limits][node1] Running...
[Configure System Limits][node2] Running...
cat <<EOF >> /etc/security/limits.conf
rabbitmq soft nofile 1024000
rabbitmq hard nofile 1024000
rabbitmq soft nproc 65535
rabbitmq hard nproc 65535
EOF
cat <<EOF >> /etc/sysctl.conf
fs.file-max = 65536
EOF
sysctl -p

cat <<EOF >> /etc/security/limits.conf
rabbitmq soft nofile 1024000
rabbitmq hard nofile 1024000
rabbitmq soft nproc 65535
rabbitmq hard nproc 65535
EOF
cat <<EOF >> /etc/sysctl.conf
fs.file-max = 65536
EOF
sysctl -p

[Configure System Limits][node1] Success
vm.swappiness = 0
net.core.somaxconn = 1024
net.ipv4.tcp_max_tw_buckets = 5000
net.ipv4.tcp_max_syn_backlog = 1024
vm.max_map_count = 262144
fs.file-max = 65536


[Configure System Limits][node2] Success
vm.swappiness = 0
net.core.somaxconn = 1024
net.ipv4.tcp_max_tw_buckets = 5000
net.ipv4.tcp_max_syn_backlog = 1024
vm.max_map_count = 262144
fs.file-max = 65536


[Configure System Limits][node3] Success
vm.swappiness = 0
net.core.somaxconn = 1024
net.ipv4.tcp_max_tw_buckets = 5000
net.ipv4.tcp_max_syn_backlog = 1024
vm.max_map_count = 262144
fs.file-max = 65536


=== Completed: Configure System Limits ===

=== Task: Check Directory ===
[Check Directory][node3] Running...
[ -d $TMP_DIR ] || mkdir -p $TMP_DIR

[Check Directory][node2] Running...
[ -d $TMP_DIR ] || mkdir -p $TMP_DIR

[Check Directory][node1] Running...
[ -d $TMP_DIR ] || mkdir -p $TMP_DIR

[Check Directory][node3] Success

[Check Directory][node1] Success

[Check Directory][node2] Success

=== Completed: Check Directory ===

=== Task: Download RabbitMQ Packages ===
[Download RabbitMQ Packages][node3] Running...
obsutil cp -r -f obs://selfhosted/pkg/rmq/ $TMP_DIR

[Download RabbitMQ Packages][node1] Running...
obsutil cp -r -f obs://selfhosted/pkg/rmq/ $TMP_DIR

[Download RabbitMQ Packages][node2] Running...
obsutil cp -r -f obs://selfhosted/pkg/rmq/ $TMP_DIR

[Download RabbitMQ Packages][node1] Success
Start at 2025-07-01 04:48:31.510484541 +0000 UTC


Parallel:      5                   Jobs:          5                   
Threshold:     50.00MB             PartSize:      auto                
VerifyLength:  false               VerifyMd5:     false               
CheckpointDir: /root/.obsutil_checkpoint     

Task id: 1381b552-4d49-4e11-8bef-a02edefd3def
OutputDir: /root/.obsutil_output         
TempFileDir: /root                         

Succeed count:      5         Failed count:       0         58.42 7.39GB/s 5/5 22.71MB/22.71MB 205ms
Succeed bytes:      22.71MB   
Metrics [max cost:194 ms, min cost:55 ms, average cost:84.20 ms, average tps:19.51, transferred size:22.71MB]

Task id: 1381b552-4d49-4e11-8bef-a02edefd3def

[Download RabbitMQ Packages][node3] Success
Start at 2025-07-01 04:48:31.502310503 +0000 UTC


Parallel:      5                   Jobs:          5                   
Threshold:     50.00MB             PartSize:      auto                
VerifyLength:  false               VerifyMd5:     false               
CheckpointDir: /root/.obsutil_checkpoint     

Task id: 7acf0536-ddfb-417a-a969-570ffc79d0cf
OutputDir: /root/.obsutil_output         
TempFileDir: /root                         

Succeed count:      5         Failed count:       0         6.37% tps:0.00 ?/s 4/5 10.53MB/22.71MB ?
Succeed bytes:      22.71MB   
Metrics [max cost:300 ms, min cost:23 ms, average cost:121.60 ms, average tps:12.86, transferred size:22.71MB]
[-------------------------------------------] 100.00% tps:36.59 210.27MB/s 5/5 22.71MB/22.71MB 309ms
Task id: 7acf0536-ddfb-417a-a969-570ffc79d0cf

[Download RabbitMQ Packages][node2] Success
Start at 2025-07-01 04:48:31.510067941 +0000 UTC


Parallel:      5                   Jobs:          5                   
Threshold:     50.00MB             PartSize:      auto                
VerifyLength:  false               VerifyMd5:     false               
CheckpointDir: /root/.obsutil_checkpoint     

Task id: a59935b7-b60d-44b1-b505-6ff1e8cf74bb
OutputDir: /root/.obsutil_output         
TempFileDir: /root                         

Succeed count:      5         Failed count:       0         22.22% tps:0.00 ?/s 4/5 5.05MB/22.71MB ?
Succeed bytes:      22.71MB   
Metrics [max cost:367 ms, min cost:66 ms, average cost:133.40 ms, average tps:10.44, transferred size:22.71MB]
[-------------------------------------------] 100.00% tps:21.98 125.46MB/s 5/5 22.71MB/22.71MB 383ms
Task id: a59935b7-b60d-44b1-b505-6ff1e8cf74bb

=== Completed: Download RabbitMQ Packages ===

=== Task: Install Erlang etc ===
[Install Erlang etc][node3] Running...
rpm -ivh $TMP_DIR_RMQ/erlang-20.3.8.26-1.el7.x86_64.rpm
rpm -qa erlang
rpm -ivh $TMP_DIR_RMQ/socat-1.7.3.2-2.el7.x86_64.rpm
rpm -qa socat

[Install Erlang etc][node2] Running...
rpm -ivh $TMP_DIR_RMQ/erlang-20.3.8.26-1.el7.x86_64.rpm
rpm -qa erlang
rpm -ivh $TMP_DIR_RMQ/socat-1.7.3.2-2.el7.x86_64.rpm
rpm -qa socat

[Install Erlang etc][node1] Running...
rpm -ivh $TMP_DIR_RMQ/erlang-20.3.8.26-1.el7.x86_64.rpm
rpm -qa erlang
rpm -ivh $TMP_DIR_RMQ/socat-1.7.3.2-2.el7.x86_64.rpm
rpm -qa socat

[Install Erlang etc][node1] Success
warning: /tmp/selfhosted/rmq/erlang-20.3.8.26-1.el7.x86_64.rpm: Header V4 RSA/SHA1 Signature, key ID 6026dfca: NOKEY
Preparing...                          ########################################
        package erlang-20.3.8.26-1.el7.x86_64 is already installed
erlang-20.3.8.26-1.el7.x86_64
Preparing...                          ########################################
        package socat-1.7.3.2-2.el7.x86_64 is already installed
socat-1.7.3.2-2.el7.x86_64

[Install Erlang etc][node3] Success
warning: /tmp/selfhosted/rmq/erlang-20.3.8.26-1.el7.x86_64.rpm: Header V4 RSA/SHA1 Signature, key ID 6026dfca: NOKEY
Preparing...                          ########################################
        package erlang-20.3.8.26-1.el7.x86_64 is already installed
erlang-20.3.8.26-1.el7.x86_64
Preparing...                          ########################################
        package socat-1.7.3.2-2.el7.x86_64 is already installed
socat-1.7.3.2-2.el7.x86_64

[Install Erlang etc][node2] Success
warning: /tmp/selfhosted/rmq/erlang-20.3.8.26-1.el7.x86_64.rpm: Header V4 RSA/SHA1 Signature, key ID 6026dfca: NOKEY
Preparing...                          ########################################
        package erlang-20.3.8.26-1.el7.x86_64 is already installed
erlang-20.3.8.26-1.el7.x86_64
Preparing...                          ########################################
        package socat-1.7.3.2-2.el7.x86_64 is already installed
socat-1.7.3.2-2.el7.x86_64

=== Completed: Install Erlang etc ===

=== Task: Install RabbitMQ ===
[Install RabbitMQ][node3] Running...
rpm -ivh $TMP_DIR_RMQ/rabbitmq-server-3.6.11-1.el7.noarch.rpm
rpm -qa rabbitmq-server
systemctl daemon-reload
mkdir -p /var/lib/rabbitmq
chown rabbitmq:rabbitmq /var/lib/rabbitmq

[Install RabbitMQ][node2] Running...
rpm -ivh $TMP_DIR_RMQ/rabbitmq-server-3.6.11-1.el7.noarch.rpm
rpm -qa rabbitmq-server
systemctl daemon-reload
mkdir -p /var/lib/rabbitmq
chown rabbitmq:rabbitmq /var/lib/rabbitmq

[Install RabbitMQ][node1] Running...
rpm -ivh $TMP_DIR_RMQ/rabbitmq-server-3.6.11-1.el7.noarch.rpm
rpm -qa rabbitmq-server
systemctl daemon-reload
mkdir -p /var/lib/rabbitmq
chown rabbitmq:rabbitmq /var/lib/rabbitmq

[Install RabbitMQ][node1] Success
warning: /tmp/selfhosted/rmq/rabbitmq-server-3.6.11-1.el7.noarch.rpm: Header V4 RSA/SHA512 Signature, key ID 6026dfca: NOKEY
Preparing...                          ########################################
        package rabbitmq-server-3.6.11-1.el7.noarch is already installed
rabbitmq-server-3.6.11-1.el7.noarch

[Install RabbitMQ][node3] Success
warning: /tmp/selfhosted/rmq/rabbitmq-server-3.6.11-1.el7.noarch.rpm: Header V4 RSA/SHA512 Signature, key ID 6026dfca: NOKEY
Preparing...                          ########################################
        package rabbitmq-server-3.6.11-1.el7.noarch is already installed
rabbitmq-server-3.6.11-1.el7.noarch

[Install RabbitMQ][node2] Success
warning: /tmp/selfhosted/rmq/rabbitmq-server-3.6.11-1.el7.noarch.rpm: Header V4 RSA/SHA512 Signature, key ID 6026dfca: NOKEY
Preparing...                          ########################################
        package rabbitmq-server-3.6.11-1.el7.noarch is already installed
rabbitmq-server-3.6.11-1.el7.noarch

=== Completed: Install RabbitMQ ===

=== Task: Configure Plugins ===
[Configure Plugins][node3] Running...
cp $TMP_DIR_RMQ/rabbitmq_delayed* /usr/lib/rabbitmq/lib/rabbitmq_server-3.6.11/plugins/
chown rabbitmq:rabbitmq /usr/lib/rabbitmq/lib/rabbitmq_server-3.6.11/plugins/rabbitmq_delayed_message_exchange*
chmod 644 /usr/lib/rabbitmq/lib/rabbitmq_server-3.6.11/plugins/rabbitmq_delayed_message_exchange*

[Configure Plugins][node1] Running...
cp $TMP_DIR_RMQ/rabbitmq_delayed* /usr/lib/rabbitmq/lib/rabbitmq_server-3.6.11/plugins/
chown rabbitmq:rabbitmq /usr/lib/rabbitmq/lib/rabbitmq_server-3.6.11/plugins/rabbitmq_delayed_message_exchange*
chmod 644 /usr/lib/rabbitmq/lib/rabbitmq_server-3.6.11/plugins/rabbitmq_delayed_message_exchange*

[Configure Plugins][node2] Running...
cp $TMP_DIR_RMQ/rabbitmq_delayed* /usr/lib/rabbitmq/lib/rabbitmq_server-3.6.11/plugins/
chown rabbitmq:rabbitmq /usr/lib/rabbitmq/lib/rabbitmq_server-3.6.11/plugins/rabbitmq_delayed_message_exchange*
chmod 644 /usr/lib/rabbitmq/lib/rabbitmq_server-3.6.11/plugins/rabbitmq_delayed_message_exchange*

[Configure Plugins][node1] Success

[Configure Plugins][node3] Success

[Configure Plugins][node2] Success

=== Completed: Configure Plugins ===

=== Task: Add Configuration File ===
[Add Configuration File][node3] Running...
cp /usr/share/doc/rabbitmq-server-3.6.11/rabbitmq.config.example /etc/rabbitmq/rabbitmq.config
chown root:rabbitmq /etc/rabbitmq/rabbitmq.config
systemctl enable rabbitmq-server
systemctl start rabbitmq-server
sleep 3

[Add Configuration File][node1] Running...
cp /usr/share/doc/rabbitmq-server-3.6.11/rabbitmq.config.example /etc/rabbitmq/rabbitmq.config
chown root:rabbitmq /etc/rabbitmq/rabbitmq.config
systemctl enable rabbitmq-server
systemctl start rabbitmq-server
sleep 3

[Add Configuration File][node2] Running...
cp /usr/share/doc/rabbitmq-server-3.6.11/rabbitmq.config.example /etc/rabbitmq/rabbitmq.config
chown root:rabbitmq /etc/rabbitmq/rabbitmq.config
systemctl enable rabbitmq-server
systemctl start rabbitmq-server
sleep 3

[Add Configuration File][node1] Success
Created symlink from /etc/systemd/system/multi-user.target.wants/rabbitmq-server.service to /usr/lib/systemd/system/rabbitmq-server.service.

[Add Configuration File][node3] Success
Created symlink from /etc/systemd/system/multi-user.target.wants/rabbitmq-server.service to /usr/lib/systemd/system/rabbitmq-server.service.

[Add Configuration File][node2] Success
Created symlink from /etc/systemd/system/multi-user.target.wants/rabbitmq-server.service to /usr/lib/systemd/system/rabbitmq-server.service.

=== Completed: Add Configuration File ===

=== Task: Use Plugin Management ===
[Use Plugin Management][node3] Running...
[Use Plugin Management][node1] Running...
rabbitmq-plugins enable rabbitmq_management
rabbitmq-plugins enable rabbitmq_delayed_message_exchange

[Use Plugin Management][node2] Running...
rabbitmq-plugins enable rabbitmq_management
rabbitmq-plugins enable rabbitmq_delayed_message_exchange

rabbitmq-plugins enable rabbitmq_management
rabbitmq-plugins enable rabbitmq_delayed_message_exchange

[Use Plugin Management][node1] Success
Plugin configuration unchanged.

Applying plugin configuration to rabbit@selfhosted-0001... nothing to do.
Plugin configuration unchanged.

Applying plugin configuration to rabbit@selfhosted-0001... nothing to do.

[Use Plugin Management][node3] Success
Plugin configuration unchanged.

Applying plugin configuration to rabbit@selfhosted-0003... nothing to do.
Plugin configuration unchanged.

Applying plugin configuration to rabbit@selfhosted-0003... nothing to do.

[Use Plugin Management][node2] Success
Plugin configuration unchanged.

Applying plugin configuration to rabbit@selfhosted-0002... nothing to do.
Plugin configuration unchanged.

Applying plugin configuration to rabbit@selfhosted-0002... nothing to do.

=== Completed: Use Plugin Management ===

=== Task: Stop RabbitMQ ===
[Stop RabbitMQ][node3] Running...
rabbitmq-plugins disable rabbitmq_management
rabbitmq-plugins disable rabbitmq_delayed_message_exchange
systemctl stop rabbitmq-server

[Stop RabbitMQ][node1] Running...
rabbitmq-plugins disable rabbitmq_management
rabbitmq-plugins disable rabbitmq_delayed_message_exchange
systemctl stop rabbitmq-server

[Stop RabbitMQ][node2] Running...
rabbitmq-plugins disable rabbitmq_management
rabbitmq-plugins disable rabbitmq_delayed_message_exchange
systemctl stop rabbitmq-server

[Stop RabbitMQ][node1] Success
The following plugins have been disabled:
  amqp_client
  cowlib
  cowboy
  rabbitmq_web_dispatch
  rabbitmq_management_agent
  rabbitmq_management

Applying plugin configuration to rabbit@selfhosted-0001... stopped 6 plugins.
The following plugins have been disabled:
  rabbitmq_delayed_message_exchange

Applying plugin configuration to rabbit@selfhosted-0001... stopped 1 plugin.

[Stop RabbitMQ][node3] Success
The following plugins have been disabled:
  amqp_client
  cowlib
  cowboy
  rabbitmq_web_dispatch
  rabbitmq_management_agent
  rabbitmq_management

Applying plugin configuration to rabbit@selfhosted-0003... stopped 6 plugins.
The following plugins have been disabled:
  rabbitmq_delayed_message_exchange

Applying plugin configuration to rabbit@selfhosted-0003... stopped 1 plugin.

[Stop RabbitMQ][node2] Success
The following plugins have been disabled:
  amqp_client
  cowlib
  cowboy
  rabbitmq_web_dispatch
  rabbitmq_management_agent
  rabbitmq_management

Applying plugin configuration to rabbit@selfhosted-0002... stopped 6 plugins.
The following plugins have been disabled:
  rabbitmq_delayed_message_exchange

Applying plugin configuration to rabbit@selfhosted-0002... stopped 1 plugin.

=== Completed: Stop RabbitMQ ===

=== Task: Sync Cookie ===
[Sync Cookie][node1] Running...
scp /var/lib/rabbitmq/.erlang.cookie node2:/var/lib/rabbitmq/.erlang.cookie
scp /var/lib/rabbitmq/.erlang.cookie node3:/var/lib/rabbitmq/.erlang.cookie

[Sync Cookie][node1] Success

=== Completed: Sync Cookie ===

=== Task: Start RabbitMQ ===
[Start RabbitMQ][node3] Running...
chown rabbitmq.rabbitmq -R /var/lib/rabbitmq/
systemctl start rabbitmq-server.service
rabbitmq-plugins enable rabbitmq_management
rabbitmq-plugins enable rabbitmq_delayed_message_exchange
sleep 3

[Start RabbitMQ][node1] Running...
chown rabbitmq.rabbitmq -R /var/lib/rabbitmq/
systemctl start rabbitmq-server.service
rabbitmq-plugins enable rabbitmq_management
rabbitmq-plugins enable rabbitmq_delayed_message_exchange
sleep 3

[Start RabbitMQ][node2] Running...
chown rabbitmq.rabbitmq -R /var/lib/rabbitmq/
systemctl start rabbitmq-server.service
rabbitmq-plugins enable rabbitmq_management
rabbitmq-plugins enable rabbitmq_delayed_message_exchange
sleep 3

[Start RabbitMQ][node1] Success
The following plugins have been enabled:
  amqp_client
  cowlib
  cowboy
  rabbitmq_web_dispatch
  rabbitmq_management_agent
  rabbitmq_management

Applying plugin configuration to rabbit@selfhosted-0001... started 6 plugins.
The following plugins have been enabled:
  rabbitmq_delayed_message_exchange

Applying plugin configuration to rabbit@selfhosted-0001... started 1 plugin.

[Start RabbitMQ][node3] Success
The following plugins have been enabled:
  amqp_client
  cowlib
  cowboy
  rabbitmq_web_dispatch
  rabbitmq_management_agent
  rabbitmq_management

Applying plugin configuration to rabbit@selfhosted-0003... started 6 plugins.
The following plugins have been enabled:
  rabbitmq_delayed_message_exchange

Applying plugin configuration to rabbit@selfhosted-0003... started 1 plugin.

[Start RabbitMQ][node2] Success
The following plugins have been enabled:
  amqp_client
  cowlib
  cowboy
  rabbitmq_web_dispatch
  rabbitmq_management_agent
  rabbitmq_management

Applying plugin configuration to rabbit@selfhosted-0002... started 6 plugins.
The following plugins have been enabled:
  rabbitmq_delayed_message_exchange

Applying plugin configuration to rabbit@selfhosted-0002... started 1 plugin.

=== Completed: Start RabbitMQ ===

=== Task: Add node2 node3 to Cluster ===
[Add node2 node3 to Cluster][node3] Running...
cat <<EOF >> /etc/hosts
10.31.0.93   selfhosted-0001
10.31.0.246  selfhosted-0002
10.31.0.68   selfhosted-0003
EOF
rabbitmqctl stop_app
rabbitmqctl join_cluster rabbit@$NODE1_HOSTNAME
rabbitmqctl start_app
sleep 3
rabbitmqctl cluster_status

[Add node2 node3 to Cluster][node2] Running...
cat <<EOF >> /etc/hosts
10.31.0.93   selfhosted-0001
10.31.0.246  selfhosted-0002
10.31.0.68   selfhosted-0003
EOF
rabbitmqctl stop_app
rabbitmqctl join_cluster rabbit@$NODE1_HOSTNAME
rabbitmqctl start_app
sleep 3
rabbitmqctl cluster_status

[Add node2 node3 to Cluster][node2] Success
Stopping rabbit application on node 'rabbit@selfhosted-0002'
Clustering node 'rabbit@selfhosted-0002' with 'rabbit@selfhosted-0001'
Starting node 'rabbit@selfhosted-0002'
Cluster status of node 'rabbit@selfhosted-0002'
[{nodes,[{disc,['rabbit@selfhosted-0001','rabbit@selfhosted-0002',
                'rabbit@selfhosted-0003']}]},
 {running_nodes,['rabbit@selfhosted-0003','rabbit@selfhosted-0001',
                 'rabbit@selfhosted-0002']},
 {cluster_name,<<"rabbit@selfhosted-0002">>},
 {partitions,[]},
 {alarms,[{'rabbit@selfhosted-0003',[]},
          {'rabbit@selfhosted-0001',[]},
          {'rabbit@selfhosted-0002',[]}]}]

[Add node2 node3 to Cluster][node3] Success
Stopping rabbit application on node 'rabbit@selfhosted-0003'
Clustering node 'rabbit@selfhosted-0003' with 'rabbit@selfhosted-0001'
Starting node 'rabbit@selfhosted-0003'
Cluster status of node 'rabbit@selfhosted-0003'
[{nodes,[{disc,['rabbit@selfhosted-0001','rabbit@selfhosted-0002',
                'rabbit@selfhosted-0003']}]},
 {running_nodes,['rabbit@selfhosted-0001','rabbit@selfhosted-0002',
                 'rabbit@selfhosted-0003']},
 {cluster_name,<<"rabbit@selfhosted-0002">>},
 {partitions,[]},
 {alarms,[{'rabbit@selfhosted-0001',[]},
          {'rabbit@selfhosted-0002',[]},
          {'rabbit@selfhosted-0003',[]}]}]

=== Completed: Add node2 node3 to Cluster ===

=== Task: RabbitMQ Users Ops ===
[RabbitMQ Users Ops][node1] Running...
rabbitmqctl list_users
rabbitmqctl add_user admin $ADMIN_PASSWORD
rabbitmqctl add_user hecom $HECOM_PASSWORD
rabbitmqctl set_user_tags admin administrator
rabbitmqctl set_permissions -p / admin ".*" ".*" ".*"
rabbitmqctl set_user_tags hecom administrator
rabbitmqctl delete_user guest

[RabbitMQ Users Ops][node1] Success
Listing users
guest   [administrator]
Creating user "admin"
Creating user "hecom"
Setting tags for user "admin" to [administrator]
Setting permissions for user "admin" in vhost "/"
Setting tags for user "hecom" to [administrator]
Deleting user "guest"

=== Completed: RabbitMQ Users Ops ===

=== Task: Create vhost ===
[Create vhost][node1] Running...
rabbitmqctl add_vhost product60
rabbitmqctl add_vhost product
rabbitmqctl add_vhost tenancy
rabbitmqctl set_permissions -p tenancy hecom ".*" ".*" ".*"
rabbitmqctl set_permissions -p product hecom ".*" ".*" ".*"

[Create vhost][node1] Success
Creating vhost "product60"
Creating vhost "product"
Creating vhost "tenancy"
Setting permissions for user "hecom" in vhost "tenancy"
Setting permissions for user "hecom" in vhost "product"

=== Completed: Create vhost ===

=== Task: Declare Exchange ===
[Declare Exchange][node1] Running...
rabbitmqctl eval 'rabbit_exchange:declare({resource, <<"product">>, exchange, <<"dead.letter.exchange">>}, topic, true, false, false, []).'
rabbitmqctl eval 'rabbit_exchange:declare({resource, <<"product">>, exchange, <<"paas.exchange">>}, topic, true, false, false, []).'
rabbitmqctl eval 'rabbit_exchange:declare({resource, <<"product">>, exchange, <<"paas.im.exchange">>}, direct, true, false, false, []).'

[Declare Exchange][node1] Success
{exchange,{resource,<<"product">>,exchange,<<"dead.letter.exchange">>},
          topic,true,false,false,[],undefined,undefined,
          {[],[]}}
{exchange,{resource,<<"product">>,exchange,<<"paas.exchange">>},
          topic,true,false,false,[],undefined,undefined,
          {[],[]}}
{exchange,{resource,<<"product">>,exchange,<<"paas.im.exchange">>},
          direct,true,false,false,[],undefined,undefined,
          {[],[]}}

=== Completed: Declare Exchange ===

=== Task: Declare and Bind Queue ===
[Declare and Bind Queue][node1] Running...
rabbitmqctl eval 'rabbit_amqqueue:declare({resource, <<"product">>, queue, <<"dead.letter.queue">>}, true, false, [], none).'
rabbitmqctl eval 'rabbit_binding:add({binding, {resource, <<"product">>, exchange, <<"dead.letter.exchange">>}, <<"dead.letter.other">>, {resource, <<"product">>, queue, <<"dead.letter.queue">>}, []}).'
rabbitmqctl eval 'rabbit_amqqueue:declare({resource, <<"product">>, queue, <<"paas.aggregation.calculate.delay.queue">>}, true, false, [], none).'
rabbitmqctl eval 'rabbit_binding:add({binding, {resource, <<"product">>, exchange, <<"paas.exchange">>}, <<"paas.es.doc.create.success.*.*">>, {resource, <<"product">>, queue, <<"paas.aggregation.calculate.delay.queue">>}, []}).'
rabbitmqctl eval 'rabbit_binding:add({binding, {resource, <<"product">>, exchange, <<"paas.exchange">>}, <<"paas.es.doc.delete.success.*.*">>, {resource, <<"product">>, queue, <<"paas.aggregation.calculate.delay.queue">>}, []}).'
rabbitmqctl eval 'rabbit_binding:add({binding, {resource, <<"product">>, exchange, <<"paas.exchange">>}, <<"paas.es.doc.update.success.*.*">>, {resource, <<"product">>, queue, <<"paas.aggregation.calculate.delay.queue">>}, []}).'
rabbitmqctl eval 'rabbit_amqqueue:declare({resource, <<"product">>, queue, <<"paas.es.data.delay.queue">>}, true, false, [], none).'
rabbitmqctl eval 'rabbit_binding:add({binding, {resource, <<"product">>, exchange, <<"paas.exchange">>}, <<"paas.es.doc.create.success.*.*">>, {resource, <<"product">>, queue, <<"paas.es.data.delay.queue">>}, []}).'
rabbitmqctl eval 'rabbit_binding:add({binding, {resource, <<"product">>, exchange, <<"paas.exchange">>}, <<"paas.es.doc.delete.success.*.*">>, {resource, <<"product">>, queue, <<"paas.es.data.delay.queue">>}, []}).'
rabbitmqctl eval 'rabbit_binding:add({binding, {resource, <<"product">>, exchange, <<"paas.exchange">>}, <<"paas.es.doc.update.success.*.*">>, {resource, <<"product">>, queue, <<"paas.es.data.delay.queue">>}, []}).'
rabbitmqctl eval 'rabbit_amqqueue:declare({resource, <<"product">>, queue, <<"paas.es.data.universal.delay.queue">>}, true, false, [], none).'  
rabbitmqctl eval 'rabbit_binding:add({binding, {resource, <<"product">>, exchange, <<"paas.exchange">>}, <<"paas.es.doc.create.success.*.*">>, {resource, <<"product">>, queue, <<"paas.es.data.universal.delay.queue">>}, []}).'  
rabbitmqctl eval 'rabbit_binding:add({binding, {resource, <<"product">>, exchange, <<"paas.exchange">>}, <<"paas.es.doc.delete.success.*.*">>, {resource, <<"product">>, queue, <<"paas.es.data.universal.delay.queue">>}, []}).'  
rabbitmqctl eval 'rabbit_binding:add({binding, {resource, <<"product">>, exchange, <<"paas.exchange">>}, <<"paas.es.doc.update.success.*.*">>, {resource, <<"product">>, queue, <<"paas.es.data.universal.delay.queue">>}, []}).'  
rabbitmqctl eval 'rabbit_amqqueue:declare({resource, <<"product">>, queue, <<"paas.es.data.universal.delay.queue">>}, true, false, [], none).'  
rabbitmqctl eval 'rabbit_binding:add({binding, {resource, <<"product">>, exchange, <<"paas.exchange">>}, <<"paas.es.doc.create.success.*.*">>, {resource, <<"product">>, queue, <<"paas.es.data.universal.delay.queue">>}, []}).'  
rabbitmqctl eval 'rabbit_binding:add({binding, {resource, <<"product">>, exchange, <<"paas.exchange">>}, <<"paas.es.doc.delete.success.*.*">>, {resource, <<"product">>, queue, <<"paas.es.data.universal.delay.queue">>}, []}).'  
rabbitmqctl eval 'rabbit_binding:add({binding, {resource, <<"product">>, exchange, <<"paas.exchange">>}, <<"paas.es.doc.update.success.*.*">>, {resource, <<"product">>, queue, <<"paas.es.data.universal.delay.queue">>}, []}).'  

[Declare and Bind Queue][node1] Success
{new,{amqqueue,{resource,<<"product">>,queue,<<"dead.letter.queue">>},
               true,false,none,[],<11091.999.0>,[],[],[],undefined,[],[],live,
               0}}
ok
{new,{amqqueue,{resource,<<"product">>,queue,
                         <<"paas.aggregation.calculate.delay.queue">>},
               true,false,none,[],<11091.1012.0>,[],[],[],undefined,[],[],
               live,0}}
ok
ok
ok
{new,{amqqueue,{resource,<<"product">>,queue,<<"paas.es.data.delay.queue">>},
               true,false,none,[],<11091.1035.0>,[],[],[],undefined,[],[],
               live,0}}
ok
ok
ok
{new,{amqqueue,{resource,<<"product">>,queue,
                         <<"paas.es.data.universal.delay.queue">>},
               true,false,none,[],<11091.1058.0>,[],[],[],undefined,[],[],
               live,0}}
ok
ok
ok
{existing,{amqqueue,{resource,<<"product">>,queue,
                              <<"paas.es.data.universal.delay.queue">>},
                    true,false,none,[],<11091.1058.0>,[],[],[],undefined,[],
                    [],live,0}}
ok
ok
ok

=== Completed: Declare and Bind Queue ===

=== Task: Set Policy ===
[Set Policy][node1] Running...
rabbitmqctl set_policy  -p "product" "AGG TTL"  "^paas.aggregation.calculate.delay.queue$" '{"ha-mode":"exactly","ha-params":1,"dead-letter-exchange":"paas.exchange", "dead-letter-routing-key":"paas.aggregation.calculate", "message-ttl":2000}'  --priority 1 --apply-to queues
rabbitmqctl set_policy  -p "product" "IM TTL"  "^paas.im.sendMsg.delay.queue$" '{"ha-mode":"exactly","ha-params":1,"dead-letter-exchange":"paas.im.exchange", "dead-letter-routing-key":"paas.im.message.send", "message-ttl":1000}'  --priority 1 --apply-to queues
rabbitmqctl set_policy  -p "product" "ES TTL"  "^paas.es.data.delay.queue$" '{"ha-mode":"exactly","ha-params":1,"dead-letter-exchange":"paas.exchange", "dead-letter-routing-key":"paas.es.doc.success.delay", "message-ttl":1000}'  --priority 1 --apply-to queues
rabbitmqctl set_policy  -p "product" "DLX_ALL"  '^(?!(^paas.aggregation.calculate.delay.queue$|^dead.letter.queue$|^paas.im.sendMsg.delay.queue$|^paas.es.data.delay.queue$|^\w{1,}-\w{1,}-\w{1,}-\w{1,}-\w{1,}$))' '{"ha-mode":"exactly","ha-params":1, "dead-letter-exchange":"dead.letter.exchange", "dead-letter-routing-key":"dead.letter.other"}'  --priority 0 --apply-to queues
rabbitmqctl set_policy  -p "product" "DLX_D"  "^dead.letter.queue$" '{"ha-mode":"exactly","ha-params":1,"ha-sync-mode":"automatic"}'  --priority 0 --apply-to queues
rabbitmqctl set_policy  -p "tenancy" "HA"  "tenancy.*" '{"ha-mode":"exactly","ha-params":1,"ha-sync-mode":"automatic"}'  --priority 0 --apply-to queues
rabbitmqctl set_policy  -p "product" "ES TTL UNIVERSAL"  "^paas.es.data.universal.delay.queue$" '{"ha-mode":"exactly","ha-params":1,"dead-letter-exchange":"paas.exchange", "dead-letter-routing-key":"paas.es.doc.success.universal.delay", "message-ttl":2000}'  --priority 1 --apply-to queues

[Set Policy][node1] Success
Setting policy "AGG TTL" for pattern "^paas.aggregation.calculate.delay.queue$" to "{\"ha-mode\":\"exactly\",\"ha-params\":1,\"dead-letter-exchange\":\"paas.exchange\", \"dead-letter-routing-key\":\"paas.aggregation.calculate\", \"message-ttl\":2000}" with priority "1"
Setting policy "IM TTL" for pattern "^paas.im.sendMsg.delay.queue$" to "{\"ha-mode\":\"exactly\",\"ha-params\":1,\"dead-letter-exchange\":\"paas.im.exchange\", \"dead-letter-routing-key\":\"paas.im.message.send\", \"message-ttl\":1000}" with priority "1"
Setting policy "ES TTL" for pattern "^paas.es.data.delay.queue$" to "{\"ha-mode\":\"exactly\",\"ha-params\":1,\"dead-letter-exchange\":\"paas.exchange\", \"dead-letter-routing-key\":\"paas.es.doc.success.delay\", \"message-ttl\":1000}" with priority "1"
Setting policy "DLX_ALL" for pattern "^(?!(^paas.aggregation.calculate.delay.queue$|^dead.letter.queue$|^paas.im.sendMsg.delay.queue$|^paas.es.data.delay.queue$|^\\w{1,}-\\w{1,}-\\w{1,}-\\w{1,}-\\w{1,}$))" to "{\"ha-mode\":\"exactly\",\"ha-params\":1, \"dead-letter-exchange\":\"dead.letter.exchange\", \"dead-letter-routing-key\":\"dead.letter.other\"}" with priority "0"
Setting policy "DLX_D" for pattern "^dead.letter.queue$" to "{\"ha-mode\":\"exactly\",\"ha-params\":1,\"ha-sync-mode\":\"automatic\"}" with priority "0"
Setting policy "HA" for pattern "tenancy.*" to "{\"ha-mode\":\"exactly\",\"ha-params\":1,\"ha-sync-mode\":\"automatic\"}" with priority "0"
Setting policy "ES TTL UNIVERSAL" for pattern "^paas.es.data.universal.delay.queue$" to "{\"ha-mode\":\"exactly\",\"ha-params\":1,\"dead-letter-exchange\":\"paas.exchange\", \"dead-letter-routing-key\":\"paas.es.doc.success.universal.delay\", \"message-ttl\":2000}" with priority "1"

=== Completed: Set Policy ===

[root@selfhosted-0001 petal]# 
```

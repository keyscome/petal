# zk

## Overview

Apache ZooKeeper 3.4.8 is a centralized coordination service used for distributed configuration, naming, synchronization, and group membership — commonly used as the metadata/coordination backend for Kafka, HBase, and similar systems. This example deploys a 3-node ZooKeeper quorum and installs **ZKUI**, a web-based ZooKeeper browser, on `node3`. It also imports initial ZooKeeper data from a JSON export and applies updates from an Excel file using `zktool`.

## Architecture

- **Topology**: 3-node quorum (`node1`, `node2`, `node3`)
- **ZooKeeper version**: 3.4.8
- **Client port**: `3000` (replaces the default `2181`)
- **Quorum ports**: `3001`/`3002` (node1), `3003`/`3004` (node2), `3005`/`3006` (node3) — leader-election and peer-communication ports
- **Install path**: `/data/zookeeper-3.4.8/` (`$ZK_HOME`)
- **Data directory**: `$ZK_HOME/data/` (contains `myid` per node and the snapshot/log data)
- **Config file**: `$ZK_CONFIG_FILE` (`zoo.cfg`)
- **ZKUI**: deployed on `node3` at `/data/zkui/` — web UI at port `9090`
- **Node IPs**: supplied via `$NODE1_IP`, `$NODE2_IP`, `$NODE3_IP`

## Files

| File | Description |
|------|-------------|
| `install.yml` | Deploys ZooKeeper quorum, configures it, imports seed data, and starts ZKUI |
| `uninstall.yml` | Stops ZooKeeper and ZKUI, then moves installation directories to a backup location |

### `install.yml`

1. **Check Directory** — Creates required base directories (`$TMP_DIR`, `$TMP_DIR_ZK`, `$DATA_HOME`, `$ZK_DATA_DIR`) on all nodes if they don't already exist.
2. **Download Zookeeper from OBS** — Fetches `zookeeper-3.4.8.tar.gz`, `zkui.tgz`, and `zktool` from OBS into `$TMP_DIR/zk/` on all nodes in parallel.
3. **Extract Zookeeper TGZ to DATA_HOME** — Extracts the ZooKeeper archive to `$DATA_HOME` on all three nodes.
4. **Update ZK quorum lines in zoo.cfg** — Iterates over the three node IPs and rewrites the `server.1`, `server.2`, and `server.3` lines in `zoo.cfg` using `sed`, assigning sequential quorum and election port pairs (3001/3002, 3003/3004, 3005/3006).
5. **Write myid for node1/2/3** — Creates `$ZK_HOME/data/myid` with the node's numeric ID (1, 2, or 3). These three tasks run on their respective nodes.
6. **Start Zookeeper** — Runs `zkServer.sh start` on all three nodes.
7. **Check after Start Zookeeper** — Waits 3 seconds, then verifies the `myid` file, checks the running ZK process, queries `zkServer.sh status` (shows leader/follower role), and confirms port 3000 is listening.
8. **Extract ZKUI** — Extracts `zkui.tgz` to `$DATA_HOME` on `node3`.
9. **Change Configuration of ZKUI** — Updates the ZKUI config file with the ZooKeeper connection string (`$ZK_STRING`) via `sed`.
10. **Start ZKUI** — Launches the ZKUI JAR (`zkui-2.0-SNAPSHOT-jar-with-dependencies.jar`) as a background process (`nohup`) on `node3`, logging to `zkui-out.log`.
11. **Check after Start ZKUI** — Waits 3 seconds, then verifies the ZKUI process and that port 9090 is listening.
12. **Import Zookeeper Data** — Runs `zktool import` on `node3` using `zookeeper_export.json` as the input, seeding the ZooKeeper namespace with initial configuration data.
13. **Update Zookeeper Data** — Downloads `zk.xlsx` from OBS and runs `zktool update` to apply incremental ZooKeeper node updates from the spreadsheet.

### `uninstall.yml`

1. **Check Directory** — Ensures the backup temp directory exists.
2. **Stop ZKServer** — Runs `zkServer.sh stop` on all three nodes.
3. **Move Zookeeper to $TMP_DIR_ZK** — Moves the `/data/zookeeper-3.4.8/` directory to the temp backup location.
4. **Check after Uninstallation of Zookeeper** — Verifies no ZooKeeper processes remain and port 3000 is no longer in use.
5. **Stop ZKUI** — Kills the ZKUI Java process on `node3`.
6. **Move ZKUI jar to $TMP_DIR_ZK** — Moves the ZKUI directory to the temp backup location.
7. **Check after Uninstallation of ZKUI** — Verifies port 9090 is no longer in use.

## Example

```bash
[root@selfhosted-0001 petal]# ./petal-linux-amd64 -file task/mw/zk/install.yml 
=== Task: Check Directory ===
[Check Directory][node3] Running...
[ -d $TMP_DIR ] || mkdir -p $TMP_DIR
[ -d $TMP_DIR_ZK ] || mkdir -p $TMP_DIR_ZK
[ -d $DATA_HOME ] || mkdir -p $DATA_HOME
[ -d $ZK_DATA_DIR ] || mkdir -p $ZK_DATA_DIR

[Check Directory][node1] Running...
[ -d $TMP_DIR ] || mkdir -p $TMP_DIR
[ -d $TMP_DIR_ZK ] || mkdir -p $TMP_DIR_ZK
[ -d $DATA_HOME ] || mkdir -p $DATA_HOME
[ -d $ZK_DATA_DIR ] || mkdir -p $ZK_DATA_DIR

[Check Directory][node2] Running...
[ -d $TMP_DIR ] || mkdir -p $TMP_DIR
[ -d $TMP_DIR_ZK ] || mkdir -p $TMP_DIR_ZK
[ -d $DATA_HOME ] || mkdir -p $DATA_HOME
[ -d $ZK_DATA_DIR ] || mkdir -p $ZK_DATA_DIR

[Check Directory][node1] Success

[Check Directory][node2] Success

[Check Directory][node3] Success

=== Completed: Check Directory ===

=== Task: Download Zookeeper from OBS ===
[Download Zookeeper from OBS][node3] Running...
[Download Zookeeper from OBS][node2] Running...
obsutil cp -r -f obs://selfhosted/pkg/zk/ $TMP_DIR
ls -l $TMP_DIR/zk

[Download Zookeeper from OBS][node1] Running...
obsutil cp -r -f obs://selfhosted/pkg/zk/ $TMP_DIR
ls -l $TMP_DIR/zk

obsutil cp -r -f obs://selfhosted/pkg/zk/ $TMP_DIR
ls -l $TMP_DIR/zk

[Download Zookeeper from OBS][node3] Success
Start at 2025-06-30 09:25:53.434307387 +0000 UTC


Parallel:      5                   Jobs:          5                   
Threshold:     50.00MB             PartSize:      auto                
VerifyLength:  false               VerifyMd5:     false               
CheckpointDir: /root/.obsutil_checkpoint     

Task id: eff6e67d-ecb1-4150-88bd-5868f9109680
OutputDir: /root/.obsutil_output         
TempFileDir: /root                         

Succeed count:      5         Failed count:       0         .26 518.80MB/s 5/5 38.91MB/38.91MB 278ms
Succeed bytes:      38.91MB   
Metrics [max cost:263 ms, min cost:146 ms, average cost:178.00 ms, average tps:14.39, transferred size:38.91MB]

Task id: eff6e67d-ecb1-4150-88bd-5868f9109680
total 39860
-rw-r-----  1 root      root       9931044 Jun 30 17:25 zktool
drwxr-xr-x  2 root      root          4096 Jun 19 16:59 zkui
-rw-r-----  1 root      root       8396415 Jun 30 17:25 zkui.tgz
drwxr-xr-x 11 nomaduser nomaduser     4096 Feb  6  2016 zookeeper-3.4.8
-rw-r-----  1 root      root      22250626 Jun 30 17:25 zookeeper-3.4.8.tar.gz
-rw-r-----  1 root      root        221889 Jun 30 17:25 zookeeper_export.json

[Download Zookeeper from OBS][node2] Success
Start at 2025-06-30 09:25:53.435706985 +0000 UTC


Parallel:      5                   Jobs:          5                   
Threshold:     50.00MB             PartSize:      auto                
VerifyLength:  false               VerifyMd5:     false               
CheckpointDir: /root/.obsutil_checkpoint     

Task id: eb1e7f35-d9cc-4008-847d-c7accb350132
OutputDir: /root/.obsutil_output         
TempFileDir: /root                         

Succeed count:      5         Failed count:       0         .20 682.63MB/s 5/5 38.91MB/38.91MB 258ms
Succeed bytes:      38.91MB   
Metrics [max cost:249 ms, min cost:136 ms, average cost:158.60 ms, average tps:15.44, transferred size:38.91MB]

Task id: eb1e7f35-d9cc-4008-847d-c7accb350132
total 39856
-rw-r-----  1 root      root       9931044 Jun 30 17:25 zktool
-rw-r-----  1 root      root       8396415 Jun 30 17:25 zkui.tgz
drwxr-xr-x 11 nomaduser nomaduser     4096 Feb  6  2016 zookeeper-3.4.8
-rw-r-----  1 root      root      22250626 Jun 30 17:25 zookeeper-3.4.8.tar.gz
-rw-r-----  1 root      root        221889 Jun 30 17:25 zookeeper_export.json

[Download Zookeeper from OBS][node1] Success
Start at 2025-06-30 09:25:53.438582563 +0000 UTC


Parallel:      5                   Jobs:          5                   
Threshold:     50.00MB             PartSize:      auto                
VerifyLength:  false               VerifyMd5:     false               
CheckpointDir: /root/.obsutil_checkpoint     

Task id: 03aa4453-05e4-4654-8c72-5f95c2b4ba3c
OutputDir: /root/.obsutil_output         
TempFileDir: /root                         

Succeed count:      5         Failed count:       0         .91 125.92MB/s 5/5 38.91MB/38.91MB 510ms
Succeed bytes:      38.91MB   
Metrics [max cost:499 ms, min cost:144 ms, average cost:263.40 ms, average tps:7.84, transferred size:38.91MB]

Task id: 03aa4453-05e4-4654-8c72-5f95c2b4ba3c
total 39876
-rw-r----- 1 root      root       9931044 Jun 30 17:25 zktool
-rw-r----- 1 root      root       8396415 Jun 30 17:25 zkui.tgz
-rw-r----- 1 root      root         17608 Jun 30 11:52 zk.xlsx
drwxr-xr-x 3 nomaduser nomaduser     4096 Jun 29 17:17 zookeeper-3.4.8
-rw-r----- 1 root      root      22250626 Jun 30 17:25 zookeeper-3.4.8.tar.gz
-rw-r----- 1 root      root        221889 Jun 30 17:25 zookeeper_export.json

=== Completed: Download Zookeeper from OBS ===

=== Task: Extract Zookeeper TGZ to DATA_HOME ===
[Extract Zookeeper TGZ to DATA_HOME][node3] Running...
[Extract Zookeeper TGZ to DATA_HOME][node1] Running...
tar zxf $TMP_DIR/zk/zookeeper-3.4.8.tar.gz -C $DATA_HOME

tar zxf $TMP_DIR/zk/zookeeper-3.4.8.tar.gz -C $DATA_HOME

[Extract Zookeeper TGZ to DATA_HOME][node2] Running...
tar zxf $TMP_DIR/zk/zookeeper-3.4.8.tar.gz -C $DATA_HOME

[Extract Zookeeper TGZ to DATA_HOME][node2] Success

[Extract Zookeeper TGZ to DATA_HOME][node3] Success

[Extract Zookeeper TGZ to DATA_HOME][node1] Success

=== Completed: Extract Zookeeper TGZ to DATA_HOME ===

=== Task: Update ZK quorum lines in zoo.cfg ===
[Update ZK quorum lines in zoo.cfg][node3] Running...
i=1
for ip in ${NODE1_IP} ${NODE2_IP} ${NODE3_IP}; do
  port1=$((3001 + 2*(i-1)))
  port2=$((port1 + 1))
  sed -i "s|^server\.${i}=.*|server.${i}=${ip}:${port1}:${port2}|" ${ZK_CONFIG_FILE}
  i=$((i + 1))
done

[Update ZK quorum lines in zoo.cfg][node2] Running...
i=1
for ip in ${NODE1_IP} ${NODE2_IP} ${NODE3_IP}; do
  port1=$((3001 + 2*(i-1)))
  port2=$((port1 + 1))
  sed -i "s|^server\.${i}=.*|server.${i}=${ip}:${port1}:${port2}|" ${ZK_CONFIG_FILE}
  i=$((i + 1))
done

[Update ZK quorum lines in zoo.cfg][node1] Running...
i=1
for ip in ${NODE1_IP} ${NODE2_IP} ${NODE3_IP}; do
  port1=$((3001 + 2*(i-1)))
  port2=$((port1 + 1))
  sed -i "s|^server\.${i}=.*|server.${i}=${ip}:${port1}:${port2}|" ${ZK_CONFIG_FILE}
  i=$((i + 1))
done

[Update ZK quorum lines in zoo.cfg][node3] Success

[Update ZK quorum lines in zoo.cfg][node1] Success

[Update ZK quorum lines in zoo.cfg][node2] Success

=== Completed: Update ZK quorum lines in zoo.cfg ===

=== Task: Write myid for node1 ===
[Write myid for node1][node1] Running...
mkdir -p $ZK_HOME/data
echo 1 > $ZK_HOME/data/myid

[Write myid for node1][node1] Success

=== Completed: Write myid for node1 ===

=== Task: Write myid for node2 ===
[Write myid for node2][node2] Running...
mkdir -p $ZK_HOME/data
echo 2 > $ZK_HOME/data/myid

[Write myid for node2][node2] Success

=== Completed: Write myid for node2 ===

=== Task: Write myid for node3 ===
[Write myid for node3][node3] Running...
mkdir -p $ZK_HOME/data
echo 3 > $ZK_HOME/data/myid

[Write myid for node3][node3] Success

=== Completed: Write myid for node3 ===

=== Task: Start Zookeeper ===
[Start Zookeeper][node3] Running...
$ZK_HOME/bin/zkServer.sh start

[Start Zookeeper][node1] Running...
$ZK_HOME/bin/zkServer.sh start

[Start Zookeeper][node2] Running...
$ZK_HOME/bin/zkServer.sh start

[Start Zookeeper][node1] Success
ZooKeeper JMX enabled by default
Using config: /data/zookeeper-3.4.8/bin/../conf/zoo.cfg
Starting zookeeper ... STARTED

[Start Zookeeper][node3] Success
ZooKeeper JMX enabled by default
Using config: /data/zookeeper-3.4.8/bin/../conf/zoo.cfg
Starting zookeeper ... STARTED

[Start Zookeeper][node2] Success
ZooKeeper JMX enabled by default
Using config: /data/zookeeper-3.4.8/bin/../conf/zoo.cfg
Starting zookeeper ... STARTED

=== Completed: Start Zookeeper ===

=== Task: Check after Start Zookeeper ===
[Check after Start Zookeeper][node3] Running...
sleep 3
cat /data/zookeeper-3.4.8/data/myid
ps -ef | grep /data/zookeeper-3.4.8 | grep -v grep
$ZK_HOME/bin/zkServer.sh status
netstat -nltp | grep 3000

[Check after Start Zookeeper][node1] Running...
[Check after Start Zookeeper][node2] Running...
sleep 3
cat /data/zookeeper-3.4.8/data/myid
ps -ef | grep /data/zookeeper-3.4.8 | grep -v grep
$ZK_HOME/bin/zkServer.sh status
netstat -nltp | grep 3000

sleep 3
cat /data/zookeeper-3.4.8/data/myid
ps -ef | grep /data/zookeeper-3.4.8 | grep -v grep
$ZK_HOME/bin/zkServer.sh status
netstat -nltp | grep 3000

[Check after Start Zookeeper][node1] Success
1
root      8818     1 17 17:25 ?        00:00:00 java -Dzookeeper.log.dir=. -Dzookeeper.root.logger=INFO,CONSOLE -cp /data/zookeeper-3.4.8/bin/../build/classes:/data/zookeeper-3.4.8/bin/../build/lib/*.jar:/data/zookeeper-3.4.8/bin/../lib/slf4j-log4j12-1.6.1.jar:/data/zookeeper-3.4.8/bin/../lib/slf4j-api-1.6.1.jar:/data/zookeeper-3.4.8/bin/../lib/netty-3.7.0.Final.jar:/data/zookeeper-3.4.8/bin/../lib/log4j-1.2.16.jar:/data/zookeeper-3.4.8/bin/../lib/jline-0.9.94.jar:/data/zookeeper-3.4.8/bin/../zookeeper-3.4.8.jar:/data/zookeeper-3.4.8/bin/../src/java/lib/*.jar:/data/zookeeper-3.4.8/bin/../conf: -Dcom.sun.management.jmxremote -Dcom.sun.management.jmxremote.local.only=false org.apache.zookeeper.server.quorum.QuorumPeerMain /data/zookeeper-3.4.8/bin/../conf/zoo.cfg
root      8848  8699  0 17:25 pts/0    00:00:00 ssh -T node3 env NODE2_IP='10.31.0.246' ZK_STRING='10.31.0.93:3000,10.31.0.246:3000,10.31.0.68:3000' TMP_DIR_ZK='/tmp/selfhosted/zk/' DATA_HOME='/data' ZK_DATA_DIR='/data/zookeeper-3.4.8/data' ZK_HOME='/data/zookeeper-3.4.8' ZKUI_CONFIG_FILE='/data/zkui/config.cfg' ZK_CONFIG_FILE='/data/zookeeper-3.4.8/conf/zoo.cfg' ZKUI_HOME='/data/zkui' NODE1_IP='10.31.0.93' TMP_DIR='/tmp/selfhosted/' NODE3_IP='10.31.0.68' bash -s
root      8849  8699  0 17:25 pts/0    00:00:00 ssh -T node1 env NODE2_IP='10.31.0.246' ZK_STRING='10.31.0.93:3000,10.31.0.246:3000,10.31.0.68:3000' TMP_DIR_ZK='/tmp/selfhosted/zk/' DATA_HOME='/data' ZK_DATA_DIR='/data/zookeeper-3.4.8/data' ZK_HOME='/data/zookeeper-3.4.8' ZKUI_CONFIG_FILE='/data/zkui/config.cfg' ZK_CONFIG_FILE='/data/zookeeper-3.4.8/conf/zoo.cfg' ZKUI_HOME='/data/zkui' NODE1_IP='10.31.0.93' TMP_DIR='/tmp/selfhosted/' NODE3_IP='10.31.0.68' bash -s
root      8850  8699  0 17:25 pts/0    00:00:00 ssh -T node2 env NODE2_IP='10.31.0.246' ZK_STRING='10.31.0.93:3000,10.31.0.246:3000,10.31.0.68:3000' TMP_DIR_ZK='/tmp/selfhosted/zk/' DATA_HOME='/data' ZK_DATA_DIR='/data/zookeeper-3.4.8/data' ZK_HOME='/data/zookeeper-3.4.8' ZKUI_CONFIG_FILE='/data/zkui/config.cfg' ZK_CONFIG_FILE='/data/zookeeper-3.4.8/conf/zoo.cfg' ZKUI_HOME='/data/zkui' NODE1_IP='10.31.0.93' TMP_DIR='/tmp/selfhosted/' NODE3_IP='10.31.0.68' bash -s
ZooKeeper JMX enabled by default
Using config: /data/zookeeper-3.4.8/bin/../conf/zoo.cfg
Mode: follower
tcp6       0      0 :::3000                 :::*                    LISTEN      8818/java           

[Check after Start Zookeeper][node2] Success
2
root     17058     1  0 Jun18 ?        00:07:03 java -Dzookeeper.log.dir=. -Dzookeeper.root.logger=INFO,CONSOLE -cp /data/zookeeper-3.4.8/bin/../build/classes:/data/zookeeper-3.4.8/bin/../build/lib/*.jar:/data/zookeeper-3.4.8/bin/../lib/slf4j-log4j12-1.6.1.jar:/data/zookeeper-3.4.8/bin/../lib/slf4j-api-1.6.1.jar:/data/zookeeper-3.4.8/bin/../lib/netty-3.7.0.Final.jar:/data/zookeeper-3.4.8/bin/../lib/log4j-1.2.16.jar:/data/zookeeper-3.4.8/bin/../lib/jline-0.9.94.jar:/data/zookeeper-3.4.8/bin/../zookeeper-3.4.8.jar:/data/zookeeper-3.4.8/bin/../src/java/lib/*.jar:/data/zookeeper-3.4.8/bin/../conf: -Dcom.sun.management.jmxremote -Dcom.sun.management.jmxremote.local.only=false org.apache.zookeeper.server.quorum.QuorumPeerMain /data/zookeeper-3.4.8/bin/../conf/zoo.cfg
ZooKeeper JMX enabled by default
Using config: /data/zookeeper-3.4.8/bin/../conf/zoo.cfg
Error contacting service. It is probably not running.
tcp6       0      0 :::3000                 :::*                    LISTEN      17058/java          

[Check after Start Zookeeper][node3] Success
3
root     24874     1 19 17:25 ?        00:00:00 java -Dzookeeper.log.dir=. -Dzookeeper.root.logger=INFO,CONSOLE -cp /data/zookeeper-3.4.8/bin/../build/classes:/data/zookeeper-3.4.8/bin/../build/lib/*.jar:/data/zookeeper-3.4.8/bin/../lib/slf4j-log4j12-1.6.1.jar:/data/zookeeper-3.4.8/bin/../lib/slf4j-api-1.6.1.jar:/data/zookeeper-3.4.8/bin/../lib/netty-3.7.0.Final.jar:/data/zookeeper-3.4.8/bin/../lib/log4j-1.2.16.jar:/data/zookeeper-3.4.8/bin/../lib/jline-0.9.94.jar:/data/zookeeper-3.4.8/bin/../zookeeper-3.4.8.jar:/data/zookeeper-3.4.8/bin/../src/java/lib/*.jar:/data/zookeeper-3.4.8/bin/../conf: -Dcom.sun.management.jmxremote -Dcom.sun.management.jmxremote.local.only=false org.apache.zookeeper.server.quorum.QuorumPeerMain /data/zookeeper-3.4.8/bin/../conf/zoo.cfg
ZooKeeper JMX enabled by default
Using config: /data/zookeeper-3.4.8/bin/../conf/zoo.cfg
Mode: leader
tcp6       0      0 :::3000                 :::*                    LISTEN      24874/java          

=== Completed: Check after Start Zookeeper ===

=== Task: Extract ZKUI ===
[Extract ZKUI][node3] Running...
tar zxf $TMP_DIR_ZK/zkui.tgz -C $DATA_HOME

[Extract ZKUI][node3] Success

=== Completed: Extract ZKUI ===

=== Task: Change Configuration of ZKUI ===
[Change Configuration of ZKUI][node3] Running...
sed -i "s|zkServer=.*|zkServer=${ZK_STRING}|g" $ZKUI_CONFIG_FILE

[Change Configuration of ZKUI][node3] Success

=== Completed: Change Configuration of ZKUI ===

=== Task: Start ZKUI ===
[Start ZKUI][node3] Running...
cd $ZKUI_HOME
nohup java -jar zkui-2.0-SNAPSHOT-jar-with-dependencies.jar \
  > zkui-out.log 2>&1 </dev/null &

[Start ZKUI][node3] Success

=== Completed: Start ZKUI ===

=== Task: Check after Start ZKUI ===
[Check after Start ZKUI][node3] Running...
sleep 3
ps -ef | grep zkui-2.0-SNAPSHOT-jar-with-dependencies.jar | grep -v grep
netstat -nltp | grep 9090

[Check after Start ZKUI][node3] Success
root     24988     1 99 17:25 ?        00:00:03 java -jar zkui-2.0-SNAPSHOT-jar-with-dependencies.jar
tcp6       0      0 :::9090                 :::*                    LISTEN      24988/java          

=== Completed: Check after Start ZKUI ===

=== Task: Import Zookeeper Data ===
[Import Zookeeper Data][node3] Running...
cd $TMP_DIR_ZK
ls -l
chmod +x zktool
export TARGET_ZK=localhost:3000
export IMPORT_FILE=zookeeper_export.json
./zktool import

[Import Zookeeper Data][node3] Success
total 39860
-rw-r-----  1 root      root       9931044 Jun 30 17:25 zktool
drwxr-xr-x  2 root      root          4096 Jun 19 16:59 zkui
-rw-r-----  1 root      root       8396415 Jun 30 17:25 zkui.tgz
drwxr-xr-x 11 nomaduser nomaduser     4096 Feb  6  2016 zookeeper-3.4.8
-rw-r-----  1 root      root      22250626 Jun 30 17:25 zookeeper-3.4.8.tar.gz
-rw-r-----  1 root      root        221889 Jun 30 17:25 zookeeper_export.json
2025/06/30 17:26:03 connected to [::1]:3000
2025/06/30 17:26:03 authenticated: id=258886782955683840, timeout=10000
2025/06/30 17:26:03 re-submitting `0` credentials after reconnect
✅ Updated: /config/product/paas/thirdPartyPlatform/qyWeChat/pageUrl = core/WelcomePage
✅ Updated: /config/product/paas/globalsettings = 
...
✅ Updated: /config/product/paas/adcolumn/main/isOpen = false
Update done
2025/06/30 17:26:04 recv loop terminated: EOF
2025/06/30 17:26:04 send loop terminated: <nil>

=== Completed: Update Zookeeper Data ===

[root@selfhosted-0001 petal]#
```

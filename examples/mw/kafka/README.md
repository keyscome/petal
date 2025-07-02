# kafka

```bash
[root@selfhosted-0001 petal]# ./petal-linux-amd64 -file task/mw/kafka/install.yml 
=== Task: Download Kafka Package ===
[Download Kafka Package][node3] Running...
obsutil cp -r -f obs://selfhosted/pkg/kafka/ $TMP_DIR
ls -l $TMP_DIR/kafka

[Download Kafka Package][node2] Running...
obsutil cp -r -f obs://selfhosted/pkg/kafka/ $TMP_DIR
ls -l $TMP_DIR/kafka

[Download Kafka Package][node1] Running...
obsutil cp -r -f obs://selfhosted/pkg/kafka/ $TMP_DIR
ls -l $TMP_DIR/kafka

[Download Kafka Package][node2] Success
Start at 2025-06-30 09:32:23.496371753 +0000 UTC


Parallel:      5                   Jobs:          5                   
Threshold:     50.00MB             PartSize:      auto                
VerifyLength:  false               VerifyMd5:     false               
CheckpointDir: /root/.obsutil_checkpoint     

Task id: 4bd99b55-281e-41da-b570-fdb0db5b1c66
OutputDir: /root/.obsutil_output         
TempFileDir: /root                         

Succeed count:      2         Failed count:       0         0 271.54MB/s 2/2 108.35MB/108.35MB 600ms
Succeed bytes:      108.35MB  
Metrics [max cost:592 ms, min cost:592 ms, average cost:296.00 ms, average tps:1.66, transferred size:108.35MB]

Task id: 4bd99b55-281e-41da-b570-fdb0db5b1c66
total 110948
-rw-r----- 1 root root 113609072 Jun 30 17:32 kafka_2.12-3.6.1.tgz

[Download Kafka Package][node3] Success
Start at 2025-06-30 09:32:23.494892542 +0000 UTC


Parallel:      5                   Jobs:          5                   
Threshold:     50.00MB             PartSize:      auto                
VerifyLength:  false               VerifyMd5:     false               
CheckpointDir: /root/.obsutil_checkpoint     

Task id: 407857db-75fe-418f-85fa-9c457c5b3653
OutputDir: /root/.obsutil_output         
TempFileDir: /root                         

Succeed count:      2         Failed count:       0         .83% tps:0.00 ?/s 1/2 90.83MB/108.35MB ?
Succeed bytes:      108.35MB  
Metrics [max cost:761 ms, min cost:761 ms, average cost:380.50 ms, average tps:1.30, transferred size:108.35MB]

Task id: 407857db-75fe-418f-85fa-9c457c5b3653
total 110948
-rw-r----- 1 root root 113609072 Jun 30 17:32 kafka_2.12-3.6.1.tgz

[Download Kafka Package][node1] Success
Start at 2025-06-30 09:32:23.49574921 +0000 UTC


Parallel:      5                   Jobs:          5                   
Threshold:     50.00MB             PartSize:      auto                
VerifyLength:  false               VerifyMd5:     false               
CheckpointDir: /root/.obsutil_checkpoint     

Task id: 862df979-f6d3-4e84-8c94-fa499bc483e8
OutputDir: /root/.obsutil_output         
TempFileDir: /root                         

Succeed count:      2         Failed count:       0         5 189.75MB/s 2/2 108.35MB/108.35MB 772ms
Succeed bytes:      108.35MB  
Metrics [max cost:764 ms, min cost:764 ms, average cost:382.00 ms, average tps:1.30, transferred size:108.35MB]

Task id: 862df979-f6d3-4e84-8c94-fa499bc483e8
total 110948
-rw-r----- 1 root root 113609072 Jun 30 17:32 kafka_2.12-3.6.1.tgz

=== Completed: Download Kafka Package ===

=== Task: Extract Kafka TGZ to DATA_HOME ===
[Extract Kafka TGZ to DATA_HOME][node3] Running...
[Extract Kafka TGZ to DATA_HOME][node1] Running...
tar zxf $TMP_DIR/kafka/kafka_2.12-3.6.1.tgz -C /data
mv /data/kafka_2.12-3.6.1 /data/kafka
ls -l /data/kafka

[Extract Kafka TGZ to DATA_HOME][node2] Running...
tar zxf $TMP_DIR/kafka/kafka_2.12-3.6.1.tgz -C /data
mv /data/kafka_2.12-3.6.1 /data/kafka
ls -l /data/kafka

tar zxf $TMP_DIR/kafka/kafka_2.12-3.6.1.tgz -C /data
mv /data/kafka_2.12-3.6.1 /data/kafka
ls -l /data/kafka

[Extract Kafka TGZ to DATA_HOME][node2] Success
total 72
drwxr-xr-x 3 root root  4096 Nov 24  2023 bin
drwxr-xr-x 3 root root  4096 Nov 24  2023 config
drwxr-xr-x 2 root root 12288 Jun 30 17:32 libs
-rw-r--r-- 1 root root 15030 Nov 24  2023 LICENSE
drwxr-xr-x 2 root root  4096 Nov 24  2023 licenses
-rw-r--r-- 1 root root 28184 Nov 24  2023 NOTICE
drwxr-xr-x 2 root root  4096 Nov 24  2023 site-docs

[Extract Kafka TGZ to DATA_HOME][node1] Success
total 72
drwxr-xr-x 3 root root  4096 Nov 24  2023 bin
drwxr-xr-x 3 root root  4096 Nov 24  2023 config
drwxr-xr-x 2 root root 12288 Jun 30 17:32 libs
-rw-r--r-- 1 root root 15030 Nov 24  2023 LICENSE
drwxr-xr-x 2 root root  4096 Nov 24  2023 licenses
-rw-r--r-- 1 root root 28184 Nov 24  2023 NOTICE
drwxr-xr-x 2 root root  4096 Nov 24  2023 site-docs

[Extract Kafka TGZ to DATA_HOME][node3] Success
total 72
drwxr-xr-x 3 root root  4096 Nov 24  2023 bin
drwxr-xr-x 3 root root  4096 Nov 24  2023 config
drwxr-xr-x 2 root root 12288 Jun 30 17:32 libs
-rw-r--r-- 1 root root 15030 Nov 24  2023 LICENSE
drwxr-xr-x 2 root root  4096 Nov 24  2023 licenses
-rw-r--r-- 1 root root 28184 Nov 24  2023 NOTICE
drwxr-xr-x 2 root root  4096 Nov 24  2023 site-docs

=== Completed: Extract Kafka TGZ to DATA_HOME ===

=== Task: Update Kafka server.properties node1 ===
[Update Kafka server.properties node1][node1] Running...
sed -i "s|^broker.id=.*|broker.id=1|" $KAFKA_CONFIG_FILE
sed -i "s|^listeners=.*|listeners=PLAINTEXT://$(hostname -I):9092|" $KAFKA_CONFIG_FILE
sed -i "s|^log.dirs=.*|log.dirs=/data/kafka/kafka-logs|" $KAFKA_CONFIG_FILE
sed -i "s|^zookeeper.connect=.*|zookeeper.connect=${NODE1_IP}:3000,${NODE2_IP}:3000,${NODE3_IP}:3000|" $KAFKA_CONFIG_FILE

[Update Kafka server.properties node1][node1] Success

=== Completed: Update Kafka server.properties node1 ===

=== Task: Update Kafka server.properties node2 ===
[Update Kafka server.properties node2][node2] Running...
sed -i "s|^broker.id=.*|broker.id=2|" $KAFKA_CONFIG_FILE
sed -i "s|^listeners=.*|listeners=PLAINTEXT://$(hostname -I):9092|" $KAFKA_CONFIG_FILE
sed -i "s|^log.dirs=.*|log.dirs=/data/kafka/kafka-logs|" $KAFKA_CONFIG_FILE
sed -i "s|^zookeeper.connect=.*|zookeeper.connect=${NODE1_IP}:3000,${NODE2_IP}:3000,${NODE3_IP}:3000|" $KAFKA_CONFIG_FILE

[Update Kafka server.properties node2][node2] Success

=== Completed: Update Kafka server.properties node2 ===

=== Task: Update Kafka server.properties node3 ===
[Update Kafka server.properties node3][node3] Running...
sed -i "s|^broker.id=.*|broker.id=3|" $KAFKA_CONFIG_FILE
sed -i "s|^listeners=.*|listeners=PLAINTEXT://$(hostname -I):9092|" $KAFKA_CONFIG_FILE
sed -i "s|^log.dirs=.*|log.dirs=/data/kafka/kafka-logs|" $KAFKA_CONFIG_FILE
sed -i "s|^zookeeper.connect=.*|zookeeper.connect=${NODE1_IP}:3000,${NODE2_IP}:3000,${NODE3_IP}:3000|" $KAFKA_CONFIG_FILE

[Update Kafka server.properties node3][node3] Success

=== Completed: Update Kafka server.properties node3 ===

=== Task: Start Kafka Server ===
[Start Kafka Server][node3] Running...
/data/kafka/bin/kafka-server-start.sh -daemon $KAFKA_CONFIG_FILE
sleep 5
ps -ef | grep kafka
netstat -tulnp | grep 9092

[Start Kafka Server][node1] Running...
/data/kafka/bin/kafka-server-start.sh -daemon $KAFKA_CONFIG_FILE
sleep 5
ps -ef | grep kafka
netstat -tulnp | grep 9092

[Start Kafka Server][node2] Running...
/data/kafka/bin/kafka-server-start.sh -daemon $KAFKA_CONFIG_FILE
sleep 5
ps -ef | grep kafka
netstat -tulnp | grep 9092

[Start Kafka Server][node1] Success
root     11681 26724  0 17:32 pts/0    00:00:00 ./petal-linux-amd64 -file task/mw/kafka/install.yml
root     11756 11681  0 17:32 pts/0    00:00:00 ssh -T node3 env NODE3_IP='10.31.0.68' KAFKA_CONFIG_FILE='/data/kafka/config/server.properties' TMP_DIR='/tmp/selfhosted' NODE1_IP='10.31.0.93' NODE2_IP='10.31.0.246' bash -s
root     11757 11681  0 17:32 pts/0    00:00:00 ssh -T node2 env NODE3_IP='10.31.0.68' KAFKA_CONFIG_FILE='/data/kafka/config/server.properties' TMP_DIR='/tmp/selfhosted' NODE1_IP='10.31.0.93' NODE2_IP='10.31.0.246' bash -s
root     11758 11681  0 17:32 pts/0    00:00:00 ssh -T node1 env NODE3_IP='10.31.0.68' KAFKA_CONFIG_FILE='/data/kafka/config/server.properties' TMP_DIR='/tmp/selfhosted' NODE1_IP='10.31.0.93' NODE2_IP='10.31.0.246' bash -s
root     12147     1 70 17:32 ?        00:00:03 java -Xmx1G -Xms1G -server -XX:+UseG1GC -XX:MaxGCPauseMillis=20 -XX:InitiatingHeapOccupancyPercent=35 -XX:+ExplicitGCInvokesConcurrent -XX:MaxInlineLevel=15 -Djava.awt.headless=true -Xloggc:/data/kafka/bin/../logs/kafkaServer-gc.log -verbose:gc -XX:+PrintGCDetails -XX:+PrintGCDateStamps -XX:+PrintGCTimeStamps -XX:+UseGCLogFileRotation -XX:NumberOfGCLogFiles=10 -XX:GCLogFileSize=100M -Dcom.sun.management.jmxremote -Dcom.sun.management.jmxremote.authenticate=false -Dcom.sun.management.jmxremote.ssl=false -Dkafka.logs.dir=/data/kafka/bin/../logs -Dlog4j.configuration=file:/data/kafka/bin/../config/log4j.properties -cp /data/kafka/bin/../libs/activation-1.1.1.jar:/data/kafka/bin/../libs/aopalliance-repackaged-2.6.1.jar:/data/kafka/bin/../libs/argparse4j-0.7.0.jar:/data/kafka/bin/../libs/audience-annotations-0.12.0.jar:/data/kafka/bin/../libs/caffeine-2.9.3.jar:/data/kafka/bin/../libs/checker-qual-3.19.0.jar:/data/kafka/bin/../libs/commons-beanutils-1.9.4.jar:/data/kafka/bin/../libs/commons-cli-1.4.jar:/data/kafka/bin/../libs/commons-collections-3.2.2.jar:/data/kafka/bin/../libs/commons-digester-2.1.jar:/data/kafka/bin/../libs/commons-io-2.11.0.jar:/data/kafka/bin/../libs/commons-lang3-3.8.1.jar:/data/kafka/bin/../libs/commons-logging-1.2.jar:/data/kafka/bin/../libs/commons-validator-1.7.jar:/data/kafka/bin/../libs/connect-api-3.6.1.jar:/data/kafka/bin/../libs/connect-basic-auth-extension-3.6.1.jar:/data/kafka/bin/../libs/connect-json-3.6.1.jar:/data/kafka/bin/../libs/connect-mirror-3.6.1.jar:/data/kafka/bin/../libs/connect-mirror-client-3.6.1.jar:/data/kafka/bin/../libs/connect-runtime-3.6.1.jar:/data/kafka/bin/../libs/connect-transforms-3.6.1.jar:/data/kafka/bin/../libs/error_prone_annotations-2.10.0.jar:/data/kafka/bin/../libs/hk2-api-2.6.1.jar:/data/kafka/bin/../libs/hk2-locator-2.6.1.jar:/data/kafka/bin/../libs/hk2-utils-2.6.1.jar:/data/kafka/bin/../libs/jackson-annotations-2.13.5.jar:/data/kafka/bin/../libs/jackson-core-2.13.5.jar:/data/kafka/bin/../libs/jackson-databind-2.13.5.jar:/data/kafka/bin/../libs/jackson-dataformat-csv-2.13.5.jar:/data/kafka/bin/../libs/jackson-datatype-jdk8-2.13.5.jar:/data/kafka/bin/../libs/jackson-jaxrs-base-2.13.5.jar:/data/kafka/bin/../libs/jackson-jaxrs-json-provider-2.13.5.jar:/data/kafka/bin/../libs/jackson-module-jaxb-annotations-2.13.5.jar:/data/kafka/bin/../libs/jackson-module-scala_2.12-2.13.5.jar:/data/kafka/bin/../libs/jakarta.activation-api-1.2.2.jar:/data/kafka/bin/../libs/jakarta.annotation-api-1.3.5.jar:/data/kafka/bin/../libs/jakarta.inject-2.6.1.jar:/data/kafka/bin/../libs/jakarta.validation-api-2.0.2.jar:/data/kafka/bin/../libs/jakarta.ws.rs-api-2.1.6.jar:/data/kafka/bin/../libs/jakarta.xml.bind-api-2.3.3.jar:/data/kafka/bin/../libs/javassist-3.29.2-GA.jar:/data/kafka/bin/../libs/javax.activation-api-1.2.0.jar:/data/kafka/bin/../libs/javax.annotation-api-1.3.2.jar:/data/kafka/bin/../libs/javax.servlet-api-3.1.0.jar:/data/kafka/bin/../libs/javax.ws.rs-api-2.1.1.jar:/data/kafka/bin/../libs/jaxb-api-2.3.1.jar:/data/kafka/bin/../libs/jersey-client-2.39.1.jar:/data/kafka/bin/../libs/jersey-common-2.39.1.jar:/data/kafka/bin/../libs/jersey-container-servlet-2.39.1.jar:/data/kafka/bin/../libs/jersey-container-servlet-core-2.39.1.jar:/data/kafka/bin/../libs/jersey-hk2-2.39.1.jar:/data/kafka/bin/../libs/jersey-server-2.39.1.jar:/data/kafka/bin/../libs/jetty-client-9.4.52.v20230823.jar:/data/kafka/bin/../libs/jetty-continuation-9.4.52.v20230823.jar:/data/kafka/bin/../libs/jetty-http-9.4.52.v20230823.jar:/data/kafka/bin/../libs/jetty-io-9.4.52.v20230823.jar:/data/kafka/bin/../libs/jetty-security-9.4.52.v20230823.jar:/data/kafka/bin/../libs/jetty-server-9.4.52.v20230823.jar:/data/kafka/bin/../libs/jetty-servlet-9.4.52.v20230823.jar:/data/kafka/bin/../libs/jetty-servlets-9.4.52.v20230823.jar:/data/kafka/bin/../libs/jetty-util-9.4.52.v20230823.jar:/data/kafka/bin/../libs/jetty-util-ajax-9.4.52.v20230823.jar:/data/kafka/bin/../libs/jline-3.22.0.jar:/data/kafka/bin/../libs/jopt-simple-5.0.4.jar:/data/kafka/bin/../libs/jose4j-0.9.3.jar:/data/kafka/bin/../libs/jsr305-3.0.2.jar:/data/kafka/bin/../libs/kafka_2.12-3.6.1.jar:/data/kafka/bin/../libs/kafka-clients-3.6.1.jar:/data/kafka/bin/../libs/kafka-group-coordinator-3.6.1.jar:/data/kafka/bin/../libs/kafka-log4j-appender-3.6.1.jar:/data/kafka/bin/../libs/kafka-metadata-3.6.1.jar:/data/kafka/bin/../libs/kafka-raft-3.6.1.jar:/data/kafka/bin/../libs/kafka-server-common-3.6.1.jar:/data/kafka/bin/../libs/kafka-shell-3.6.1.jar:/data/kafka/bin/../libs/kafka-storage-3.6.1.jar:/data/kafka/bin/../libs/kafka-storage-api-3.6.1.jar:/data/kafka/bin/../libs/kafka-streams-3.6.1.jar:/data/kafka/bin/../libs/kafka-streams-examples-3.6.1.jar:/data/kafka/bin/../libs/kafka-streams-scala_2.12-3.6.1.jar:/data/kafka/bin/../libs/kafka-streams-test-utils-3.6.1.jar:/data/kafka/bin/../libs/kafka-tools-3.6.1.jar:/data/kafka/bin/../libs/kafka-tools-api-3.6.1.jar:/data/kafka/bin/../libs/lz4-java-1.8.0.jar:/data/kafka/bin/../libs/maven-artifact-3.8.8.jar:/data/kafka/bin/../libs/metrics-core-2.2.0.jar:/data/kafka/bin/../libs/metrics-core-4.1.12.1.jar:/data/kafka/bin/../libs/netty-buffer-4.1.100.Final.jar:/data/kafka/bin/../libs/netty-codec-4.1.100.Final.jar:/data/kafka/bin/../libs/netty-common-4.1.100.Final.jar:/data/kafka/bin/../libs/netty-handler-4.1.100.Final.jar:/data/kafka/bin/../libs/netty-resolver-4.1.100.Final.jar:/data/kafka/bin/../libs/netty-transport-4.1.100.Final.jar:/data/kafka/bin/../libs/netty-transport-classes-epoll-4.1.100.Final.jar:/data/kafka/bin/../libs/netty-transport-native-epoll-4.1.100.Final.jar:/data/kafka/bin/../libs/netty-transport-native-unix-common-4.1.100.Final.jar:/data/kafka/bin/../libs/osgi-resource-locator-1.0.3.jar:/data/kafka/bin/../libs/paranamer-2.8.jar:/data/kafka/bin/../libs/pcollections-4.0.1.jar:/data/kafka/bin/../libs/plexus-utils-3.3.1.jar:/data/kafka/bin/../libs/reflections-0.10.2.jar:/data/kafka/bin/../libs/reload4j-1.2.25.jar:/data/kafka/bin/../libs/rocksdbjni-7.9.2.jar:/data/kafka/bin/../libs/scala-collection-compat_2.12-2.10.0.jar:/data/kafka/bin/../libs/scala-java8-compat_2.12-1.0.2.jar:/data/kafka/bin/../libs/scala-library-2.12.18.jar:/data/kafka/bin/../libs/scala-logging_2.12-3.9.4.jar:/data/kafka/bin/../libs/scala-reflect-2.12.18.jar:/data/kafka/bin/../libs/slf4j-api-1.7.36.jar:/data/kafka/bin/../libs/slf4j-reload4j-1.7.36.jar:/data/kafka/bin/../libs/snappy-java-1.1.10.5.jar:/data/kafka/bin/../libs/swagger-annotations-2.2.8.jar:/data/kafka/bin/../libs/trogdor-3.6.1.jar:/data/kafka/bin/../libs/zookeeper-3.8.3.jar:/data/kafka/bin/../libs/zookeeper-jute-3.8.3.jar:/data/kafka/bin/../libs/zstd-jni-1.5.5-1.jar kafka.Kafka /data/kafka/config/server.properties
root     12245 11761  0 17:32 ?        00:00:00 grep kafka
tcp6       0      0 :::9092                 :::*                    LISTEN      12147/java          

[Start Kafka Server][node3] Success
root     25880     1 78 17:32 ?        00:00:03 java -Xmx1G -Xms1G -server -XX:+UseG1GC -XX:MaxGCPauseMillis=20 -XX:InitiatingHeapOccupancyPercent=35 -XX:+ExplicitGCInvokesConcurrent -XX:MaxInlineLevel=15 -Djava.awt.headless=true -Xloggc:/data/kafka/bin/../logs/kafkaServer-gc.log -verbose:gc -XX:+PrintGCDetails -XX:+PrintGCDateStamps -XX:+PrintGCTimeStamps -XX:+UseGCLogFileRotation -XX:NumberOfGCLogFiles=10 -XX:GCLogFileSize=100M -Dcom.sun.management.jmxremote -Dcom.sun.management.jmxremote.authenticate=false -Dcom.sun.management.jmxremote.ssl=false -Dkafka.logs.dir=/data/kafka/bin/../logs -Dlog4j.configuration=file:/data/kafka/bin/../config/log4j.properties -cp /data/kafka/bin/../libs/activation-1.1.1.jar:/data/kafka/bin/../libs/aopalliance-repackaged-2.6.1.jar:/data/kafka/bin/../libs/argparse4j-0.7.0.jar:/data/kafka/bin/../libs/audience-annotations-0.12.0.jar:/data/kafka/bin/../libs/caffeine-2.9.3.jar:/data/kafka/bin/../libs/checker-qual-3.19.0.jar:/data/kafka/bin/../libs/commons-beanutils-1.9.4.jar:/data/kafka/bin/../libs/commons-cli-1.4.jar:/data/kafka/bin/../libs/commons-collections-3.2.2.jar:/data/kafka/bin/../libs/commons-digester-2.1.jar:/data/kafka/bin/../libs/commons-io-2.11.0.jar:/data/kafka/bin/../libs/commons-lang3-3.8.1.jar:/data/kafka/bin/../libs/commons-logging-1.2.jar:/data/kafka/bin/../libs/commons-validator-1.7.jar:/data/kafka/bin/../libs/connect-api-3.6.1.jar:/data/kafka/bin/../libs/connect-basic-auth-extension-3.6.1.jar:/data/kafka/bin/../libs/connect-json-3.6.1.jar:/data/kafka/bin/../libs/connect-mirror-3.6.1.jar:/data/kafka/bin/../libs/connect-mirror-client-3.6.1.jar:/data/kafka/bin/../libs/connect-runtime-3.6.1.jar:/data/kafka/bin/../libs/connect-transforms-3.6.1.jar:/data/kafka/bin/../libs/error_prone_annotations-2.10.0.jar:/data/kafka/bin/../libs/hk2-api-2.6.1.jar:/data/kafka/bin/../libs/hk2-locator-2.6.1.jar:/data/kafka/bin/../libs/hk2-utils-2.6.1.jar:/data/kafka/bin/../libs/jackson-annotations-2.13.5.jar:/data/kafka/bin/../libs/jackson-core-2.13.5.jar:/data/kafka/bin/../libs/jackson-databind-2.13.5.jar:/data/kafka/bin/../libs/jackson-dataformat-csv-2.13.5.jar:/data/kafka/bin/../libs/jackson-datatype-jdk8-2.13.5.jar:/data/kafka/bin/../libs/jackson-jaxrs-base-2.13.5.jar:/data/kafka/bin/../libs/jackson-jaxrs-json-provider-2.13.5.jar:/data/kafka/bin/../libs/jackson-module-jaxb-annotations-2.13.5.jar:/data/kafka/bin/../libs/jackson-module-scala_2.12-2.13.5.jar:/data/kafka/bin/../libs/jakarta.activation-api-1.2.2.jar:/data/kafka/bin/../libs/jakarta.annotation-api-1.3.5.jar:/data/kafka/bin/../libs/jakarta.inject-2.6.1.jar:/data/kafka/bin/../libs/jakarta.validation-api-2.0.2.jar:/data/kafka/bin/../libs/jakarta.ws.rs-api-2.1.6.jar:/data/kafka/bin/../libs/jakarta.xml.bind-api-2.3.3.jar:/data/kafka/bin/../libs/javassist-3.29.2-GA.jar:/data/kafka/bin/../libs/javax.activation-api-1.2.0.jar:/data/kafka/bin/../libs/javax.annotation-api-1.3.2.jar:/data/kafka/bin/../libs/javax.servlet-api-3.1.0.jar:/data/kafka/bin/../libs/javax.ws.rs-api-2.1.1.jar:/data/kafka/bin/../libs/jaxb-api-2.3.1.jar:/data/kafka/bin/../libs/jersey-client-2.39.1.jar:/data/kafka/bin/../libs/jersey-common-2.39.1.jar:/data/kafka/bin/../libs/jersey-container-servlet-2.39.1.jar:/data/kafka/bin/../libs/jersey-container-servlet-core-2.39.1.jar:/data/kafka/bin/../libs/jersey-hk2-2.39.1.jar:/data/kafka/bin/../libs/jersey-server-2.39.1.jar:/data/kafka/bin/../libs/jetty-client-9.4.52.v20230823.jar:/data/kafka/bin/../libs/jetty-continuation-9.4.52.v20230823.jar:/data/kafka/bin/../libs/jetty-http-9.4.52.v20230823.jar:/data/kafka/bin/../libs/jetty-io-9.4.52.v20230823.jar:/data/kafka/bin/../libs/jetty-security-9.4.52.v20230823.jar:/data/kafka/bin/../libs/jetty-server-9.4.52.v20230823.jar:/data/kafka/bin/../libs/jetty-servlet-9.4.52.v20230823.jar:/data/kafka/bin/../libs/jetty-servlets-9.4.52.v20230823.jar:/data/kafka/bin/../libs/jetty-util-9.4.52.v20230823.jar:/data/kafka/bin/../libs/jetty-util-ajax-9.4.52.v20230823.jar:/data/kafka/bin/../libs/jline-3.22.0.jar:/data/kafka/bin/../libs/jopt-simple-5.0.4.jar:/data/kafka/bin/../libs/jose4j-0.9.3.jar:/data/kafka/bin/../libs/jsr305-3.0.2.jar:/data/kafka/bin/../libs/kafka_2.12-3.6.1.jar:/data/kafka/bin/../libs/kafka-clients-3.6.1.jar:/data/kafka/bin/../libs/kafka-group-coordinator-3.6.1.jar:/data/kafka/bin/../libs/kafka-log4j-appender-3.6.1.jar:/data/kafka/bin/../libs/kafka-metadata-3.6.1.jar:/data/kafka/bin/../libs/kafka-raft-3.6.1.jar:/data/kafka/bin/../libs/kafka-server-common-3.6.1.jar:/data/kafka/bin/../libs/kafka-shell-3.6.1.jar:/data/kafka/bin/../libs/kafka-storage-3.6.1.jar:/data/kafka/bin/../libs/kafka-storage-api-3.6.1.jar:/data/kafka/bin/../libs/kafka-streams-3.6.1.jar:/data/kafka/bin/../libs/kafka-streams-examples-3.6.1.jar:/data/kafka/bin/../libs/kafka-streams-scala_2.12-3.6.1.jar:/data/kafka/bin/../libs/kafka-streams-test-utils-3.6.1.jar:/data/kafka/bin/../libs/kafka-tools-3.6.1.jar:/data/kafka/bin/../libs/kafka-tools-api-3.6.1.jar:/data/kafka/bin/../libs/lz4-java-1.8.0.jar:/data/kafka/bin/../libs/maven-artifact-3.8.8.jar:/data/kafka/bin/../libs/metrics-core-2.2.0.jar:/data/kafka/bin/../libs/metrics-core-4.1.12.1.jar:/data/kafka/bin/../libs/netty-buffer-4.1.100.Final.jar:/data/kafka/bin/../libs/netty-codec-4.1.100.Final.jar:/data/kafka/bin/../libs/netty-common-4.1.100.Final.jar:/data/kafka/bin/../libs/netty-handler-4.1.100.Final.jar:/data/kafka/bin/../libs/netty-resolver-4.1.100.Final.jar:/data/kafka/bin/../libs/netty-transport-4.1.100.Final.jar:/data/kafka/bin/../libs/netty-transport-classes-epoll-4.1.100.Final.jar:/data/kafka/bin/../libs/netty-transport-native-epoll-4.1.100.Final.jar:/data/kafka/bin/../libs/netty-transport-native-unix-common-4.1.100.Final.jar:/data/kafka/bin/../libs/osgi-resource-locator-1.0.3.jar:/data/kafka/bin/../libs/paranamer-2.8.jar:/data/kafka/bin/../libs/pcollections-4.0.1.jar:/data/kafka/bin/../libs/plexus-utils-3.3.1.jar:/data/kafka/bin/../libs/reflections-0.10.2.jar:/data/kafka/bin/../libs/reload4j-1.2.25.jar:/data/kafka/bin/../libs/rocksdbjni-7.9.2.jar:/data/kafka/bin/../libs/scala-collection-compat_2.12-2.10.0.jar:/data/kafka/bin/../libs/scala-java8-compat_2.12-1.0.2.jar:/data/kafka/bin/../libs/scala-library-2.12.18.jar:/data/kafka/bin/../libs/scala-logging_2.12-3.9.4.jar:/data/kafka/bin/../libs/scala-reflect-2.12.18.jar:/data/kafka/bin/../libs/slf4j-api-1.7.36.jar:/data/kafka/bin/../libs/slf4j-reload4j-1.7.36.jar:/data/kafka/bin/../libs/snappy-java-1.1.10.5.jar:/data/kafka/bin/../libs/swagger-annotations-2.2.8.jar:/data/kafka/bin/../libs/trogdor-3.6.1.jar:/data/kafka/bin/../libs/zookeeper-3.8.3.jar:/data/kafka/bin/../libs/zookeeper-jute-3.8.3.jar:/data/kafka/bin/../libs/zstd-jni-1.5.5-1.jar kafka.Kafka /data/kafka/config/server.properties
root     25960 25497  0 17:32 ?        00:00:00 grep kafka
tcp6       0      0 :::9092                 :::*                    LISTEN      25880/java          

[Start Kafka Server][node2] Success
root      5288     1 76 17:32 ?        00:00:03 java -Xmx1G -Xms1G -server -XX:+UseG1GC -XX:MaxGCPauseMillis=20 -XX:InitiatingHeapOccupancyPercent=35 -XX:+ExplicitGCInvokesConcurrent -XX:MaxInlineLevel=15 -Djava.awt.headless=true -Xloggc:/data/kafka/bin/../logs/kafkaServer-gc.log -verbose:gc -XX:+PrintGCDetails -XX:+PrintGCDateStamps -XX:+PrintGCTimeStamps -XX:+UseGCLogFileRotation -XX:NumberOfGCLogFiles=10 -XX:GCLogFileSize=100M -Dcom.sun.management.jmxremote -Dcom.sun.management.jmxremote.authenticate=false -Dcom.sun.management.jmxremote.ssl=false -Dkafka.logs.dir=/data/kafka/bin/../logs -Dlog4j.configuration=file:/data/kafka/bin/../config/log4j.properties -cp /data/kafka/bin/../libs/activation-1.1.1.jar:/data/kafka/bin/../libs/aopalliance-repackaged-2.6.1.jar:/data/kafka/bin/../libs/argparse4j-0.7.0.jar:/data/kafka/bin/../libs/audience-annotations-0.12.0.jar:/data/kafka/bin/../libs/caffeine-2.9.3.jar:/data/kafka/bin/../libs/checker-qual-3.19.0.jar:/data/kafka/bin/../libs/commons-beanutils-1.9.4.jar:/data/kafka/bin/../libs/commons-cli-1.4.jar:/data/kafka/bin/../libs/commons-collections-3.2.2.jar:/data/kafka/bin/../libs/commons-digester-2.1.jar:/data/kafka/bin/../libs/commons-io-2.11.0.jar:/data/kafka/bin/../libs/commons-lang3-3.8.1.jar:/data/kafka/bin/../libs/commons-logging-1.2.jar:/data/kafka/bin/../libs/commons-validator-1.7.jar:/data/kafka/bin/../libs/connect-api-3.6.1.jar:/data/kafka/bin/../libs/connect-basic-auth-extension-3.6.1.jar:/data/kafka/bin/../libs/connect-json-3.6.1.jar:/data/kafka/bin/../libs/connect-mirror-3.6.1.jar:/data/kafka/bin/../libs/connect-mirror-client-3.6.1.jar:/data/kafka/bin/../libs/connect-runtime-3.6.1.jar:/data/kafka/bin/../libs/connect-transforms-3.6.1.jar:/data/kafka/bin/../libs/error_prone_annotations-2.10.0.jar:/data/kafka/bin/../libs/hk2-api-2.6.1.jar:/data/kafka/bin/../libs/hk2-locator-2.6.1.jar:/data/kafka/bin/../libs/hk2-utils-2.6.1.jar:/data/kafka/bin/../libs/jackson-annotations-2.13.5.jar:/data/kafka/bin/../libs/jackson-core-2.13.5.jar:/data/kafka/bin/../libs/jackson-databind-2.13.5.jar:/data/kafka/bin/../libs/jackson-dataformat-csv-2.13.5.jar:/data/kafka/bin/../libs/jackson-datatype-jdk8-2.13.5.jar:/data/kafka/bin/../libs/jackson-jaxrs-base-2.13.5.jar:/data/kafka/bin/../libs/jackson-jaxrs-json-provider-2.13.5.jar:/data/kafka/bin/../libs/jackson-module-jaxb-annotations-2.13.5.jar:/data/kafka/bin/../libs/jackson-module-scala_2.12-2.13.5.jar:/data/kafka/bin/../libs/jakarta.activation-api-1.2.2.jar:/data/kafka/bin/../libs/jakarta.annotation-api-1.3.5.jar:/data/kafka/bin/../libs/jakarta.inject-2.6.1.jar:/data/kafka/bin/../libs/jakarta.validation-api-2.0.2.jar:/data/kafka/bin/../libs/jakarta.ws.rs-api-2.1.6.jar:/data/kafka/bin/../libs/jakarta.xml.bind-api-2.3.3.jar:/data/kafka/bin/../libs/javassist-3.29.2-GA.jar:/data/kafka/bin/../libs/javax.activation-api-1.2.0.jar:/data/kafka/bin/../libs/javax.annotation-api-1.3.2.jar:/data/kafka/bin/../libs/javax.servlet-api-3.1.0.jar:/data/kafka/bin/../libs/javax.ws.rs-api-2.1.1.jar:/data/kafka/bin/../libs/jaxb-api-2.3.1.jar:/data/kafka/bin/../libs/jersey-client-2.39.1.jar:/data/kafka/bin/../libs/jersey-common-2.39.1.jar:/data/kafka/bin/../libs/jersey-container-servlet-2.39.1.jar:/data/kafka/bin/../libs/jersey-container-servlet-core-2.39.1.jar:/data/kafka/bin/../libs/jersey-hk2-2.39.1.jar:/data/kafka/bin/../libs/jersey-server-2.39.1.jar:/data/kafka/bin/../libs/jetty-client-9.4.52.v20230823.jar:/data/kafka/bin/../libs/jetty-continuation-9.4.52.v20230823.jar:/data/kafka/bin/../libs/jetty-http-9.4.52.v20230823.jar:/data/kafka/bin/../libs/jetty-io-9.4.52.v20230823.jar:/data/kafka/bin/../libs/jetty-security-9.4.52.v20230823.jar:/data/kafka/bin/../libs/jetty-server-9.4.52.v20230823.jar:/data/kafka/bin/../libs/jetty-servlet-9.4.52.v20230823.jar:/data/kafka/bin/../libs/jetty-servlets-9.4.52.v20230823.jar:/data/kafka/bin/../libs/jetty-util-9.4.52.v20230823.jar:/data/kafka/bin/../libs/jetty-util-ajax-9.4.52.v20230823.jar:/data/kafka/bin/../libs/jline-3.22.0.jar:/data/kafka/bin/../libs/jopt-simple-5.0.4.jar:/data/kafka/bin/../libs/jose4j-0.9.3.jar:/data/kafka/bin/../libs/jsr305-3.0.2.jar:/data/kafka/bin/../libs/kafka_2.12-3.6.1.jar:/data/kafka/bin/../libs/kafka-clients-3.6.1.jar:/data/kafka/bin/../libs/kafka-group-coordinator-3.6.1.jar:/data/kafka/bin/../libs/kafka-log4j-appender-3.6.1.jar:/data/kafka/bin/../libs/kafka-metadata-3.6.1.jar:/data/kafka/bin/../libs/kafka-raft-3.6.1.jar:/data/kafka/bin/../libs/kafka-server-common-3.6.1.jar:/data/kafka/bin/../libs/kafka-shell-3.6.1.jar:/data/kafka/bin/../libs/kafka-storage-3.6.1.jar:/data/kafka/bin/../libs/kafka-storage-api-3.6.1.jar:/data/kafka/bin/../libs/kafka-streams-3.6.1.jar:/data/kafka/bin/../libs/kafka-streams-examples-3.6.1.jar:/data/kafka/bin/../libs/kafka-streams-scala_2.12-3.6.1.jar:/data/kafka/bin/../libs/kafka-streams-test-utils-3.6.1.jar:/data/kafka/bin/../libs/kafka-tools-3.6.1.jar:/data/kafka/bin/../libs/kafka-tools-api-3.6.1.jar:/data/kafka/bin/../libs/lz4-java-1.8.0.jar:/data/kafka/bin/../libs/maven-artifact-3.8.8.jar:/data/kafka/bin/../libs/metrics-core-2.2.0.jar:/data/kafka/bin/../libs/metrics-core-4.1.12.1.jar:/data/kafka/bin/../libs/netty-buffer-4.1.100.Final.jar:/data/kafka/bin/../libs/netty-codec-4.1.100.Final.jar:/data/kafka/bin/../libs/netty-common-4.1.100.Final.jar:/data/kafka/bin/../libs/netty-handler-4.1.100.Final.jar:/data/kafka/bin/../libs/netty-resolver-4.1.100.Final.jar:/data/kafka/bin/../libs/netty-transport-4.1.100.Final.jar:/data/kafka/bin/../libs/netty-transport-classes-epoll-4.1.100.Final.jar:/data/kafka/bin/../libs/netty-transport-native-epoll-4.1.100.Final.jar:/data/kafka/bin/../libs/netty-transport-native-unix-common-4.1.100.Final.jar:/data/kafka/bin/../libs/osgi-resource-locator-1.0.3.jar:/data/kafka/bin/../libs/paranamer-2.8.jar:/data/kafka/bin/../libs/pcollections-4.0.1.jar:/data/kafka/bin/../libs/plexus-utils-3.3.1.jar:/data/kafka/bin/../libs/reflections-0.10.2.jar:/data/kafka/bin/../libs/reload4j-1.2.25.jar:/data/kafka/bin/../libs/rocksdbjni-7.9.2.jar:/data/kafka/bin/../libs/scala-collection-compat_2.12-2.10.0.jar:/data/kafka/bin/../libs/scala-java8-compat_2.12-1.0.2.jar:/data/kafka/bin/../libs/scala-library-2.12.18.jar:/data/kafka/bin/../libs/scala-logging_2.12-3.9.4.jar:/data/kafka/bin/../libs/scala-reflect-2.12.18.jar:/data/kafka/bin/../libs/slf4j-api-1.7.36.jar:/data/kafka/bin/../libs/slf4j-reload4j-1.7.36.jar:/data/kafka/bin/../libs/snappy-java-1.1.10.5.jar:/data/kafka/bin/../libs/swagger-annotations-2.2.8.jar:/data/kafka/bin/../libs/trogdor-3.6.1.jar:/data/kafka/bin/../libs/zookeeper-3.8.3.jar:/data/kafka/bin/../libs/zookeeper-jute-3.8.3.jar:/data/kafka/bin/../libs/zstd-jni-1.5.5-1.jar kafka.Kafka /data/kafka/config/server.properties
root      5364  4905  0 17:32 ?        00:00:00 grep kafka
tcp6       0      0 :::9092                 :::*                    LISTEN      5288/java           

=== Completed: Start Kafka Server ===

[root@selfhosted-0001 petal]#
```

# jdk8

```bash
[root@selfhosted-0001 petal]# ./petal-linux-amd64 -file task/mw/jdk8/install.yml 
=== Task: Install JDK 1.8 ===
[Install JDK 1.8][node2] Running...
yum install -y java-1.8.0-openjdk
yum install -y java-1.8.0-openjdk-devel

[Install JDK 1.8][node1] Running...
yum install -y java-1.8.0-openjdk
yum install -y java-1.8.0-openjdk-devel

[Install JDK 1.8][node3] Running...
yum install -y java-1.8.0-openjdk
yum install -y java-1.8.0-openjdk-devel

[Install JDK 1.8][node2] Success
Loaded plugins: fastestmirror
Loading mirror speeds from cached hostfile
 * base: mirrors.huaweicloud.com
 * extras: mirrors.huaweicloud.com
 * updates: mirrors.huaweicloud.com
Resolving Dependencies
--> Running transaction check
---> Package java-1.8.0-openjdk.x86_64 1:1.8.0.392.b08-2.el7_9 will be installed
--> Processing Dependency: java-1.8.0-openjdk-headless(x86-64) = 1:1.8.0.392.b08-2.el7_9 for package: 1:java-1.8.0-openjdk-1.8.0.392.b08-2.el7_9.x86_64
--> Running transaction check
---> Package java-1.8.0-openjdk-headless.x86_64 1:1.8.0.392.b08-2.el7_9 will be installed
--> Finished Dependency Resolution

Dependencies Resolved

================================================================================
 Package                          Arch   Version                    Repo   Size
================================================================================
Installing:
 java-1.8.0-openjdk               x86_64  1:1.8.0.392.b08-2.el7_9  updates 318 k
Installing for dependencies:
 java-1.8.0-openjdk-headless      x86_64  1:1.8.0.392.b08-2.el7_9  updates  32 M

Transaction Summary
================================================================================
Install  1 Package (+1 Dependent package)

Total download size: 32 M
Installed size: 110 M
Downloading packages:
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  Installing : 1:java-1.8.0-openjdk-headless-1.8.0.392.b08-2.el7_9.x86_64  1/2
  Installing : 1:java-1.8.0-openjdk-1.8.0.392.b08-2.el7_9.x86_64           2/2
  Verifying  : 1:java-1.8.0-openjdk-headless-1.8.0.392.b08-2.el7_9.x86_64  1/2
  Verifying  : 1:java-1.8.0-openjdk-1.8.0.392.b08-2.el7_9.x86_64           2/2

Installed:
  java-1.8.0-openjdk.x86_64 1:1.8.0.392.b08-2.el7_9

Dependency Installed:
  java-1.8.0-openjdk-headless.x86_64 1:1.8.0.392.b08-2.el7_9

Complete!
Loaded plugins: fastestmirror
Loading mirror speeds from cached hostfile
 * base: mirrors.huaweicloud.com
 * extras: mirrors.huaweicloud.com
 * updates: mirrors.huaweicloud.com
Resolving Dependencies
--> Running transaction check
---> Package java-1.8.0-openjdk-devel.x86_64 1:1.8.0.392.b08-2.el7_9 will be installed
--> Finished Dependency Resolution

Dependencies Resolved

================================================================================
 Package                      Arch    Version                    Repo      Size
================================================================================
Installing:
 java-1.8.0-openjdk-devel     x86_64  1:1.8.0.392.b08-2.el7_9   updates  9.8 M

Transaction Summary
================================================================================
Install  1 Package

Total download size: 9.8 M
Installed size: 39 M
Downloading packages:
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  Installing : 1:java-1.8.0-openjdk-devel-1.8.0.392.b08-2.el7_9.x86_64    1/1
  Verifying  : 1:java-1.8.0-openjdk-devel-1.8.0.392.b08-2.el7_9.x86_64    1/1

Installed:
  java-1.8.0-openjdk-devel.x86_64 1:1.8.0.392.b08-2.el7_9

Complete!

[Install JDK 1.8][node1] Success
Loaded plugins: fastestmirror
Loading mirror speeds from cached hostfile
 * base: mirrors.huaweicloud.com
 * extras: mirrors.huaweicloud.com
 * updates: mirrors.huaweicloud.com
Package 1:java-1.8.0-openjdk-1.8.0.392.b08-2.el7_9.x86_64 already installed and latest version
Nothing to do
Package 1:java-1.8.0-openjdk-devel-1.8.0.392.b08-2.el7_9.x86_64 already installed and latest version
Nothing to do

[Install JDK 1.8][node3] Success
Loaded plugins: fastestmirror
Loading mirror speeds from cached hostfile
 * base: mirrors.huaweicloud.com
 * extras: mirrors.huaweicloud.com
 * updates: mirrors.huaweicloud.com
Package 1:java-1.8.0-openjdk-1.8.0.392.b08-2.el7_9.x86_64 already installed and latest version
Nothing to do
Package 1:java-1.8.0-openjdk-devel-1.8.0.392.b08-2.el7_9.x86_64 already installed and latest version
Nothing to do

=== Completed: Install JDK 1.8 ===

=== Task: Check Java Version after Installation ===
[Check Java Version after Installation][node3] Running...
java -version
rpm -qa |grep java-1.8.0-openjdk
rpm -qa |grep java-1.8.0-openjdk-devel

[Check Java Version after Installation][node1] Running...
java -version
rpm -qa |grep java-1.8.0-openjdk
rpm -qa |grep java-1.8.0-openjdk-devel

[Check Java Version after Installation][node2] Running...
java -version
rpm -qa |grep java-1.8.0-openjdk
rpm -qa |grep java-1.8.0-openjdk-devel

[Check Java Version after Installation][node3] Success
openjdk version "1.8.0_392"
OpenJDK Runtime Environment (build 1.8.0_392-b08)
OpenJDK 64-Bit Server VM (build 25.392-b08, mixed mode)
java-1.8.0-openjdk-1.8.0.392.b08-2.el7_9.x86_64
java-1.8.0-openjdk-headless-1.8.0.392.b08-2.el7_9.x86_64
java-1.8.0-openjdk-devel-1.8.0.392.b08-2.el7_9.x86_64

[Check Java Version after Installation][node1] Success
openjdk version "1.8.0_392"
OpenJDK Runtime Environment (build 1.8.0_392-b08)
OpenJDK 64-Bit Server VM (build 25.392-b08, mixed mode)
java-1.8.0-openjdk-1.8.0.392.b08-2.el7_9.x86_64
java-1.8.0-openjdk-headless-1.8.0.392.b08-2.el7_9.x86_64
java-1.8.0-openjdk-devel-1.8.0.392.b08-2.el7_9.x86_64

[Check Java Version after Installation][node2] Success
openjdk version "1.8.0_392"
OpenJDK Runtime Environment (build 1.8.0_392-b08)
OpenJDK 64-Bit Server VM (build 25.392-b08, mixed mode)
java-1.8.0-openjdk-1.8.0.392.b08-2.el7_9.x86_64
java-1.8.0-openjdk-headless-1.8.0.392.b08-2.el7_9.x86_64
java-1.8.0-openjdk-devel-1.8.0.392.b08-2.el7_9.x86_64

=== Completed: Check Java Version after Installation ===
```

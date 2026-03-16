# obsutil

```bash
[root@selfhosted-0001 petal]# ./petal-linux-amd64 -file task/base/obsutil/install.yml 
=== Task: Download OBS Util Package ===
[Download OBS Util Package][node3] Running...
wget https://obs-community-intl.obs.ap-southeast-1.myhuaweicloud.com/obsutil/current/obsutil_linux_amd64.tar.gz

[Download OBS Util Package][node1] Running...
wget https://obs-community-intl.obs.ap-southeast-1.myhuaweicloud.com/obsutil/current/obsutil_linux_amd64.tar.gz

[Download OBS Util Package][node2] Running...
wget https://obs-community-intl.obs.ap-southeast-1.myhuaweicloud.com/obsutil/current/obsutil_linux_amd64.tar.gz

[Download OBS Util Package][node3] Success
--2025-07-01 08:30:12--  https://obs-community-intl.obs.ap-southeast-1.myhuaweicloud.com/obsutil/current/obsutil_linux_amd64.tar.gz
Resolving obs-community-intl.obs.ap-southeast-1.myhuaweicloud.com... 49.4.88.212
Connecting to obs-community-intl.obs.ap-southeast-1.myhuaweicloud.com|49.4.88.212|:443... connected.
HTTP request sent, awaiting response... 200 OK
Length: 11223628 (10.7M) [application/gzip]
Saving to: 'obsutil_linux_amd64.tar.gz'

     0K .......... .......... .......... .......... ..........  0%  9.62M 1s
   ...
10922K .......... ....                                        100%  89.4M=0.3s

2025-07-01 08:30:12 (38.2 MB/s) - 'obsutil_linux_amd64.tar.gz' saved [11223628/11223628]

[Download OBS Util Package][node1] Success
--2025-07-01 08:30:12--  https://obs-community-intl.obs.ap-southeast-1.myhuaweicloud.com/obsutil/current/obsutil_linux_amd64.tar.gz
Resolving obs-community-intl.obs.ap-southeast-1.myhuaweicloud.com... 49.4.88.212
Connecting to obs-community-intl.obs.ap-southeast-1.myhuaweicloud.com|49.4.88.212|:443... connected.
HTTP request sent, awaiting response... 200 OK
Length: 11223628 (10.7M) [application/gzip]
Saving to: 'obsutil_linux_amd64.tar.gz'

     0K .......... .......... .......... .......... ..........  0%  8.45M 1s
   ...
10922K .......... ....                                        100%  91.7M=0.3s

2025-07-01 08:30:12 (37.1 MB/s) - 'obsutil_linux_amd64.tar.gz' saved [11223628/11223628]

[Download OBS Util Package][node2] Success
--2025-07-01 08:30:12--  https://obs-community-intl.obs.ap-southeast-1.myhuaweicloud.com/obsutil/current/obsutil_linux_amd64.tar.gz
Resolving obs-community-intl.obs.ap-southeast-1.myhuaweicloud.com... 49.4.88.212
Connecting to obs-community-intl.obs.ap-southeast-1.myhuaweicloud.com|49.4.88.212|:443... connected.
HTTP request sent, awaiting response... 200 OK
Length: 11223628 (10.7M) [application/gzip]
Saving to: 'obsutil_linux_amd64.tar.gz'

     0K .......... .......... .......... .......... ..........  0%  7.83M 1s
   ...
10922K .......... ....                                        100%  85.2M=0.3s

2025-07-01 08:30:13 (35.8 MB/s) - 'obsutil_linux_amd64.tar.gz' saved [11223628/11223628]

=== Completed: Download OBS Util Package ===

=== Task: Extract OBS Util Package ===
[Extract OBS Util Package][node3] Running...
tar zxf obsutil_linux_amd64.tar.gz
mv obsutil_linux_amd64*/obsutil /usr/bin/

[Extract OBS Util Package][node1] Running...
tar zxf obsutil_linux_amd64.tar.gz
mv obsutil_linux_amd64*/obsutil /usr/bin/

[Extract OBS Util Package][node2] Running...
tar zxf obsutil_linux_amd64.tar.gz
mv obsutil_linux_amd64*/obsutil /usr/bin/

[Extract OBS Util Package][node3] Success

[Extract OBS Util Package][node1] Success

[Extract OBS Util Package][node2] Success

=== Completed: Extract OBS Util Package ===

=== Task: Check OBS Util Cli ===
[Check OBS Util Cli][node3] Running...
which obsutil

[Check OBS Util Cli][node1] Running...
which obsutil

[Check OBS Util Cli][node2] Running...
which obsutil

[Check OBS Util Cli][node3] Success
/usr/bin/obsutil

[Check OBS Util Cli][node1] Success
/usr/bin/obsutil

[Check OBS Util Cli][node2] Success
/usr/bin/obsutil

=== Completed: Check OBS Util Cli ===
```

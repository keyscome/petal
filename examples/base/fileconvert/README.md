# fileconvert

## Overview

This example installs a full document and media conversion stack on application nodes. It combines Apache OpenOffice 4.1.7, LibreOffice 6.4.4.2 (with the Chinese language pack and SDK), Google Chrome 119, ChromeDriver 119, and Chinese system fonts — all of which are required by a file-conversion microservice (`paas-file-convert-server`) that converts uploaded documents (Word, Excel, PDF) and renders web pages to images. A scheduled health-check script is set up via cron to keep the service running.

## Architecture

- **Topology**: 2 nodes in parallel (`node1`, `node2`)
- **Package source**: All packages are pre-bundled in `fileconvert.tgz` on OBS (no internet access required at install time)
- **Install paths**:
  - Apache OpenOffice: `/opt/openoffice4/`
  - LibreOffice: `/opt/libreoffice6.4/`
  - Google Chrome: `/usr/bin/google-chrome-stable`
  - ChromeDriver: `/usr/bin/chromedriver`
  - Fonts: `/usr/share/fonts/WenQuanYi/`, `/usr/share/fonts/SimSun/`
- **App directory**: `/data/hqcrm/paas-file-convert-server/`
- **Cron job**: Commented-out entry (`*/7 * * * *`) for `check.sh` to monitor the conversion service
- **System locale**: Set to `zh_CN.UTF-8` (required for rendering Chinese characters)
- **Dependencies**: `vulkan`, `liberation-fonts`, `xorg-x11-fonts-75dpi`, `wkhtmltox`

## Files

| File | Description |
|------|-------------|
| `install.yml` | Downloads and installs the entire document conversion stack, configures locale, and sets up the health-check cron job |

### `install.yml`

1. **Download Fileconvet Base Environment Package** — Fetches `fileconvert.tgz` from OBS object storage onto both nodes using `obsutil cp`.
2. **Extract Fileconvet Base Environment Package** — Extracts the tarball, producing a `fileconvert/` directory containing all offline packages.
3. **Install Apache_OpenOffice_4.1.7_Linux_x86-64_install-rpm_zh-CN** — Extracts the OpenOffice archive and installs the core RPM first (`openoffice-4.1.7-9800.x86_64.rpm`), then all component RPMs (base, calc, draw, impress, writer, math) via glob. Removes the extracted RPM directory after installation.
4. **Install LibreOffice_6.4.4.2_Linux_x86-64_rpm** — Extracts and installs all LibreOffice 6.4.4.2 component RPMs from the main package archive, then removes the RPMs directory.
5. **Install LibreOffice_6.4.4.2_Linux_x86-64_rpm_langpack_zh-CN** — Extracts and installs the Simplified Chinese language pack for LibreOffice 6.4.4.2.
6. **Install LibreOffice_6.4.4.2_Linux_x86-64_rpm_sdk** — Extracts and installs the LibreOffice 6.4.4.2 SDK (development headers and UNO bridge), required by the conversion microservice for programmatic document operations.
7. **Install google-chrome-stable-119.0.6045.105-1.x86_64** — Installs prerequisite packages (`xorg-x11-fonts-75dpi`, `wkhtmltox`, `vulkan`, `liberation-fonts`) before installing the pinned Chrome RPM. Uses both `rpm -ivh` and `yum localinstall` to resolve any remaining dependencies.
8. **Install chromedriver-linux64-v119** — Extracts `chromedriver-linux64-v119.zip` and copies the `chromedriver` binary to `/usr/bin/`, making it available for Selenium-based web-rendering tasks.
9. **Install linux-usr-share-fonts** — Extracts Chinese font files from `linux-usr-share-fonts.zip` into `/usr/share/fonts/`, then rebuilds the font cache with `mkfontscale`, `mkfontdir`, and `fc-cache`. Verifies installed fonts with `fc-list`.
10. **Configure Some System ENV** — Appends `export LANG="zh_CN.UTF-8"` to `/etc/profile` and writes a full locale configuration block to `/etc/sysconfig/i18n`, enabling Chinese character rendering in the conversion service.
11. **Configuration Check Script** — Creates the app home directory, copies `check.sh` to `/data/hqcrm/paas-file-convert-server/`, makes it executable, and adds a commented-out cron entry (`# */7 * * * *`) to `crontab` for periodic health monitoring. The entry is commented out to allow operators to enable it manually after verifying the service.

## Example

```bash
[root@selfhosted-0001 petal]# ./petal-linux-amd64 -file task/base/fileconvert/install.yml 
=== Task: Download Fileconvet Base Environment Package ===
[Download Fileconvet Base Environment Package][node2] Running...
obsutil cp obs://selfhosted/pkg/fileconvert/fileconvert.tgz .

[Download Fileconvet Base Environment Package][node1] Running...
obsutil cp obs://selfhosted/pkg/fileconvert/fileconvert.tgz .

[Download Fileconvet Base Environment Package][node2] Success
Start at 2025-07-01 09:00:05.123456789 +0000 UTC


Parallel:      5                   Jobs:          5                   
Threshold:     50.00MB             PartSize:      auto                
VerifyLength:  false               VerifyMd5:     false               
CheckpointDir: /root/.obsutil_checkpoint     

Task id: a1b2c3d4-e5f6-7890-abcd-ef1234567890
OutputDir: /root/.obsutil_output         
TempFileDir: /root                         

Succeed count:      1         Failed count:       0         
Succeed bytes:      512.34MB  
Metrics [max cost:3241 ms, min cost:3241 ms, average cost:1620.50 ms, average tps:0.15, transferred size:512.34MB]
[------------------------------------------------------] 100.00% tps:0.00 ?/s 1/1 512.34MB/512.34MB 3.241s
Task id: a1b2c3d4-e5f6-7890-abcd-ef1234567890

[Download Fileconvet Base Environment Package][node1] Success
Start at 2025-07-01 09:00:05.120000000 +0000 UTC


Parallel:      5                   Jobs:          5                   
Threshold:     50.00MB             PartSize:      auto                
VerifyLength:  false               VerifyMd5:     false               
CheckpointDir: /root/.obsutil_checkpoint     

Task id: b2c3d4e5-f6a7-8901-bcde-f12345678901
OutputDir: /root/.obsutil_output         
TempFileDir: /root                         

Succeed count:      1         Failed count:       0         
Succeed bytes:      512.34MB  
Metrics [max cost:3358 ms, min cost:3358 ms, average cost:1679.00 ms, average tps:0.14, transferred size:512.34MB]
[------------------------------------------------------] 100.00% tps:0.00 ?/s 1/1 512.34MB/512.34MB 3.358s
Task id: b2c3d4e5-f6a7-8901-bcde-f12345678901

=== Completed: Download Fileconvet Base Environment Package ===

=== Task: Extract Fileconvet Base Environment Package ===
[Extract Fileconvet Base Environment Package][node2] Running...
tar zxf fileconvert.tgz

[Extract Fileconvet Base Environment Package][node1] Running...
tar zxf fileconvert.tgz

[Extract Fileconvet Base Environment Package][node2] Success

[Extract Fileconvet Base Environment Package][node1] Success

=== Completed: Extract Fileconvet Base Environment Package ===

=== Task: Install Apache_OpenOffice_4.1.7_Linux_x86-64_install-rpm_zh-CN ===
[Install Apache_OpenOffice_4.1.7_Linux_x86-64_install-rpm_zh-CN][node2] Running...
tar zxf fileconvert/Apache_OpenOffice_4.1.7_Linux_x86-64_install-rpm_zh-CN.tar.gz
rpm -ivh zh-CN/RPMS/openoffice-4.1.7-9800.x86_64.rpm
rpm -ivh zh-CN/RPMS/*.rpm
rm -fr zh-CN/RPM

[Install Apache_OpenOffice_4.1.7_Linux_x86-64_install-rpm_zh-CN][node1] Running...
tar zxf fileconvert/Apache_OpenOffice_4.1.7_Linux_x86-64_install-rpm_zh-CN.tar.gz
rpm -ivh zh-CN/RPMS/openoffice-4.1.7-9800.x86_64.rpm
rpm -ivh zh-CN/RPMS/*.rpm
rm -fr zh-CN/RPM

[Install Apache_OpenOffice_4.1.7_Linux_x86-64_install-rpm_zh-CN][node2] Success
Preparing...                          ################################# [100%]
Updating / installing...
   1:openoffice-4.1.7-9800            ################################# [100%]
Preparing...                          ################################# [100%]
Updating / installing...
   1:openoffice4-4.1.7-9800           ################################# [100%]
   2:openoffice4-base-4.1.7-9800      ################################# [100%]
   3:openoffice4-calc-4.1.7-9800      ################################# [100%]
   4:openoffice4-draw-4.1.7-9800      ################################# [100%]
   5:openoffice4-impress-4.1.7-9800   ################################# [100%]
   6:openoffice4-writer-4.1.7-9800    ################################# [100%]
   7:openoffice4-math-4.1.7-9800      ################################# [100%]

[Install Apache_OpenOffice_4.1.7_Linux_x86-64_install-rpm_zh-CN][node1] Success
Preparing...                          ################################# [100%]
Updating / installing...
   1:openoffice-4.1.7-9800            ################################# [100%]
Preparing...                          ################################# [100%]
Updating / installing...
   1:openoffice4-4.1.7-9800           ################################# [100%]
   2:openoffice4-base-4.1.7-9800      ################################# [100%]
   3:openoffice4-calc-4.1.7-9800      ################################# [100%]
   4:openoffice4-draw-4.1.7-9800      ################################# [100%]
   5:openoffice4-impress-4.1.7-9800   ################################# [100%]
   6:openoffice4-writer-4.1.7-9800    ################################# [100%]
   7:openoffice4-math-4.1.7-9800      ################################# [100%]

=== Completed: Install Apache_OpenOffice_4.1.7_Linux_x86-64_install-rpm_zh-CN ===

=== Task: Install LibreOffice_6.4.4.2_Linux_x86-64_rpm ===
[Install LibreOffice_6.4.4.2_Linux_x86-64_rpm][node2] Running...
tar zxf fileconvert/LibreOffice_6.4.4.2_Linux_x86-64_rpm.tar.gz
rpm -ivh LibreOffice_6.4.4.2_Linux_x86-64_rpm/RPMS/*.rpm
rm -fr LibreOffice_6.4.4.2_Linux_x86-64_rpm/RPMS

[Install LibreOffice_6.4.4.2_Linux_x86-64_rpm][node1] Running...
tar zxf fileconvert/LibreOffice_6.4.4.2_Linux_x86-64_rpm.tar.gz
rpm -ivh LibreOffice_6.4.4.2_Linux_x86-64_rpm/RPMS/*.rpm
rm -fr LibreOffice_6.4.4.2_Linux_x86-64_rpm/RPMS

[Install LibreOffice_6.4.4.2_Linux_x86-64_rpm][node2] Success
Preparing...                          ################################# [100%]
Updating / installing...
   1:libreoffice6.4-base-6.4.4.2-2    ################################# [  3%]
   2:libreoffice6.4-calc-6.4.4.2-2    ################################# [  6%]
   3:libreoffice6.4-draw-6.4.4.2-2    ################################# [ 10%]
   4:libreoffice6.4-impress-6.4.4.2-2 ################################# [ 13%]
   5:libreoffice6.4-writer-6.4.4.2-2  ################################# [ 16%]
   6:libreoffice6.4-math-6.4.4.2-2    ################################# [ 19%]
   7:libreoffice6.4-6.4.4.2-2         ################################# [100%]

[Install LibreOffice_6.4.4.2_Linux_x86-64_rpm][node1] Success
Preparing...                          ################################# [100%]
Updating / installing...
   1:libreoffice6.4-base-6.4.4.2-2    ################################# [  3%]
   2:libreoffice6.4-calc-6.4.4.2-2    ################################# [  6%]
   3:libreoffice6.4-draw-6.4.4.2-2    ################################# [ 10%]
   4:libreoffice6.4-impress-6.4.4.2-2 ################################# [ 13%]
   5:libreoffice6.4-writer-6.4.4.2-2  ################################# [ 16%]
   6:libreoffice6.4-math-6.4.4.2-2    ################################# [ 19%]
   7:libreoffice6.4-6.4.4.2-2         ################################# [100%]

=== Completed: Install LibreOffice_6.4.4.2_Linux_x86-64_rpm ===

=== Task: Install LibreOffice_6.4.4.2_Linux_x86-64_rpm_langpack_zh-CN ===
[Install LibreOffice_6.4.4.2_Linux_x86-64_rpm_langpack_zh-CN][node2] Running...
tar zxf fileconvert/LibreOffice_6.4.4.2_Linux_x86-64_rpm_langpack_zh-CN.tar.gz
rpm -ivh LibreOffice_6.4.4.2_Linux_x86-64_rpm_langpack_zh-CN/RPMS/*.rpm
rm -fr LibreOffice_6.4.4.2_Linux_x86-64_rpm_langpack_zh-CN/RPMS

[Install LibreOffice_6.4.4.2_Linux_x86-64_rpm_langpack_zh-CN][node1] Running...
tar zxf fileconvert/LibreOffice_6.4.4.2_Linux_x86-64_rpm_langpack_zh-CN.tar.gz
rpm -ivh LibreOffice_6.4.4.2_Linux_x86-64_rpm_langpack_zh-CN/RPMS/*.rpm
rm -fr LibreOffice_6.4.4.2_Linux_x86-64_rpm_langpack_zh-CN/RPMS

[Install LibreOffice_6.4.4.2_Linux_x86-64_rpm_langpack_zh-CN][node2] Success
Preparing...                          ################################# [100%]
Updating / installing...
   1:libreoffice6.4-langpack-zh_CN-6.4.4.2-2  ################################# [100%]

[Install LibreOffice_6.4.4.2_Linux_x86-64_rpm_langpack_zh-CN][node1] Success
Preparing...                          ################################# [100%]
Updating / installing...
   1:libreoffice6.4-langpack-zh_CN-6.4.4.2-2  ################################# [100%]

=== Completed: Install LibreOffice_6.4.4.2_Linux_x86-64_rpm_langpack_zh-CN ===

=== Task: Install LibreOffice_6.4.4.2_Linux_x86-64_rpm_sdk ===
[Install LibreOffice_6.4.4.2_Linux_x86-64_rpm_sdk][node2] Running...
tar zxf fileconvert/LibreOffice_6.4.4.2_Linux_x86-64_rpm_sdk.tar.gz
rpm -ivh LibreOffice_6.4.4.2_Linux_x86-64_rpm_sdk/RPMS/*.rpm
rm -fr LibreOffice_6.4.4.2_Linux_x86-64_rpm_sdk/RPMS/

[Install LibreOffice_6.4.4.2_Linux_x86-64_rpm_sdk][node1] Running...
tar zxf fileconvert/LibreOffice_6.4.4.2_Linux_x86-64_rpm_sdk.tar.gz
rpm -ivh LibreOffice_6.4.4.2_Linux_x86-64_rpm_sdk/RPMS/*.rpm
rm -fr LibreOffice_6.4.4.2_Linux_x86-64_rpm_sdk/RPMS/

[Install LibreOffice_6.4.4.2_Linux_x86-64_rpm_sdk][node2] Success
Preparing...                          ################################# [100%]
Updating / installing...
   1:libreoffice6.4-sdk-6.4.4.2-2     ################################# [100%]

[Install LibreOffice_6.4.4.2_Linux_x86-64_rpm_sdk][node1] Success
Preparing...                          ################################# [100%]
Updating / installing...
   1:libreoffice6.4-sdk-6.4.4.2-2     ################################# [100%]

=== Completed: Install LibreOffice_6.4.4.2_Linux_x86-64_rpm_sdk ===

=== Task: Install google-chrome-stable-119.0.6045.105-1.x86_64 ===
[Install google-chrome-stable-119.0.6045.105-1.x86_64][node2] Running...
rpm -ivh fileconvert/xorg-x11-fonts-75dpi-7.5-9.el7.noarch.rpm
rpm -ivh fileconvert/wkhtmltox-0.12.6-1.centos7.x86_64.rpm
yum install vulkan -y
yum install liberation-fonts -y
rpm -ivh fileconvert/google-chrome-stable-119.0.6045.105-1.x86_64.rpm
yum localinstall fileconvert/google-chrome-stable-119.0.6045.105-1.x86_64.rpm -y

[Install google-chrome-stable-119.0.6045.105-1.x86_64][node1] Running...
rpm -ivh fileconvert/xorg-x11-fonts-75dpi-7.5-9.el7.noarch.rpm
rpm -ivh fileconvert/wkhtmltox-0.12.6-1.centos7.x86_64.rpm
yum install vulkan -y
yum install liberation-fonts -y
rpm -ivh fileconvert/google-chrome-stable-119.0.6045.105-1.x86_64.rpm
yum localinstall fileconvert/google-chrome-stable-119.0.6045.105-1.x86_64.rpm -y

[Install google-chrome-stable-119.0.6045.105-1.x86_64][node2] Success
Preparing...                          ################################# [100%]
Updating / installing...
   1:xorg-x11-fonts-75dpi-7.5-9.el7   ################################# [100%]
Preparing...                          ################################# [100%]
Updating / installing...
   1:wkhtmltox-1:0.12.6-1.centos7     ################################# [100%]
Loaded plugins: fastestmirror
Loading mirror speeds from cached hostfile
Package vulkan-1.1.97.0-1.el7.x86_64 already installed and latest version
Nothing to do
Loaded plugins: fastestmirror
Loading mirror speeds from cached hostfile
Resolving Dependencies
--> Running transaction check
---> Package liberation-fonts.noarch 1:1.07.2-16.el7 will be installed
--> Finished Dependency Resolution
...
Complete!
Preparing...                          ################################# [100%]
Updating / installing...
   1:google-chrome-stable-119.0.6045  ################################# [100%]
Loaded plugins: fastestmirror
Examining fileconvert/google-chrome-stable-119.0.6045.105-1.x86_64.rpm: google-chrome-stable-119.0.6045.105-1.x86_64
Marking fileconvert/google-chrome-stable-119.0.6045.105-1.x86_64.rpm as an update to google-chrome-stable-119.0.6045.105-1.x86_64
Nothing to do

[Install google-chrome-stable-119.0.6045.105-1.x86_64][node1] Success
Preparing...                          ################################# [100%]
Updating / installing...
   1:xorg-x11-fonts-75dpi-7.5-9.el7   ################################# [100%]
Preparing...                          ################################# [100%]
Updating / installing...
   1:wkhtmltox-1:0.12.6-1.centos7     ################################# [100%]
Loaded plugins: fastestmirror
Loading mirror speeds from cached hostfile
Package vulkan-1.1.97.0-1.el7.x86_64 already installed and latest version
Nothing to do
Loaded plugins: fastestmirror
Loading mirror speeds from cached hostfile
Resolving Dependencies
--> Running transaction check
---> Package liberation-fonts.noarch 1:1.07.2-16.el7 will be installed
--> Finished Dependency Resolution
...
Complete!
Preparing...                          ################################# [100%]
Updating / installing...
   1:google-chrome-stable-119.0.6045  ################################# [100%]
Loaded plugins: fastestmirror
Examining fileconvert/google-chrome-stable-119.0.6045.105-1.x86_64.rpm: google-chrome-stable-119.0.6045.105-1.x86_64
Marking fileconvert/google-chrome-stable-119.0.6045.105-1.x86_64.rpm as an update to google-chrome-stable-119.0.6045.105-1.x86_64
Nothing to do

=== Completed: Install google-chrome-stable-119.0.6045.105-1.x86_64 ===

=== Task: Install chromedriver-linux64-v119 ===
[Install chromedriver-linux64-v119][node2] Running...
unzip fileconvert/chromedriver-linux64-v119.zip
cp chromedriver-linux64/chromedriver  /usr/bin/

[Install chromedriver-linux64-v119][node1] Running...
unzip fileconvert/chromedriver-linux64-v119.zip
cp chromedriver-linux64/chromedriver  /usr/bin/

[Install chromedriver-linux64-v119][node2] Success
Archive:  fileconvert/chromedriver-linux64-v119.zip
  inflating: chromedriver-linux64/chromedriver

[Install chromedriver-linux64-v119][node1] Success
Archive:  fileconvert/chromedriver-linux64-v119.zip
  inflating: chromedriver-linux64/chromedriver

=== Completed: Install chromedriver-linux64-v119 ===

=== Task: Install linux-usr-share-fonts ===
[Install linux-usr-share-fonts][node2] Running...
unzip fileconvert/linux-usr-share-fonts.zip -D /usr/share/fonts
mkfontscale
mkfontdir
fc-cache
fc-list

[Install linux-usr-share-fonts][node1] Running...
unzip fileconvert/linux-usr-share-fonts.zip -D /usr/share/fonts
mkfontscale
mkfontdir
fc-cache
fc-list

[Install linux-usr-share-fonts][node2] Success
Archive:  fileconvert/linux-usr-share-fonts.zip
  inflating: /usr/share/fonts/WenQuanYi/WenQuanYiMicroHei.ttf
  inflating: /usr/share/fonts/WenQuanYi/WenQuanYiMicroHeiMono.ttf
  inflating: /usr/share/fonts/WenQuanYi/WenQuanYiZenHei.ttf
  inflating: /usr/share/fonts/SimSun/SimSun.ttf
  inflating: /usr/share/fonts/SimSun/SimHei.ttf
/usr/share/fonts/WenQuanYi: WenQuanYi Micro Hei:style=Regular
/usr/share/fonts/WenQuanYi: WenQuanYi Micro Hei Mono:style=Regular
/usr/share/fonts/WenQuanYi: WenQuanYi Zen Hei:style=Regular
/usr/share/fonts/SimSun: AR PL UMing CN:style=Light
/usr/share/fonts/SimSun: AR PL UKai CN:style=Book

[Install linux-usr-share-fonts][node1] Success
Archive:  fileconvert/linux-usr-share-fonts.zip
  inflating: /usr/share/fonts/WenQuanYi/WenQuanYiMicroHei.ttf
  inflating: /usr/share/fonts/WenQuanYi/WenQuanYiMicroHeiMono.ttf
  inflating: /usr/share/fonts/WenQuanYi/WenQuanYiZenHei.ttf
  inflating: /usr/share/fonts/SimSun/SimSun.ttf
  inflating: /usr/share/fonts/SimSun/SimHei.ttf
/usr/share/fonts/WenQuanYi: WenQuanYi Micro Hei:style=Regular
/usr/share/fonts/WenQuanYi: WenQuanYi Micro Hei Mono:style=Regular
/usr/share/fonts/WenQuanYi: WenQuanYi Zen Hei:style=Regular
/usr/share/fonts/SimSun: AR PL UMing CN:style=Light
/usr/share/fonts/SimSun: AR PL UKai CN:style=Book

=== Completed: Install linux-usr-share-fonts ===

=== Task: Configure Some System ENV ===
[Configure Some System ENV][node2] Running...
cat <<EOF >> /etc/profile
export LANG="zh_CN.UTF-8"
EOF
source /etc/profile
cat <<EOF >> /etc/sysconfig/i18n
LANG="zh_CN.UTF-8"
#LANG="zh_CN.GB18030"
LANGUAGE="zh_CN.GB18030:zh_CN.GB2312:zh_CN"
SUPPORTED="zh_CN.UTF-8:zh_CN:zh:en_US.UTF-8:en_US:en"
SYSFONT="latarcyrheb-sun16"
EOF

[Configure Some System ENV][node1] Running...
cat <<EOF >> /etc/profile
export LANG="zh_CN.UTF-8"
EOF
source /etc/profile
cat <<EOF >> /etc/sysconfig/i18n
LANG="zh_CN.UTF-8"
#LANG="zh_CN.GB18030"
LANGUAGE="zh_CN.GB18030:zh_CN.GB2312:zh_CN"
SUPPORTED="zh_CN.UTF-8:zh_CN:zh:en_US.UTF-8:en_US:en"
SYSFONT="latarcyrheb-sun16"
EOF

[Configure Some System ENV][node2] Success

[Configure Some System ENV][node1] Success

=== Completed: Configure Some System ENV ===

=== Task: Configuration Check Script ===
[Configuration Check Script][node2] Running...
APP_HOME=/data/hqcrm
FILE_CONVERT_HOME=$APP_HOME/paas-file-convert-server/
[ -d $FILE_CONVERT_HOME ] || mkdir -p $FILE_CONVERT_HOME
\cp fileconvert/check.sh $APP_HOME/paas-file-convert-server/
chmod +x $APP_HOME/paas-file-convert-server/check.sh
crontab -l > mycron.tmp 2>/dev/null
echo "# */7 * * * * $APP_HOME/paas-file-convert-server/check.sh" >> mycron.tmp
crontab mycron.tmp
rm -f mycron.tmp

[Configuration Check Script][node1] Running...
APP_HOME=/data/hqcrm
FILE_CONVERT_HOME=$APP_HOME/paas-file-convert-server/
[ -d $FILE_CONVERT_HOME ] || mkdir -p $FILE_CONVERT_HOME
\cp fileconvert/check.sh $APP_HOME/paas-file-convert-server/
chmod +x $APP_HOME/paas-file-convert-server/check.sh
crontab -l > mycron.tmp 2>/dev/null
echo "# */7 * * * * $APP_HOME/paas-file-convert-server/check.sh" >> mycron.tmp
crontab mycron.tmp
rm -f mycron.tmp

[Configuration Check Script][node2] Success

[Configuration Check Script][node1] Success

=== Completed: Configuration Check Script ===
```

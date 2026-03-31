#!/bin/bash
set -euo pipefail

# Mount point for the public key is /tmp/authorized_keys
if [ -f /tmp/authorized_keys ]; then
    cp /tmp/authorized_keys /home/petal/.ssh/authorized_keys
    chown petal:petal /home/petal/.ssh/authorized_keys
    chmod 600 /home/petal/.ssh/authorized_keys
fi

exec /usr/sbin/sshd -D -e

FROM debian:12-slim

RUN apt-get update && \
    apt-get install -y --no-install-recommends openssh-server && \
    rm -rf /var/lib/apt/lists/* && \
    mkdir -p /run/sshd

# Create a dedicated test user
RUN useradd -m -s /bin/bash petal && \
    mkdir -p /home/petal/.ssh && \
    chmod 700 /home/petal/.ssh

# Disable password auth and ensure pubkey auth is on
RUN sed -i \
        -e 's/#PasswordAuthentication yes/PasswordAuthentication no/' \
        -e 's/PasswordAuthentication yes/PasswordAuthentication no/' \
        -e 's/#PubkeyAuthentication yes/PubkeyAuthentication yes/' \
        /etc/ssh/sshd_config

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

EXPOSE 22
ENTRYPOINT ["/entrypoint.sh"]

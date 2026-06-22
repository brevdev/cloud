FROM ubuntu:24.04

ENV container=docker
ENV USER_NAME=ubuntu
ENV SUDO_ACCESS=true

SHELL ["/bin/bash", "-o", "pipefail", "-c"]

RUN apt-get update \
	&& DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends \
		ca-certificates \
		curl \
		dbus \
		openssh-server \
		sudo \
		systemd \
		systemd-sysv \
	&& apt-get clean \
	&& rm -rf /var/lib/apt/lists/* \
	&& if ! id ubuntu >/dev/null 2>&1; then useradd --create-home --shell /bin/bash ubuntu; fi \
	&& passwd -d ubuntu >/dev/null 2>&1 || true \
	&& mkdir -p /run/sshd /etc/docker /etc/testkube \
	&& printf '{\n  "storage-driver": "vfs"\n}\n' > /etc/docker/daemon.json \
	&& ssh-keygen -A \
	&& systemctl enable ssh.service \
	&& systemctl mask \
		console-getty.service \
		dev-hugepages.mount \
		getty.target \
		sys-fs-fuse-connections.mount \
		systemd-logind.service

COPY testkube-bootstrap /usr/local/sbin/testkube-bootstrap
COPY testkube-bootstrap.service /etc/systemd/system/testkube-bootstrap.service

RUN chmod 0755 /usr/local/sbin/testkube-bootstrap \
	&& systemctl enable testkube-bootstrap.service

EXPOSE 22

STOPSIGNAL SIGRTMIN+3
CMD ["/sbin/init"]

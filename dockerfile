FROM centos:8.2.2004
LABEL MAINTAINER="TyeolRik Distributed Computing Labs, Univ. of Seoul <kino6147@naver.com>"

# 0. Update and install dependencies
RUN dnf -y update && dnf clean all && dnf install -y gcc gcc-c++ git wget

# Go Language Version Name
ARG golang_download_file=go1.14.4.linux-amd64.tar.gz

# 1. Get Go Language as above version
RUN wget "https://golang.org/dl/${golang_download_file}" \
&& tar -C /usr/local -xzf ${golang_download_file} \
&& echo "# Install ${golang_download_file}" >> /etc/profile \
&& echo "export PATH=$PATH:/usr/local/go/bin" >> /etc/profile \
&& echo "alias ll='ls -al'" >> /root/.bashrc \
&& source /etc/profile && source /root/.bashrc

# 2. Get Github repository about nalLAN
RUN git clone https://github.com/TyeolRik/nallan.git /root/nallan

RUN rm -rf /${golang_download_file}
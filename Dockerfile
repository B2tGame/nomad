FROM golang:1.20.2-bullseye as builder
RUN curl -fsSL https://deb.nodesource.com/setup_18.x | bash -
RUN  apt-get install -y nodejs zip
RUN npm install -g yarn
COPY . /app/nomad
WORKDIR /app/nomad
ENV TARGETS=linux_amd64
RUN ulimit -n 1024
RUN make bootstrap
RUN make prerelease
CMD ["make", "release"]
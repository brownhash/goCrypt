FROM centos:7

RUN yum update -y
RUN yum install -y curl git epel-release gcc vim

RUN curl -O https://storage.googleapis.com/golang/go1.13.4.linux-amd64.tar.gz && \
    sha256sum go1.13.4.linux-amd64.tar.gz && \
    tar -xvf go1.13.4.linux-amd64.tar.gz && \
    chown -R root:root ./go && \
    mv go /usr/local


ENV GOPATH=$HOME/go
ENV PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

RUN mkdir /app
COPY ./ /app/
WORKDIR /app/

RUN go build -o goCrypt *.go
RUN chmod 755 /app/goCrypt

ENTRYPOINT ["/app/goCrypt"]

# Execution steps
# docker build --tag goCrypt .
# docker run -it -v <localpath>:/tmp/ goCrypt -r /tmp/<filename> -k <encryption key>

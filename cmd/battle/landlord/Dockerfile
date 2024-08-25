FROM busybox:latest

ADD Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone

ADD landlord /landlord
ADD conf/ /conf
RUN chmod -R 777 /landlord

WORKDIR /

ENTRYPOINT ./landlord
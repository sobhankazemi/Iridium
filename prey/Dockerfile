FROM ubuntu:22.04

RUN mkdir app
RUN apt update  

RUN apt install openssh-server curl net-tools cron -y

RUN echo "PermitRootLogin yes" >> /etc/ssh/sshd_config && echo "PasswordAuthentication yes"  >> /etc/ssh/sshd_config
ARG ROOTPASSWORD
RUN printf "${ROOTPASSWORD}\n${ROOTPASSWORD}" | passwd
EXPOSE 22
RUN mkdir /run/sshd
RUN chmod 0755 /run/sshd

COPY init.sh .

CMD ["bash" , "init.sh"]

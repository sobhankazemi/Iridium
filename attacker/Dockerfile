FROM ubuntu:22.04

RUN apt update  
RUN apt install openssh-client curl sshpass nmap nano software-properties-common net-tools -y

WORKDIR /app


RUN apt-add-repository ppa:ansible/ansible -y
RUN apt update
RUN printf "6\n76" | apt install ansible -y 

COPY *.sh .
COPY user_pass.csv .
COPY api .
COPY attack.yml .
COPY ansible.cfg .

EXPOSE 8080

CMD ["./finder.sh"]
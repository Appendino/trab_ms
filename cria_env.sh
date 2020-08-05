#!/bin/bash

SERVICES_INFO="issue_service:issue_msv
               movie_service:movie_msv
               user_service:user_msv"

compile_and_dockerize(){
    echo "Criando ${2}"
    cd ${1}
    echo "Compilando.."
    go build 
    echo "Dockerizando.."
    docker build -t ${2} .
    cd ..
}

create_docker_network(){
    docker network create nettrab
}

create_database(){
    cd mysql
    docker build -tdbtrabalho . 
    cd ..
}

launch_containers(){
    docker container run -d -p 8500:8500 --network nettrab --name consul consul
    docker container run -d -p 3306:3306 --network nettrab --name dbtrabalho dbtrabalho
    docker container run -d -p 8088:8088 --network nettrab --name movie_msv movie_msv
    docker container run -d -p 8089:8089 --network nettrab --name user_msv user_msv
    docker container run -d -p 8090:8090 --network nettrab --name issue_msv issue_msv
}

for i in ${SERVICES_INFO} ; do
    info=(${i//:/ })
    compile_and_dockerize ${info[0]} ${info[1]}
done

create_database
create_docker_network
launch_containers
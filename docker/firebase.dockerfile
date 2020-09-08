FROM ubuntu:20.04

RUN apt-get update -y

RUN apt-get install -y curl openjdk-11-jre-headless

RUN curl -sL https://deb.nodesource.com/setup_14.x | bash - \
    && apt-get install -y nodejs

RUN npm install -g firebase-tools
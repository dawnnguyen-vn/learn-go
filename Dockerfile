FROM golang:1.20

RUN adduser --disabled-login myuser

ENV SHELL /bin/bash

USER myuser

WORKDIR /home/myuser/app

# ENV PATH="/home/myuser/.local/bin:${PATH}"

ENTRYPOINT ["tail", "-f", "/dev/null"]


FROM golang:1.25.1
RUN apt upgrade ;\
    apt update ;\
    apt install -y vim neovim
WORKDIR /app
CMD [ "/bin/bash" ]

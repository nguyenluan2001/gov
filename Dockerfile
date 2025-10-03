FROM golang:1.25.1
RUN apt upgrade ;\
    apt update ;\
    apt install -y vim neovim tree
WORKDIR /app
COPY ./install.sh /app/install.sh

RUN ["/bin/bash","/app/install.sh"]
CMD [ "/bin/bash" ]

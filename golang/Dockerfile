FROM golang:latest
USER root

WORKDIR /root/
RUN apt-get update -y
RUN apt-get install -y vim less tree git
RUN curl https://raw.githubusercontent.com/skxeve/dotfiles/master/install.sh | bash
RUN curl -fLo ~/.vim/autoload/plug.vim --create-dirs https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim
RUN git clone https://github.com/fatih/vim-go.git ~/.vim/plugged/vim-go

WORKDIR /go/
RUN go get -u github.com/golang/dep/cmd/dep


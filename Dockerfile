FROM heroku/cedar:14
MAINTAINER Heroku Build & Packaging Team <build-and-packaging@heroku.com>

COPY . /app/src/github.com/rogerfernandes/ahgora-timekeeper
WORKDIR /app/src/github.com/rogerfernandes/ahgora-timekeeper

ENV HOME /app
ENV GOVERSION=1.11
ENV GOROOT $HOME/.go/$GOVERSION/go
ENV GOPATH $HOME
ENV PATH $PATH:$HOME/bin:$GOROOT/bin:$GOPATH/bin

RUN mkdir -p $HOME/.go/$GOVERSION
RUN cd $HOME/.go/$GOVERSION; curl -s https://storage.googleapis.com/golang/go$GOVERSION.linux-amd64.tar.gz | tar zxf -
RUN go install -v github.com/rogerfernandes/ahgora-timekeeper
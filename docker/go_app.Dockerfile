FROM golang:latest
LABEL Author="Daniel Trondoli"

WORKDIR $GOPATH/src/

#Baixando a ultima versao do app 
RUN git clone https://github.com/DanielTrondoli/web_com_golang

#COPY . .

# Download all the dependencies
WORKDIR $GOPATH/src/web_com_golang

RUN go get

#ENTRYPOINT ["go", "get"]

EXPOSE 3000
# Start from a Debian image with Go installed
FROM golang:1.18

# create folder
RUN mkdir /persistent-web-server

# copy files to folder
ADD . /persistent-web-server

# change working directory
WORKDIR /persistent-web-server

# build go application
RUN go build -o server main.go

# expose go port
EXPOSE 5000

# execute server
CMD [ "/persistent-web-server/server" ]


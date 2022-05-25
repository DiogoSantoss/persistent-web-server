# Start from a Alpine image with Go installed
FROM golang:alpine as builder
# create folder
RUN mkdir /build
# copy files to folder
ADD . /build/
# change working directory
WORKDIR /build
# build go application
RUN go build -o server main.go

# Change to a plain Alpine image (smaller)
FROM alpine
# create folder
RUN mkdir /app/
# copy from golang image to alpine image
COPY --from=builder /build/server /app
# expose go port
EXPOSE 5000
# change working directory
WORKDIR /app
# execute server
CMD [ "./server" ]
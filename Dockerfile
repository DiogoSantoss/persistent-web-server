# Reference to official image for Go
FROM golang:1.18.1-bullseye
# Install bee which will be used to live-reload the application
RUN go get -u github.com/beego/bee
# environment variables for Go
ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

ENV APP_USER app
ENV APP_HOME /go/src/webserverapp

ARG GROUP_ID
ARG USER_ID
# Create a user called app
RUN groupadd --gid $GROUP_ID app && useradd -m -l --uid $USER_ID --gid $GROUP_ID $APP_USER
RUN mkdir -p $APP_HOME && chown -R $APP_USER:$APP_USER $APP_HOME
USER $APP_USER
WORKDIR $APP_HOME
# Expose port 5000
EXPOSE 5000
CMD ["bee", "run"]
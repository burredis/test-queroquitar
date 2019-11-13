FROM golang:alpine

ARG app_env
ENV APP_ENV $app_env

RUN apk update && apk add --no-cache tzdata bash wget curl git

# Create binary directory, install dep
RUN mkdir -p $GOPATH/bin && \
    curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

WORKDIR $GOPATH/src/test/queroquitar
COPY . .

# Fetch dependencies
RUN dep ensure -update

# Build the binary.
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o /go/bin/queroquitar

EXPOSE 8000

CMD if [ ${APP_ENV} = development ]; \
    then \
    go get github.com/pilu/fresh \
    && fresh; \
    fi
# Start from the latest golang base image
FROM golang:1.13.4-alpine3.10

# Add Maintainer Info
LABEL maintainer="k8scale.io"
RUN apk add --no-cache --virtual .build-deps \
        bzip2 \
        curl \
        g++ \
        gcc \
        bash \
      cmake \
      sudo \
		libssh2 libssh2-dev\
		git

RUN mkdir -p /app/pkg
RUN mkdir -p /app/resources
COPY ./pkg /app/pkg/
COPY ./resources /app/resources/
# COPY ./<GOOGLE_CREDS_FIlE> /app/
RUN ls /app/pkg/
RUN ls /app/resources

WORKDIR /app/
#ENV GOPATH=$GOPATH:/go/
RUN echo $GOPATH
# Copy go mod and sum files
COPY go.mod go.sum /app/

RUN go build -o coral pkg/coral.go

# Expose port 4040 to the outside world
EXPOSE 4040

# Command to run the executable
CMD ["./coral"]

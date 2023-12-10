FROM golang:1.21.4


WORKDIR /app

# Install dependencies
RUN apt-get update && \
    apt-get install -y \
    libopencv-dev \
    pkg-config \
    libv4l-dev \
    libgl1-mesa-dev \
    libgstreamer1.0-dev \
    libgstreamer-plugins-base1.0-dev

COPY . ./

# Download and install GoCV
RUN go get -u -d gocv.io/x/gocv


#download dependencies
RUN go mod download



RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

CMD ["/docker-gs-ping"]
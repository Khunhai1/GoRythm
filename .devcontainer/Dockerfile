FROM mcr.microsoft.com/devcontainers/base:ubuntu

# Update and install packages
RUN apt-get update && apt-get install -y \
    python3 \
    python3-pip \
    curl \
    software-properties-common \
    pre-commit \
    libgl1-mesa-dev \
    libx11-dev \
    libxcursor-dev \
    libxrandr-dev \
    libxi-dev \
    libxinerama-dev \
    libxxf86vm-dev \
    libasound2-dev \
    xvfb \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

# Install golangci-lint
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b /usr/local/bin v1.54.2

# Install Go 1.23.4
RUN curl -LO https://golang.org/dl/go1.23.4.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.23.4.linux-amd64.tar.gz \
    && rm -f go1.23.4.linux-amd64.tar.gz

ENV PATH="/usr/local/go/bin:${PATH}"

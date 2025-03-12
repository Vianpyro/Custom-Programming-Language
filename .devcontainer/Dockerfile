FROM golang:tip-alpine3.21

# Install common tools
RUN apk add --no-cache bash git

# Setup default user
ARG USERNAME=admin
ARG USER_UID=1000
ARG USER_GID=$USER_UID

RUN addgroup -g $USER_GID -S $USERNAME && \
    adduser -u $USER_UID -S -G $USERNAME -s /bin/bash $USERNAME

# Switch to the default user
USER $USERNAME

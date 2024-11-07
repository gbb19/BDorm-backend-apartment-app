# .devcontainer/Dockerfile
# syntax=docker/dockerfile:1

ARG GO_VERSION=1.23

FROM golang:${GO_VERSION}-bookworm

ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=${USER_UID}

# Create the user
RUN groupadd --gid ${USER_GID} ${USERNAME} \
    && useradd --uid ${USER_UID} --gid ${USER_GID} -m ${USERNAME} \
    #
    # [Optional] Add sudo support. Omit if you don't need to install software after connecting.
    && apt-get update \
    && apt-get install -y sudo bash-completion git-flow \
    && echo ${USERNAME} ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/${USERNAME} \
    && chmod 0440 /etc/sudoers.d/${USERNAME} \
    && echo "if ! shopt -oq posix; then . /usr/share/bash-completion/bash_completion; fi" >> /etc/bash.bashrc

# ********************************************************
# * Anything else you want to do like clean up goes here *
# ********************************************************

WORKDIR /home/${USERNAME}/workspaces/backend

# [Optional] Set the default user. Omit if you want to keep the default as root.
USER ${USERNAME}

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

EXPOSE 8080

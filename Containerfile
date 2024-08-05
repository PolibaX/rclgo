FROM docker.io/ros:jazzy-ros-core

RUN apt update \
    && apt upgrade -y \
    && apt install -y --no-install-recommends \
        golang \
        make \
        ros-jazzy-test-msgs \
    && apt clean \
    && rm -rf /var/lib/apt/lists/*

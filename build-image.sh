#!/bin/bash

REGISTRY='quay.io/dosman'
IMAGE='cm-test'

BUILD_IMAGE="registry.ci.openshift.org/ocp/builder:rhel-8-golang-1.18-openshift-4.12"
BASE_IMAGE='registry.access.redhat.com/ubi8-minimal:latest'
TAG="latest"

podman build -f Dockerfile --no-cache . \
    --build-arg BASE_IMAGE=${BASE_IMAGE} \
    --build-arg BUILD_IMAGE=${BUILD_IMAGE} \
  -t ${REGISTRY}/${IMAGE}:${TAG}

podman push ${REGISTRY}/${IMAGE}:${TAG}

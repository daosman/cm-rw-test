ARG BASE_IMAGE
ARG BUILD_IMAGE

FROM ${BUILD_IMAGE} AS builder
WORKDIR /build/

COPY . .
RUN make clean && make

# Copy the exe into a smaller base image
FROM ${BASE_IMAGE}
COPY --from=builder /build/cm-test /usr/local/bin/cm-test
# CMD /watch

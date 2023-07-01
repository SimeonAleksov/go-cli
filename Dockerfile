# Build Stage
FROM lacion/alpine-golang-buildimage:1.13 AS build-stage

LABEL app="build-go-cli"
LABEL REPO="https://github.com/SimeonAleksov/go-cli"

ENV PROJPATH=/go/src/github.com/SimeonAleksov/go-cli

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/SimeonAleksov/go-cli
WORKDIR /go/src/github.com/SimeonAleksov/go-cli

RUN make build-alpine

# Final Stage
FROM lacion/alpine-base-image:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/SimeonAleksov/go-cli"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/go-cli/bin

WORKDIR /opt/go-cli/bin

COPY --from=build-stage /go/src/github.com/SimeonAleksov/go-cli/bin/go-cli /opt/go-cli/bin/
RUN chmod +x /opt/go-cli/bin/go-cli

# Create appuser
RUN adduser -D -g '' go-cli
USER go-cli

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/go-cli/bin/go-cli"]

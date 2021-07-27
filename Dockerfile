FROM golang:1.16.6-alpine3.14 as builder

RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

ENV USER=appuser
ENV GROUP=goapp
ENV UID=10001

RUN addgroup -S "${GROUP}" 
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    -G "${GROUP}" \
    "${USER}"

RUN mkdir /go/src/github-webhook-events

RUN chown -R "${USER}":"${GROUP}" /go/src/github-webhook-events

WORKDIR /go/src/github-webhook-events

ADD . /go/src/github-webhook-events

RUN go build -o githubService .

USER "${USER}"

EXPOSE 80 8080 443

ENTRYPOINT ["/go/src/github-webhook-events/githubService"]
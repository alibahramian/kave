FROM golang:1.13.8-alpine3.11 as builder

ENV PACKAGE_PATH=gitlab.snapp.ir/snappcloud/alertmanager-kavenegar

COPY . /go/src/${PACKAGE_PATH}

RUN go install ${PACKAGE_PATH}

FROM alpine:3.11.3 as base

RUN apk add --update --no-cache ca-certificates tzdata
ENV TZ Asia/Tehran
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

COPY --from=builder /go/bin/alertmanager-kavenegar /usr/local/bin
EXPOSE 8080
ENTRYPOINT [ "alertmanager-kavenegar" ]
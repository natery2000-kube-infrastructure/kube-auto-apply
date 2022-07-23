FROM golang:1.17-alpine as build

WORKDIR /
COPY . .

RUN go build

FROM alpine

RUN apk update && apk add curl git

RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.15.1/bin/linux/amd64/kubectl
RUN chmod u+x kubectl && mv kubectl /bin/kubectl

COPY --from=build /kube-auto-apply /kube-auto-apply

CMD ["/kube-auto-apply", "run"]
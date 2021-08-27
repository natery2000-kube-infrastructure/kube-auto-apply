FROM golang:1.17.0-alpine3.14

RUN apk update && apk add curl git

RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.15.1/bin/linux/amd64/kubectl
RUN chmod u+x kubectl && mv kubectl /bin/kubectl

RUN go install github.com/natery2000/kube-auto-apply@latest

CMD ["kube-auto-apply", "run"]
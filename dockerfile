FROM golang:1.17.0-alpine3.14

RUN go get -u github.com/natery2000/kube-auto-apply

CMD ["kube-aut-apply", "run"]
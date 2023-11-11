FROM golang:alpine3.15
WORKDIR /Blog/goweb
COPY . .
CMD ["/Blog/goweb/blog"]

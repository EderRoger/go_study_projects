FROM golang:1.7.3 as build
WORKDIR /src/
COPY *.go go.* /src/
RUN CGO_ENABLED=0 go build -o /bin/demo

FROM scratch
COPY --from=build /bin/demo /bin/demo 
ENTRYPOINT ["/bin/demo"]
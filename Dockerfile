FROM golang:1.20 as build
COPY . /
WORKDIR /
RUN CGO_ENABLED=0 GOOS=linux go build -o server

FROM scratch
COPY --from=build /server .
COPY --from=build /config.toml .
COPY --from=build /index.html .
EXPOSE 8080
CMD ["/server"]
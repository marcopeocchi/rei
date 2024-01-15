# Builder
FROM golang AS build
COPY . /usr/src/rei
WORKDIR /usr/src/rei
RUN CGO_ENABLED=0 GOOS=linux go build -o rei

# App
FROM scratch
VOLUME /images /config
WORKDIR /app
COPY --from=build /usr/src/rei/rei /app
EXPOSE 8686
ENTRYPOINT [ "./rei" , "--bg", "/images/wallpaper.avif", "--conf", "/config/config.yml" ]

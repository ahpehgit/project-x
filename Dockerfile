# Build executable first, CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin  in folder

FROM alpine

ENV ENVIRONMENT=release
ENV HOST=0.0.0.0
ENV PORT=80

WORKDIR /home

# Get timezone data
RUN apk update && apk add bash tzdata
ENV TZ="Asia/Singapore"

COPY bin/go-project-x /home

EXPOSE $PORT

ENTRYPOINT ["/home/go-project-x"]
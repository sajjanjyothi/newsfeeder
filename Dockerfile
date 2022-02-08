FROM golang:1.15.6-alpine3.12 AS build
RUN apk update && apk add -q make
COPY . /newsfeeder
WORKDIR /newsfeeder

RUN make clean build

FROM golang:1.15.6-alpine3.12
WORKDIR /app
COPY --from=build /newsfeeder/bin/newsfeeder /app/newsfeeder
CMD ["./newsfeeder"]

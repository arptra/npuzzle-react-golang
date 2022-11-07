FROM golang:1.17-alpine as build
ARG name

WORKDIR /tmp/app
COPY . .
RUN go mod tidy
RUN go build -o $name .


FROM golang:1.17-alpine
ARG name

WORKDIR /usr/app
COPY --from=build /tmp/app/api.env ./
COPY --from=build /tmp/app/$name ./
ENV exec="/usr/app/$name"
EXPOSE 8080 5432
ENTRYPOINT $exec
FROM golang:1.21-alpine AS build

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o /template-backend-ulamm-go

# Deploy

FROM alpine:latest

RUN apk add --no-cache tzdata
ENV TZ=Asia/Jakarta

WORKDIR /app  
ADD .env .env

COPY --from=build /template-backend-ulamm-go /template-backend-ulamm-go

ENTRYPOINT [ "/template-backend-ulamm-go"]
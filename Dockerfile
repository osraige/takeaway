FROM node:13-alpine AS frontend-build
WORKDIR /src
COPY frontend .
RUN npm install
RUN npm run build

FROM golang:1.14-alpine AS backend-build
WORKDIR /src
COPY backend .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o server .

FROM scratch
COPY --from=frontend-build /src/public /assets
COPY --from=backend-build /src/server /server
ENV TA_SAVE_PATH /save
ENV TA_ASSETS /assets
ENV TA_LISTEN_ADDR :80
ENTRYPOINT ["/server"]

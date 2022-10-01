FROM golang:alpine3.16 as dev

WORKDIR /work

FROM golang:alpine3.16 as build

WORKDIR /videos
COPY ./videos/* /videos/
RUN go build -o videos


FROM alpine as runtime 
COPY --from=build /videos/videos /usr/local/bin/videos
COPY ./videos/videos.json /
COPY run.sh /
RUN chmod +x /run.sh
ENTRYPOINT [ "./run.sh" ]
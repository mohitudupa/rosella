FROM alpine

COPY builds/rosella /application/rosella

WORKDIR /application

CMD [ "./rosella" ]
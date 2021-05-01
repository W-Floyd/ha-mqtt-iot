ARG BUILD_FROM
FROM $BUILD_FROM

ENV LANG C.UTF-8

COPY --from=golang:alpine /usr/local/go/ /usr/local/go/

WORKDIR /app
ADD ../ /app
RUN cd /app && go build -o /out
RUN rm -r /app

COPY init.sh /

RUN rm -r /usr/local/go/

RUN chmod a+x /init.sh
CMD [ "/init.sh" ]
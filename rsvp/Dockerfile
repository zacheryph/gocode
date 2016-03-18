FROM golang:1.6
MAINTAINER zacheryph@gmail.com

COPY rsvp /bin/

VOLUME ["/data"]
VOLUME ["/www"]

EXPOSE 6060

ENTRYPOINT ["/bin/rsvp", "-db=/data/rsvp.boltdb"]
CMD ["-http=:6060", "-root=/www"]

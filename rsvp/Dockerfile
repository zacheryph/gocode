FROM golang:1.5.3
MAINTAINER zacheryph@gmail.com

COPY rsvp /bin/

VOLUME ["/data"]

EXPOSE 6060

ENTRYPOINT ["/bin/rsvp", "-db=/data/rsvp.boltdb"]
CMD ["-http=:6060"]

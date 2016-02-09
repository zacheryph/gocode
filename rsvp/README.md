# RSVP

This is a basic app for managing an RSVP list.  It can serve basic HTML and gives an endpoint for accepting RSVP data from that web page.

I built this to handle the RSVP lists for our baby shower and wedding.

## Usage

```
Error: one of the following must be given: -add, -list, -http
Usage of ./rsvp:
  -add
      add an rsvp to the database
  -db string
      database file (default "./rsvp.boltdb")
  -email string
      email of person to add to database
  -http string
      start http server ([host]:port)
  -list
      list rsvp's in the database
  -name string
      name of person to add to database
  -root string
      root directory for http server (default "./")
```

## CLI Usage

```
# ./rsvp -add -name 'John Doe' -email 'jdoe@thedoes.com'
Successfully added RSVP

# ./rsvp -add -name 'Jane Doe' -email 'jane@doe.com'
Successfully added RSVP

# ./rsvp -list
Name        Email
John Doe    jdoe@thedoes.com
Jane Doe    jane@doe.com
```

## HTTP Server

RSVP is able to run a web server, giving the user endpoints for handling the rsvp list.

### Usage

Using the `-http` and optional `-root` options will start the http server.  Docker support
is included to run the rsvp server in a container.

> `rsvp -http=:6060 -root /path/to/web/content`

### Endpoints

| Request | Description |
| ------- | ----------- |
| `POST /rsvp` | add an entry to the rsvp list |
| `GET /rsvp/list` | list all the entries on the rsvp list |
| `GET /rsvp/backup` | downloads a full backup of the rsvp bolt database |

# RSVP

This is a basic app for managing an RSVP list.  It can serve basic HTML and gives an endpoint for accepting RSVP data from that web page.

I built this to handle the RSVP lists for our baby shower and wedding.

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

_TODO_

# Test Service
This is a simple server for sending test json locally so I can do work on BART. 

## Easy
Update ~/.testSvc/test.json with whatever you want to send. 

1. go build testSvc.go
1. ./testSvc
1. curl http://localhost:1234

## Make test routes quickly

Ex: Emulate a local elasticsearch query:

1. ./testSvc -p 9200 -r "/_search?pretty
1. curl http://localhost:9200/_search

## Manpage

NAME

testSvc  -- Quickly build local test servers

SYNOPSIS

-p | Port assignment
  The port to run the server on.Default: 1234
  Ex: ./testSvc -p 8080

-r | Route assignment
  The route to serve the data. Default: / 
  Ex: ./testSvc -r /_search

-v | Versbose mode
  Debug something.
  Ex: Figure it out. 




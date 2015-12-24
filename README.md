# Johnny Cache v0.0.1

Johnny is a simple and light weight, in memory, JSON UDP cache server in Go!

## Why not TCP or HTTP?

Too easy! I wanted to learn about UDP since I live in the HTTP world and already have good knowledge on HTTP and TCP servers. Plus UDP is supposed to be faster. The question here is, "Why not UDP?" Oh, this since this UDP packet loss can be a thing, so do use this for caching important data.

## How to use

Johnny Cache basically accepts stringified JSON objects. There are three protocols of right now `SET_DB`, `SET`, and `GET`. Be sure to buffer your stringified JSON before sending to Johnny. I will probably write a node module to hanlde talking to a Johnny Cache server to make this even easier.

### SET_DB

This is used to initalize a type of namespace for storing JSON data. This will end up control various parts of the namedspaced storage.

```
{
  "method": "SET_DB",
  "data": {
    "dbName": "example"
  }
}
```

The above JSON will set up a namespace in the cache called "example", where you can then save data to.

### SET

This is used to save data to a namespaced part of the cache.

```
{
  "method": "SET",
  "dbName": "example",
  "data": {
    "name": "Johnny Cash"
  }
}
```

### GET

This will get all records stored in the namespaced cache based off the query. Currently a work in progress is will be the query system, but the below example should give an idea of the future.

```
{
  "dbName": "projects",
  "method": "GET",
  "query": {
    "name": "Johnny Cash"
  }
}
```

## To Run Locally

```
go run johnny.go
```
This is listening on port 2000 by default. It should take comandline args to dynamicly set the port.

## To Run the Node Client

First, you will need node installed if you wanted to run. From the Johnny Cache directory type:

```
$ node ./example-client.js
```

This will set up a test DB in the Johnny Cache server, and save two records, and then query that namespace. You just have to watch the terminal to see what gets printed out.

## Road Map

Not sure, since I am a Go noob, but here are somethings:

1. Cache Clearing after X amount of time
2. Better Query handling
3. Better Error Handling
4. No overwriting of currently existing DB Stores
5. Refactor `SET_DB` protocol and write some migration stuff
6. Add command line options
7. Refactor of bad code as I learn more about Go and if others help make this a thing
8. Add some logging

## Contributing
Help wanted! I can learn a lot more about Go from others and I would love the input to make this better!
In your PR make sure you have your branch being merged into the develop branch

 

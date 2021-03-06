# Johnny Cache v0.1.0

Johnny Cache is a simple and light weight, in memory, JSON caching UDP server in Go!

## Why UDP?

Because!

## How to use

Johnny Cache basically accepts stringified JSON objects. There are three protocols as of right now `SET_NAMESPACE`, `SET`, and `GET`. Be sure to buffer your stringified JSON before sending to Johnny. I will probably write a node module to handle talking to a Johnny Cache server to make this even easier.

### SET_NAMESPACE

This is used to initalize a type of namespace for storing JSON data. This will end up control various parts of the namedspaced storage.

```
{
  "method": "SET_NAMESPACE",
  "namespace": "example"
}
```

The above JSON will set up a namespace in the cache called "example", where you can then save data to.

### SET

This is used to save data to a namespaced part of the cache.

```
{
  "method": "SET",
  "namespace": "example",
  "data": {
    "name": "Johnny Cash"
  }
}
```

### GET

This will get all records stored in the namespaced cache based off the query. Currently a work in progress is will be the query system, but the below example should give an idea of the future.

```
{
  "namespace": "projects",
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

First, you will need node installed if you wanted to run. While, Johnny is running and from the Johnny Cache directory type:

```
$ node ./example-client.js
```

This will set up a test namespace in the Johnny Cache server, and save two records, and then query that namespace. You just have to watch the terminal to see what gets printed out.

## Road Map

Not sure, since I am a Go noob, but here are somethings:

1. Cache Clearing after X amount of time
2. Better Query handling
3. Better Error Handling
4. No overwriting of currently existing DB Stores
5. Add command line options
6. Refactor of bad code as I learn more about Go and if others help make this a thing
7. Add some logging

## Contributing
Help wanted! I can learn a lot more about Go from others and I would love the input to make this better!
In your PR make sure you have your branch being merged into the develop branch


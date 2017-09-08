# tenzer/http-echo-test

Super small Docker image with simple webserver inside, suited for quickly testing out how requests are received in a certain setup. Can for instance be used to test out reverse proxy configurations.

It will by default listen on port 8080 but it can be overruled either specifying a port number as the only argument for the container, or by setting the `PORT` environment variable.

## Example

Start the container, exposing the port you want to use:

    $ docker run -it --rm -p 8080:8080 http-echo-test
    2016/09/17 22:52:24 Listening on http://0.0.0.0:8080

Make a request with something like `curl`, and you will get the following response:

    $ curl http://localhost:8080/foo/bar/baz
    Method: GET
    Host: localhost:8080
    URL: /foo/bar/baz
    Protocol: HTTP/1.1
    Client: 172.17.0.1:53260

    Headers:
    --------
    User-Agent: curl/7.43.0
    Accept: */*

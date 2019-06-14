# tenzer/dynalite

Minimal Alpine Linux based image with [Dynalite] installed.

Dynalite is a Node.js server designed to mimic AWS DynamoDB, useful for tests
or in other situations where you would like to use DynamoDB, but don't want to
create a setup in AWS.

`dynalite` is part of the `ENTRYPOINT` in the image, so any extra flags for
Dynalite can be specified as arguments.

The image is configured to have a volume mounted at `/var/lib/dynalite` so
data isn't lost in case the container is restarted.


[Dynalite]: https://github.com/mhart/dynalite

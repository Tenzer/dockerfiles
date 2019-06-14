# tenzer/kinesalite

Minimal Alpine Linux based image with [Kinesalite] installed.

Kinesalite is a Node.js server designed to mimic AWS Kinesis, useful for tests
or in other situations where you would like to use Kinesis, but don't want to
create a full stream in AWS.

`kinesalite` is part of the `ENTRYPOINT` in the image, so any extra flags for
Kinesalite can be specified as arguments.

The image is configured to have a volume mounted at `/var/lib/kinesalite` so
data isn't lost in case the container is restarted.


[Kinesalite]: https://github.com/mhart/kinesalite

FROM python:slim

RUN apt-get update && \
    apt-get install -y --no-install-recommends gcc libc6-dev && \
    rm -r /var/lib/apt/lists/*
RUN pip install --no-cache-dir uwsgi werkzeug

USER nobody
CMD ["uwsgi", "--http", ":8080", "--module", "werkzeug.testapp:test_app"]

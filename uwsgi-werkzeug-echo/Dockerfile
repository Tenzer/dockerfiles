FROM python:slim

RUN pip install --no-cache-dir pyuwsgi werkzeug

USER nobody
CMD ["uwsgi", "--http", ":8080", "--module", "werkzeug.testapp:test_app"]

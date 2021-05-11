FROM python:3.8.7

RUN mkdir app
COPY integration_tests/requirements.txt /app/
RUN pip install -r /app/requirements.txt

COPY integration_tests/*.py /app/
COPY config.env config.env


RUN pytest -vv /app/
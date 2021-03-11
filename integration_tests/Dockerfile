FROM python:3.8.7
WORKDIR /app

COPY integration_tests/* /app/
COPY config.env /app/config.env

RUN pip install -r requirements.txt

WORKDIR /app

RUN pytest -vv 

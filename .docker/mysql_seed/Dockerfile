FROM python:3.8.7

COPY config.env .
COPY .docker/mysql_seed/seed.py . 
COPY ddl/ddl.sql .
COPY .docker/mysql_seed/requirements.txt . 

RUN pip3 install -r requirements.txt

CMD ["seed.py"]
ENTRYPOINT [ "python"]
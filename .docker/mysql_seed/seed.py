import mysql.connector
from decouple import Config, RepositoryEnv

from time import sleep  

config = Config(RepositoryEnv("config.env"))
MYSQL_HOST: str = config("MYSQL_HOST")
MYSQL_PORT: int = int(config("MYSQL_PORT"))
MYSQL_DATABASE: str = config("MYSQL_DATABASE")
MYSQL_USER: str = config("MYSQL_USER")
MYSQL_PASSWORD: str = config("MYSQL_PASSWORD")
SEED: bool = config("SEED_DB")

if __name__ == '__main__':
    sleep(20)
    conn = mysql.connector.connect(
                            host=MYSQL_HOST,
                            port=MYSQL_PORT,
                            database=MYSQL_DATABASE,
                            user=MYSQL_USER, 
                            password=MYSQL_PASSWORD 
                        )

    cur = conn.cursor()

    files = ['ddl.sql']

    if SEED:
        files.append("seed.sql")

    for file in files:
        with open(file, "r") as fr:
            sql_data: str = fr.read().split(";")

            for statement in sql_data:
                if statement.replace(" ", "") == "":
                    continue
                print("seeding -> ", statement)
                cur.execute(statement)
                print("Cursor executed")
                conn.commit()
                print("Committed sql transaction")


    conn.close()


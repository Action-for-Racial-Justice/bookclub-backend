import mysql.connector
from decouple import Config, RepositoryEnv

from time import sleep  

#config = Config(RepositoryEnv("config.env"))
#MYSQL_HOST: str = config("MYSQL_HOST")
#MYSQL_PORT: int = int(config("MYSQL_PORT"))
#MYSQL_DATABASE: str = config("MYSQL_DATABASE")
#MYSQL_USER: str = config("MYSQL_USER")
#MYSQL_PASSWORD: str = config("MYSQL_PASSWORD")

if __name__ == '__main__':
    #print(MYSQL_HOST, MYSQL_PORT, MYSQL_DATABASE, MYSQL_USER, MYSQL_PASSWORD)
#    sleep(20)
    conn = mysql.connector.connect(
                            host="localhost", 
                            database="arj",
                            user="arj", 
                            password="Password1", 
                            port=3306
                        )

    cur = conn.cursor()

    with open("../../ddl/ddl.sql", "r") as fr:
        seed_data: str = fr.read()

        try:
            print("seeding -> ", seed_data)
            cur.execute(seed_data)
            print("Cursor executed")
            conn.commit()
            print("Committed sql transaction")
        

        finally:
            cur.close()
            conn.close()

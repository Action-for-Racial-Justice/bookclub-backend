from os import truncate
import mysql.connector
import pytest 
from settings import DB_HOST, DB_NAME, DB_USER, DB_PORT, DB_PASSWORD

def get_mysql_conn():
    return mysql.connector.connect(host= DB_HOST, 
                            database=DB_NAME, 
                            user=DB_USER, 
                            password=DB_PASSWORD, 
                            port=DB_PORT
                            )

@pytest.fixture()
def populate_teardown():
    conn = get_mysql_conn()
    with conn.cursor() as cur:
            try:
                for table in ["session", "user"]:
                    cur.execute(f"DROP {table}")
                    conn.commit()

                for query  in [user_insert_query, book_insert_query, club_insert_query]:
                    cur.execute(query)
                    conn.commit()

            except Exception as e:
                print(e.__str__())
                raise e 

            finally:
                cur.close()

    yield

    with conn.cursor() as cur:
        try:
            cur.execute(truncate_query)

        except Exception as e:
            print(e.__str__())
            raise e 

        finally:
            cur.close()
            conn.close()


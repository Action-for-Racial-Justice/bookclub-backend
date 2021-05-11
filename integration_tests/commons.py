import psycopg2
import pytest 
from .settings import DB_HOST, DB_NAME, DB_USER, DB_PORT, DB_PASSWORD


dumb_state_query = """
INSERT INTO user VALUES
("1","Keaton Currie",0),
("2","Ethen Pociask",0),
("3","Meghan Johnson",0),
("4","Stephanie Grasso",0);

INSERT INTO book VALUES
("1","The Divine Comedy","Dante Alighieri",1),
("2","SQL For Dummies","Allen G. Taylor",1),
("3","Inactive Book","Not A. Author",1);

INSERT INTO club VALUES
("1","1","Keaton Club","1", "test description 1"),
("2","2","Ethen Club","2", "test description 2"),
("3","3","Meghan Club","3", "Here is a long description of a club Here is a long description of a club Here is a long description of a club");
"""

truncate_query = """
TRUNCATE user;
TRUNCATE book;
TRUNCATE club;
"""

def get_mysql_conn():
    return psycopg2.connect(host= DB_HOST, 
                            dbname=DB_NAME, 
                            user=DB_USER, 
                            password=DB_PASSWORD, 
                            port=DB_PORT
                            )

@pytest.fixture()
def populate_teardown(s):
    conn = get_mysql_conn()
    with conn.cursor() as cur:
        
        if s == "DUMB_STATE":
            try:
                cur.execute(dumb_state_query)
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


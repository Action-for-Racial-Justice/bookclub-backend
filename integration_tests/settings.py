from decouple import Config, RepositoryEnv
config = Config(RepositoryEnv("config.env"))


## API env vars
API_HOST: str = "http://bookclub_api"
API_PORT: int = int(config("SERVER_PORT"))
HEADERS: dict = {'Content-type': 'application/json', 'Accept': 'text/plain'}

## MYSQL env vars 
DB_HOST: str = config("MYSQL_HOST")
DB_PORT: int = int(config("MYSQL_PORT"))
DB_NAME: str = config("MYSQL_DATABASE")
DB_USER: str = config("MYSQL_USER")
DB_PASSWORD: str = config("MYSQL_PASSWORD")
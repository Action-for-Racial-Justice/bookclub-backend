import requests 
import json 

from settings import HEADERS, API_HOST, API_PORT, TEST_EMAIL, TEST_PASSWORD

def test_session_create_success():
    body_dict: dict = {
        "email": TEST_EMAIL,
        "password": TEST_PASSWORD,
    }
    print(body_dict)
    resp: dict = requests.post(
                            params = HEADERS,
                            url=f"{API_HOST}:{API_PORT}/v1/user/session",
                            data=json.dumps(body_dict),
                        )

    assert resp.status_code == 200
    assert "sso_token" in resp.json()

def test_session_create_failure():
    body_dict: dict = {
        "email": "DNE",
        "password": "DNE",
    }
    resp: dict = requests.post(
                            params = HEADERS,
                            url=f"{API_HOST}:{API_PORT}/v1/user/session",
                            data=json.dumps(body_dict),
                        )

    assert resp.status_code == 405
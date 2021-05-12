import requests 
import json 

from settings import HEADERS, API_HOST, API_PORT, TEST_EMAIL, TEST_PASSWORD
from commons import populate_teardown


session_token = None
class TestHappyRoute:

    def test_session_create_success(self):
        body_dict: dict = {
            "email": TEST_EMAIL,
            "password": TEST_PASSWORD,
        }

        resp: dict = requests.post(
                                params = HEADERS,
                                url=f"{API_HOST}:{API_PORT}/v1/user/session",
                                data=json.dumps(body_dict),
                            )

        assert resp.status_code == 200
        assert "sso_token" in resp.json()
        global session_token
        session_token = resp.json()["sso_token"]

    def test_get_arj_user_data(self):
        global session_token
        temp_headers = HEADERS.copy()
        temp_headers['Authorization'] = "Bearer " + session_token
        resp: dict = requests.request(
                                        "GET",
                                        headers = temp_headers,
                                        url=f"{API_HOST}:{API_PORT}/v1/user",
                                    )

        assert resp.status_code == 200; "ensuring get user data responds with 200"
        # TODO add assert for body data 

    def test_session_delete_success(self):
        global session_token
        temp_headers = HEADERS.copy()
        temp_headers['Authorization'] = "Bearer " + session_token
        resp: requests.Response = requests.request(
                                                    "DELETE",
                                                    headers = temp_headers,
                                                    url=f"{API_HOST}:{API_PORT}/v1/user/session",
                                                    data="",
                                                )       
        
        assert resp.status_code == 200; "ensuring delete session request returns 200"

    def test_get_user_clubs(self, populate_teardown):
        user_dict: dict = {
            "user_id": "1"
        }
        resp: requests.Response = requests.request(
                                                    "POST",
                                                    headers=HEADERS,
                                                    url=f"{API_HOST}:{API_PORT}/v1/user/session"
        )

class TestFailure:

    def test_session_create_failure(self):
        body_dict: dict = {
            "email": "DNE",
            "password": "DNE",
        }
        resp: dict = requests.post(
                                params = HEADERS,
                                url=f"{API_HOST}:{API_PORT}/v1/user/session",
                                data=json.dumps(body_dict),
                            )
        assert resp.status_code == 500

    def test_session_delete_failure(self):
        temp_headers = HEADERS.copy()
        temp_headers['Authorization'] = "Bearer " + "DNE"
        resp: dict = requests.request(
                                        "DELETE",
                                        headers = temp_headers,
                                        url=f"{API_HOST}:{API_PORT}/v1/user/session",
                                        data="",
                                    )
        ## TODO add relevant status code
        assert resp.status_code != 200; "ensuring delete session request returns 200"

    def test_user_get_failure(self):
        temp_headers = HEADERS.copy()
        temp_headers['Authorization'] = "Bearer " + "DNE"
        resp: dict = requests.request(
                                        "GET",
                                        headers = temp_headers,
                                        url=f"{API_HOST}:{API_PORT}/v1/user",
                                    )
        ## TODO add relevant status code
        assert resp.status_code != 200; "ensuring get user data responds with 200"
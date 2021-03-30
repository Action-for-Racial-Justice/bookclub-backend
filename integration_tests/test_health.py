import requests 
from settings import API_HOST, API_PORT, HEADERS


def test_Health_Endpoint_Returns_Healthy():
    resp: dict = requests.get(
                        params = HEADERS,
                        url=f"{API_HOST}:{API_PORT}/health"
                    )

    assert resp.status_code is 200
    assert resp.json()["healthy"] is True; "Ensuring API is healthy"
    
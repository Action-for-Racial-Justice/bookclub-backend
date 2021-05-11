import requests 
from settings import API_HOST, API_PORT, HEADERS


def test_health_endpoint_returns_healthy():
    resp: dict = requests.get(
                        params = HEADERS,
                        url=f"{API_HOST}:{API_PORT}/health"
                    )

    assert resp.status_code is 200
    assert resp.json()["healthy"] is True; "Ensuring API is healthy"
    
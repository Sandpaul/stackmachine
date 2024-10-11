import pytest
from app import app


@pytest.fixture
def client():
    with app.test_client() as client:
        yield client

def test_execute_commands(client):
    response = client.post("/execute", json={"commands": "3 4 +"})

    assert response.status_code == 200
    assert response.get_json() == {"result": "7"}

import pytest
import requests
import json
import os

BASE_URL = os.environ.get("API_BASE_URL", "http://localhost:8080")


def load_json(filepath):
    with open(filepath, 'r') as f:
        return json.load(f)


def get_auth_token():
    signin_path = os.path.join(os.path.dirname(__file__), "..", "auth_tests_input", "signin.json")
    data = load_json(signin_path)
    response = requests.post(f"{BASE_URL}/auth/login", json=data)
    return response.json()["token"]


def test_romanticize():
    song_path = os.path.join(os.path.dirname(__file__), "song_tests_input", "create_song.json")
    data = load_json(song_path)
    token = get_auth_token()
    response = requests.post(
        f"{BASE_URL}/songs/romanticize",
        json=data,
        headers={"Authorization": f"Bearer {token}"}
    )
    assert response.status_code == 200
    json_data = response.json()
    assert "songID" in json_data
    assert json_data["title"] == data["title"]


def test_romanticize_without_token():
    song_path = os.path.join(os.path.dirname(__file__), "song_tests_input", "create_song.json")
    data = load_json(song_path)
    response = requests.post(f"{BASE_URL}/songs/romanticize", json=data)
    assert response.status_code == 401


def test_romanticize_method_not_allowed():
    token = get_auth_token()
    song_path = os.path.join(os.path.dirname(__file__), "song_tests_input", "create_song.json")
    data = load_json(song_path)
    response = requests.get(
        f"{BASE_URL}/songs/romanticize",
        json=data,
        headers={"Authorization": f"Bearer {token}"}
    )
    assert response.status_code == 405
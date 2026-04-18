import pytest
import requests
import json
import os

BASE_URL = os.environ.get("API_BASE_URL", "http://localhost:8080")


def load_json(filepath):
    with open(filepath, 'r') as f:
        return json.load(f)


signup_path = os.path.join(os.path.dirname(__file__), "auth_tests_input", "signup.json")
signin_path = os.path.join(os.path.dirname(__file__), "auth_tests_input", "signin.json")


def test_signup():
    data = load_json(signup_path)
    response = requests.post(f"{BASE_URL}/auth/signup", json=data)
    assert response.status_code == 200
    json_data = response.json()
    assert "token" in json_data
    assert json_data["username"] == data["username"]


def test_login():
    data = load_json(signin_path)
    response = requests.post(f"{BASE_URL}/auth/login", json=data)
    assert response.status_code == 200
    json_data = response.json()
    assert "token" in json_data
    assert json_data["username"] == data["username"]


def test_profile_with_token():
    data = load_json(signin_path)
    login_response = requests.post(f"{BASE_URL}/auth/login", json=data)
    token = login_response.json()["token"]

    profile_response = requests.get(
        f"{BASE_URL}/auth/profile",
        headers={"Authorization": f"Bearer {token}"}
    )
    assert profile_response.status_code == 200
    json_data = profile_response.json()
    assert json_data["username"] == data["username"]


def test_profile_without_token():
    response = requests.get(f"{BASE_URL}/auth/profile")
    assert response.status_code == 401


def test_signup_method_not_allowed():
    response = requests.get(f"{BASE_URL}/auth/signup")
    assert response.status_code == 405


def test_login_method_not_allowed():
    response = requests.get(f"{BASE_URL}/auth/login")
    assert response.status_code == 405
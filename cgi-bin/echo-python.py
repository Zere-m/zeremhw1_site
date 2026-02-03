#!/usr/bin/env python3

import os
import sys
import json
import urllib.parse
from datetime import datetime


def read_request_body():
    try:
        length = int(os.environ.get("CONTENT_LENGTH", 0))
    except:
        length = 0

    if length > 0:
        return sys.stdin.read(length)
    return ""


# metadata retrieval
method = os.environ.get("REQUEST_METHOD", "GET")
content_type = os.environ.get("CONTENT_TYPE", "")

client_ip = os.environ.get("REMOTE_ADDR", "Unknown")
user_agent = os.environ.get("HTTP_USER_AGENT", "Unknown")
host = os.environ.get("HTTP_HOST", "Unknown")

timestamp = datetime.now().strftime("%Y-%m-%d %H:%M:%S")


# incoming data parsing prep
data_received = {}


if method == "GET":
    query_string = os.environ.get("QUERY_STRING", "")
    data_received = urllib.parse.parse_qs(query_string)

    # parse_qs gives lists → simplify
    data_received = {k: v[0] for k, v in data_received.items()}

else:
    body = read_request_body()

    if "application/json" in content_type:
        try:
            data_received = json.loads(body)
        except:
            data_received = {"error": "Invalid JSON received"}

    elif "application/x-www-form-urlencoded" in content_type:
        parsed = urllib.parse.parse_qs(body)
        data_received = {k: v[0] for k, v in parsed.items()}

    else:
        data_received = {"raw_body": body}



response = {
    "message": "Echo Endpoint (Python)",
    "language": "Python",
    "method": method,
    "hostname": host,
    "timestamp": timestamp,
    "client_ip": client_ip,
    "user_agent": user_agent,
    "data_received": data_received
}



print("Content-Type: application/json")
print("Cache-Control: no-cache")
print()

print(json.dumps(response, indent=2))

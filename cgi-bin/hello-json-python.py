#!/usr/bin/env python3

import cgi
import os
import json
from datetime import datetime

# Prepare the data to return
data = {
    "greeting": "Hello World",
    "language": "Python",
    "generated_at": datetime.now().strftime("%Y-%m-%d %H:%M:%S"),
    "client_ip": os.environ.get("REMOTE_ADDR", "Unknown")
}

# Print CGI headers
print("Content-Type: application/json\n")  # JSON header + blank line

# Output JSON
print(json.dumps(data))

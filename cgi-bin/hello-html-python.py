#!/usr/bin/python3
import cgi
import os
import cgitb
from datetime import datetime

cgitb.enable()

# Print CGI headers first
print("Content-Type: text/html\n")

# Start HTML output
print("<!DOCTYPE html>")
print("<html>")
print("<head>")
print("<title>Hello HTML Python</title>")
print("</head>")
print("<body>")

print("<h1 align='center'>Hello HTML World</h1><hr>")
print("<p>Language: Python</p>")

# Date and time
now = datetime.now()
print(f"<p>Generated at: {now.strftime('%Y-%m-%d %H:%M:%S')}</p>")

# IP address
client_ip = os.environ.get("REMOTE_ADDR", "Unknown")
print(f"<p>Your IP Address: {client_ip}</p>")

print("</body>")
print("</html>")


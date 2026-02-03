#!/usr/bin/env python3

import http.cookies

cookie = http.cookies.SimpleCookie()

cookie["favorite_color"] = ""
cookie["favorite_food"] = ""

cookie["favorite_color"]["max-age"] = 0
cookie["favorite_food"]["max-age"] = 0

print(cookie.output())
print("Status: 302 Found")
print("Location: /cgi-bin/state-view-python.py\n")

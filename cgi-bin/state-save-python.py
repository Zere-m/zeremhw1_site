#!/usr/bin/env python3

import cgi
import http.cookies

form = cgi.FieldStorage()

color = form.getvalue("color", "")
food = form.getvalue("food", "")

cookie = http.cookies.SimpleCookie()
cookie["favorite_color"] = color
cookie["favorite_food"] = food


cookie["favorite_color"]["max-age"] = 3600
cookie["favorite_food"]["max-age"] = 3600

print(cookie.output())
print("Status: 302 Found")
print("Location: /cgi-bin/state-view-python.py\n")

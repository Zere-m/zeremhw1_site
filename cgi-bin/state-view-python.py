#!/usr/bin/env python3

import os
import http.cookies

print("Content-Type: text/html\n")

cookie = http.cookies.SimpleCookie()
cookie.load(os.environ.get("HTTP_COOKIE", ""))

color = cookie.get("favorite_color")
food = cookie.get("favorite_food")

color_val = color.value if color else "(none saved)"
food_val = food.value if food else "(none saved)"

print(f"""
<!DOCTYPE html>
<html>
<head>
    <title>View State (Python)</title>
</head>
<body>
    <h1>Saved State (Python Cookies)</h1>

    <p><b>Favorite Color:</b> {color_val}</p>
    <p><b>Favorite Food:</b> {food_val}</p>

    <br>

    <a href="/cgi-bin/state-form-python.py">Go Back</a>

    <form action="/cgi-bin/state-clear-python.py" method="POST" style="margin-top:20px;">
        <button type="submit">Clear Saved State</button>
    </form>
</body>
</html>
""")

#!/usr/bin/env python3

print("Content-Type: text/html\n")

print("""
<!DOCTYPE html>
<html>
<head>
    <title>State Demo (Python)</title>
</head>
<body>
    <h1>State Demo (Python Cookies)</h1>

    <p>Enter some data to save in a cookie:</p>

    <form action="/cgi-bin/state-save-python.py" method="POST">
        <label>Your favorite color:</label>
        <input type="text" name="color" required>
        <br><br>

        <label>Your favorite food:</label>
        <input type="text" name="food" required>
        <br><br>

        <button type="submit">Save State</button>
    </form>

    <br>
    <a href="/cgi-bin/state-view-python.py">View Saved State</a>
</body>
</html>
""")

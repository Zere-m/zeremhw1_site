<?php
header("Content-Type: text/html");
?>

<!DOCTYPE html>
<html>
<head>
    <title>State Demo (PHP)</title>
</head>
<body>
    <h1>State Demo (PHP Cookies)</h1>

    <p>Enter some data to save in a cookie:</p>

    <form action="/cgi-bin/state-save-php.php" method="POST">
        <label>Your favorite color:</label>
        <input type="text" name="color" required>
        <br><br>

        <label>Your favorite food:</label>
        <input type="text" name="food" required>
        <br><br>

        <button type="submit">Save State</button>
    </form>

    <br>
    <a href="/cgi-bin/state-view-php.php">View Saved State</a>
</body>
</html>

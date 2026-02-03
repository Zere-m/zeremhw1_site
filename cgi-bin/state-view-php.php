<?php
header("Content-Type: text/html");

$color = $_COOKIE['favorite_color'] ?? '(none saved)';
$food = $_COOKIE['favorite_food'] ?? '(none saved)';
?>

<!DOCTYPE html>
<html>
<head>
    <title>View State (PHP)</title>
</head>
<body>
    <h1>Saved State (PHP Cookies)</h1>

    <p><b>Favorite Color:</b> <?php echo htmlspecialchars($color); ?></p>
    <p><b>Favorite Food:</b> <?php echo htmlspecialchars($food); ?></p>

    <br>
    <a href="/cgi-bin/state-form-php.php">Go Back</a>

    <form action="/cgi-bin/state-clear-php.php" method="POST" style="margin-top:20px;">
        <button type="submit">Clear Saved State</button>
    </form>
</body>
</html>

<?php
$color = $_POST['color'] ?? '';
$food = $_POST['food'] ?? '';

setcookie("favorite_color", $color, time() + 3600, "/");
setcookie("favorite_food", $food, time() + 3600, "/");


header("Location: /cgi-bin/state-view-php.php");
exit;
?>

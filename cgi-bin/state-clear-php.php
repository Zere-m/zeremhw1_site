<?php
// Clear cookies
setcookie("favorite_color", "", time() - 3600, "/");
setcookie("favorite_food", "", time() - 3600, "/");

// Redirect back to view page
header("Location: /cgi-bin/state-view-php.php");
exit;
?>

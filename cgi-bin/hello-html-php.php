<?php

$date = date("Y-m-d H:i:s");
$ip = $_SERVER["REMOTE_ADDR"];
?>

<!DOCTYPE html>
<html>
<head>
  <title>Hello HTML World (PHP)</title>
</head>
<body>

  <h1 align="center">Hello HTML World</h1>
  <hr>

  <p><b>Language:</b> PHP</p>

  <p>This page was generated at: <?php echo $date; ?></p>

  <p>Your IP Address is: <?php echo $ip; ?></p>

</body>
</html>

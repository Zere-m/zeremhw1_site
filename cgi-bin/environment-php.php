<?php


header("Content-Type: application/json");
header("Cache-Control: no-cache");

// server argument has environment and reques
echo json_encode($_SERVER, JSON_PRETTY_PRINT);
?>

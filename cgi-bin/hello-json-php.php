<?php

header("Content-Type: application/json");
header("Cache-Control: no-cache");

$data = [
    "greeting" => "Hello World",
    "language" => "PHP",
    "generated_at" => date("Y-m-d H:i:s"),
    "client_ip" => $_SERVER["REMOTE_ADDR"]
];

echo json_encode($data, JSON_PRETTY_PRINT);
?>

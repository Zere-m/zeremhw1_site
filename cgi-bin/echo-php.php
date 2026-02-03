<?php


header("Content-Type: application/json");
header("Cache-Control: no-cache");


$method = $_SERVER["REQUEST_METHOD"];
$host = $_SERVER["HTTP_HOST"];
$timestamp = date("Y-m-d H:i:s");

$client_ip = $_SERVER["REMOTE_ADDR"];
$user_agent = $_SERVER["HTTP_USER_AGENT"] ?? "Unknown";

$content_type = $_SERVER["CONTENT_TYPE"] ?? "";



$data_received = [];

// case 1: GET 
if ($method === "GET") {
    $data_received = $_GET;
}

// case 2: POST 
else if ($method === "POST") {

    // JSON body
    if (strpos($content_type, "application/json") !== false) {
        $raw_body = file_get_contents("php://input");
        $data_received = json_decode($raw_body, true);

        if ($data_received === null) {
            $data_received = ["error" => "Invalid JSON received"];
        }
    }

    // x-www-form-urlencoded
    else {
        $data_received = $_POST;
    }
}

// case 3: PUT/DELETE 
else if ($method === "PUT" || $method === "DELETE") {

    $raw_body = file_get_contents("php://input");

    if (strpos($content_type, "application/json") !== false) {
        $data_received = json_decode($raw_body, true);

        if ($data_received === null) {
            $data_received = ["error" => "Invalid JSON received"];
        }
    }
    else {
        // Parse urlencoded manually
        parse_str($raw_body, $data_received);
    }
}

// in case of failure
else {
    $data_received = ["error" => "Unsupported HTTP method"];
}



$response = [
    "message" => "Echo Endpoint (PHP)",
    "language" => "PHP",
    "method" => $method,
    "hostname" => $host,
    "timestamp" => $timestamp,
    "client_ip" => $client_ip,
    "user_agent" => $user_agent,
    "data_received" => $data_received
];



echo json_encode($response, JSON_PRETTY_PRINT);

?>

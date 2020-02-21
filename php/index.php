<?php

$base_url = "http://api.wepdf.io/";

//GET request
$curl = curl_init();

curl_setopt_array($curl, array(
    CURLOPT_URL => $base_url . "render?apikey=YOUR_API_KEY&url=https://en.wikipedia.org/wiki/PDF&landscape=true",
    CURLOPT_RETURNTRANSFER => true,
    CURLOPT_HTTPHEADER => array('Content-Type:application/json'),
));

$response = curl_exec($curl);
file_put_contents('wepdf_render.pdf', $response);


//POST request
$curl = curl_init();
$file_render = file_get_contents("YOUR_FILE.html");
curl_setopt_array($curl, array(
    CURLOPT_URL => $base_url . "render?apikey=YOUR_API_KEY&landscape=true",
    CURLOPT_RETURNTRANSFER => true,
    CURLOPT_POST => true,
    CURLOPT_POSTFIELDS => $file_render,
    CURLOPT_HTTPHEADER => array('Content-Type:text/html'),
));

$response = curl_exec($curl);
file_put_contents('wepdf_convert.pdf', $response);
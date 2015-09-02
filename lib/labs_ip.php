<?php
$ip = $_SERVER["HTTP_X_FORWARDED_FOR"];
if ($ip) {
	$ip_list = explode(",", $ip);
	$ip = end($ip_list);
} else {
	$ip = $_SERVER['REMOTE_ADDR'];
}
if ($ip != getenv("LABS_IP")) {
	header('HTTP/1.0 403 Forbidden');
    die('Direct Access Forbidden');
}
?>
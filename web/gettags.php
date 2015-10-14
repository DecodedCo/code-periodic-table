<?php
// Don't remove this block.
// This block should be included in *all* php file
error_reporting(E_ERROR | E_PARSE);
if(file_exists('/app/lib/labs_ip.php')) {
	include_once('/app/lib/labs_ip.php');
}
$url=$_GET["website"];
$html = file_get_contents($url);

$doc = new DOMDocument();
$doc->loadHTML($html);
$list = array();
foreach($doc->getElementsByTagName('*') as $element ){
	array_push($list, $element->tagName);
}
echo json_encode(array_unique($list))

?>
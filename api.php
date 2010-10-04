<?php

header('Content-type: text/javascript');
header("Cache-Control: no-cache, must-revalidate");

$uri = trim($_SERVER['REQUEST_URI']);

$second_indexof = stripos($uri, '/', 1);
$third_indexof = stripos($uri, '/', $second_indexof+1);

$method = substr($uri, 1, $second_indexof-1);
$callback = substr($uri, $second_indexof+1, $third_indexof-$second_indexof-1);
$url = substr($uri, $third_indexof+1);

$source = preg_replace('/(\\n|\\r)/', '\\n', file_get_contents($url));
if($method == "p" || $method == "proxy") {
  // do nothing, it's fine as is
  echo "$callback({\"data\": \"$source\"})";
} else if($method == "r" || $method == "rewrap") {
  // strip the current json kernel off, biznatch
  $paren_indexof = stripos($source, '(');
  $source = substr($source, $paren_indexof+1, strlen($source)-($paren_indexof)-2);
  echo "$callback($source)";
}

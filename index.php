<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN"
"http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">

<html xmlns="http://www.w3.org/1999/xhtml">
  <head>
    <title>jsonp.org: Public JSONP Wrapper Service</title>
    <style type="text/css">
      html, body {
        font-family: Arial;
      }
    </style>
  </head>
  <body>
    <h1>Welcome to jsonp.org: Public JSONP Wrapper Service</h1>
    
    <p>    
    Proxy Usage:
    <pre>
    http://api.jsonp.org/p/&lt;callback&gt;/&lt;url&gt;
    </pre>
    </p>
    <p>
    Rewrapper Usage:
    <pre>
    http://api.jsonp.org/r/&lt;callback&gt;/&lt;url&gt;
    </pre>
    </p>

    <p>
    jQuery <a href="http://jsonp.org/jquery.jsonp.js">jsonp.org plugin</a>:
    <pre>
    $.jsonp("http://www.google.com/robots.txt", function(data) {
      alert("Google's robots.txt: " + data);
    });
    </pre>
    </p>

    <p>
    I created this because I was unable to get several urls working
    with jsonpwrapper.com or YQL. There's absolutely no error checking and I'm constantly changing the "Api" and URL structure, and obviously the jQuery plugin as well. It is intended that this will be stabilized in the future, at which point this message will be removed.</p>
  
  </body>
</html>

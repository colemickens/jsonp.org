from google.appengine.ext import webapp
from google.appengine.ext.webapp.util import run_wsgi_app
import re, urllib, urllib2
import django.utils.simplejson as json
from google.appengine.api.urlfetch import DownloadError
from validate_jsonp import is_valid_jsonp_callback_value

class Handler(webapp.RequestHandler):
  def get(self):
    self.response.headers['Content-Type'] = "text/javascript"

    args = re.match(r"^/(\w+)/(\w+)/(.*)", self.request.path)
    
    if args is None:
      self.response.out.write("Invalid request")
      return
    
    method = args.group(1)
    callback = args.group(2)
    url = urllib.unquote(args.group(3))

    if not is_valid_jsonp_callback_value(callback):
      self.response.out.write("Invalid JSONP Callback")
      return

    try:
      sock = urllib2.urlopen(url)
      content = sock.read()
      sock.close()
      
      if method == 'p':
        data = content
        status = "ok"
      elif method == 'r':
        match = re.match("^(.*)\((.*)\)$", content)
        if match is None:
          data = "Failed to rewrap content"
          status = "bad"
        else:
          data = match.group(2)
          status = "ok"

    except:
      status = "bad"
      data = "Bad URL: %s" % e

    jsondata = json.dumps( { 'status':status, 'data':data } );

    response = '%s(%s)' % (callback, jsondata)

    self.response.out.write(response)
    return

application = webapp.WSGIApplication([
  ("/r/.*", Handler),
  ("/p/.*", Handler),
  ], debug=True
 )

def main():
  run_wsgi_app(application)

if __name__ == "__main__":
  main()

(function( $ ){

jsonpReqCount = 0;

jsonpInternal = function(url, success_cb, rewrap) {
	var data;

	function success() {
			success_cb.call(this, data);
	}

	var jsonpCB  = "jsonp" + (++jsonpReqCount);
	window[ jsonpCB ] = function( tmp ) {
		if(rewrap) {
			data = tmp;
		} else {
			data = tmp.data;
		}

		success();
		window[ jsonpCB ] = undefined;

		try {
			delete window[ jsonpCB ];
		} catch(e) { }

		if (head) {
			head.removeChild(jsonpScript);
		}
	};

	var jsonpScript = document.createElement("script");
	var head = document.getElementsByTagName("head")[0] || document.documentElement;

	if(rewrap == true) {
		jsonpScript.src = "http://jsonp.org/api?cb="+jsonCB+"&method=rewrap&url="+encodeURIComponent(url);
	} else {
		// defaults to method=proxy
		// saves a few characters for url length
		jsonpScript.src = "http://jsonp.org/api?cb="+jsonCB+"+&url=" + encodeURIComponent(url);
	}
	head.insertBefore(jsonpScript, head.firstChild);
	return undefined;
}

jQuery.extend({
	jsonpRewrap: function(url, success_cb) {
		return jsonpInternal(url, success_cb, true);
	},
	jsonpProxy: function(url, success_cb) {
		return jsonpInternal(url, success_cb, false);
	}
});

})( jQuery );

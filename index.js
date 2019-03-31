const React = require('preact');

class App extends React.Component {
  render() {
    return <h1>Hello, World!</h1>;
  }
}

React.render(<App></App>, document.querySelector('#content'));


https://www.taniarascia.com/how-to-connect-to-an-api-with-javascript/
https://scotch.io/tutorials/how-to-use-the-javascript-fetch-api-to-get-data
https://www.w3schools.com/js/js_object_methods.asp

var xmlhttp = new XMLHttpRequest();
	xmlhttp.onreadystatechange = function() {
	if (this.readyState == 4 && this.status == 200) {
    var myArr = JSON.stringify(this.responseText);
    document.getElementById("ID").innerHTML = myArr[0];
		}
	};
xmlhttp.open("GET", "http://localhost:8081/industries", true);
xmlhttp.send();
	
	//============================
	
fetch("http://localhost:8081/industries")
	//.then((resp) => resp.json())
	.then(function(resp)
	{
		var obj = JSON.parse(resp);
		alert(obj);
	}
	)
	alert("jjbj")
	
	
HTTPS
https://blog.kowalczyk.info/article/Jl3G/https-for-free-in-go-with-little-help-of-lets-encrypt.html
https://stackoverflow.com/questions/46522749/how-to-solve-redirect-has-been-blocked-by-cors-policy-no-access-control-allow
https://github.com/mholt/certmagic

Chrome soln: https://chrome.google.com/webstore/detail/allow-control-allow-origi/nlfbmbojpeacfghkpbjhddihlkkiljbi
https://medium.com/@matryer/the-http-handler-wrapper-technique-in-golang-updated-bc7fbcffa702#.e4k81jxd3
    
	
	var dir string
	myRouter.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))
	srv := &http.Server{
        Handler:      myRouter,
        Addr:         "127.0.0.1:8081",
        // Good practice: enforce timeouts for servers you create!
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }
	
	
	func makeServerFromMux(mux *http.ServeMux) *http.Server {
    // set timeouts so that a slow or malicious client doesn't
    // hold resources forever
    return &http.Server{
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 5 * time.Second,
        IdleTimeout:  120 * time.Second,
        Handler:      mux,
    }
}

func makeHTTPToHTTPSRedirectServer() *http.Server {
    handleRedirect := func(w http.ResponseWriter, r *http.Request) {
        newURI := "https://" + r.Host + r.URL.String()
        http.Redirect(w, r, newURI, http.StatusFound)
    }
    mux := &http.ServeMux{}
    mux.HandleFunc("/", handleRedirect)
    return makeServerFromMux(mux)
}
	
	httpSrv := makeHTTPToHTTPSRedirectServer()
    //httpSrv.Handler = m.HTTPHandler(httpServer.Handler)
	httpSrv.Handler = http.Handler(httpSrv.Handler)
	
	Industries = append(Industries, Industry{ID: "11", Name: "Agricultura", Children: &Industry{ID: "111", Name: "Siembra de soya", Children: &Industry{ID: "1111", Name: "Siembra de soya"}}})
	Industries = append(Industries, Industry{ID: "97", Name: "Otros", Children: &Industry{ID: "971", Name: "Otros perecederos", Children: &Industry{ID: "972", Name: "Otros servicios", Children: &Industry{ID: "9721", Name: "Servicios varios"}}}})

	http.HandleFunc("/", homePage)
	http.HandleFunc("/industries", http.HandlerFunc(returnAllArticles))
	http.HandleFunc("/industries/{id:[0-9]+}", http.HandlerFunc(returnSingleArticle))
	http.HandleFunc("/createindustries/{id}", http.HandlerFunc(CreateEntry))
    http.HandleFunc("/deleteindustries/{id}", http.HandlerFunc(DeleteEntry))

 	httpSrv.Addr = ":80"
    fmt.Printf("Starting HTTP server on %s\n", httpSrv.Addr)
    err := httpSrv.ListenAndServe()
    if err != nil {
        log.Fatalf("httpSrv.ListenAndServe() failed with %s", err)
    }

	
    //log.Fatal(srv.ListenAndServe())
	
		
		//magic, err := certmagic.Manage([]string{"localhost"})
	//http.ListenAndServe(":80", magic.HTTPChallengeHandler(myRouter))
	
	//certmagic.HTTPS([]string{"localhost"}, myRouter)
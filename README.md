# GoWebserver
A webserver (somewhat similar to my Python webserver, but with less features and with better performance) written in Go/GoLang (It's even my first project written in Go/GoLang). That supports HTTP and HTTPS, to host a website on your device or local network.

# Instructions
Run the server with 'go run main.go' or 'go build -o <output_filename> main.go'. And it will configure the server from the configuration data.

# Configuration
The server would run from reading the 'server_config.json' file. It would configure and bind the server from what's configured on the configuration file.

* "use_global_ip": Set to true or false, if you want the server to listen on all interfaces (localhost, lan, etc.).
* "server_ip": Specify an IP address (localhost or the IP address of your device) to bind the server with [Required if 'use_global_ip' is set to false].
* "server_port": Specify a port number for the server to bind to [Required].
* "server_protocol": Specify 'HTTP' or 'HTTPS' [Required].
* "https_cert": Specify the filename of the HTTPS certificate [Required if using HTTPS].
* "https_key": Specify the filename of the HTTPS key [Required if using HTTPS].
* "html_doc": Specify the filename of an HTML Docuement to load [Not Recquired, if the main HTML document is named 'index.html'].

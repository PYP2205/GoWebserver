/*
Go Webserver

Programmed by: Paramon Yevstigneyev

Programmed in: Go v1.19.3 (64-Bit/AMD64)

Description:
A simple webserver for hosting any web apps or projects, that is cross platform and easy to configure.
In the .json file provided, you just set things such as the webserver's ip address, port number, protocol,
https key and certificate file (recquired if the protocol is set to HTTPS), and the name of the HTML document
(If none is specified, then it will default to 'index.html')/
*/

package main

import (
	"fmt"
	// Used to get the hostname of the user's machine
	"net"
	// Used to parse the configuration file
	"encoding/json"
	// Used to read bytes of the configuration file stream
	"io/ioutil"
	// Used to log any potential errors
	"log"
	// Used to make a HTTP or HTTPS webserver, and listen for requests
	"net/http"
	// Used to get the IP address of the machine,
	// Used to open the configuration file
	"os"
	// Used to compare strings and capitalize them when reading the configuration file (so things aren't case senstive)
	"strings"
	// Used to get the time of the request when it was sent by a user.
	"time"
)

// A struct for storing the configuration data from the .json file.
type ServerProfile struct {
	Use_Global_IP bool `json:"use_global_ip"`
	Server_IP string `json:"server_ip"`
	Server_Port int `json:"server_port"`
	Server_Protocol string `json:"server_protocol"`
	HTTPS_Cert string `json:"https_cert"`
	HTTPS_Key string `json:"https_key"`
	HTML_Doc string `json:"html_doc"`
}
var Profile ServerProfile

// A function for parsing the configuration data set by the user.
func Parse_Config(config_file string) {
	config, err := os.Open(config_file)
	if err != nil {
		log.Fatal(err)
	}
	byte_val, _ := ioutil.ReadAll(config)
	// Stores the configuration data into the struct.
	json.Unmarshal(byte_val, &Profile)
	config.Close()
}

// A function for handling HTTP and HTTPS requests.
func RequestHandler(w http.ResponseWriter, r *http.Request)  {
	// Prints the protocol, method, ip address and port number of the client, the time the request was sent, and the request data.
	log.Printf("\n\t%s %s %s -- [%s] %s\n", r.Proto, r.Method, r.RemoteAddr, time.Now(), r.URL)
	// Serves the HTML document file
	http.ServeFile(w, r, r.URL.Path[1:])
}

func Server_URL() string {
	if Profile.Use_Global_IP {
		user_hostname, err := os.Hostname()
		if err != nil {
			log.Fatalf("\nError: %e\n", err)
		}
		host_ip, err := net.LookupIP(user_hostname)
		if err != nil {
			log.Fatalf("\nError: %e\n", err)
		}
		server_addr := fmt.Sprintf("0.0.0.0:%d", Profile.Server_Port)
		fmt.Printf("\n\tGo Webserver:\n\tListening on %d\n\t%s://localhost:%d\n\n", Profile.Server_Port, strings.ToLower(Profile.Server_Protocol), Profile.Server_Port)
		for i := 0;i < len(host_ip);i++ {

			// If a host address is an IPv4 Address, then it will print the url.
			if host_ip[i].To4() != nil {
				fmt.Printf("\n\t%s://%s:%d\n\n", strings.ToLower(Profile.Server_Protocol), host_ip[i], Profile.Server_Port)
			
			// If a host address is an IPv6 Address, then it will print the url with [] around the IP address.
			} else {
				fmt.Printf("\n\t%s://[%s]:%d\n\n", strings.ToLower(Profile.Server_Protocol), host_ip[i], Profile.Server_Port)
			}
		}
		return server_addr
	} else {
		var ip_add net.IP = net.IP(Profile.Server_IP)
		if ip_add.To4() != nil {
			server_addr := fmt.Sprintf("%s:%d", Profile.Server_IP, Profile.Server_Port)
		fmt.Printf("\n\tGo Webserver:\n\tListening on %d\n\t%s://%s\n\n", Profile.Server_Port, strings.ToLower(Profile.Server_Protocol), server_addr)
		return server_addr
		} else {
			server_addr := fmt.Sprintf("[%s]:%d", Profile.Server_IP, Profile.Server_Port)
			fmt.Printf("\n\tGo Webserver:\n\tListening on %d\n\t%s://%s\n\n", Profile.Server_Port, strings.ToLower(Profile.Server_Protocol), server_addr)
			return server_addr;
		}
		

	}
}
func main() {
	// Parses the configuration file and stores its data onto the struct.
	Parse_Config("server_config.json")
    http.HandleFunc("/", RequestHandler)
	switch (strings.ToUpper(Profile.Server_Protocol)) {
		
		// If the user specifies 'HTTP' as the server protocol, then it would start the server and listen for HTTP Requsts.
		case "HTTP":
			server_addr := Server_URL()
			log.Fatal(http.ListenAndServe(server_addr, nil))
		
		// If the user spcifies 'HTTPS' as the server protocol, then it will start the server and listen for HTTPS Requests.
		case "HTTPS":
			server_addr := Server_URL()
			log.Fatal(http.ListenAndServeTLS(server_addr, "cert.pem", "key.pem", nil))
			
		// If the user enters an invalid protocol, then it will log the error as fatal onto the console, and end the program.
		default:
			log.Fatalf("\nError: %s is not a valid protocol\n", Profile.Server_Protocol)
	}
}

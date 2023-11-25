##### Author: Ali Umar
##### Date: 2023-11-28
##### Version: 1.0
##### Description: Some short little build scripts to get me going as I am developing on windows 

.PHONY : web server clean

web : 
	cd app && http-server -p 9000 -s -o . -c-1

server : 
	cd server && go run main.go serve

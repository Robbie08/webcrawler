# webcrawler
This webcrawler application will scrape links from a given source, cache them and then repeat the process. 

### Getting Started
* First make sure you have Go set up on your machine. You can [download Go by clicking on this link](https://golang.org/doc/tutorial/getting-started) and following the instruction. 
* Once Go is installed on your machine, place this repo under the <b>src directory</b> in your <b>go PATH</b>.

### Running the App
* Use the command line to cd into the webcrawler directory and type the command 
`make build` to build and create the executable.
* Upon success you can now run the application by using the command `make run`.
* By default a sever will spin up on `localhost:8080`

### Using the App
This Application is currently under v0.1 and this crawler can only scrape through directed websites we set.

* Once the application is running, you can open up [Postman](https://www.postman.com/) or a web browser.
* To Access The Home Page: `http://localhost:8080/`
* To Spin up Crawler : `http://localhost:8080/run`
* To Shutdown Server Gracefully: `http://localhost:8080/shutdown`

Results from webcrawler will appear in your terminal where the server is running.


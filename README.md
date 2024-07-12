# week11_assignment_group_D
setting up environment
     used https://golang.org/doc/install for go language
     https://git-scm.com/book/en/v2/Getting-Started-Installing-Git for git 
     and https://docs.docker.com/get-docker/ for docker set up as our installation guide 

creating the API server
     initialized a new Go module by running the command "go mod init week11_assignment_group_D"
     created a "main.go"
        used "mux := http.NewServeMux()" to create a new HTTP request multiplexer (mux). A mux is like a router that directs incoming HTTP requests to different handler functions based on the URL path.
        "log.Println("Starting server on :3001")" tells the server is starting on port 3001.
        "if err != nil { log.Fatalf("Error starting server: %s\n", err) }:" if there is an error starting the server then above mentioned error message is displayed and the program terminates.

version control and git
     created a github repository "week11_assignment_group_D" 
     then cloned it to our systems 
     after cloning made the required changes like initialing go module, creating an api server and dockerfile
     staged the changes and then pushed them from our pcs

create a dockerfile 
     created a dockerfile 
         FROM golang:1.22 this commands sets the version of go language that will be used as base in a docker image
         WORKDIR /app this commands sets a folder inside docker which will be containing all the other files used in creating image
         COPY . . this command copies all the files in the current repository to /app folder
         RUN command compiles the code and create a working file "main"
         EXPOSE 3001 tells the docker which port will be used
         CMD ["./main"] command tells the docker to run which file to run when the container is executed
build and push docker image
     used "docker build -t mariyabenson/dockergoassignment4:01 ." to create an image 
     used "docker run -p 3001:3001 mariyabenson/dockergoassignment4:01" to run the image

# Minifyr - a simple URL shortener!
## postgreSQL Docker image
This project uses postgres:14, other versions MAY work. <br>
Other versions might work.<br>
To setup the docker image simply run the following command:<br>
```$ docker run --name [NAME-OF-CONTAINER] -e POSTGRES_USER=[ENTER-USER-HERE] -e POSTGRES_PASSWORD=[ENTER-PWD-HERE] -d postgres:14```<br>
```-name``` flag is optional, by leaving it out Docker will randomly generate an UUID
## Projects and Packages
This project uses the same repository for backend and frontend and as such two sets of packages need to be installed.<br>
Navigate to the desired folder and run ```npm install``` for both and you should be good to start using the app.<br>
<br>
If there are some GO packages missing you can run ```go mod tidy``` trough the console while inside the backend repository.

# BookLibrary 
> <i>Your local of  books library for home</i>
<br>

Usage
---
> <b>local_ip:port</b> <i>is used for requests.</i>

- ```http://local_ip:port``` <i> <-- web interface</i>
- ```curl -X GET local_ip:port/books``` <-- <i>show all books</i>
- ```curl -X POST local_ip:port/books/upload \ -F "file@=name.txt"``` <i><-- upload a file</i>
- ```curl -X GET local_ip:port/ping``` <i> <--check status</i>

Installation
--
> <i>You will need to install golang to run the project.</i><br>
> <i>You can download it <a href="https://go.dev/doc/install">here</a></i>

* <i>Clone the repository</i><br>
```git clone https://github.com/upsetting-god/BookLibrary.git```

* <i>Go to the folder</i><br>
```cd BookLibrary```

* <i>Launch the project</i><br>
```go run main.go```


Settings
---

<b>You can configure the server settings in the config.yaml file.</b>

```yaml
server:
  port: 8080 # Specify the port you want to use, default is 8080

allowed_ex: # Allowed files extensions
  - ".pdf"
  - ".txt"
  - ".html"
  - ".fb2"
```
Screenshots
---


<i>Server<img width="860" height="803" alt="dark_theme" src="https://raw.githubusercontent.com/upsetting-god/BookLibrary/main/pic/program.pic"/></i>
<br>
<br>
<i>Light Theme<img width="860" height="803" alt="dark_theme" src="https://raw.githubusercontent.com/upsetting-god/BookLibrary/main/pic/light_theme.png"/></i>
<br>
<br>
<i>Dark Theme<img width="860" height="803" alt="dark_theme" src="https://raw.githubusercontent.com/upsetting-god/BookLibrary/main/pic/dark_theme.png"/></i>














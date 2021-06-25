# Embedding static files

Starting with Go 1.16 static files can be embedded directly into the built binary. This example demonstrates how to achieve that with a modern frontend build system by embedding the files in [frontend/build/](frontend/build/) and serving them from a web server.

You can try it by running the code and heading to http://localhost:8080/index.html

[See the code &raquo;](embed.go)

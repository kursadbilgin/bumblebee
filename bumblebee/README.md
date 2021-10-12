# Running & Tests
- For running application `go run main.go` 
    - Application will start at port 8080
- For running test, `go test ./Test -v`

# Technical Details
- Golang/Gin is used as application framework
- Redis is caching web link and deep link

# API Details
```sh 

POST /v1/api/convert-link-to-deeplink          # Converts link to deep link
POST /v1/api/convert-deeplink-to-link          # Converts deep link to link

```
- You can find Postman export for sample requests in `postman_collection.json`  or `swagger.yaml`
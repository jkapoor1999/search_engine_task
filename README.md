# search_engine_task

## *STEP 1: CLONE THE REPOSITORY*
`git clone https://github.com/jkapoor1999/search_engine_task`

## *STEP 2: RUNNING THE APPLICATION*
```
cd "search_engine_task"

make build

make up
```
## *STEP 3: TESTING THE APPLICATION*
Use any API testing applications such as Postman, Insomnia or ThunderClient

All available routes
1. To `insert` a webpage in MongoDB : POST Request**
    `http://localhost:4000/v1/insert`
     Add this to the JSON body
```
      {
        "title": "Page 1",
        "keywords": [
          "Ford",
          "Car",
          "Review"
        ]
      }
 ```
 2.To `get` back the computed answer from MongoDB : GET Request
     `http://localhost:4000/v1/get`
     Add this in the JSON body
```
     {    
       "user_keywords": [
                "Ford",      
                "Review"    
              ]
      }
```
3. To perform a `healthcheck` on the server: GET Request
    `http://localhost:4000/v1/get`
    
## *STEP 4: TO STOP THE APPLICATION*
`make down`
      
    
    

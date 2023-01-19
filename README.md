# search_engine_task

## *STEP 1: CLONE THE REPOSITORY*
`git clone https://github.com/jkapoor1999/search_engine_task`

## *STEP 2: RUNNING DOCKER*
```
cd "search_engine_task"

docker-compose build

docker-compose up
```
## *STEP 3: TESTING API*
Use any API testing applications such as Postman, Insomnia or ThunderClient

All available routes
1. To store webpage in MongoDB : POST Request**
    `http://localhost:4000/v1/savepage`
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
 2.To get back the computed answer from MongoDB : GET Request
     `http://localhost:4000/v1/getresult`
     Add this in the JSON body
```
     {    
       "user_keywords": [
                "Ford",      
                "Review"    
              ]
      }
```
## *STEP 4: TO STOP THE CONTAINER*
`docker-compose down`
      
    
    

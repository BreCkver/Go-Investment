# Title: Investment
## Overview: 

When we have financial resources from new investors, we like to ask ourselves what is the best option to invest them.
Which is the best option to invest them. We have 3 amounts of credits that we give to our customers clients ($300, $500 and $700). When the investment money comes in, we want to determine  how many credits of each amount we could allocate with that money, without having 1 peso left over.


#### Use Cases
The app coverage the next use cases..
* The investment's amount is a multiple of 100 
* If it is not possible to assign a value, all values must receive to zero
* Summary Stored in database
* Get the applications' summary
* Services folder has a unit test with coverage from 96.7%

#### Out of Scope
* If it is not possible to assign a value, app doesn't show details
* Summary doesn't show historical
* Summary is global, no by user
* Integration test
* GoRoutines to improve the performance 
* Methods description for major understanding

## Project structure
The architecture in this moment is simple, so it has the next structure:
data                -> The files contain methods for interacting with databases.
handlers            -> The file contain handler, ports
services            -> pThe entire business logic of the application.
models              -> program models 
api                 -> Here we store http-server settings, etc.

## Libraries used directly
- Gorilla           -> Handler request
- Mongo-driver      -> Mongo access


## Prerequisites
1. Git
2. Docker
3. Docker Compose to user Linux 

## How to compile and run 
1. Make a new direcotory
2. Move to new directory
3. Donwload code with git, such as:  
    git clone https://github.com/BreCkver/Go-Investment.git
4. Open PowerShell or shell, in it move to new directory 
5. Execute the commands:
	docker-compose up
6. Access to project from http://localhost:8081/


## Publish
- The app was published in heroku, the link is:
    https://go-investment.herokuapp.com/


## Examples
    ### `credit-assignment`
        - url: https://go-investment.herokuapp.com/credit-assignment
        - body
        ```json 
                {
                    "investment" : 1500
                }
        ```
    
    ### `statistics`
        - url: https://go-investment.herokuapp.com/statistics
        - body: empty
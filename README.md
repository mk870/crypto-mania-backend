# Crypto Mania  Server (Backend)
<img src="https://i.ibb.co/51PDVwk/gin.png" alt="gin" border="0" width="400" align="center"> 

## Project Summary 
* This  is a Restful Gin application with CRUD operations that allow users to create an account to the cryptomania app, save or add crypto coins to a personal watchlist, delete those crypto coins and login and out using Go jwt authentication.
* It uses  Gin framework.
* The app has 8 endpoints namely : /user, /logout, /verifification, /refresh-token, /login, /coin, /coins,and /coin/id.
* Uses Go Jwt to secure these endpoints.
* Uses Gin GORM to persist data to a postgreSQL database.


### **Resources Used**
***
**Go Version**: 1.19.4

**Dependencies**: Jwt Token, GORM, Go Mail, GIN, postgreSQL-Driver, Go-Cors and Google uuid.  
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=flat&logo=go&logoColor=white) 	![JWT](https://img.shields.io/badge/JWT-black?style=flat&logo=JSON%20web%20tokens) 	![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=flat&logo=postgresql&logoColor=white)

**For Web Framework Requirements**: go.mod

**APIs**: None

### **EndPoints Building**
***
Built 4 Controllers, one for user account creation and the other for user data interaction.
#### **User Account Creation Endpoints:** 
* **/user (PostMethod)**: Takes in firstname, lastname, password and email for user signup. A Jwt token is created as an authentication tool, its stored on the database and also sent by java mail to user email for verification. The password is encrypted using BCryptPasswordEncoder.

* **/verification  (GetMethod)**: validates the email token against the one on the database, once verified the account is enabled. 
* **/login  (GetMethod)**: A Jwt token is created and returned if user login credentials are valid. 

#### **User Data Interaction Endpoints:**  
* **/coin (PostMethod)**:  saves user crypto coins to the database with all the coin's features like marketCap, alltime highs, rank and number of coins in circulation. 
* **/coins (GetMethod)**:  retrieves all the saved crypto coins of a client from the database.
* **/coin/id (DeleteMethod)** : deletes a specific crypto currency by id from the database.  

### **Data Storage**
Used GORM (ORM) to persist and retrieve data from a postgreSQL database.  
Built 3 models: 
* User Model to store app users.
* VerificationToken Model to store signup verification tokens.
* Coin Model to store crypto coins of users. 



### **Productionization**
***
In this step I deployed the postgreSQL database to AWS via 3rd party and deployed the Gin app to Railway Cloud.

**Live Implemantation:** [Crypto-Mania](https://crypto-mania-react.vercel.app/)
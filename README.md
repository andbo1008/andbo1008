Test bank application 

For convenience, all URLs associated with the user were taken to the "user" group, and to work with the account and transactions, they were taken to the "account" nested group.

Using an example, I will show which routes are responsible for what:
USER
For create a user need to send POST request "http://localhost:1313/user" in json file:
         name, lastname, email(will be unique), password
For get a user need to send GET request  "http://localhost:1313/user/id our user"

ACCOUNT
For create an account to send POST request "http://localhost:1313/user/id/account" (
    json file :
    currency:only 1 currency "USD" , "MXN" ,"COP" , 
    total : default (0.00 ))
    
For get all user accounts need to send GET request "http://localhost:1313/user/id/account"

Transaction
To make a transaction need add a "transaction" to the url
POST "http://localhost:1313/user/id/account/transaction"
json: email :(for find someone whoь to send currency),
      currensy:  1 currency "USD" , "MXN" ,"COP",
      total: wich sum sended


To view the list of transactions
Send GET request "http://localhost:1313/user/id/account/transaction"
And you can see all transactions what do user




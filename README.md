# Splitwise
Sharing is caring

## Table of Contents
* [Team Members](#team-members)
* [Functional Requirement](#Funtional-Requirement)
* [Scope](#Scope)
* [Design Diagram](#Design Diagram)
* [REST endpoints](#REST endpoints)
* [Build And Run](#Build And Run)
* [Test Data](#Test Data)
* [Possible Improvements](#Possible Improvements)

### <a name="team-members"></a>Team Members
* "Hemant Yadav" <github# hyadav5>

### <a name="Funtional-Requirement"></a>Functional Requirement
Below are the functional requirement expected from the application:
1) Ability to add user
2) Ability to add group
3) Ability to add expense
4) Ability to add group expense
5) Run settlement
### <a name="Scope"></a>Scope
Application support all the above mentioned functional requirement. The current design is limited with the below scopes:
1) All the resources like User, Group, Expense have ID field, but it is not used. This ID field can be used once we have database or front end in place.
2) Exhustive error handling is not done on the http layer. As it was out of scope.
3) Endpoints are not REST compliant.
4) REST endpoints are exposed by the application just to ease the job of end user to make requests through curl/Postman.
5) Split is only limited to EQUAL division.

### <a name="Design Diagram"></a>Design Diagram
![Alt text](./documentation/splitwise.jpg?raw=true "Title")

### <a name="REST endpoints"></a>REST endpoints
**/adduser**  = Add a user

Body: {
"Name": "sundar",
"Contact": "700097246",
"Email": "sundar@gmail.com"
}

**/addgroup** = Add a group

{
"Name": "mandli",
"Users": ["hemant", "shyam", "pandey"]
}

**/users** = Show all users

**/addexpense** = Add an expense

{
"Type": "EQUAL",
"Amount": 50000,
"PaidBy": "pandey",
"Among": ["hemant", "shyam"]
}

**/addgroupexpense** = Add a group expense

Body: {
"Amount": 5000,
"PaidBy": "pandey",
"GroupName": "flatmates"
}

**/runsettlements** = Run settlement


### <a name="Build And Run"></a>Build And Run
Available make commands:

    make clean

    make build

    make test

Default:

    make
**Run:**

The generated binary can be run simply like a script. It will open a REST serving port at 8080.

    ./splitwise
### <a name="Test Data"></a>Test Data
Once the application run, its is loaded with the below test data.
1) Add 4 users: hemant, kd, pandey, shyam
2) Add a group with the name: "flatmates"
3) Add an expense of 1000 paid by "hemant" and distributed among above 4 equally.
4) Add an expense of 1000 paid by "kd" and distributed among above 4 equally.
5) Add a group expense of 2000 paid by "hemant" and distributed among group member.
5) Run settlement detailing the pending amount among them.

### <a name="Possible Improvements"></a>Possible Improvements
In the interest of time, following design consideration are ignored:
1) We can use Swagger UI and REST endpoints, this will ease our http request and response handling.
2) A middleware should have been added before the APIs to take care of request validation, authentication and authorization.
3) Different split types can be added like with exact amount, percentage, shares etc.
4) A ledger can be added to keep track of active expenses and archived expenses history.







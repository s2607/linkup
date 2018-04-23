#Linkup

A small microservice for accepting and analysing responses to simple questions in order to recomend charities. 

##Minor notes

The criterion object is being replace with the "validresponse" object and here is what that looks like: 
Valid response objects share a global key pool and other objects will have a relation table to link them to their valid responses.
services will have a list of validresponses that must be selected, in the future it would be good to invert this because if admin users add validresponses to a question then a manually inverted list may not catch everything it was supposed to.

Automated tests are missing for some of the objects and features that where added later on. There are no automated UI tests but the UI was written entierly without the use of javascript which makes it easy to write them with shell and curl (or your favorite tools.) Use curl -X to submit form fields and don't forget to keep a cookie jar.

See the SDS for how to derive the ERD from the UML class diagram, some objects will not store all their fields in order to prevent race conditions that would be created by the "service" interface being used at the same time as the "admin" interface. 

##valid response

question -> valid responses replace criterion collection (validation criteria)

PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
DROP TABLE service;
CREATE TABLE service (key integer primary key asc,name string,description string,url string);
INSERT INTO service VALUES(2,'SNAP','Supplemental Nutrition Assistance Program','https://www.fns.usda.gov/snap/');
INSERT INTO service VALUES(4,'Liberty','college','');
DROP TABLE criterion;
CREATE TABLE criterion (inv bool,conj bool,qkey int,key integer primary key asc,isnil bool,aval int,lval bool,qtype int,bval int,regex string,dec bool,exc bool,pos bool,apresent boo, bpresent bool);
INSERT INTO criterion VALUES(0,0,0,1,0,0,0,NULL,0,'male|female|other',0,0,0,0,0);
INSERT INTO criterion VALUES(0,1,5,2,0,0,0,NULL,26124,'',0,0,0,0,0);
INSERT INTO criterion VALUES(1,1,8,3,0,0,0,NULL,1,'',0,0,0,0,0);
INSERT INTO criterion VALUES(0,1,20,4,0,0,1,NULL,0,'',0,0,0,0,0);
INSERT INTO criterion VALUES(0,0,0,5,0,0,0,NULL,0,'freshmen|sophomore|junior|senior',0,0,0,0,0);
INSERT INTO criterion VALUES(0,0,0,6,0,1,0,NULL,5,'',0,0,0,0,0);
INSERT INTO criterion VALUES(0,0,0,7,0,1,0,NULL,5,'',0,1,0,0,0);
INSERT INTO criterion VALUES(1,0,0,8,0,1,0,NULL,5,'',0,0,0,0,0);
DROP TABLE servicescriterion;
CREATE TABLE servicescriterion (okey int,ikey int, PRIMARY KEY ( ikey, okey ));
INSERT INTO servicescriterion VALUES(2,2);
INSERT INTO servicescriterion VALUES(2,3);
INSERT INTO servicescriterion VALUES(4,4);
DROP TABLE questionscriterion;
CREATE TABLE questionscriterion (okey int,ikey int, PRIMARY KEY ( ikey, okey ));
INSERT INTO questionscriterion VALUES(3,1);
INSERT INTO questionscriterion VALUES(25,5);
DROP TABLE servicesquestion;
CREATE TABLE servicesquestion (okey int,ikey int, PRIMARY KEY ( ikey, okey ));
INSERT INTO servicesquestion VALUES(2,5);
INSERT INTO servicesquestion VALUES(2,8);
INSERT INTO servicesquestion VALUES(4,20);
INSERT INTO servicesquestion VALUES(4,21);
INSERT INTO servicesquestion VALUES(4,25);
DROP TABLE question;
CREATE TABLE question (key integer primary key asc,prompt string,qtype int);
INSERT INTO question VALUES(2,'Are you married?',2);
INSERT INTO question VALUES(3,'Do you own a vehical?',2);
INSERT INTO question VALUES(4,'What is your median income?',1);
INSERT INTO question VALUES(5,'What is your yearly household income? (please round to integer dollars)',1);
INSERT INTO question VALUES(6,'Do you receive any other support from other beneficiaries?',2);
INSERT INTO question VALUES(7,'Do you have any dependants?',2);
INSERT INTO question VALUES(8,'How many children do you have?',1);
INSERT INTO question VALUES(9,'How old is your youngest child?',1);
INSERT INTO question VALUES(10,'How old is your oldest child?',1);
INSERT INTO question VALUES(11,'How many people live with you?',1);
INSERT INTO question VALUES(12,'Which country are you  a primary resident of?',0);
INSERT INTO question VALUES(13,'Do you have a job?',2);
INSERT INTO question VALUES(14,'Have you served in the military?',2);
INSERT INTO question VALUES(15,'What branch of the military have you served in?',0);
INSERT INTO question VALUES(16,'How many years have you served in the military',1);
INSERT INTO question VALUES(17,'Do you own a home?',2);
INSERT INTO question VALUES(18,'What is the value of your home?',1);
INSERT INTO question VALUES(19,'How much is the mortgage on the home?',2);
INSERT INTO question VALUES(20,'Do you go to school here?',2);
INSERT INTO question VALUES(21,'Are your grades good?',2);
INSERT INTO question VALUES(25,'what is your class?',0);
INSERT INTO question VALUES(26,'what is your age?',1);
COMMIT;
/*

 File: ./SQLyounitsqldb.SQL
 Date: 25-Jun-2023
 Author: D#

 CREATE SCHEMA younitschema under younitdb

 */

/*
 #1 complex (root)
*/


/* CREATE SCHEMA younitschema; */


CREATE TABLE IF NOT EXISTS complex (
	id MEDIUMINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	name          VARCHAR ( 255 ) NOT NULL,
	streetnumber  VARCHAR ( 255 ) NOT NULL,
	streetname    VARCHAR ( 255 ) NOT NULL,
	upnumber      VARCHAR ( 255 ) UNIQUE NOT NULL,

	recordversion INT NOT NULL,
	createdBy VARCHAR ( 255 ) NOT NULL,
	updatedBy VARCHAR ( 255 ) NOT NULL,
	updatedOn TIMESTAMP NOT NULL DEFAULT current_timestamp,
	createdOn TIMESTAMP NOT NULL DEFAULT current_timestamp
);

CREATE TABLE IF NOT EXISTS unit (
	id MEDIUMINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	FKcomplexID   MEDIUMINT   NOT NULL,
	name          VARCHAR ( 255 ) NOT NULL,
	streetnumber  VARCHAR ( 255 ) NOT NULL,
	streetname    VARCHAR ( 255 )  NOT NULL,

    FOREIGN KEY (FKcomplexID)
      REFERENCES complex (id),

	recordversion INT NOT NULL,
	createdBy     VARCHAR ( 255 ) NOT NULL,
	updatedBy     VARCHAR ( 255 ) NOT NULL,
	updatedOn     TIMESTAMP NOT NULL DEFAULT current_timestamp,
  createdOn TIMESTAMP NOT NULL DEFAULT current_timestamp
);


/*
 #3 resident (root)
*/
CREATE TABLE IF NOT EXISTS resident (
	id MEDIUMINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	firstName VARCHAR ( 255 )  NOT NULL,
	lastName VARCHAR ( 255 ) NOT NULL,
	dateofBirth VARCHAR ( 255 ) NOT NULL,

	recordversion INT NOT NULL,
	createdBy VARCHAR ( 255 ) NOT NULL,
	updatedBy VARCHAR ( 255 ) NOT NULL,
	updatedOn TIMESTAMP NOT NULL DEFAULT current_timestamp,
	createdOn TIMESTAMP NOT NULL DEFAULT current_timestamp
);

/*
 #4 unitresident
*/
CREATE TABLE IF NOT EXISTS unitresident (
	id MEDIUMINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	unitresidentID serial PRIMARY KEY,
	FKresidentID serial NOT NULL,

	FOREIGN KEY (FKresidentID)
			REFERENCES resident (id),

	recordversion INT NOT NULL,
	createdBy VARCHAR ( 255 ) NOT NULL,
	updatedBy VARCHAR ( 255 ) NOT NULL,
	updatedOn TIMESTAMP NOT NULL DEFAULT current_timestamp,
	createdOn TIMESTAMP NOT NULL DEFAULT current_timestamp

);


/*
 #5 landlord (root)
*/
CREATE TABLE IF NOT EXISTS landlord (
	id MEDIUMINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	landlordID serial PRIMARY KEY,
	firstName VARCHAR ( 255 )  NOT NULL,
	lastName VARCHAR ( 255 )  NOT NULL,
	dateofBirth VARCHAR ( 255 )  NOT NULL,

	recordversion INT NOT NULL,
	createdBy VARCHAR ( 255 ) NOT NULL,
	updatedBy VARCHAR ( 255 ) NOT NULL,
	updatedOn TIMESTAMP NOT NULL DEFAULT current_timestamp,
	createdOn TIMESTAMP NOT NULL DEFAULT current_timestamp
);


/*
 #6 landlordownunit
*/
CREATE TABLE IF NOT EXISTS landlordownunit (
	id MEDIUMINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	landlordownunit serial PRIMARY KEY,
	FKlandlordID INT NOT NULL,
	FKunitID INT NOT NULL,

	FOREIGN KEY (FKlandlordID)
			REFERENCES landlord (id),
	FOREIGN KEY (FKunitID)
			REFERENCES unit (id),

	recordversion INT NOT NULL,
	createdBy VARCHAR ( 255 ) NOT NULL,
	updatedBy VARCHAR ( 255 ) NOT NULL,
	updatedOn TIMESTAMP NOT NULL DEFAULT current_timestamp,
	createdOn TIMESTAMP NOT NULL DEFAULT current_timestamp

);


/*
 #7 propertymanager (root)
*/
CREATE TABLE IF NOT EXISTS propertymanager (
	id MEDIUMINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	firstName VARCHAR ( 255 )  NOT NULL,
	lastName VARCHAR ( 255 )  NOT NULL,
	dateofBirth VARCHAR ( 255 )  NOT NULL,

	recordversion INT NOT NULL,
	createdBy VARCHAR ( 255 ) NOT NULL,
	updatedBy VARCHAR ( 255 ) NOT NULL,
	updatedOn TIMESTAMP NOT NULL DEFAULT current_timestamp,
	createdOn TIMESTAMP NOT NULL DEFAULT current_timestamp
);


/*
 #8 stratamanager (root)
*/
CREATE TABLE IF NOT EXISTS stratamanager (
	id MEDIUMINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	firstName VARCHAR ( 255 )  NOT NULL,
	lastName VARCHAR ( 255 )  NOT NULL,
	dateofBirth VARCHAR ( 255 )  NOT NULL,

	recordversion INT NOT NULL,
	createdBy VARCHAR ( 255 ) NOT NULL,
	updatedBy VARCHAR ( 255 ) NOT NULL,
	updatedOn TIMESTAMP NOT NULL DEFAULT current_timestamp,
	createdOn TIMESTAMP NOT NULL DEFAULT current_timestamp
);

/*
 #9 communitytitlescheme
*/
CREATE TABLE IF NOT EXISTS communitytitlescheme (
	id MEDIUMINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR ( 255 )  NOT NULL,

	recordversion INT NOT NULL,
	createdBy VARCHAR ( 255 ) NOT NULL,
	updatedBy VARCHAR ( 255 ) NOT NULL,
	updatedOn TIMESTAMP NOT NULL DEFAULT current_timestamp,
	createdOn TIMESTAMP NOT NULL DEFAULT current_timestamp
);

/*
 #10 complexassignedstratamanager
*/
CREATE TABLE IF NOT EXISTS complexassignedstratamanager (
	id MEDIUMINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	FKcomplexID INT NOT NULL,
	FKstratamanagerID INT NOT NULL,
	startdate TIMESTAMP NOT NULL,
	enddate TIMESTAMP,

	FOREIGN KEY (FKcomplexID)
			REFERENCES complex (id),
	FOREIGN KEY (FKstratamanagerID)
			REFERENCES stratamanager (id),

	recordversion INT NOT NULL,
	createdBy VARCHAR ( 255 ) NOT NULL,
	updatedBy VARCHAR ( 255 ) NOT NULL,
	updatedOn TIMESTAMP NOT NULL DEFAULT current_timestamp,
	createdOn TIMESTAMP NOT NULL DEFAULT current_timestamp


);

/*
 #11 landlordpropertymanagerunit
*/
CREATE TABLE IF NOT EXISTS landlordpropertymanagerunit (
	id MEDIUMINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	FKlandlordID                    INT NOT NULL,
	FKpropertymanagerID             INT NOT NULL,
	FKunitID                        INT NOT NULL,

	startdate TIMESTAMP,
	enddate TIMESTAMP,

	FOREIGN KEY (FKlandlordID)
		REFERENCES landlord (id),
	FOREIGN KEY (FKpropertymanagerID)
		REFERENCES propertymanager (id),
	FOREIGN KEY (FKunitID)
		REFERENCES unit (id),

	recordversion INT NOT NULL,
	createdBy VARCHAR ( 255 ) NOT NULL,
	updatedBy VARCHAR ( 255 ) NOT NULL,
	updatedOn TIMESTAMP NOT NULL DEFAULT current_timestamp,
	createdOn TIMESTAMP NOT NULL DEFAULT current_timestamp


);

/*
 #12 stratamanagercommunitytitlescheme
*/


CREATE TABLE IF NOT EXISTS stratamanagercommunitytitlescheme (
	id MEDIUMINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	FKcommunitytitleschemeID                INT NOT NULL,
	FKstratamanagerID                       INT NOT NULL,

	startdate TIMESTAMP,
	enddate TIMESTAMP,

	FOREIGN KEY (FKcommunitytitleschemeID)
		REFERENCES communitytitlescheme (id),
	FOREIGN KEY (FKstratamanagerID)
		REFERENCES stratamanager (id),

	recordversion INT NOT NULL,
	createdBy VARCHAR ( 255 ) NOT NULL,
	updatedBy VARCHAR ( 255 ) NOT NULL,
	updatedOn TIMESTAMP NOT NULL DEFAULT current_timestamp,
	createdOn TIMESTAMP NOT NULL DEFAULT current_timestamp

);


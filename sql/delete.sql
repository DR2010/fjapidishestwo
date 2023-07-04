/*
 there is a sequence when deleting tables
*/

MySQL 9.0 - way to grant stuff
GRANT ALL PRIVILEGES ON *.* TO 'kevin'@'%' WITH GRANT OPTION;
FLUSH PRIVILEGES;


CREATE USER arthuryounit WITH PASSWORD 'Y0un!t@rthur';

CREATE DATABASE younitdb;
CREATE SCHEMA younitschema;
CREATE USER younituser PASSWORD 'younituserp@ssword';
ALTER ROLE younituser WITH PASSWORD 'younituserp@ssword';

GRANT USAGE ON SCHEMA younitschema TO younituser;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA younitschema TO younituser;
GRANT ALL PRIVILEGES ON DATABASE younitdb TO younituser;

ALTER DEFAULT PRIVILEGES
FOR USER younituser
IN SCHEMA younitschema
GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO younituser;

/*
 The right to drop an object, or to alter its definition in any way, is not treated as a grantable privilege;
 it is inherent in the owner, and cannot be granted or revoked.
 (However, a similar effect can be obtained by granting or revoking membership in the role that
 owns the object; see below.) The owner implicitly has all grant options for the object, too.

 */

DROP TABLE IF EXISTS stratamanagercommunitytitlescheme;
DROP TABLE IF EXISTS Landlordownunit;
DROP TABLE IF EXISTS unitresident;
DROP TABLE IF EXISTS ComplexAssignedStrataManager;
DROP TABLE IF EXISTS LandlordPropertyManagerUnit;
DROP TABLE IF EXISTS CommunityTitleScheme;
DROP TABLE IF EXISTS PropertyManager;

/*
    click CTRL twice to get multiple cursors in multiple lanes
 */

DROP TABLE IF EXISTS unit;
DROP TABLE IF EXISTS complex;
DROP TABLE IF EXISTS resident;
DROP TABLE IF EXISTS landlord;
DROP TABLE IF EXISTS StrataManager;

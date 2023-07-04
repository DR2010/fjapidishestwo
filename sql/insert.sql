
# Database Insert

```sql
insert into younitdb.younitschema.complex
(complexid, name, streetnumber, streetname, upnumber, recordversion, createdby, updatedby)
values (1, 'first complex','9', 'Juaquim Street', '291', 1, 'younittest', 'younittest');

insert into younitdb.younitschema.complex
(complexid, name, streetnumber, streetname, upnumber, recordversion, createdby, updatedby)
values (2, 'second complex','329', 'Humber Street', '138', 1, 'younittest', 'younittest');

insert into younitdb.younitschema.complex
(complexid, name, streetnumber, streetname, upnumber, recordversion, createdby, updatedby)
values (3, 'third complex','421', 'Medusa Street', '634', 1, 'younittest', 'younittest');


insert into younitdb.younitschema.unit
(complexid, name, streetnumber, streetname, upnumber, recordversion, createdby, updatedby)
values (1, 'first complex','9', 'Juaquim Street', '291', 1, 'younittest', 'younittest');

insert into younitdb.younitschema.resident
(complexid, name, streetnumber, streetname, upnumber, recordversion, createdby, updatedby)
values (1, 'first complex','9', 'Juaquim Street', '291', 1, 'younittest', 'younittest');

insert into younitdb.younitschema.unitresident
(complexid, name, streetnumber, streetname, upnumber, recordversion, createdby, updatedby)
values (1, 'first complex','9', 'Juaquim Street', '291', 1, 'younittest', 'younittest');

insert into younitdb.younitschema.landlord
(complexid, name, streetnumber, streetname, upnumber, recordversion, createdby, updatedby)
values (1, 'first complex','9', 'Juaquim Street', '291', 1, 'younittest', 'younittest');

insert into younitdb.younitschema.landlordownunit
(complexid, name, streetnumber, streetname, upnumber, recordversion, createdby, updatedby)
values (1, 'first complex','9', 'Juaquim Street', '291', 1, 'younittest', 'younittest');
```
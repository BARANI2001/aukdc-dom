With new_qpk as(INSERT INTO Honorarium("TransactionID", "FacultyID", "CourseCode", "InitialAmount", "FinalAmount", "TypeID", "CreatedTime") VALUES (1234, 234, 'XC5254', 40, 50, 1 ,NOW()::timestamp(0)) RETURNING Honorarium."TransactionID") INSERT INTO "Question Paper/Key"("TransactionID","TypeID","QuestionPaperCount","KeyCount", "KeyRate", "QuestionPaperRate") VALUES((SELECT "TransactionID" FROM new_qpk), 1,6,7,8,9) returning "TransactionID";


INSERT INTO Course("CourseCode","Title","Regulation") VALUES('XC5254','DBMS','2019');
Insert Into Faculty("FacultyID","Name","PhoneNumber","Email","Facultyaukdcdom=# Insert Into Faculty("FacultyID","Name","PhoneNumber","Email","FacultyType","Department","Designation","Password","PANID","PanPicture","ExtensionNumber","Esign") VALUES(1234,'hi',9876543210,'fac@gmail.com','Permanent','math','Teaching Fellow','pass','asd','asoidj',1234567890,'asdfijsdf');


SELECT * FROM Users WHERE email = $1;

INSERT INTO Users (Username, Email,  Password, Role) 
VALUES ($1, $2, $3, $4);


INSERT INTO Rooms (Id, Owner_id)
VALUES ($1, $2) RETURNING Id, Owner_id;

SELECT * FROM Rooms WHERE Id = $1;
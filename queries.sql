
SELECT * FROM Users WHERE email = $1;

INSERT INTO Users (Username, Email,  Password, Role) 
VALUES ($1, $2, $3, $4);


INSERT INTO Rooms (Id, Room_Name, Owner_id)
VALUES ($1, $2, $3) RETURNING Id, Room_Name, Owner_id;

SELECT * FROM Rooms WHERE Id = $1;

SELECT id, room_name, owner_id FROM Rooms WHERE Owner_id = $1 ORDER BY created_at DESC;

INSERT INTO RoomStudent (Room_Id, Student_Name) 
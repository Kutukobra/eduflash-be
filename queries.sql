
SELECT * FROM Users WHERE email = $1;

INSERT INTO Users (Username, Email,  Password, Role) 
VALUES ($1, $2, $3, $4);
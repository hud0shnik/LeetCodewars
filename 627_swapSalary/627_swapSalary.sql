/*
Апдейт + кейс
немного не понял почему задача называется (Swap Salary) если просят поменять не зарплату а пол
*/

UPDATE Salary SET sex = CASE 
WHEN sex='m' THEN 'f'
WHEN sex='f' THEN 'm'
END
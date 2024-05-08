/*
Лефт джоин + юзинг
*/

SELECT unique_id, name FROM Employees LEFT JOIN EmployeeUNI USING(id)
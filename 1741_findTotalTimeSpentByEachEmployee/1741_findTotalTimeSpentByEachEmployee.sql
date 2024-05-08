/*
event_day выводится как day
Для вычисления времени по дню используется функция SUM
Для групирования по сотрудникам и дням исользуется GROUP BY emp_id, day
*/

SELECT emp_id, event_day AS day, SUM(out_time-in_time) AS total_time 
FROM Employees GROUP BY emp_id, day
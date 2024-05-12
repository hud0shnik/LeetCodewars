/*
Суммируем все операции пользователей чтобы вычислить текущий баланс
После этого отфильтровываем пользователей по балансу
*/
SELECT name, sum(amount) as balance FROM users LEFT JOIN
Transactions USING(account)
GROUP BY account, name
HAVING sum(amount) > 10000
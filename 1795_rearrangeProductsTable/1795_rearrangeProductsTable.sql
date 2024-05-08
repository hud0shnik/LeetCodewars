/*
Три селекта
каждый проходится по ненулевым значениям цен
с помощью юнионов три селекта объединяются в одну таблицу
сортировку добавил для красоты, LeetCode аксептит реквест и без неё
*/

SELECT product_id, 'store1' AS store, store1 AS price FROM products WHERE store1 IS NOT NULL
UNION
SELECT product_id, 'store2' AS store, store2 AS price FROM products WHERE store2 IS NOT NULL
UNION
SELECT product_id, 'store3' AS store, store3 AS price FROM products WHERE store3 IS NOT NULL
ORDER BY product_id, store
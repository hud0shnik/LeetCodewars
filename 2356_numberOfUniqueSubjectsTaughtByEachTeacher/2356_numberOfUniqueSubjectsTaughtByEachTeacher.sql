-- Считаем уникальные предметы и групируем по учителям
SELECT teacher_id, count(DISTINCT subject_id) as cnt 
FROM Teacher GROUP BY teacher_id
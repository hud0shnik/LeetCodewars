-- Классический селект со строковой функцией
SELECT tweet_id FROM Tweets WHERE char_length(content) > 15
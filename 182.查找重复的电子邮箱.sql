--
-- @lc app=leetcode.cn id=182 lang=mysql
--
-- [182] 查找重复的电子邮箱
--
# Write your MySQL query statement below

SELECT
    Email
FROM
    Person
GROUP BY 
    Email
HAVING
    COUNT(Email) > 1

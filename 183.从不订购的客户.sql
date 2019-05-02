--
-- @lc app=leetcode.cn id=183 lang=mysql
--
-- [183] 从不订购的客户
--
# Write your MySQL query statement below

SELECT
    Name AS Customers
FROM
    Customers
WHERE
    Id NOT IN (SELECT
            CustomerId
        FROM
            Orders
    )


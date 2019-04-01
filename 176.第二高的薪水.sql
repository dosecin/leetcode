-- [176] 第二高的薪水 
-- 
-- https://leetcode-cn.com/problems/second-highest-salary/description/
-- 
-- Tags: database 
-- 
-- Langs: mssql mysql oraclesql 
-- 
-- * database
-- * Easy (29.10%)
-- * Total Accepted: 24.4K
-- * Total Submissions: 83.4K
-- * Testcase Example: '{"headers": {"Employee": ["Id", "Salary"]}, "rows": {"Employee": [[1, 100], [2, 200], [3, 300]]}}'
-- 
-- 编写一个 SQL 查询，获取 Employee 表中第二高的薪水（Salary） 。
-- 
-- +----+--------+
-- | Id | Salary |
-- +----+--------+
-- | 1 | 100 |
-- | 2 | 200 |
-- | 3 | 300 |
-- +----+--------+
-- 
-- 
-- 例如上述 Employee 表，SQL查询应该返回 200 作为第二高的薪水。如果不存在第二高的薪水，那么查询应返回 null。
-- 
-- +---------------------+
-- | SecondHighestSalary |
-- +---------------------+
-- | 200 |
-- +---------------------+

-- Write your MySQL query statement below
SELECT MAX(e.Salary) AS SecondHighestSalary FROM Employee AS e,
(SELECT MAX(e1.Salary) AS MaxSalary FROM Employee AS e1) AS e1
WHERE e.Salary <> e1.MaxSalary
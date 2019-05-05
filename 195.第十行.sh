#
# @lc app=leetcode.cn id=195 lang=bash
#
# [195] 第十行
#
# Read from the file file.txt and output the tenth line to stdout.

linenum=$(cat file.txt | wc -l)
if [ $linenum -ge 10 ]
then
    head -n 10 file.txt | tail -n 1
fi

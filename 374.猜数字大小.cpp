/*
 * @lc app=leetcode.cn id=374 lang=cpp
 *
 * [374] 猜数字大小
 */
// Forward declaration of guess API.
// @param num, your guess
// @return -1 if my number is lower, 1 if my number is higher, otherwise return 0
int guess(int num);

class Solution {
public:
    int guessNumber(int n) {
        int min = 0, max = n;
        while (min <= max) {
            int mid = min + (max-min)/2;
            int g = guess(mid);
            if (g == 0) return mid;
            if (g == 1)
                min = mid+1;
            else 
                max = mid-1;
        }
        return min;
    }
};


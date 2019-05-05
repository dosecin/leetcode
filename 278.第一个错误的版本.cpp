/*
 * @lc app=leetcode.cn id=278 lang=cpp
 *
 * [278] 第一个错误的版本
 */
// Forward declaration of isBadVersion API.
bool isBadVersion(int version);

class Solution {
public:
    int firstBadVersion(int n) {
        int left = 1, right = n;
        while (true) {
            int mid = left + (right-left)/2;
            if (mid == left) return isBadVersion(mid)?mid:mid+1;
            if (isBadVersion(mid)) {
                right = mid;
            } else {
                left = mid;
            }
        }
    }
};


//using a sliding window approch 
int lengthOfLongestSubstring(char* s) 
{
    int map[128] = {0}; //because we know ASCII has 128 standard chars, we make an array that covers this range

    int right = 0; //current char we are looking at (window’s expanding end)
    int left = 0; //starting index of the current substring window (window’s shrinking start)
    
    int maxLen = 0; //track the longest substring we have found so far

    //null terminator '\0' has value zero
    //SO, this loop keeps running until s[start] is '\0'
    while(s[right]) 
    {
        //include current character s[right] in the window by incrementing its count
        map[s[right]]++; 

        //if the current character count is more than 1, we have a duplicate
        //shrink the window from the left until the duplicate is removed
        while (map[s[right]] > 1) 
        {
            //decrement count of character at left and move left pointer forward
            map[s[left]]--;
            left++;
        }

        //calculate current window length
        int windowLen = right - left + 1;

        //update maxLen if current window is longer
        if (maxLen < windowLen)
            maxLen = windowLen;

        //move right pointer forward to expand the window
        right++;
    }

    //return the maximum length found
    return maxLen;
}

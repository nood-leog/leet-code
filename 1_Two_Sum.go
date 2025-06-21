func twoSum(nums []int, target int) []int {

    myMap := make(map[int]int) //make a map

    //loop through each number in the nums array
    for i, num := range nums { 
        complement := target - num
        if prevIndex, ok := myMap[complement]; ok {
            return []int{prevIndex, i}
        }
        myMap[num] = i
    }
    return nil 
}

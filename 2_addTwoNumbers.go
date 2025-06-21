func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    // Create a starting point for the new linked list.
    // This dummy head makes it easier to build the list step by step.
    dummyHead := &ListNode{}

    // This pointer will move along the new list as we add numbers.
    current := dummyHead

    // This keeps track of any value that "carries over" when digits add up to 10 or more.
    carry := 0

    // Keep adding numbers as long as there is something in l1, l2, or a leftover carry.
    for l1 != nil || l2 != nil || carry != 0 {

        // Get the value from l1 if there is one; otherwise use 0.
        var x int
        if l1 != nil {
            x = l1.Val
            // Move l1 to the next digit.
            l1 = l1.Next
        }

        // Get the value from l2 if there is one; otherwise use 0.
        var y int
        if l2 != nil {
            y = l2.Val
            // Move l2 to the next digit.
            l2 = l2.Next
        }

        // Add the two digits plus any carry from before.
        sum := x + y + carry

        // Update the carry: if sum is 10 or more, carry will be 1; else 0.
        carry = sum / 10

        // Create a new node for this digit (the ones place of sum).
        current.Next = &ListNode{Val: sum % 10}

        // Move current to the new last node we just created.
        current = current.Next
    }

    // Return the new list, but skip the dummy head node.
    return dummyHead.Next
}

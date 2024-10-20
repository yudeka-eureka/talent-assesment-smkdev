# Answer Explaination

## Number 2: Balance Bracket case
Let's break down the time and space complexity of the solution provided in detail.

### Time Complexity Analysis:
1. **Iterating through the input string**: The code iterates over each character in the input string exactly once. This loop has a time complexity of 𝑂(𝑛), where 𝑛 is the number of characters in the string.

2. **Operations inside the loop**:
- For each character in the string, there is either a push operation (when an opening bracket is encountered) or a pop operation (when a matching closing bracket is encountered).

- Both push and pop operations on a stack take 𝑂(1) time.

- Additionally, looking up values in the bracketPairs map and comparing the characters takes 𝑂(1) time.

    Since both push and pop operations, as well as map lookups, are constant-time operations, the overall work inside the loop is constant per character.

3. **Final check (stack empty check)**: After the loop finishes, the final check to determine if the stack is empty is also 𝑂(1).

### Total Time Complexity:
Since we are iterating through the string once and performing constant-time operations for each character, the total time complexity is 𝑂(𝑛), where 𝑛 is the length of the input string.

### Space Complexity Analysis:
1. **Stack usage**: The maximum number of elements that can be pushed onto the stack is proportional to the number of opening brackets in the string. In the worst case, if the string consists of only opening brackets (e.g., `"{{{{[[[["`), the stack will hold all of them. Therefore, the maximum space used by the stack is 𝑂(𝑛).

2. **Bracket map**: The `bracketPairs` map is a constant-sized map with three entries, so its space complexity is 𝑂(1).

3. **Input storage**: The input string is stored externally, and in Go, strings are immutable, so there's no additional space complexity related to modifying or copying the input.

### Total Space Complexity:
The dominant term for space complexity is the space required by the stack, which can grow linearly with the size of the input. Therefore, the overall space complexity is 𝑂(𝑛), where 𝑛 is the length of the input string.

### Conclusion:
- **Time Complexity**: 𝑂(𝑛), where 𝑛 is the length of the input string.
- **Space Complexity**: 𝑂(𝑛), due to the stack used to track unmatched opening brackets.

This algorithm is efficient both in terms of time and space for the problem of determining if a string of brackets is balanced.
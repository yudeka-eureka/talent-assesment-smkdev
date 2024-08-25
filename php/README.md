# Task No. 2 explanation

1. Declaring variables, defining each closing and corresponding open bracket pairs.
2. Create loops for each string position depends on string length
      
        for ($i = 0; $i < strlen($str); $i++)

3. skip non-bracket characters (e.g., spaces) :
    
        if (!in_array($char, array_merge(array_keys($brackets), array_values($brackets)))) {
          continue;
        }

  Skipping character that is not in the brackets array.

4. If $char is an opening bracket (which means it is one of the values in the $brackets array), it is pushed onto the $stack.

        // If it's an opening bracket, push it onto the stack
        if (in_array($char, $brackets)) {
            array_push($stack, $char);
        }

5. If $char is a closing bracket (i.e., one of the keys in $brackets):
  Check if the stack is empty: If the stack is empty, there is no matching opening bracket for this closing bracket, so the function returns "NO".
  Check for matching pair: If the stack is not empty, the top element (obtained by array_pop($stack)) is removed and checked against the corresponding opening bracket (using $brackets[$char]). If they don't match, the function returns "NO".
    
        else {
          if (empty($stack) || array_pop($stack) !== $brackets[$char]) {
            return "NO";
          }
        }

6.  After iterating through the entire string:
    If the stack is empty (count($stack) == 0), all the brackets were matched and closed properly, so the function returns "YES".
    If the stack is not empty, it means there are unmatched opening brackets, so the function returns "NO".

        return count($stack) == 0 ? "YES" : "NO";

# 1. Weighted Strings (Score: 20)

The alphabet, from ‘a’ to ‘z’, is assigned weights corresponding to their ordinal positions. For instance, ‘a’ has a weight of 1, ‘b’ has a weight of 2, and so on, with ‘z’ having a weight of 26. The weight of a string is calculated by summing the weights of all its characters. Consider the string “gits” > gits = 7 + 9 + 20 + 19 = 55.

Example:
Input: Given a string “abbbcccd” and an array of queries “[1, 3, 9, 8]”. The queries are used to determine the status of the numbers in the queries based on the following rules:

	1.	If a number in the queries is equal to the weight of a character or substring, return “Yes”.
	2.	If a number in the queries is different from the weight of a character or substring, return “No”.

a = 1, b = 2, bb = 4, c = 3, cc = 6, ccc = 9, d = 4

Output: [Yes, Yes, Yes, No]

Explanation:

	1.	1 => Yes, 3 => Yes, 9 => Yes, and 8 => No.
	2.	Based on the character and substring weighting of “abbbcccd”, the status of the queries “[1, 3, 9, 8]” is “[Yes, Yes, Yes, No]”.

Rule: For repeated and consecutive characters, the weighting should be specified for each substring from the first occurrence to the n-th occurrence. Example: bbccc => b, bb, c, cc, ccc. Task: Create a function to solve the Weighted Strings problem!

---

# 2. Balanced Bracket (Score: 30)

Determining whether a string of brackets is balanced is a fundamental problem in computer science, particularly in parsing and validation of expressions. A balanced string of brackets consists of pairs of opening and closing brackets, where each opening bracket has a corresponding closing bracket of the same type, and the brackets are nested properly.

Given a string of brackets, determine whether it is balanced. The string may contain different types of brackets, including parentheses '()', square brackets '[]', and curly braces '{}'. The brackets may be separated by whitespace or not.

**Example 1:**
Input: `{ ( [ ] ) }` Output: YES Explanation: Each bracket is balanced, with every opening bracket having a corresponding closing bracket of the same type. *opening : { } opening : ( ) opening : [ ] opening : ( )*

**Example 2:**
Input: `{ [ ( ] ) }` Output: NO Explanation: The string `{ [ ( ] ) }` is not balanced for the characters enclosed by the curly braces '{}', namely '[ ( ]'.

**Example 3:**
Input: `{ ( ( [ ] ) [ ] ) [ ] }` Output: YES Explanation: Each bracket is balanced, with every opening bracket having a corresponding closing bracket of the same type, even though the bracket structure is irregular.

**Rules:**
1. Allowed bracket/character types: `( )`, `{ }`, or `[ ]`.
2. Brackets can be separated with or without whitespace.
3. Check for matching pairs of opening and closing brackets and return a string "YES" or "NO" accordingly.

**Task:**
1. Create a function to find Balanced Brackets with the *Lowest Complexity*!
2. Determine the complexity of your code. *Explain* the details of the complexity analysis for answer No. 2 and include it in the *README* of your repository!

---

# 3. Highest Palindrome (Score: 50)

Given a string representing a number, the goal is to find the highest possible palindrome that can be formed by changing at most 'k' digits in the string. A palindrome is a number that reads the same backward as forward.

You are given a string 's' representing a number and an integer 'k'. The task is to find the highest palindrome that can be formed by changing at most 'k' digits in the string 's'.

**Example 1:**
Input:
s: 3943
k: 1
Palindrome:
1. 3943 => 3993
2. 3943 => 3443
Output: 3993
Explanation: Among the possible palindromes obtained, the highest palindrome is 3993 since 3993 > 3443.

**Example 2:**
Input: s: 932239 k: 2
Palindrome:
1. 932239 => palindrome
2. Further Replacement (k = 2): => 992299
Output: 992299
Explanation: Among the possible palindromes obtained, the highest palindrome is 992299 since 992299 > 932239.

**Rules:**
1. If a palindrome cannot be formed from the given string even after 'k' replacements, and the string does not represent a valid number, output '-1'.
2. Do not use built-in functions or tools for searching, filtering, or sorting.
3. Avoid using loops.
4. Only use recursion to solve the problem.

**Task:**
Create a recursive function to solve the *Highest Palindrome* problem!


```
sequenceDiagram
    autonumber
    participant User
    participant Admin
    participant System
    participant Database

    %% User Login
    User->>System: Login(username, password)
    System->>Database: Validate credentials
    Database-->>System: Credentials valid/invalid
    System-->>User: Login success/failure

    %% Submit Documents
    Admin->>System: Submit Document (PDF)
    System->>Database: Save document
    Database-->>System: Document saved
    System-->>Admin: Submission success

    %% Change Document Status
    Admin->>System: Change Document Status (publish)
    System->>Database: Update document status
    Database-->>System: Status updated
    System-->>Admin: Status change success

    %% Delete Documents
    Admin->>System: Delete Document (document_id)
    System->>Database: Remove document
    Database-->>System: Document removed
    System-->>Admin: Deletion success

    %% Import Users from CSV
    Admin->>System: Import Users (CSV)
    System->>Database: Save users
    Database-->>System: Users saved
    System-->>Admin: Import success

    %% Export Users to CSV
    Admin->>System: Export Users (CSV)
    System->>Database: Retrieve users
    Database-->>System: Users data
    System-->>Admin: Export success (CSV file)

    %% Export User Read Analytics
    Admin->>System: Export User Read Analytics
    System->>Database: Retrieve analytics data
    Database-->>System: Analytics data
    System-->>Admin: Export success (CSV file)

    %% View Dashboard of User Read Analytics
    Admin->>System: View Analytics Dashboard
    System->>Database: Retrieve analytics data
    Database-->>System: Analytics data
    System-->>Admin: Display dashboard

    %% Publish Documents by Department
    Admin->>System: Publish Document (by department)
    System->>Database: Update document visibility
    Database-->>System: Document visibility updated
    System-->>Admin: Publish success

    %% Publish Documents by Level
    Admin->>System: Publish Document (by level)
    System->>Database: Update document visibility
    Database-->>System: Document visibility updated
    System-->>Admin: Publish success

    %% View List of Documents
    User->>System: View Document List
    System->>Database: Retrieve document list
    Database-->>System: Document list
    System-->>User: Display document list

    %% View Notifications
    User->>System: View Notifications
    System->>Database: Retrieve notifications
    Database-->>System: Notifications list
    System-->>User: Display notifications

    %% View Banner Notifications
    User->>System: View Banner Notifications
    System->>Database: Retrieve banner notifications
    Database-->>System: Banner notifications
    System-->>User: Display banner notifications

    %% Mark Notifications as Read
    User->>System: Mark Notification as Read (notification_id)
    System->>Database: Update notification status
    Database-->>System: Status updated
    System-->>User: Mark as read success

    %% View Document Details
    User->>System: View Document Details (document_id)
    System->>Database: Retrieve document details
    Database-->>System: Document details
    System-->>User: Display document details

    %% Comment on Documents
    User->>System: Comment on Document (document_id, comment)
    System->>Database: Save comment
    Database-->>System: Comment saved
    System-->>User: Comment success

    %% Print Documents with Watermark
    User->>System: Print Document (document_id)
    System->>Database: Retrieve document
    Database-->>System: Document data
    System-->>User: Print document with watermark

    %% Office-only Login
    User->>System: Login (office network)
    System->>Database: Validate network and credentials
    Database-->>System: Validation result
    System-->>User: Login success/failure

    %% View Organization Structure
    User->>System: View Organization Structure
    System->>Database: Retrieve org structure
    Database-->>System: Organization structure data
    System-->>User: Display organization structure

    %% Setup Organization Structure
    Admin->>System: Setup Organization Structure
    System->>Database: Save org structure
    Database-->>System: Structure saved
    System-->>Admin: Setup success

    %% View User Profile
    User->>System: View Profile
    System->>Database: Retrieve profile data
    Database-->>System: Profile data
    System-->>User: Display profile

    %% Search Documents
    User->>System: Search Document (name/number)
    System->>Database: Search documents
    Database-->>System: Search results
    System-->>User: Display search results

    %% Restrict Printing of Documents
    User->>System: Print Document (restricted)
    System-->>User: Print restriction message
```
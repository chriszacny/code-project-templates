"""
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.

Example 1:

Input: s = "()"
Output: true

Example 2:

Input: s = "()[]{}"
Output: true

Example 3:

Input: s = "(]"
Output: false
 
Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.
"""


import unittest


class AppTests(unittest.TestCase):
    def test_basic(self):
        # self.assertEqual(is_valid_parens("(){}[]"), True)
        # self.assertEqual(is_valid_parens("(){"), False)
        # self.assertEqual(is_valid_parens("(}"), False)
        # self.assertEqual(is_valid_parens("()"), True)
        # self.assertEqual(is_valid_parens("(]"), False)
        # self.assertEqual(is_valid_parens("{[]}"), True)
        self.assertEqual(is_valid_parens("()))"), False)


def is_match(left, right):
    table = { "(": ")", "[": "]", "{": "}" }
    #print(f'left is: {left}')
    if left in table.keys() and table[left] == right:
        return True
    return False


def is_valid_parens(input_str: str) -> bool:
    """
    Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

    An input string is valid if:

    Open brackets must be closed by the same type of brackets.
    Open brackets must be closed in the correct order.
    Every close bracket has a corresponding open bracket of the same type.

    ()[]{}
    ()[{}]
    (([[]]))
    (}){ - false
    """
    
    # Check if len(str) is even
    if len(input_str) % 2 != 0:
        return False
    
    stack = []
    prev = None
    for _, v in enumerate(input_str):
        #print(f'v is {v}')
        if prev == None or not is_match(prev, v):
            stack.append(v)
            prev = v
        else:
            if len(stack) > 0:
                stack.pop()
            if len(stack) > 0:
                prev = stack[-1]
            else:
                prev = None
    if len(stack) == 0:
        return True
    return False


if __name__ == "__main__":
    unittest.main()

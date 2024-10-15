const highestPalindrome = require('./highestPalindrome');

test('string 3943 and k = 1 highest palindrome to equal 3993', () => {
  expect(highestPalindrome("3943",1)).toBe('3993');
});

test('string 092282 and k = 3 highest palindrome to equal 992299', () => {
  expect(highestPalindrome("092282",3)).toBe('992299');
});

test('string 5566 and k = 1 highest palindrome to equal -1', () => {
  expect(highestPalindrome("5566",1)).toBe('-1');
});

test('string 932239 and k = 2 highest palindrome to equal 992299', () => {
  expect(highestPalindrome("932239",2)).toBe('992299');
});

test('string 11331 and k = 4 highest palindrome to equal 99399', () => {
  expect(highestPalindrome("11331",4)).toBe('99399');
});

test('string A1341 and k = 1 highest palindrome to equal -1', () => {
  expect(highestPalindrome("A1341",1)).toBe('-1');
});
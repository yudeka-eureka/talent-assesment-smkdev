const highestPalindrome = require('./highestPalindrome')

test('string 3993 and k = 1 highest palindrome to equal 3993', () => {
    expect(highestPalindrome("3993",1)).toBe('3993');
});

test('string 932239 and k = 2 highest palindrome to equal 992299', () => {
    expect(highestPalindrome("932239",2)).toBe('992299');
});

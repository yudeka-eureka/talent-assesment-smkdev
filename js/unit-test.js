describe('isPalindrome function', function() {
  it('should return true for a palindrome string', function() {
    expect(isPalindrome('madam')).toBe(true);
  });

  it('should return false for a non-palindrome string', function() {
    expect(isPalindrome('hello')).toBe(false);
  });

  it('should return true for a single character string', function() {
    expect(isPalindrome('a')).toBe(true);
  });

  it('should return true for an empty string', function() {
    expect(isPalindrome('')).toBe(true);
  });

  it('should return false for a string with different cases', function() {
    expect(isPalindrome('Madam')).toBe(false);
  });

  it('should return true for a string with spaces and punctuation', function() {
    expect(isPalindrome('A man, a plan, a canal: Panama')).toBe(true);
  });
});

const weightedString = require('./weightedStrings')

test('weightedString with input "abbcccd" and [1, 3, 9, 8] to equal ["YES", "YES", "YES", "NO"]', () => {
    expect(weightedString("abbcccd", [1, 3, 9, 8])).toEqual(["YES", "YES", "YES", "NO"]);
});

test('weightedString with input "aaabbbcccddd" and [5, 9, 7, 8, 12, 5] to equal ["NO", "YES", "NO", "YES", "YES", "NO"]', () => {
    expect(weightedString("aaabbbcccddd", [5, 9, 7, 8, 12, 5])).toEqual(["NO", "YES", "NO", "YES", "YES", "NO"]);
});

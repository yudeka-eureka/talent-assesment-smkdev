const weightedString = require('./weightedStrings')

test('weightedString with input "abbcccd" and [1, 3, 9, 8] to equal [true, true, true, false]', () => {
    expect(weightedString("abbcccd", [1, 3, 9, 8])).toEqual([true, true, true, false]);
});

test('weightedString with input "aaabbbcccddd" and [5, 9, 7, 8, 12, 5] to equal [false, true, false, true, true, false]', () => {
    expect(weightedString("aaabbbcccddd", [5, 9, 7, 8, 12, 5])).toEqual([false, true, false, true, true, false]);
});

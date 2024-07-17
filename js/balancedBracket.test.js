const balanceBracket = require('./balancedBracket');

test('adds { [ ( ) ] } to equal true', () => {
    expect(balanceBracket("{ [ ( ) ] }")).toBe(true);
});

test('adds { [ ( ) ] } to equal false', () => {
    expect(balanceBracket("{ [ ( ) ] }")).toBe(false);
});

test('adds { [ ( ] ) } to equal true', () => {
    expect(balanceBracket("{ [ ( ) ] }")).toBe(true);
});

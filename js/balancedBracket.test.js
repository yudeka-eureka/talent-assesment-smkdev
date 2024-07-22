const balanceBracket = require('./balancedBracket');

test('adds { [ ( ] ) } to equal NO', () => {
    expect(balanceBracket("{ [ ( ] ) }")).toBe("NO");
});

test('adds { ( ( [ ] ) [ ] ) [ ] } to equal YES', () => {
    expect(balanceBracket("{ ( ( [ ] ) [ ] ) [ ] }")).toBe("YES");
});

test('adds {[()]} to equal YES', () => {
    expect(balanceBracket("{[()]}")).toBe("YES");
});

test('adds {(([|])[])[]} to equal NO', () => {
    expect(balanceBracket("{(([|])[])[]}")).toBe("NO");
});
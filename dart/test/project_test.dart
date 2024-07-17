import 'package:project/project.dart';
import 'package:test/test.dart';

void main() {


  test('balanceBracket', () {
    expect(balanceBracket("{ [ ( ) ] }"), 'YES');
    expect(balanceBracket("{ [ ( ] ) }"), 'NO');
    expect(balanceBracket("{ ( ( [ ] ) [ ] ) [ ] }"), 'YES');
  });



  test('highestPalindrome', () {
    expect(highestPalindrome('3943', 1), '3993');
    expect(highestPalindrome('932239', 2), '992299');
  });


  test('weightedStrings', () {
    expect(weightedStrings('abbcccd', [1, 3, 9, 8]), ['YES', 'YES', 'YES', 'NO']);
    expect(weightedStrings('aaabbbcccddd', [5, 9, 7, 8, 12, 5]), ['NO', 'YES', 'NO', 'YES', 'YES', 'NO']);
  });

}

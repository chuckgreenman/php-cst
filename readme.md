# PHP Concrete Syntax Tree

Concrete Syntax Trees are a little different from the Abstract Syntax Trees that many languages implement.  While ASTs are great for static analysis, they are not great for code generation or modification.

Concrete Syntax Trees can help us programmatically update code while keeping the syntax of the original file, leaving us with cleaner diffs and more exact code changes.

## Parts
- Lexer - Converts PHP code into tokens
- Parser - Converts tokens into nodes
- Node Types - The types of nodes that can be parsed
- Printer - Converts nodes into PHP code
- Projection Options - Options for modifying the tree


## Lexer

In the Lexer we are taking the raw PHP code and converting it into tokens, the parser will use those tokens to build the tree. The following is adapted from the [PHP language reference manual](https://www.php.net/manual/en/tokens.php), and includes information about whether it is implemented in the PHP Concrete Syntax Tree (CST) or not.  The tokens that are not implemented in the CST are marked with ❌.

| Token | Syntax | Reference | Implemented in PHP CST |
|-------|--------|-----------|------------------------|
| T_ABSTRACT | abstract | Class Abstraction | ❌ |
| T_AMPERSAND_FOLLOWED_BY_VAR_OR_VARARG | & | Type declarations (available as of PHP 8.1.0) | ❌ |
| T_AMPERSAND_NOT_FOLLOWED_BY_VAR_OR_VARARG | & | Type declarations (available as of PHP 8.1.0) | ❌ |
| T_AND_EQUAL | &= | assignment operators | ❌ |
| T_ARRAY | array() | array(), array syntax | ❌ |
| T_ARRAY_CAST | (array) | type-casting | ❌ |
| T_AS | as | foreach | ❌ |
| T_ATTRIBUTE | #[ | attributes (available as of PHP 8.0.0) | ❌ |
| T_BAD_CHARACTER |  | anything below ASCII 32 except \t (0x09), \n (0x0a) and \r (0x0d) (available as of PHP 7.4.0) | ❌ |
| T_BOOLEAN_AND | && | logical operators | ❌ |
| T_BOOLEAN_OR | \|\| | logical operators | ❌ |
| T_BOOL_CAST | (bool) or (boolean) | type-casting | ❌ |
| T_BREAK | break | break | ❌ |
| T_CALLABLE | callable | callable | ❌ |
| T_CASE | case | switch | ❌ |
| T_CATCH | catch | Exceptions | ❌ |
| T_CLASS | class | classes and objects | ❌ |
| T_CLASS_C | \_\_CLASS\_\_ | magic constants | ❌ |
| T_CLONE | clone | classes and objects | ❌ |
| T_CLOSE_TAG | ?> or %> | escaping from HTML | ❌ |
| T_COALESCE | ?? | comparison operators | ❌ |
| T_COALESCE_EQUAL | ??= | assignment operators (available as of PHP 7.4.0) | ❌ |
| T_COMMENT | // or #, and /* */ | comments | ❌ |
| T_CONCAT_EQUAL | .= | assignment operators | ❌ |
| T_CONST | const | class constants | ❌ |
| T_CONSTANT_ENCAPSED_STRING | "foo" or 'bar' | string syntax | ❌ |
| T_CONTINUE | continue | continue | ❌ |
| T_CURLY_OPEN | {$ | advanced variable string interpolation | ❌ |
| T_DEC | -- | incrementing/decrementing operators | ❌ |
| T_DECLARE | declare | declare | ❌ |
| T_DEFAULT | default | switch | ❌ |
| T_DIR | \_\_DIR\_\_ | magic constants | ❌ |
| T_DIV_EQUAL | /= | assignment operators | ❌ |
| T_DNUMBER | 0.12, etc. | floating point numbers | ❌ |
| T_DO | do | do..while | ❌ |
| T_DOC_COMMENT | /** */ | PHPDoc style comments | ❌ |
| T_DOLLAR_OPEN_CURLY_BRACES | ${ | basic variable string interpolation | ❌ |
| T_DOUBLE_ARROW | => | array syntax | ❌ |
| T_DOUBLE_CAST | (real), (double) or (float) | type-casting | ❌ |
| T_DOUBLE_COLON | :: | see T_PAAMAYIM_NEKUDOTAYIM below | ❌ |
| T_ECHO | echo | echo | ❌ |
| T_ELLIPSIS | ... | function arguments | ❌ |
| T_ELSE | else | else | ❌ |
| T_ELSEIF | elseif | elseif | ❌ |
| T_EMPTY | empty | empty() | ❌ |
| T_ENCAPSED_AND_WHITESPACE | " $a" | constant part of string with variables | ❌ |
| T_ENDDECLARE | enddeclare | declare, alternative syntax | ❌ |
| T_ENDFOR | endfor | for, alternative syntax | ❌ |
| T_ENDFOREACH | endforeach | foreach, alternative syntax | ❌ |
| T_ENDIF | endif | if, alternative syntax | ❌ |
| T_ENDSWITCH | endswitch | switch, alternative syntax | ❌ |
| T_ENDWHILE | endwhile | while, alternative syntax | ❌ |
| T_ENUM | enum | Enumerations (available as of PHP 8.1.0) | ❌ |
| T_END_HEREDOC |  | heredoc syntax | ❌ |
| T_EVAL | eval() | eval() | ❌ |
| T_EXIT | exit or die | exit(), die() | ❌ |
| T_EXTENDS | extends | extends, classes and objects | ❌ |
| T_FILE | \_\_FILE\_\_ | magic constants | ❌ |
| T_FINAL | final | Final Keyword | ❌ |
| T_FINALLY | finally | Exceptions | ❌ |
| T_FN | fn | arrow functions (available as of PHP 7.4.0) | ❌ |
| T_FOR | for | for | ❌ |
| T_FOREACH | foreach | foreach | ❌ |
| T_FUNCTION | function | functions | ❌ |
| T_FUNC_C | \_\_FUNCTION\_\_ | magic constants | ❌ |
| T_GLOBAL | global | variable scope | ❌ |
| T_GOTO | goto | goto | ❌ |
| T_HALT_COMPILER | \_\_halt_compiler() | \_\_halt_compiler | ❌ |
| T_IF | if | if | ❌ |
| T_IMPLEMENTS | implements | Object Interfaces | ❌ |
| T_INC | ++ | incrementing/decrementing operators | ❌ |
| T_INCLUDE | include | include | ❌ |
| T_INCLUDE_ONCE | include_once | include_once | ❌ |
| T_INLINE_HTML |  | text outside PHP | ❌ |
| T_INSTANCEOF | instanceof | type operators | ❌ |
| T_INSTEADOF | insteadof | Traits | ❌ |
| T_INTERFACE | interface | Object Interfaces | ❌ |
| T_INT_CAST | (int) or (integer) | type-casting | ❌ |
| T_ISSET | isset() | isset() | ❌ |
| T_IS_EQUAL | == | comparison operators | ❌ |
| T_IS_GREATER_OR_EQUAL | >= | comparison operators | ❌ |
| T_IS_IDENTICAL | === | comparison operators | ❌ |
| T_IS_NOT_EQUAL | != or <> | comparison operators | ❌ |
| T_IS_NOT_IDENTICAL | !== | comparison operators | ❌ |
| T_IS_SMALLER_OR_EQUAL | <= | comparison operators | ❌ |
| T_LINE | \_\_LINE\_\_ | magic constants | ❌ |
| T_LIST | list() | list() | ❌ |
| T_LNUMBER | 123, 012, 0x1ac, etc. | integers | ❌ |
| T_LOGICAL_AND | and | logical operators | ❌ |
| T_LOGICAL_OR | or | logical operators | ❌ |
| T_LOGICAL_XOR | xor | logical operators | ❌ |
| T_MATCH | match | match (available as of PHP 8.0.0) | ❌ |
| T_METHOD_C | \_\_METHOD\_\_ | magic constants | ❌ |
| T_MINUS_EQUAL | -= | assignment operators | ❌ |
| T_MOD_EQUAL | %= | assignment operators | ❌ |
| T_MUL_EQUAL | *= | assignment operators | ❌ |
| T_NAMESPACE | namespace | namespaces | ❌ |
| T_NAME_FULLY_QUALIFIED | \App\Namespace | namespaces (available as of PHP 8.0.0) | ❌ |
| T_NAME_QUALIFIED | App\Namespace | namespaces (available as of PHP 8.0.0) | ❌ |
| T_NAME_RELATIVE | namespace\Namespace | namespaces (available as of PHP 8.0.0) | ❌ |
| T_NEW | new | classes and objects | ❌ |
| T_NS_C | \_\_NAMESPACE\_\_ | namespaces | ❌ |
| T_NS_SEPARATOR | \ | namespaces | ❌ |
| T_NUM_STRING | "$a[0]" | numeric array index inside string | ❌ |
| T_OBJECT_CAST | (object) | type-casting | ❌ |
| T_OBJECT_OPERATOR | -> | classes and objects | ❌ |
| T_NULLSAFE_OBJECT_OPERATOR | ?-> | classes and objects | ❌ |
| T_OPEN_TAG | <?php, <? or <% | escaping from HTML | ❌ |
| T_OPEN_TAG_WITH_ECHO | <?= or <%= | escaping from HTML | ❌ |
| T_OR_EQUAL | \|= | assignment operators | ❌ |
| T_PAAMAYIM_NEKUDOTAYIM | :: | scope resolution. Also defined as T_DOUBLE_COLON. | ❌ |
| T_PLUS_EQUAL | += | assignment operators | ❌ |
| T_POW | ** | arithmetic operators | ❌ |
| T_POW_EQUAL | **= | assignment operators | ❌ |
| T_PRINT | print | print | ❌ |
| T_PRIVATE | private | classes and objects | ❌ |
| T_PRIVATE_SET | private(set) | property hooks (available as of PHP 8.4.0) | ❌ |
| T_PROPERTY_C | \_\_PROPERTY\_\_ | magic constants | ❌ |
| T_PROTECTED | protected | classes and objects | ❌ |
| T_PROTECTED_SET | protected(set) | property hooks (available as of PHP 8.4.0) | ❌ |
| T_PUBLIC | public | classes and objects | ❌ |
| T_PUBLIC_SET | public(set) | property hooks (available as of PHP 8.4.0) | ❌ |
| T_READONLY | readonly | classes and objects (available as of PHP 8.1.0) | ❌ |
| T_REQUIRE | require | require | ❌ |
| T_REQUIRE_ONCE | require_once | require_once | ❌ |
| T_RETURN | return | returning values | ❌ |
| T_SL | << | bitwise operators | ❌ |
| T_SL_EQUAL | <<= | assignment operators | ❌ |
| T_SPACESHIP | <=> | comparison operators | ❌ |
| T_SR | >> | bitwise operators | ❌ |
| T_SR_EQUAL | >>= | assignment operators | ❌ |
| T_START_HEREDOC | <<< | heredoc syntax | ❌ |
| T_STATIC | static | variable scope | ❌ |
| T_STRING | parent, self, etc. | identifiers, e.g. keywords like parent and self, function names, class names and more are matched. See also T_CONSTANT_ENCAPSED_STRING. | ❌ |
| T_STRING_CAST | (string) | type-casting | ❌ |
| T_STRING_VARNAME | "${a | variable variables to interpolate in a string | ❌ |
| T_SWITCH | switch | switch | ❌ |
| T_THROW | throw | Exceptions | ❌ |
| T_TRAIT | trait | Traits | ❌ |
| T_TRAIT_C | \_\_TRAIT\_\_ | \_\_TRAIT\_\_ | ❌ |
| T_TRY | try | Exceptions | ❌ |
| T_UNSET | unset() | unset() | ❌ |
| T_UNSET_CAST | (unset) | type-casting | ❌ |
| T_USE | use | namespaces | ❌ |
| T_VAR | var | classes and objects | ❌ |
| T_VARIABLE | $foo | variables | ❌ |
| T_WHILE | while | while, do..while | ❌ |
| T_WHITESPACE | \t \r\n |  | ❌ |
| T_XOR_EQUAL | ^= | assignment operators | ❌ |
| T_YIELD | yield | generators | ❌ |
| T_YIELD_FROM | yield from | generators | ❌ |

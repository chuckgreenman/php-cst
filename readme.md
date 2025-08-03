# PHP Concrete Syntax Tree

Concrete Syntax Trees are a little different from the Abstract Syntax Trees that many languages implement.  While ASTs are great for static analysis, they are not great for code generation or modification.

Concrete Syntax Trees can help us programmatically update code while keeping the syntax of the original file, leaving us with cleaner diffs and more exact code changes.

## Parts
- Parser - Converts PHP code into tokens
- Lexer - Converts tokens into nodes
- Node Types - The types of nodes that can be parsed
- Printer - Converts nodes into PHP code
- Projection Options - Options for modifying the tree

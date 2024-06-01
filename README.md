# Lexical analysis

Lexical tokenization is conversion of a text into meaningful lexical tokens belonging to categories defined by a "lexer" program.

what does it parse?

- Modern Syntax
- Static Types
- Implicit & Emplicit Type Ineterface
- Array and String Literals
- Control flow and Imports


# Tokenization
- Tokenization is the process in which you transform a stream of characters into meaningful array of tokens


# Abstract Syntax Tree (AST)
An AST is a data structure which represents the program's structure. AST's are easy to traverse and have many uses.

# Producing an AST
We will construct the AST by processing the tokens we received from the lexer.
This process of building AST from tokens is called ==Parsing==.

# Pratt parsing and how it works:

1. **Operator Precedence**:
   Pratt parsing relies on the concept of operator precedence, which defines the order of operations (e.g., multiplication before addition) in an expression. Each operator is assigned a precedence level that determines how tightly it binds to its operands.

2. **Parsing Approach**:
   - Pratt parsing is a form of recursive descent parsing where each operator has a parsing function associated with it. These parsing functions determine how to parse expressions involving the respective operator based on the current token and the surrounding context.
   - The parsing functions are organized according to operator precedence. Higher precedence operators have parsing functions that handle tighter binding and evaluate expressions accordingly.

3. **Key Components**:
   - **Tokenization**: The input expression is first tokenized into a sequence of tokens (e.g., numbers, operators).
   - **Parsing Functions**: Pratt parsing defines parsing functions for each operator based on its precedence and associativity. These functions handle the parsing of expressions involving the respective operator.
   - **Recursive Descent**: The parsing process is recursive in nature, where parsing functions call each other recursively to handle nested expressions and operators.

4. **Advantages**:
   - Pratt parsing is flexible and extensible, allowing for easy addition of new operators with different precedence levels.
   - It efficiently handles operator precedence and associativity without the need for complex parsing tables or state machines.

5. **Example**:
   Consider the expression `2 + 3 * 4`. In Pratt parsing:
   - Tokenization yields tokens `[2, +, 3, *, 4]`.
   - The parsing functions for `+` and `*` operators are invoked based on their precedence.
   - The `*` operator (higher precedence) binds more tightly, so `3 * 4` is evaluated first.
   - The result of `3 * 4` is then combined with `2` using the `+` operator to produce the final result `14`.

Pratt parsing is commonly used in simple expression evaluators and compilers, particularly for languages with operator-based syntax. It provides an elegant and efficient approach to parsing expressions while respecting operator precedence and associativity rules.

good read: [Simple but powerful pratt parsing](https://matklad.github.io/2020/04/13/simple-but-powerful-pratt-parsing.html)

# Binding Power
Binding power (also known as precedence) determines how tightly an operator binds to its operands. Higher binding power means the operator binds more tightly, thus getting evaluated earlier in the expression. For example, in the expression 3 + 4 * 5, the multiplication operator * has higher binding power than the addition operator +, so 4 * 5 is evaluated first.

# NUD (Null Denotation)
NUD stands for `Null Denotation`. It is a function that defines how to parse expressions that start with a specific token. NUD is used for tokens that appear at the beginning of an expression. For instance:

A number might have a NUD function that simply returns the numeric value.
A parenthesis might have a NUD function that parses an enclosed expression until a closing parenthesis is found.

example: prefix or unary expressions.

# LED (Left Denotation)
LED stands for `Left Denotation`. It is a function that defines how to parse expressions that continue with a specific token, given the left-hand side (LHS) of the expression. LED is used for infix operators (like + or *) and postfix operators. For example:

An addition operator + might have a LED function that parses the right-hand side (RHS) of the addition expression and combines it with the LHS.
Example of Pratt Parsing

example: infix and postfix expressions

# Lookup Table
A **lookup table** is a data structure used to store information about the behavior of tokens, specifically their binding powers and the functions used to parse them (NUD and LED functions). This table allows the parser to quickly access the appropriate actions for each token type.

### Key Elements of a Lookup Table in Pratt Parsing

1. **Token Type:** Identifies the type of token (e.g., number, plus sign, parenthesis).
2. **Binding Power:** Indicates the precedence of the token, which helps in determining how tightly the token binds to its operands.
3. **NUD Function:** Defines how the parser should handle this token if it appears at the start of an expression.
4. **LED Function:** Defines how the parser should handle this token if it appears in the middle of an expression, given the left-hand side.

### Benefits of Using a Lookup Table

- **Efficiency:** Quick access to token handling logic based on the token type.
- **Flexibility:** Easy to update or extend the parser by modifying the lookup table without changing the core parsing logic.
- **Clarity:** Clear separation of concerns, with token behavior defined in one place, making the parser code easier to understand and maintain.

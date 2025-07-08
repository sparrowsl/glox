# Lox in Go 🦫 (glox)

This is an implementation of the [Lox](http://craftinginterpreters.com/) language from the book **Crafting Interpreters** by Bob Nystrom — written in **Go**.

The goal is to build a full interpreter from scratch in Go, following the first part of the book, which builds a **tree-walk interpreter**.

---

## 📚 Book Link

👉 [Crafting Interpreters](http://craftinginterpreters.com/)

---

## 🛠️ Project Structure

- `scanner/` – tokenization of source code
- `parser/` – builds an AST from tokens
- `interpreter/` – walks the AST and executes
- `ast/` – AST node definitions
- `main.go` – REPL and file runner

---

## 📦 Goals

- Clean, idiomatic Go implementation
- Modular and testable components
- Full feature parity with the Java version from the book

---

## 🚀 Run

```bash
go run main.go path/to/file.lox

# or

go build ./...
./lox path/to/file.lox
```


## ✅ Features (Part I – Tree-Walk Interpreter)

Progress on implementing features from the first half of the book:

### 📖 Chapter 1–4: Introduction & Tools

- [ ] Project setup
- [ ] REPL scaffolding
- [ ] File runner
- [ ] Error reporting

---

### ✂️ Chapter 5: Scanning

- [ ] Define token types
- [ ] Implement Scanner structure
- [ ] Recognize lexemes
- [ ] Single-character tokens (`( ) { } , . - + ; *`)
- [ ] One or two character tokens (`!=`, `==`, `<=`, `>=`)
- [ ] Literals
  - [ ] Strings
  - [ ] Numbers
  - [ ] Identifiers
- [ ] Keywords
  - [ ] `and`, `class`, `else`, `false`, `for`, `fun`, `if`, `nil`, `or`, `print`, `return`, `super`, `this`, `true`, `var`, `while`
- [ ] Whitespace and comments
- [ ] Error handling (unexpected characters)
- [ ] Token stringification for debugging

---

### 🌳 Chapter 6: Representing Code (AST)

- [ ] Define expression types
  - [ ] Binary
  - [ ] Grouping
  - [ ] Literal
  - [ ] Unary
  - [ ] Variable
  - [ ] Assignment
  - [ ] Logical
  - [ ] Call
- [ ] Visitor pattern in Go
- [ ] AST Printer for debugging

---

### 🧠 Chapter 7–9: Parsing Expressions

- [ ] Pratt parser (recursive descent)
- [ ] Implement grammar rules
  - [ ] Primary
  - [ ] Unary
  - [ ] Factor
  - [ ] Term
  - [ ] Comparison
  - [ ] Equality
  - [ ] Logic
  - [ ] Assignment
- [ ] Parse grouping and literals
- [ ] Parse binary expressions
- [ ] Parse unary expressions
- [ ] Associativity and precedence
- [ ] Parse function calls
- [ ] Parse error handling (synchronization)

---

### 🧾 Chapter 10: Statements and State

- [ ] Statement parsing
  - [ ] Print statement
  - [ ] Expression statement
  - [ ] Variable declarations
  - [ ] Block statements
- [ ] AST node definitions for statements
- [ ] Parse `var` declarations
- [ ] Parse `print` and `expression` statements

---

### 🔁 Chapter 11: Control Flow

- [ ] `if` statements
- [ ] `while` loops
- [ ] `for` loops (desugar into `while`)
- [ ] `else` branches
- [ ] Block scoping and nesting
- [ ] Logical operators (`and`, `or`)

---

### ⚙️ Chapter 12: Functions

- [ ] Function declaration parsing
- [ ] Return statements
- [ ] Function calling and arity checks
- [ ] First-class function values
- [ ] Store functions in environment

---

### 🔒 Chapter 13: Resolving and Scoping

- [ ] Lexical environment structure
- [ ] Variable resolution
- [ ] Proper scoping for blocks
- [ ] Error handling for undefined variables
- [ ] Handle local vs global variables

---

### 🧵 Chapter 14: Closures

- [ ] Capture variables from outer scopes
- [ ] Maintain closures across returns
- [ ] Nested function definitions
- [ ] Free variable resolution

---

### 🏷️ Chapter 15: Classes (Early Object Support)

- [ ] Class declarations
- [ ] `this` keyword
- [ ] Method definitions
- [ ] Instances and property access

---

### 🧬 Chapter 16: Inheritance

- [ ] `super` keyword
- [ ] Method resolution order
- [ ] Subclass constructor chaining
- [ ] Method overriding

---

### 🧮 Chapter 17–18: Interpreter Core

- [ ] Evaluate expressions
  - [ ] Literals, grouping, unary, binary
  - [ ] Variables and assignments
  - [ ] Logical expressions
- [ ] Execute statements
  - [ ] Print
  - [ ] Var
  - [ ] Expression
  - [ ] Block
  - [ ] Control flow (`if`, `while`, `for`)
  - [ ] Function calls and returns
- [ ] Environment tracking
- [ ] Visitor methods for each node type
- [ ] Support for truthiness and equality
- [ ] Native functions (e.g., `clock()`)

---

### ⚠️ Chapter 19: Error Handling

- [ ] Runtime errors with line numbers
- [ ] Proper interpreter panic recovery
- [ ] User-facing messages
- [ ] Continue REPL after errors

---

## 🧪 Testing & Tooling

- [ ] Unit tests for Scanner
- [ ] Unit tests for Parser
- [ ] Evaluation tests for Interpreter
- [ ] AST Printer validation
- [ ] Golden test outputs

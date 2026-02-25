# learning-go-by-writing-another-lisp

Small Go playground for language/tooling experiments.

## What You Were Working On

You were building a tiny Lisp-style compiler front-end in `pkg/lcom`:

- `tokenize.go`: converts source text into tokens (`paren`, `number`, `name`, `string`)
- `parse.go`: converts tokens into an AST (`NumberLiteral`, `StringLiteral`, `CallExpression`)
- `cmd/lcom/main.go`: CLI entry point for trying the flow end-to-end

## When You Worked On This

`pkg/lcom` timeline from git history:

- `2024-05-06 23:56:12 +0900` - commit `3a23a88` - "compilering"
- `2024-05-07 09:04:04 +0900` - commit `3c76a51` - "started parser"

Archive note:

- There are also uncommitted local changes in this working tree, which happened after the May 2024 commits (exact timestamps are not stored by git until committed).

## lcom Status Snapshot

- Parser/tokenizer core functions exist and are wired.
- Tests in `pkg/lcom/test` look mid-refactor:
  - They use `package lcom_test` but call unexported functions/fields.
  - Some parse tests pass the wrong input types (using `Node` where `[]Token` is expected).
  - `TestParseExpression` is still empty.

## Quick Resume Checklist

1. Decide test style:
   - Internal tests (`package lcom`) to test unexported helpers directly, or
   - External tests (`package lcom_test`) and export parser/tokenizer API.
2. Fix parser tests to match function signatures.
3. Add `TestParseExpression` coverage for nested calls and closing-paren handling.
4. Run:
   - `go test ./...`
   - `go test -cover ./...`

# segsql

[![codecov](https://codecov.io/github/ismaildurmaz/segsql/branch/master/graph/badge.svg?token=O7NC5M5R3E)](https://codecov.io/github/ismaildurmaz/segsql)

> **Composable, segment-based SQL builder for Go**  
> Build SQL queries from independent, reusable segments with first-class support for parameters and subqueries.

`segsql` is a **low-level, ORM-free SQL builder** for Go.  
It allows you to construct SQL queries using **independent, composable segments** such as `SELECT`, `FROM`, `WHERE`, `JOIN`, `GROUP BY`, and `ORDER BY`.

The main goal is to make **dynamic SQL generation**:
- readable
- safe
- composable
- fully under your control

---

## ✨ Features

- 🧩 Segment-based query building
- 🔁 Fully composable & reusable query parts
- 🔐 Safe parameter handling (SQL injection resistant)
- 🔂 First-class subquery support
- 🧠 Fluent, readable API
- 🧱 Not an ORM — full SQL control
- 🧪 Ideal for `sqlc`, manual SQL, and complex filtering logic

---

## 🚫 What segsql is NOT

`segsql` intentionally does **not** provide:

- ❌ ORM functionality
- ❌ Entity / struct mapping
- ❌ Schema or migrations
- ❌ Query execution or connection handling

> `segsql` only **builds SQL**.  
> Executing it is your responsibility.

---

## 📦 Installation

```bash
go get github.com/ismaildurmaz/segsql
```

---

## 🚀 Quick Example

```go
sql, err := sqlbuilder.SelectQuery().
    Select().Fields("u.id", "u.name").ToSelectQuery().
    From().TableWithAlias("users", "u").ToSelectQuery().
    Where().
        Eq("u.status", "active").
        And().
        In("u.role", "admin", "editor").
    ToSelectQuery().
    ToSQL()
```

---

## 🧠 Design Philosophy

- **Segments over monoliths**
- **Explicit over implicit**
- **SQL-first**
- **Builder, not executor**

---

## 📄 License

MIT License © Ismail Durmaz

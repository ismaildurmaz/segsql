# segsql

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=ismaildurmaz_segsql&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=ismaildurmaz_segsql)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=ismaildurmaz_segsql&metric=bugs)](https://sonarcloud.io/summary/new_code?id=ismaildurmaz_segsql)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=ismaildurmaz_segsql&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=ismaildurmaz_segsql)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=ismaildurmaz_segsql&metric=coverage)](https://sonarcloud.io/summary/new_code?id=ismaildurmaz_segsql)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=ismaildurmaz_segsql&metric=duplicated_lines_density)](https://sonarcloud.io/summary/new_code?id=ismaildurmaz_segsql)

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

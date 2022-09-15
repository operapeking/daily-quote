# daily-quote

```sql
CREATE TABLE IF NOT EXISTS "quote" (
        "uid" INTEGER PRIMARY KEY AUTOINCREMENT,
        "content" TEXT NULL UNIQUE,
        "from" TEXT,
		"author" TEXT
    );
```
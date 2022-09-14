package main

import "database/sql"

var quote1 = Quote{1, "欲买桂花同载酒，只可惜故人，何日再见了。", "钟离"}
var quote2 = Quote{2, "与君相别离，不知何日是归期，我如朝露转瞬晞。", "花散里"}
var quote3 = Quote{3, "摩拉天然是货币，货币不天然是摩拉。", "钟离"}

var db *sql.DB

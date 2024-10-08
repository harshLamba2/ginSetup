1. Where()
Purpose: The .Where() method is used to add conditions to your SQL queries. It's equivalent to the WHERE clause in SQL and allows you to filter records based on specific criteria.

Usage: You can pass conditions as raw SQL, structs, or maps.

Using Raw SQL(Code):{
db.Where("name = ?", "John").Find(&users)
This will generate: SELECT * FROM users WHERE name = 'John';
}

Using Structs(Code):{
db.Where(&User{Name: "John", Age: 25}).Find(&users)
This will generate: SELECT * FROM users WHERE name = 'John' AND age = 25;
}

Using Maps(Code):{
db.Where(map[string]interface{}{"name": "John", "age": 25}).Find(&users)
This will generate the same query as the struct.
}

Conclusion:
db.Where("age > ?", 30).Find(&users)




2. .Find()
Purpose: The .Find() method is used to retrieve records from the database based on the conditions you define. It fetches multiple records that match the query.

Without conditions(Code):{

db.Find(&users)
This retrieves all records from the users table: SELECT * FROM users;
}

With conditions (e.g., .Where())(Code):{

db.Where("name = ?", "John").Find(&users)
This fetches users with the name "John": SELECT * FROM users WHERE name = 'John';
}

Conclusion:
var users []User
db.Find(&users)  // Fetch all users

Note: If you want to retrieve only one record, you should use .First() or .Last() instead of .Find().




3. .Preload()
Purpose: The .Preload() method is used for eager loading related associations. It's useful when you want to retrieve related models (e.g., HasMany, HasOne, BelongsTo, ManyToMany) in the same query. Instead of performing separate queries for each association, .Preload() loads them all in a single query.

Usage:

Single Preload(Code):{

db.Preload("Orders").Find(&users)
This will load users and their related Orders.
}

Multiple Preloads(Code):{

db.Preload("Orders").Preload("Profile").Find(&users)
}

Preload with conditions(Code):{

db.Preload("Orders", "status = ?", "completed").Find(&users)
This loads users and only their Orders where the status is "completed".
}

Conclusion:
var users []User
db.Preload("Orders").Find(&users)  // Fetch users along with their orders

Use Case: If a user has multiple related Orders, you can preload them to avoid the N+1 query problem (where a separate query is made for each user’s orders).




4. .Joins()
Purpose: The .Joins() method is used for joining related tables, similar to SQL JOIN. It's typically used when you want to fetch records based on relationships between multiple tables.

Basic Join(Code):{

db.Joins("JOIN orders ON orders.user_id = users.id").Where("orders.status = ?", "completed").Find(&users)
This performs a SQL JOIN between the users and orders tables.
}

With Associations: If you're using GORM associations, you can join related tables automatically(Code):{

db.Joins("Orders").Find(&users)
}

Left Join(Code):{

db.Joins("LEFT JOIN orders ON orders.user_id = users.id").Find(&users)
}

Conclusion:
db.Joins("JOIN orders ON orders.user_id = users.id").Where("orders.status = ?", "completed").Find(&users)
Difference from .Preload(): While .Preload() is for eager loading (loading related models), .Joins() is used to join tables directly in the query. You use .Joins() when you want to write more specific SQL-like queries, while .Preload() is for ORM-based associations.


<!-- Other Less Frequenctly Used Methds -->


1. First()
Retrieves the first record that matches the conditions.

db.First(&user)


2. Last()
Retrieves the last record that matches the conditions.

db.Last(&user)


3. Save()
Saves the value into the database (both insert and update).

db.Save(&user)


4. Create()
Inserts a new record into the database.

db.Create(&user)


5. Delete()
Deletes a record that matches the condition.

db.Delete(&user)


6. Update() / .Updates()
Updates a specific column or multiple columns in the record.

db.Model(&user).Update("name", "John")
db.Model(&user).Updates(User{Name: "John", Age: 25})


7. Scan()
Scans the result into a destination (often a struct).

db.Model(&user).Scan(&result)


8. Select()
Specifies which fields to retrieve or modify.

db.Select("name", "age").Find(&users)


9. Order()
Specifies the order of the results (e.g., ascending or descending).

db.Order("age desc").Find(&users)


10. Group()
Used for grouping records.

db.Model(&user).Group("name").Find(&results)


11. Having()
Used with .Group() to specify conditions for grouped records.

db.Group("name").Having("count(name) > ?", 1).Find(&results)


12. Limit()
Limits the number of records retrieved.

db.Limit(10).Find(&users)


13. Offset()
Skips a specified number of records.

db.Offset(5).Find(&users)


14. Count()
Counts the number of records that match the query.

var count int64
db.Model(&User{}).Where("age > ?", 20).Count(&count)


15. Raw()
Executes a raw SQL query.

db.Raw("SELECT * FROM users WHERE name = ?", "John").Scan(&result)


16. Exec()
Executes raw SQL for non-query statements (like updates or deletes).

db.Exec("UPDATE users SET age = ? WHERE name = ?", 30, "John")


17. Distinct()
Selects distinct values.

db.Distinct("name").Find(&users)


18. Pluck()
Retrieves a specific column as a slice.

db.Model(&User{}).Pluck("name", &names)


19. Unscoped()
Allows soft-deleted records to be included in the query.

db.Unscoped().Where("age > ?", 30).Find(&users)


20. Association()
For managing associations (e.g., Has Many, Belongs To).

db.Model(&user).Association("Orders").Find(&orders)
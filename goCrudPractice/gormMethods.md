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


1. First()- done
Retrieves the first record that matches the conditions.

db.First(&user)


2. Last()- done
Retrieves the last record that matches the conditions.

db.Last(&user)


3. Save()- done
Saves the value into the database (both insert and update).

db.Save(&user)


4. Create()- done
Inserts a new record into the database.

db.Create(&user)


5. Delete()- done
Deletes a record that matches the condition.

db.Delete(&user)


6. Update() / .Updates()- done //UpdateColumn()/UpdateColumns(): Similar to Update() and Updates() but bypass hooks.
Updates a specific column or multiple columns in the record.

db.Model(&user).Update("name", "John")
db.Model(&user).Updates(User{Name: "John", Age: 25})


7. Scan() -done
Scans the result into a destination (often a struct).

db.Model(&user).Scan(&result)


8. Select()- done
Specifies which fields to retrieve or modify.

db.Select("name", "age").Find(&users)


9. Order()- done
Specifies the order of the results (e.g., ascending or descending).

db.Order("age desc").Find(&users)


10. Group()- done
Used for grouping records.

db.Model(&user).Group("name").Find(&results)


11. Having()- done
Used with .Group() to specify conditions for grouped records.

db.Group("name").Having("count(name) > ?", 1).Find(&results)


12. Limit()- done
Limits the number of records retrieved.

db.Limit(10).Find(&users)


13. Offset()- done
Skips a specified number of records.

db.Offset(5).Find(&users)


14. Count()
Counts the number of records that match the query.

var count int64
db.Model(&User{}).Where("age > ?", 20).Count(&count)


15. Raw()- done
Executes a raw SQL query.

db.Raw("SELECT * FROM users WHERE name = ?", "John").Scan(&result)


16. Exec()- done
Executes raw SQL for non-query statements (like updates or deletes).

db.Exec("UPDATE users SET age = ? WHERE name = ?", 30, "John")


17. Distinct()- done
Selects distinct values.

db.Distinct("name").Find(&users)


18. Pluck()- done
Retrieves a specific column as a slice.

db.Model(&User{}).Pluck("name", &names)


19. Unscoped()- didn't work
Allows soft-deleted records to be included in the query.

db.Unscoped().Where("age > ?", 30).Find(&users)

SOFT DELETE
    Soft delete is a data management technique used in databases to mark records as deleted without actually removing them from the database. This approach allows for better data integrity and recovery options.
    In a typical users table, a soft delete might look like this:

    EXAMPLE
    id	name	    email	            DeletedAt
    1	John Doe	john@example.com	NULL
    2	Jane Doe	jane@example.com	2024-10-01 10:00:00

    > User 1 is active (not deleted).
    > User 2 is soft-deleted (the DeletedAt field has a timestamp).


20. Association() BIG BIG SUBJJECT TO BE USED WITH MULTI TABLE OPERATIONS
{
    Association() method is used to manage relationships (associations) between models. It provides an API to query, add, remove, replace, and clear associations between records

    CODE{
           package main

            import (
                "gorm.io/driver/sqlite"
                "gorm.io/gorm"
            )

            type User struct {
                ID    uint
                Name  string
                Email string
                // Has many relationships
                Orders []Order
            }

            type Order struct {
                ID     uint
                Amount float64
                UserID uint
            }

            func main() {
                // Initialize the database
                db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
                if err != nil {
                    panic("failed to connect database")
                }

                // Auto migrate the models
                db.AutoMigrate(&User{}, &Order{})

                // Sample data
                user := User{Name: "John", Email: "john@example.com"}
                order1 := Order{Amount: 100.50}
                order2 := Order{Amount: 200.75}

                // Create user and associate orders
                db.Create(&user)
                db.Model(&user).Association("Orders").Append(&order1, &order2)
            }

        METHODS:{
             > Append: Adds new records to the association
             db.Model(&user).Association("Orders").Append(&order1, &order2)

            >// Returns all orders associated with the user
            var orders []Order
            db.Model(&user).Association("Orders").Find(&orders)

            // Replace all associated orders with new ones
            newOrder := Order{Amount: 300.00}
            db.Model(&user).Association("Orders").Replace(&newOrder)

            // Remove order1 from the user's associations (does not delete from the database)
            db.Model(&user).Association("Orders").Delete(&order1)

            // Clear all associations of orders for the user
            db.Model(&user).Association("Orders").Clear()

            // Get the number of associated orders
            count := db.Model(&user).Association("Orders").Count()
            fmt.Println("Number of orders:", count)


            var user User
            db.Preload("Orders").First(&user)
            // This will load the user along with their associated orders

            Preload: GORM's Preload function is used to eager load associated data (like related tables) when you fetch a record. In this case, it means that when you load the User, GORM will also fetch the Orders that belong to that user at the same time.


            >GORM is basically creating two SQL queries for you behind the scenes:
            -SELECT * FROM users LIMIT 1; //Fetch the User's Orders (related to the user’s ID):
            -SELECT * FROM orders WHERE user_id = <user's ID>;

            Summary
            >The Association() method allows you to manipulate relationships between records.
            >You can add (Append), remove (Delete), replace (Replace), and query (Find, Count) associated records.
            >It’s particularly useful for managing hasMany, belongsTo, hasOne, and many2many relationships in GORM.

        }

    }
}

db.Model(&user).Association("Orders").Find(&orders)




<!-- UPDATE METHODS -->

1. FirstOrCreate() [might not be a gorm function]: Insert if not found, update if exists. //useful function that combines two actions: finding a record and updating it if it exists, or creating a new record if it doesn't. It is typically used when you want to either update an existing record or insert a new one, based on certain conditions.
2. FirstOrUpdate(): Update if found, do nothing if not.
3. Omit(): Exclude specific fields from being updated.
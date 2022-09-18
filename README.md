# The Payroll System implemented in Go    
Robert C. Martin - Agile Software Development, Principles, Patterns, and Practices (or Agile Principles, Patterns, and Practices in C#) 


# Notes:
### AddEmployeeTransaction  
Strategy pattern used instead of Template Method pattern (because Go doesn't have inheritance)

### ChangeEmployeeTransaction  
Strategy pattern used instead of Template Method pattern (because Go doesn't have inheritance)

### PayrollDatabase
Global instance is used like in the book. I personally don't prefer global state, but kept it for consistency purposes.

### PayTransaction
Tests have not been implemented
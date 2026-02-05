# Data Structures - Class Diagrams

## Stack Implementation

```mermaid
classDiagram
    class Stacker~T~ {
        <<interface>>
        +Size() int64
        +Find(value T) bool
        +Pop() (T, error)
        +Push(value T)
        +Top() (T, error)
        +Bottom() (T, error)
    }

    class Sizer~T~ {
        <<interface>>
        +Size() int64
    }

    class Finder~T~ {
        <<interface>>
        +Find(value T) bool
    }

    class stack~T~ {
        -size int64
        -top node~T~
        +Pop() (T, error)
        +Push(value T)
        +Find(value T) bool
        +Top() (T, error)
        +Bottom() (T, error)
        +Size() int64
    }

    class node~T~ {
        +value T
        +next node~T~
    }

    Stacker~T~ <|.. Sizer~T~
    Stacker~T~ <|.. Finder~T~
    Stacker~T~ <|.. stack~T~
    stack~T~ *-- node~T~ : top
```

## Relationships Overview

```mermaid
erDiagram
    Stacker ||--|| Sizer : extends
    Stacker ||--|| Finder : extends
    stack ||--o{ Stacker : implements
    stack ||--o{ node : contains
    node ||--o{ node : linked list
```

## Implementation Details

### Interfaces

- **`Stacker[T]`**: Main interface for stack operations
  - Extends `Sizer[T]` and `Finder[T]`
  - Provides core stack functionality: Push, Pop, Top, Bottom

- **`Sizer[T]`**: Interface for size operations
  - Single method: `Size() int64`

- **`Finder[T]`**: Interface for search operations
  - Single method: `Find(value T) bool`

### Concrete Types

- **`stack[T]`**: Generic stack implementation
  - Uses linked list internally
  - Type constraint: `T comparable`
  - Private implementation (lowercase name)

- **`node[T]`**: Linked list node
  - Generic type with value and next pointer
  - Forms the underlying data structure

### Key Design Patterns

1. **Interface Segregation**: Separate interfaces for different concerns
2. **Generic Programming**: Type-safe implementation with constraints
3. **Encapsulation**: Private implementation with public interface
4. **Linked List**: Efficient O(1) push/pop operations

## Type Constraints

All generic types use the constraint `T comparable` to enable:
- Equality comparisons in `Find()` method
- Proper type safety
- Zero value handling

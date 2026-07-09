# Simple Values and Basic Types in Go

Go is a statically typed language, which means every value has a specific type determined at compile time. This directory introduces Go's primary, built-in "simple" data types that form the building blocks of any program.

---

## 1. Integers (Signed and Unsigned)

Integers are whole numbers. Go divides integers into signed (positive, negative, and zero) and unsigned (positive and zero only) types, and provides various sizes depending on the memory required.

### Signed Integers (Can be positive, negative, or zero)

| Type | Size | Range | Common Use Case |
| :--- | :--- | :--- | :--- |
| **`int8`** | 1 byte (8 bits) | `-128` to `127` | Tiny counters, low-memory buffers. |
| **`int16`** | 2 bytes (16 bits) | `-32,768` to `32,767` | Medium-small numbers. |
| **`int32`** | 4 bytes (32 bits) | `-2,147,483,648` to `2,147,483,647` | Large counts, games, general 32-bit arithmetic. |
| **`int64`** | 8 bytes (64 bits) | `-9,223,372,036,854,775,808` to `9,223,372,036,854,775,807` | Unix timestamps, massive numbers, database IDs. |
| **`int`** | Platform-dependent | 32-bit or 64-bit | **Default choice** for standard whole numbers. |

### Unsigned Integers (Positive numbers and zero only)

| Type | Size | Range | Common Use Case |
| :--- | :--- | :--- | :--- |
| **`uint8`** | 1 byte (8 bits) | `0` to `255` | Storing pixel colors (RGB), raw binary data. |
| **`uint16`** | 2 bytes (16 bits) | `0` to `65,535` | Network port numbers (e.g., `8080`). |
| **`uint32`** | 4 bytes (32 bits) | `0` to `4,294,967,295` | IP addresses, hashing algorithms. |
| **`uint64`** | 8 bytes (64 bits) | `0` to `18,446,744,073,709,551,615` | Large unsigned values, cryptography. |
| **`uint`** | Platform-dependent | 32-bit or 64-bit | Machine-word sized counts. |
| **`uintptr`** | Platform-dependent | Fits a memory pointer | Low-level pointer arithmetic and C-language bindings. |

> [!NOTE]
> `int` and `uint` are distinct types from `int32` / `int64` even if they have the same size on your machine. You cannot perform operations between them without explicit casting (e.g., `int64(myInt)`).

---

## 2. Floating-Point Numbers

Floats represent numbers that contain decimal points or fractions. Go supports two precision formats based on the IEEE-754 standard:

| Type | Size | Precision (Significant Digits) | Common Use Case |
| :--- | :--- | :--- | :--- |
| **`float32`** | 4 bytes (32 bits) | `~7 decimal digits` | Graphics buffers (OpenGL/Vulkan), machine learning vectors where memory usage is critical. |
| **`float64`** | 8 bytes (64 bits) | `~15 decimal digits` | **Default choice** for calculations, scientific computing, financial decimals (though floats should generally be avoided for exact money values). |

### Key Differences: `float32` vs `float64`
* **Precision**: `float64` is significantly more precise than `float32`. Calculations using `float32` accumulate rounding errors much faster.
* **Compatibility**: Go's math package functions (like `math.Sqrt`, `math.Sin`, etc.) accept and return `float64`. Working with `float32` requires constant type casting.
* **Performance**: Modern 64-bit CPUs handle `float64` calculations just as fast as `float32`.

---

## 3. Complex Numbers

For advanced mathematics (e.g., engineering, physics, signal processing), Go natively supports complex numbers consisting of a real and an imaginary part:

* **`complex64`**: Real and imaginary parts are both `float32` values.
* **`complex128`**: Real and imaginary parts are both `float64` values (this is the **default** when using built-in functions like `complex(real, imag)`).

---

## 4. Strings (`string`)

A string is an immutable sequence of characters, commonly used to represent text (e.g., `"Hello, Go!"`).
* **Double Quotes (`"..."`)**: Standard string literals. Supports escape sequences like `\n` (newline) and `\t` (tab).
* **Backticks (`` `...` ``)**: Raw string literals. They do not interpret escape sequences and can span multiple lines.
* **UTF-8 Encoding**: Go strings natively support characters from any language, emojis, and symbols because they are UTF-8 encoded by default.

---

## 5. Booleans (`bool`)

Booleans represent logical truth values and can only have one of two states:
* **`true`**
* **`false`**
They are primarily used in conditional statements (`if`/`else`) and loops.

---

## 6. Specialized Type Aliases

Go has two special aliases for integer types that are widely used when dealing with characters and text:

* **`byte`**: An alias for `uint8`. Represents raw binary data or a single ASCII character.
* **`rune`**: An alias for `int32`. A rune represents a single Unicode character or code point (e.g., `'A'`, `'ই'`, `'😊'`). Runes are enclosed in **single quotes**.

---

## Default "Zero Values"

In Go, variables declared without an initial value are automatically assigned their **zero value**:

| Type Group | Zero Value |
|------------|------------|
| Numeric (`int`, `float`, `complex`) | `0` / `0.0` / `(0+0i)` |
| Booleans (`bool`) | `false` |
| Strings (`string`) | `""` (empty string) |

---

## How to Run the Code in this Directory

To execute the example program demonstrating these simple values, run:

```bash
go run main.go
```

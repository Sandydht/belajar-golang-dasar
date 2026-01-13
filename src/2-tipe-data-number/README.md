# Tipe Data Integer (1)
| Tipe Data | Nilai Minimum | Nilai Maksimum |
|-----------|---------------|----------------|
| int8 | -128 | 127 |
| int16 | -32768 | 32767 |
| int32 | -21474836448 | 21474836447 |
| int64 | -9223372036854775808 | 9223372036854775807 |

# Tipe Data Integer (2)
| Tipe Data | Nilai Minimum | Nilai Maksimum |
|-----------|---------------|----------------|
| uint8 | 0 | 255 |
| uint16 | 0 | 65535 |
| uint32 | 0 | 4294967295 |
| uint64 | 0 | 18446744073709551615 |

# Tipe Data Floating Point
| Tipe Data | Nilai Minimum | Nilai Maksimum |
|-----------|---------------|----------------|
| float32 | 1.18x10<sup>-38</sup> | 3.4x10<sup>38</sup> |
| float64 | 2.23x10<sup>-308</sup> | 1.80x10<sup>308</sup> |
| complex64 | complex numbers with float32 real and imaginary parts. |
| complex128 | complex numbers with float64 real and imaginary parts. |

# Alias
| Tipe Data | Alias Untuk | 
|-----------|-------------|
| byte | uint8 |
| rune | int32 |
| int | Minimal int32 |
| uint | Minimal uint32 |

# Kode Program Number
```go
package main

import "fmt"

func main() {
  fmt.println("Satu = ", 1)
  fmt.println("Dua = ", 2)
  fmt.println("Tiga Koma Lima = ", 3.5)
}
```


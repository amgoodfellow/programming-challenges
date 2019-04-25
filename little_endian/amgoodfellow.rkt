#lang racket

; Algorithm in question
(define (num->little-endian str base)
  (cond
    [(= base 16) (apply string-append (reverse (string->two-chars (validate-input str))))]
    [(= base 10)
     (hex->dec (num->little-endian (validate-input (dec->hex (string->number str))) 16))]
    [else "Invalid input"]))

; Helper method that splits strings into two char lists
(define (string->two-chars str)
  (if (zero? (string-length str)) null
    (cons (substring str 0 2) (string->two-chars (substring str 2)))))

; Helper method converting from base 10 to base 16
(define (dec->hex num)
  (if (< num 10) (number->string num)
  (string-append (dec->hex (floor (/ num 16)))
          (hex-value (remainder num 16)))))

; Mapping of base 10 values to base 16 ones
(define (hex-value num)
  (cond
    [(= num 15) "F"]
    [(= num 14) "E"]
    [(= num 13) "D"]
    [(= num 12) "C"]
    [(= num 11) "B"]
    [(= num 10) "A"]
    [(<= num 10) (number->string num)]))

; Make sure that hex values have even length
(define (validate-input str)
  (if (zero? (remainder (string-length str) 2))
      str
      (string-append "0" str)))

; Lazy, inefficient, and easy way of converting back to base 10
(define (hex->dec num)
  (number->string (string->number (string-append "#x" num))))

; Test cases
(num->little-endian "35582" 10)
(num->little-endian "8AFE" 16)
(num->little-endian "332" 16)


(defun get-num (str)
  (/ (apply '+ (map 'list (lambda (c) (char-int c)) (coerce str 'list))) (length str)))

;;;; This code calculates PI 
;;;; The formular is from Kelallur Nilakantha Somayaji (1444-1544)

;;;; Run this with: sbcl --script pi.lsp

(proclaim '(optimize speed))

(defun pi_step (cnt val num e) 
  (if (not (= 0 cnt))
    (pi_step (- cnt 1) 
        (+ val (/ num 
          (- (expt e 3) e)
        ))
        (* num -1)
        (+ e 2)
    )
    val
  )
)
(defvar p (pi_step 20000 3 4 3))
(format t "~% pi=~d " (float p))
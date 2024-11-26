syntax match hhlCall "%c"
syntax match hhlRstSent "%rst_sent"
syntax match hhlRstRcvd "%rst_rcvd"
syntax match hhlState "%state"
syntax match hhlTime "%t"
syntax match hhlFreq "%f"
syntax match hhlDate "%d"
syntax match hhlMode "%m"
syntax match hhlMySotaRef "%my_sota_ref"
syntax match hhlMyCall "%my_call"
syntax match hhlComment "\" [^%].*"

hi! def link hhlComment Comment
hi! def link hhlCall  Identifier
hi! def link hhlRstSent  Identifier
hi! def link hhlRstRcvd  Identifier
hi! def link hhlState  Identifier
hi! def link hhlTime  Identifier
hi! def link hhlFreq  Identifier
hi! def link hhlDate  Identifier
hi! def link hhlMode  Identifier
hi! def link hhlMySotaRef  Identifier
hi! def link hhlMyCall  Identifier
hi! def link hhlMyCall  Identifier


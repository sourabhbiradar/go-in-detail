# defer

Deferred functions' arguements are evaluated immediatelyy but are executed at last.

If there are multiple deferred funcs , the execution takes place as LIFO.

defer is used to close opened files or to recover from panic.

defer are not executed when runtime error,panic or os.Exit() are incurred.

Transaction log format
Sequence number
    A unique record ID, in monotonically increasing order.
Event type
    A descriptor of the type of action taken; this can be PUT or DELETE.
Key
    A string containing the key affected by this transaction.
Value
    If the event is a PUT, the value of the transaction.
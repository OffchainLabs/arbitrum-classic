from .struct import Struct
from . import queue_tup
from .. import value

inbox_ctx = Struct("inbox_ctx", [
    ("queue", queue_tup.typ),
    ("nodeAlreadySeen", value.TupleType())
])
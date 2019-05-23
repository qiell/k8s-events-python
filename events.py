from ctypes import *

lib = cdll.LoadLibrary("./events.so")
lib.GetEvents.restype = c_char_p
print(lib.GetEvents())
import sys

identity = lambda x: x

class Slot(object):
    def __init__(self, value=identity, vitality=10000):
        self.vitality = vitality
        self.value = value

    def set_vita(self, number):
        self.vitality = number
        return True

    def get_vita(self, number):
        return self.vitality

    def set_val(self, value):
        self.value = value
        return True

    def get_val(self, value):
        return self.value

def valid_slot(thing, slots):
    if isinstance(thing, int):
        if 0 <= thing < len(slots):
            return thing
        else:
            raise TypeError('Slot index is out of bounds')
    else:
        raise TypeError('Wrong type')

def live_slot(thing, slots):
    val_slot(thing, slots)
    if slots[thing]['vitality'] > 0:
        thing
    else:
        raise TypeError('Dead slot')


def val_bound(thing, top=65535, bot=-1, ref=int):
    fail = False;
    if isintance(thing, ref):
        if type(top) == int and thing >= top:
            raise TypeError('Value too large')
        if type(bot) == int and thing < bot:
            raise TypeError('Value too small')
        return thing
    else:
        raise TypeError('Wrong type')

def catch(call, *args, **kwargs):
    try:
        call(*args, **kwargs)
        return True
    except:
        return False

def set_slot(i, value, slots):
    slots[i] = value
    return True

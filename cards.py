import lib

__all__ = ['cards']

cards = {
    'I':        lambda x, **kwargs: lib.identity(x),
    'zero':     lambda **kwargs: 0,
    'succ':     lambda x, **kwargs: lib.val_bound(x, bot=False) + 1,
    'dbl':      lambda x, **kwargs: min(x * 2, 65535),
    'get':      lambda x, pslots, **kwargs: pslots[lib.live_slot(x, pslots)],
    'put':      lambda x, **kwargs: lib.identity.
    'S':        lambda f, **kwargs:
                    lambda g:
                        lambda x: f(x)(g(x)),
    'K':        lambda x, **kwargs:
                    lambda y: x,
    'inc':      lambda x, pslots, **kwargs:
                    lib.valid_slot(x, pslots) \
                    and lib.catch(lib.live_slot, x, pslots) \
                    and pslots[x].set_vita(min(pslots[x].get_vita+1), 65535) \
                    and lib.identity,
    'dec':      lambda x, oslots, **kwargs:
                    lib.valid_slot(x, oslots) \
                    and lib.catch(lib.live_slot, x, oslots) \
                    and oslots[x].set_vita(max(oslots[x].get_vita-1), 0) \
                    and lib.identity,
    'attack':   lambda i, **kwargs:
                    lambda j:
                        lambda n: lib.identity,
    'help':   lambda i, **kwargs:
                    lambda j:
                        lambda n: lib.identity,
    'copy':     lambda x, pslots, oslots: x,
    'zombie':   lambda x, pslots, oslots: x,
}

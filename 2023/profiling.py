import datetime
import cProfile


def profile_it(func):
    def profiled_func(*args, **kwargs):
        profile = cProfile.Profile()
        try:
            profile.enable()
            result = func(*args, **kwargs)
            profile.disable()
            return result
        finally:
            profile.dump_stats(
                f"{func.__name__}_profile_{datetime.datetime.now().isoformat()}.txt"
            )

    return profiled_func

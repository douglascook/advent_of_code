import datetime
import cProfile


def profile_it(func, repeats=1):
    def profiled_func(*args, **kwargs):
        profile = cProfile.Profile()
        try:
            profile.enable()
            for i in range(repeats):
                result = func(*args, **kwargs)
            profile.disable()
            return result
        finally:
            profile.dump_stats(
                f"./profiling/{func.__name__}_profile_{datetime.datetime.now().isoformat()}.txt"
            )

    return profiled_func

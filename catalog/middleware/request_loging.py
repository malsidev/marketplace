import time

from fastapi import Request

async def requests_log(request: Request, call_next):
    start_time = time.time()
    response = await call_next(request)

    duration = time.time() - start_time

    print(
        f"{request.method}| "
        f"id {request.headers.get("X-User-ID")} |"
        f"{request.url.path} |"
        f"{response.status_code} |"
        f"{duration:.3f}s"
    )

    return response
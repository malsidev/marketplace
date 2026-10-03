import json
import time
from database.kafka_client import producer
from fastapi import Request


async def requests_log(request: Request, call_next):
    start_time = time.time()

    response = await call_next(request)

    duration = time.time() - start_time

    user_id = request.headers.get("X-User-ID")

    print(
        f"{request.method} | "
        f"id {user_id} | "
        f"{request.url.path} | "
        f"{response.status_code} | "
        f"{duration:.3f}s"
    )

    message = {
        "event": "action",
        "UserID": int(user_id),
        "Method": request.method,
        "Path": request.url.path,
        "StatusCode": response.status_code
    }
    try:
        await producer.send(
            "log_users",
            json.dumps(message).encode("utf-8"),
        )
    except Exception as e:
        print(f"Kafka error: {e}")
    return response
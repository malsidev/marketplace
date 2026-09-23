from fastapi import FastAPI 

from middleware.request_loging import requests_log
from routers.init import setup_routers
from services.minio import create_bucket_if_not_exists

app  = FastAPI()

@app.on_event("startup")
async def startup():
    create_bucket_if_not_exists()

app.middleware("http")(requests_log)
setup_routers(app)

@app.get("/")
async def root():
    return {"message": "Hello World"}


if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8001)

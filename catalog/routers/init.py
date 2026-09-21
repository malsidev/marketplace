from fastapi import FastAPI

from routers.catalog import router as catalog_router

def setup_routers(app: FastAPI):
    app.include_router(catalog_router)